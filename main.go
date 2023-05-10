package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:]
	clusterName := ""
	kubectlArgs := make([]string, 0)

	for i := 0; i < len(args); i++ {
		if args[i] == "-m" {
			clusterName = args[i+1]
			i++
		} else {
			kubectlArgs = append(kubectlArgs, args[i])
		}
	}

	if clusterName == "" {
		fmt.Println("No cluster name provided, exiting.")
		os.Exit(1)
	}

	useContextCmd := exec.Command("kubectl", "config", "use-context", clusterName)
	useContextCmd.Stdout = nil
	useContextCmd.Stderr = nil

	err := useContextCmd.Run()
	if err != nil {
		fmt.Printf("Error switching context: %v\n", err)
		os.Exit(1)
	}

	kubectlCmd := exec.Command("kubectl", kubectlArgs...)
	kubectlCmd.Stdout = os.Stdout
	kubectlCmd.Stderr = os.Stderr

	err = kubectlCmd.Run()
	if err != nil {
		fmt.Printf("Error executing kubectl command: %v\n", err)
		os.Exit(1)
	}
}
