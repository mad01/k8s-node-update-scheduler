package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func cmdScheduleNodes() *cobra.Command {
	var kubeconfig, selector, fromWindow, toWindow string
	var reboot bool
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
			nodes, err := client.getNodes(selector)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			err = client.annotateNodes(nodes)
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
	command.Flags().BoolVar(&reboot, "reboot", false, "set reboot flag to true")

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
