package spheron

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var folderName string

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a folder to Spheron",
	Long:  `Upload a folder to Spheron`,
	Run: func(cmd *cobra.Command, args []string) {

		organizationId := viper.GetString("organization")

		if(folderName == "") {
			folderName = SanitizeInput("Enter Folder Name: ")
		}

		projectName := SanitizeInput("Enter Project Name: ")

		protocolSelection, err := SanitizeFixedSelect(protocolList, "Select Protocol: ")
		if err != nil {
			panic(err)
		}

		protocol := GetProtocolEnum(protocolSelection)

		// create a variable to store FileContent
		var files []spheron.FileContent

		// loop through all the files in the folder

		err = filepath.Walk(folderName, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				fmt.Println(path)

				// read in the file 

				fileData, err := os.Open(path)

				if err != nil {
					panic(err)
				}

				// determine file type 

				fileType, err := GetFileContentType(fileData)

				if err != nil {
					panic(err)
				}

				// read file content

				fileContent, err := ioutil.ReadFile(path)

				// add the file to FileContent
				files = append(files, spheron.FileContent{
					Fname: path,
					Ftype: fileType,
					Fcontent: fileContent,
				})

				if err != nil {
					panic(err)
				}
			}
			return nil
			})

		if err != nil {
			panic(err)
		}

		fmt.Println("Uploading folder ", folderName, " to project ", projectName, " in organization ", organizationId, " using protocol ", protocol)
		

		// upload files to Spheron

		uploaded, err := spheron.UploadFiles(organizationId, projectName, protocol, files)

		if err != nil {
			panic(err)
		}

		fmt.Println("Uploaded ", len(files), " files successfully!")
		fmt.Println("Site Preview: ", uploaded)

	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.Flags().StringVarP(&folderName, "folder", "f", "", "Folder to upload")
}
