package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/kyokomi/emoji"
	"github.com/spf13/cobra"
)

const (
	Version = "0.1"
)

var kubeconfig string

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	ExitfIfError(checkDependencies(), "Cannot operate")

	rootCmd := &cobra.Command{
		Use:           "carrier",
		Short:         "Carrier cli",
		Long:          `carrier cli is the official command line interface for Carrier PaaS `,
		Version:       fmt.Sprintf("%s", Version),
		SilenceErrors: true,
	}

	RegisterInstall(rootCmd)

	ExitIfError(ensureKubeConfig())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func ensureKubeConfig() error {
	if kubeconfig = os.Getenv("KUBECONFIG"); kubeconfig == "" {
		return errors.New("KUBECONFIG environment variable not set!")
	}

	return nil
}

func checkDependencies() error {
	ok := true

	dependencies := []struct {
		CommandName string
	}{
		{CommandName: "kubectl"},
		{CommandName: "helm"},
	}

	for _, dependency := range dependencies {
		_, err := exec.LookPath(dependency.CommandName)
		if err != nil {
			fmt.Println(emoji.Sprintf(":fire:Not found: %s", dependency.CommandName))
			ok = false
		}
	}

	if ok {
		return nil
	}

	return errors.New("Please check your PATH, some of our dependencies were not found")
}
