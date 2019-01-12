package version

import (
	"fmt"
	"github.com/alimy/mir-music/cmd"
	"github.com/spf13/cobra"
)

var (
	// App version
	Version = "0.0.0"

	// Api Version
	ApiVersion = "v1"

	// GitHash Value will be set during build
	GitHash = "Not provided"

	// BuildTime Value will be set during build
	BuildTime = "Not provided"

	// version sub-command
	versionCmd = &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s (APIVersion:%s)\nBuildTime:%s\nBuildGitSHA:%s\n",
				Version, ApiVersion, BuildTime, GitHash)
		},
	}
)

func init() {
	// Register version sub-command
	cmd.Register(versionCmd)
}
