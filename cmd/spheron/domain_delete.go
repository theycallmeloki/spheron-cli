package spheron

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theycallmeloki/spheron-cli/pkg/spheron"
)

var deleteDomainId string

var deleteDomainCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a domain",
	Long:  `Delete a domain`,
	Run: func(cmd *cobra.Command, args []string) {
		projectId := viper.GetString("project")
		
		if(projectId == "") {
			panic("No project selected. Please select a project with `spheronctl set project`")
		}

		if(deleteDomainId == "") {
			deleteDomainId = SanitizeInput("Enter domain ID to delete: ")
		}

		deletedDomain, err := spheron.DeleteDomain(projectId, deleteDomainId)

		if err != nil {
			panic(err)
		}

		if(deletedDomain) {
			println("Domain deleted successfully")
		} else {
			println("Domain not deleted")
		}
	},
}

func init() {
	domainCmd.AddCommand(deleteDomainCmd)
	deleteDomainCmd.Flags().StringVarP(&deleteDomainId, "domain", "d", "", "Domain ID to delete")
}

