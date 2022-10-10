package spheron

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version string = "-DEV_BUILD"

var rootCmd = &cobra.Command{
    Use:  "spheronctl",
	Version: version,
    Short: "spheron - a CLI to manage your spheron deployments",
    Long: `

   _____       _                          
  / ____|     | |                         
 | (___  _ __ | |__   ___ _ __ ___  _ __  
  \___ \| '_ \| '_ \ / _ \ '__/ _ \| '_ \ 
  ____) | |_) | | | |  __/ | | (_) | | | |
 |_____/| .__/|_| |_|\___|_|  \___/|_| |_|
        | |                               
        |_|                                                                                              
   
Spheron CLI - v` + version,
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
        os.Exit(1)
    }
}