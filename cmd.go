package main

import (
	"fmt"
	"os"

	"k8s.io/api/core/v1"

	"github.com/spf13/cobra"
)

func cmdScheduleNodes() *cobra.Command {
	var kubeconfig, selector, fromWindow, toWindow string
	var reboot, outofdateNodes bool
	var command = &cobra.Command{
		Use:   "schedule",
		Short: "schedule nodes for update",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			client, err := newKube(kubeconfig, fromWindow, toWindow, reboot)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			var nodeslist *v1.NodeList
			if outofdateNodes == true {
				nodes, err := client.getNodesNotMatchingMasterVersion(selector)
				nodeslist = nodes
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}
			} else {
				nodes, err := client.getNodes(selector)
				nodeslist = nodes
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}
			}
			err = client.annotateNodes(nodeslist)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		},
	}

	command.Flags().StringVar(&kubeconfig, "kube.config", "", "path to kube config")
	command.Flags().StringVar(&selector, "selector", "", "lable selector")
	command.Flags().StringVar(&fromWindow, "schedule.fromWindow", "", "schedule from \"hh:mm AM/PM\" time format to start updates")
	command.Flags().StringVar(&toWindow, "schedule.toWindow", "", "schedule to \"hh:mm AM/PM\" time format to stop updates")
	command.Flags().BoolVar(&reboot, "terminate", false, "flag for termination to true")
	command.Flags().BoolVar(&outofdateNodes, "out.of.date.nodes", false, "if set all nodes with a older kubelet version then the master will be terminated")

	return command
}

func cmdVersion() *cobra.Command {
	var command = &cobra.Command{
		Use:   "version",
		Short: "get version",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(getVersion())
		},
	}
	return command
}

func runCmd() error {
	var rootCmd = &cobra.Command{Use: "k8s-node-update-scheduler"}
	rootCmd.AddCommand(cmdVersion())
	rootCmd.AddCommand(cmdScheduleNodes())

	err := rootCmd.Execute()
	if err != nil {
		return fmt.Errorf("%v", err.Error())
	}
	return nil
}
