package spheron

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
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

var domainTypeList = []string{"Domain", "Subdomain"}

// a function that switches case for a selected domain type value on the CLI
func GetDomainTypeEnum(domainType string) string {
	switch domainType {
		case "Domain":
			return "domain"
		case "Subdomain":
			return "subdomain"
		default:
			// TODO: Check if this is the right way to handle this
			return "domain"
	}
}

var protocolList = []string{"Arweave", "Skynet", "Filecoin", "Pinata"}

// a function that switches case for a selected protocol value on the CLI
func GetProtocolEnum(protocol string) string {
	switch protocol {
		case "Arweave":
			return "arweave"
		case "Skynet":
			return "skynet"
		case "Filecoin":
			return "ipfs-filecoin"
		case "Pinata":
			return "ipfs-pinata"
		default:
			// TODO: Check if this is the right way to handle this
			return "ipfs-pinata"
	}
}


var frameworkList = []string{"No Framework - Simple Javascript App", "Create React App", "Vue", "Angular", "Next.js", "Preact", "Nuxt 2", "Docusaurus", "Hugo", "Eleventy", "Svelte", "Gatsby", "Sanity", "Ionic-React", "Vite", "Scully", "Stencil", "Brunch", "Ionic-Angular"}

// a function that switches case for a selected framework value on the CLI
func GetFrameworkEnum(framework string) string {
	switch framework {
		case "No Framework - Simple Javascript App":
			return "static"
		case "Create React App":
			return "react"
		case "Vue":
			return "vue"
		case "Angular":
			return "angular"
		case "Next.js":
			return "next"
		case "Preact":
			return "preact"
		case "Nuxt 2":
			return "nuxt2"
		case "Docusaurus":
			return "docusaurus"
		case "Hugo":
			return "hugo"
		case "Eleventy":
			return "eleventy"
		case "Svelte":
			return "svelte"
		case "Gatsby":
			return "gatsby"
		case "Sanity":
			return "sanity"
		case "Ionic-React":
			return "ionicreact"
		case "Vite":
			return "vite"
		case "Scully":
			return "scully"
		case "Stencil":
			return "stencil"
		case "Brunch":
			return "brunch"
		case "Ionic-Angular":
			return "ionicangular"
		default:
			return "static"
	}
}

var nodeList = []string{"Node V12 LTS", "Node V14 LTS", "Node V16 LTS"}

// a function that switches case for a selected node value on the CLI
func GetNodeEnum(node string) string {
	switch node {
		case "Node V12 LTS":
			return "V_12"
		case "Node V14 LTS":
			return "V_14"
		case "Node V16 LTS":
			return "V_16"
		default:
			return "V_14"
	}
}

// TODO: Populate this list extensively with all the possible values
var providerList = []string{"Github"}

// a function that switches case for a selected provider value on the CLI
func GetProviderEnum(provider string) string {
	switch provider {
		case "Github":
			return "GITHUB"
		default:
			return "GITHUB"
	}
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

// function to detect type of file being read

func GetFileContentType(ouput *os.File) (string, error) {

   // to sniff the content type only the first
   // 512 bytes are used.

   buf := make([]byte, 512)

   _, err := ouput.Read(buf)

   if err != nil {
      return "", err
   }

   // the function that actually does the trick
   contentType := http.DetectContentType(buf)

      return contentType, nil
}