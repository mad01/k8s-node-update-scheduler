package main

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"encoding/json"
	"fmt"
	"strings"

	"github.com/blang/semver"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/strategicpatch"

	"github.com/mad01/k8s-node-terminator/pkg/annotations"
)

func k8sGetClientConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}

func k8sGetClient(kubeconfig string) (*kubernetes.Clientset, error) {
	config, err := k8sGetClientConfig(kubeconfig)
	if err != nil {
		return nil, err
	}

	// Construct the Kubernetes client
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func newKube(kubeconfig, fromTime, toTime string, reboot bool) (*Kube, error) {
	client, err := k8sGetClient(kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("failed to get client: %v", err.Error())
	}
	a, err := newAnnotations(fromTime, toTime)
	a.reboot = fmt.Sprintf("%v", reboot)
	if err != nil {
		return nil, fmt.Errorf("failed to create new kube: %v", err.Error())
	}
	k := Kube{
		client:      client,
		annotations: a,
	}
	return &k, nil
}

// Kube kubernetes connection struct
type Kube struct {
	client      *kubernetes.Clientset
	annotations *Annotations
}

func (k *Kube) getNodes(selector string) (*v1.NodeList, error) {
	nodes, err := k.client.Core().Nodes().List(metav1.ListOptions{LabelSelector: selector})
	if err != nil {
		return nil, fmt.Errorf("failed to get nodes list %v", err.Error())
	}
	return nodes, nil
}

func (k *Kube) getKubeletVersion(node *v1.Node) (*semver.Version, error) {
	rawString := node.Status.NodeInfo.KubeletVersion
	versionString := strings.Replace(rawString, "v", "", -1)

	version, err := semver.Parse(versionString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse version %v", err.Error())
	}
	return &version, nil
}

func (k *Kube) getNodesNotMatchingMasterVersion(selector string) (*v1.NodeList, error) {
	masters, err := k.getNodes("kubernetes.io/role=master")
	if err != nil {
		return nil, fmt.Errorf("failed to get master nodes %v", err.Error())
	}
	var masterVersion *semver.Version
	if len(masters.Items) >= 1 {
		masterNode := masters.Items[0]
		version, err := k.getKubeletVersion(&masterNode)
		masterVersion = version
		if err != nil {
			return nil, fmt.Errorf("failed to get kubelet version %v", err.Error())
		}
	} else {
		return nil, fmt.Errorf("no masters found")
	}

	nodeList, err := k.getNodes(selector)
	if err != nil {
		return nil, fmt.Errorf("failed to get nodes %v", err.Error())
	}

	var filteredList v1.NodeList
	for _, node := range nodeList.Items {
		nodeVersion, err := k.getKubeletVersion(&node)
		if err != nil {
			return nil, fmt.Errorf("failed to get kubelet version %v", err.Error())
		}
		// check if nodeversion < masterversion
		if nodeVersion.LT(*masterVersion) {
			filteredList.Items = append(filteredList.Items, node)
		}
	}
	return &filteredList, nil
}

func (k *Kube) annotateNodes(nodeList *v1.NodeList) error {
	fmt.Printf("adding annotations %v=\"%v\" %v=\"%v\" %v=\"%v\"\n",
		annotations.NodeAnnotationReboot, k.annotations.reboot,
		annotations.NodeAnnotationFromWindow, k.annotations.timeWindow.FromString(),
		annotations.NodeAnnotationToWindow, k.annotations.timeWindow.ToString(),
	)
	for _, node := range nodeList.Items {
		err := k.annotatePatchNode(&node)
		if err != nil {
			return fmt.Errorf("failed to annotation node %v %v", node.GetName(), err.Error())
		}
	}
	return nil
}

func (k *Kube) annotatePatchNode(node *v1.Node) error {
	oldData, err := json.Marshal(node)
	if err != nil {
		return fmt.Errorf("failed to Marshal old node %v", err.Error())
	}

	nodeCopy := node.DeepCopy()
	a := nodeCopy.GetAnnotations()
	a[annotations.NodeAnnotationFromWindow] = k.annotations.timeWindow.FromString()
	a[annotations.NodeAnnotationToWindow] = k.annotations.timeWindow.ToString()
	a[annotations.NodeAnnotationReboot] = k.annotations.reboot
	nodeCopy.SetAnnotations(a)

	newData, err := json.Marshal(nodeCopy)
	if err != nil {
		return fmt.Errorf("failed to Marshal new node %v", err.Error())
	}

	patchBytes, err := strategicpatch.CreateTwoWayMergePatch(oldData, newData, v1.Node{})
	if err != nil {
		return fmt.Errorf("failed to create patch %v", err.Error())
	}

	_, err = k.client.Core().Nodes().Patch(node.GetName(), types.StrategicMergePatchType, patchBytes)
	if err != nil {
		return fmt.Errorf("failed to patch node %v %v", node.GetName(), err.Error())
	}

	fmt.Printf("annotated node %v\n", node.GetName())
	return nil
}
