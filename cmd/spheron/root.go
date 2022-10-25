package spheron

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var version string = "-DEV_BUILD"
var spheronAsciiArt = `
   _____       _                          
  / ____|     | |                         
 | (___  _ __ | |__   ___ _ __ ___  _ __  
  \___ \| '_ \| '_ \ / _ \ '__/ _ \| '_ \ 
  ____) | |_) | | | |  __/ | | (_) | | | |
 |_____/| .__/|_| |_|\___|_|  \___/|_| |_|
        | |                               
        |_|                               
`
var ConfigName string = "spheron"
var ConfigType string = "json"
var ConfigDir string
var ConfigPath string

var rootCmd = &cobra.Command{
    Use:  "spheronctl",
	Version: version,
    Short: "spheron - a CLI to manage your spheron deployments",
    Long: `
` + spheronAsciiArt + `
Spheron CLI - v` + version,
    PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
			return initializeConfig(cmd)
		},
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
        os.Exit(1)
    }
}

func initializeConfig(cmd *cobra.Command) error {
    
    // create a spheron configuration file if it doesn't exist
    SetSpheronConfigFile()

    if(!FileExists(ConfigPath)) {
        fmt.Println(spheronAsciiArt)
        fmt.Println("Welcome to SpheronCTL!")
        fmt.Println("\n")
        fmt.Println("No configuration file found.")
        fmt.Println("Start by configuring spheronctl with your secret API key.")
        fmt.Println("\n")
        fmt.Println("Example usage: \n")
        fmt.Println("spheronctl configure")
        fmt.Println("spheronctl configure --secret=<YOUR_SECRET_API_KEY>")
        fmt.Println("\n")
    }

	// Attempt to read the config file, gracefully ignoring errors
	// caused by a config file not being found. Return an error
	// if we cannot parse the config file.
	if err := viper.ReadInConfig(); err != nil {
		// It's okay if there isn't a config file
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

    return nil
}
