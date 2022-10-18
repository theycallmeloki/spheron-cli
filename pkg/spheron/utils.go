package spheron

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// create config file function
func SetSpheronConfigFile() {
	// get home directory
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// xdg config directory
	ConfigDir = filepath.Join(home, ".config", ConfigName)

	// create config file
	ConfigPath = filepath.Join(ConfigDir, ConfigName + "." + ConfigType)
	
	viper.SetConfigFile(ConfigPath)
	viper.SetConfigName(ConfigName)
	viper.SetConfigType(ConfigType)
	viper.AddConfigPath(ConfigDir)
	viper.SetDefault("secret", "")
	viper.SetDefault("organization", "")
	
}

// sanitized input function
func SanitizeInput(query string) string {
	fmt.Print(query)
	reader := bufio.NewReader(os.Stdin)
	inputLine, _ := reader.ReadString('\n')
	inputLine = strings.TrimSpace(inputLine)
	return inputLine
}

// is file exist function, used to look for existing config file
func FileExists(path string) bool {
    _, err := os.Stat(path)
    return !os.IsNotExist(err)
}

// function to read in the local configuration

func ReadLocalConfig() *Config {
	
	err := viper.ReadInConfig()

	if err == nil {
		c := Config{}
		c.Secret = viper.GetString("secret")
		c.Organization = viper.GetString("organization")
		return &c
	} else {
		e := Config{}
		return &e
	}

}

// function to write the configuration in viper memory to the local config file
func WriteLocalConfig() {

	_, fErr := os.Stat(ConfigDir)
	if !os.IsExist(fErr) {
		err := os.Mkdir(ConfigDir, 0700) // since it's a user directory, we don't need to worry about group or other permissions
		if !os.IsExist(err) && err != nil {
			fmt.Println(err)
		}
	}

	_, err := os.Stat(ConfigPath)

	if !os.IsExist(err) {
		if _, err := os.Create(ConfigPath); err != nil { // perm 0666
			fmt.Println(err)
		}
	}
	
	if err := viper.WriteConfig(); err != nil {
		fmt.Println(err)
	}

}