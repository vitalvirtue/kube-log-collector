/*
Copyright © 2024 NAME HERE ozkherdem@gmail.com

*/
package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/vitalvirtue/kube-log-collector/internal/collector"
    "github.com/vitalvirtue/kube-log-collector/internal/kubernetes"
    "github.com/vitalvirtue/kube-log-collector/pkg/types"
)

var rootCmd = &cobra.Command{
    Use:   "kube-log-collector",
    Short: "Collect logs from Kubernetes pods",
    Long: `kube-log-collector is a CLI tool to collect logs from Kubernetes pods
           based on namespace or pod labels.`,
    Run: func(cmd *cobra.Command, args []string) {
        namespace, _ := cmd.Flags().GetString("namespace")
        label, _ := cmd.Flags().GetString("label")
        output, _ := cmd.Flags().GetString("output")

        options := types.CollectorOptions{
            Namespace:  namespace,
            PodLabel:   label,
            OutputFile: output,
        }

        client, err := kubernetes.NewClient("")
        if err != nil {
            fmt.Printf("Error creating Kubernetes client: %v\n", err)
            os.Exit(1)
        }

        collector := collector.NewCollector(client, options)
        if err := collector.Collect(); err != nil {
            fmt.Println("Error:", err)
            os.Exit(1)
        }
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    rootCmd.Flags().StringP("namespace", "n", "", "Kubernetes namespace")
    rootCmd.Flags().StringP("label", "l", "", "Pod label selector")
    rootCmd.Flags().StringP("output", "o", "collected_logs.txt", "Output file")
}
