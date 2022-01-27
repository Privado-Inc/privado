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
	fmt.Println("> Requesting license for: ", email)

	runImageOptions := &docker.RunImageOptions{Args: []string{"auth", email}}
	err := docker.RunImageWithArgs(runImageOptions)
	if err != nil {
		exit(fmt.Sprintf("Received error: \n%s", err), true)
	}
}

func init() {
	rootCmd.AddCommand(authCmd)
}
