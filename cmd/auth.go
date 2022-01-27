package cmd

import (
	"fmt"

	"github.com/Privado-Inc/privado/pkg/docker"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth <email>",
	Short: "Request license for Privado CLI",
	Args:  cobra.ExactArgs(1),
	Run:   auth,
}

func auth(cmd *cobra.Command, args []string) {
	email := args[0]
	fmt.Println("This is the entered arg: ", email)

	if len(args) < 1 {
		exit(fmt.Sprintln("Expected Argument: <email>"), true)
	}

	err := docker.RunImageWithArgs([]string{"auth", email}, &docker.ContainerVolumes{}, &docker.ContainerPorts{})
	if err != nil {
		exit(fmt.Sprintf("Received error: \n%s", err), true)
	}
}

func init() {
	rootCmd.AddCommand(authCmd)
}
