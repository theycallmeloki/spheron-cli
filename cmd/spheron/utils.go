package spheron

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/manifoldco/promptui"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
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

// sanitized select multiple and return Selected as string array function
func SanitizeSelectMultiple(items []string, initPrompt string, repeatPrompt string) ([]string, error) {
	repeat := true

	var selectedItems []string

	for repeat {

		envPrompt := promptui.Select{
		Label: initPrompt,
		Items: items,
		}

		_, result, err := envPrompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return selectedItems, err
		}

		fmt.Printf("You chose %q\n", result)
		selectedItems = append(selectedItems, result)

		againPrompt := promptui.Select{
			Label: repeatPrompt,
			Items: []string{"Yes", "No"},
		}

		_, again, err := againPrompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return selectedItems, err
		}
		
		fmt.Printf("You chose %q\n", again)

		if again == "No" {
			repeat = false
		}

	}

	return selectedItems, nil

}

// sanitizied input multiple and return inputs as string array function 
func SanitizeInputMultiple(initPrompt string, repeatPrompt string) ([]string, error) {
	repeat := true

	var inputs []string

	for repeat {

		input := SanitizeInput(initPrompt)
		inputs = append(inputs, input)

		againPrompt := promptui.Select{
			Label: repeatPrompt,
			Items: []string{"Yes", "No"},
		}

		_, again, err := againPrompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return inputs, err
		}

		fmt.Printf("You chose %q\n", again)

		if again == "No" {
			repeat = false
		}

	}

	return inputs, nil

}

// sanitized fixed select and return selected input as string function
func SanitizeFixedSelect(items []string, prompt string) (string, error) {
	envPrompt := promptui.Select{
		Label: prompt,
		Items: items,
	}

	_, result, err := envPrompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	fmt.Printf("You chose %q\n", result)

	return result, nil

}

// is file exist function, used to look for existing config file
func FileExists(path string) bool {
    _, err := os.Stat(path)
    return !os.IsNotExist(err)
}

// contains checks if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// function to read in the local configuration

func ReadLocalConfig() *spheron.Config {
	
	err := viper.ReadInConfig()

	if err == nil {
		c := spheron.Config{}
		c.Secret = viper.GetString("secret")
		c.Organization = viper.GetString("organization")
		return &c
	} else {
		e := spheron.Config{}
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