package command

import (
	"errors"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

func init() {
	userCmd.Flags().BoolP("web", "w", false, "Open the User profile in the browser")
	RootCmd.AddCommand(userCmd)
}

var userCmd = &cobra.Command{
	Use:   "user <command>",
	Short: "See how another GitHub user expresses themself!",
	Long:  "View a user's profile README. Profile README.md files are located in the user's eponymously named repo.",
	Example: heredoc.Doc(`
	$ gh user octocat
	`),
	Args:        cobra.ExactArgs(1),
	RunE:        userView,
	Annotations: map[string]string{"IsCore": "true"}}

func userView(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("You must supply a GitHub username.")
	}
	return repoView(cmd, []string{fmt.Sprintf("%s/%s", args[0], args[0])})
}
