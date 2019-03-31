package version

import (
	"fmt"
	"github.com/alimy/mir-music/cmd/core"
	"github.com/spf13/cobra"

	appVer "github.com/alimy/mir-music/version"
)

func init() {
	// version sub-command
	versionCmd := &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s (APIVersion:%s)\nBuildTime:%s\nBuildGitSHA:%s\n",
				appVer.Version, appVer.ApiVersion, appVer.BuildTime, appVer.GitHash)
		},
	}

	// Register version sub-command
	core.Register(versionCmd)
}
