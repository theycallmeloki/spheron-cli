package spheron

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
	"os"

	"github.com/spf13/viper"
)

var SPHERON_BASE_URL string = "https://api-v2.spheron.network"

// get scope endpoint
func GetScope() (ScopeResponse, error) {
	url := SPHERON_BASE_URL + "/v1/api-keys/scope"
 	method := "GET"

	client := &http.Client {Timeout: 10 * time.Second}
	req, err := http.NewRequest(method, url, nil)

	var scope ScopeResponse

	if err != nil {
		fmt.Println(err)
		return scope, err
	}
	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return scope, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&scope)

	if(scope.Error){
		return scope, errors.New(scope.Message)
	}

	return scope, nil
}

// get organization endpoint 
func GetOrganization(organizationId string) (Organization, error) {
	url := SPHERON_BASE_URL + "/v1/organization/" + organizationId
	method := "GET"

	client := &http.Client {Timeout: 10 * time.Second}
	req, err := http.NewRequest(method, url, nil)

	var organization Organization

	if err != nil {
		fmt.Println(err)
		return organization, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return organization, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&organization)


	return organization, nil
}

// put organization endpoint
func PutOrganization(organizationId string, organizationName string, organizationUsername string, organizationImage string) (bool, error) {
	url := SPHERON_BASE_URL + "/v1/organization/" + organizationId
	method := "PUT"

	payload := strings.NewReader(`{
	"name": "` + organizationName + `",
	"username": "` + organizationUsername + `",
	"image": "` + organizationImage + `"
	}`)

	client := &http.Client {Timeout: 10 * time.Second}
	req, err := http.NewRequest(method, url, payload)
	req.Header.Set("Content-Type", "application/json")

	var confirmedChange bool 

	if err != nil {
		fmt.Println(err)
		return confirmedChange, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return confirmedChange, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&confirmedChange)
	
	return confirmedChange, nil
}

// get organization projects endpoint
func GetOrganizationProjects(organizationId string) ([] Projects, error) {
	url := SPHERON_BASE_URL + "/v1/organization/" + organizationId + "/projects"
	method := "GET"

	client := &http.Client {Timeout: 10 * time.Second}
	req, err := http.NewRequest(method, url, nil)

	var projects OrganizationProjectsResponse

	if err != nil {
		fmt.Println(err)
		// return projects key in response
		return projects.Projects, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return projects.Projects, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&projects)

	return projects.Projects, nil
}

// get organization overdue status endpoint 
func GetOrganizationOverdue(organizationId string) (OrganizationOverdueResponse, error) {
	url := SPHERON_BASE_URL + "/v1/organization/" + organizationId + "/overdue"
	method := "GET"

	client := &http.Client {Timeout: 10 * time.Second}
	req, err := http.NewRequest(method, url, nil)

	var overdue OrganizationOverdueResponse

	if err != nil {
		fmt.Println(err)
		return overdue, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return overdue, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&overdue)

	if(overdue.Error){
		return overdue, errors.New(overdue.Message)
	}

	return overdue, nil
}

// get organization projects count endpoint (possibly unused by the cli)
func GetOrganizationProjectsCount(organizationId string) (ProjectCountResponse, error) {
	url := SPHERON_BASE_URL + "/v1/organization/" + organizationId + "/projects/count"
	method := "GET"

	client := &http.Client {Timeout: 10 * time.Second}
	req, err := http.NewRequest(method, url, nil)

	var projectsCount ProjectCountResponse

	if err != nil {
		fmt.Println(err)
		return projectsCount, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return projectsCount, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&projectsCount)

	if(projectsCount.Error){
		return projectsCount, errors.New(projectsCount.Message)
	}

	return projectsCount, nil
}

// delete a member of the organization endpoint (possibly unused by the cli)
func DeleteOrganizationMember(organizationId string, memberId string) (DeleteOrganizationMemberResponse, error) {
	url := SPHERON_BASE_URL + "/v1/organization/" + organizationId + "/member"
	method := "DELETE"

	payload := strings.NewReader(`{
		userId: "` + memberId + `"
	}`)

	client := &http.Client {Timeout: 10 * time.Second}
	req, err := http.NewRequest(method, url, payload)

	var deleteOrganizationMember DeleteOrganizationMemberResponse

	if err != nil {
		fmt.Println(err)
		return deleteOrganizationMember, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return deleteOrganizationMember, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&deleteOrganizationMember)

	if(deleteOrganizationMember.Error) {
		return deleteOrganizationMember, errors.New(deleteOrganizationMember.Message)
	}

	return deleteOrganizationMember, nil
}

// get coupons associated with organization endpoint
func GetOrganizationCoupons(organizationId string) ([] Coupons, error) {
	url := SPHERON_BASE_URL + "/v1/organization/" + organizationId + "/coupons"
	method := "GET"

	client := &http.Client {Timeout: 10 * time.Second}
	req, err := http.NewRequest(method, url, nil)

	var coupons CouponsResponse

	if err != nil {
		fmt.Println(err)
		return coupons.Coupons, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return coupons.Coupons, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&coupons)

	if(coupons.Error){
		return coupons.Coupons, errors.New(coupons.Message)
	}

	return coupons.Coupons, nil
}

// get organization invites endpoint 
func GetOrganizationInvites(organizationId string) ([] Invites, error) {
	url := SPHERON_BASE_URL + "/v1/organization/" + organizationId + "/invites"
	method := "GET"

	client := &http.Client {Timeout: 10 * time.Second}
	req, err := http.NewRequest(method, url, nil)

	var invites InvitesResponse

	if err != nil {
		fmt.Println(err)
		return invites.Invites, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return invites.Invites, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&invites)

	if(invites.Error){
		return invites.Invites, errors.New(invites.Message)
	}

	return invites.Invites, nil
}

// invite a member to the organization endpoint
func InviteOrganizationMember(organizationId string, email string) (Invites, error) {
	url := SPHERON_BASE_URL + "/v1/organization/" + organizationId + "/invites"
	method := "POST"

	
    payload := strings.NewReader(`{
    "userEmail": "` + email + `"
	}`)

	client := &http.Client {Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, url, payload)

	var inviteOrganizationMember CreateInviteResponse

	if err != nil {
		fmt.Println(err)
		return inviteOrganizationMember.Invites, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return inviteOrganizationMember.Invites, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&inviteOrganizationMember)

	if(inviteOrganizationMember.Error) {
		return inviteOrganizationMember.Invites, errors.New(inviteOrganizationMember.Message)
	}

	return inviteOrganizationMember.Invites, nil
}

// get project endpoint
func GetProject(projectId string) (Projects, error) {
	url := SPHERON_BASE_URL + "/v1/project/" + projectId
	method := "GET"

	client := &http.Client {Timeout: 10 * time.Second}
	req, err := http.NewRequest(method, url, nil)

	var project Projects

	if err != nil {
		fmt.Println(err)
		return project, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return project, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&project)

	return project, nil
}

// get project deployments endpoint
func GetProjectDeployments(projectId string) ([] Deployment, error) {
	url := SPHERON_BASE_URL + "/v1/project/" + projectId + "/deployments"
	method := "GET"

	client := &http.Client {Timeout: 10 * time.Second}
	req, err := http.NewRequest(method, url, nil)

	var deployments [] Deployment

	if err != nil {
		fmt.Println(err)
		return deployments, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return deployments, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&deployments)

	return deployments, nil
}

// get project deployments count endpoint
func GetProjectDeploymentsCount(projectId string) (ProjectDeploymentCountResponse, error) {
	url := SPHERON_BASE_URL + "/v1/project/" + projectId + "/deployments/count"
	method := "GET"

	client := &http.Client {Timeout: 10 * time.Second}
	req, err := http.NewRequest(method, url, nil)

	var deploymentsCount ProjectDeploymentCountResponse

	if err != nil {
		fmt.Println(err)
		return deploymentsCount, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return deploymentsCount, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&deploymentsCount)

	if(deploymentsCount.Error) {
		return deploymentsCount, errors.New(deploymentsCount.Message)
	}

	return deploymentsCount, nil
}

// post environment variables endpoint
func PostEnvironmentVariables(projectId string, environmentVariables [] EnvironmentVariables) ([] EnvironmentVariables, error) {
	url := SPHERON_BASE_URL + "/v1/project/" + projectId + "/environment-variables"
	method := "POST"

	environmentVariablesPayload := CreateEnvironmentVariablesPayload {
		EnvironmentVariables: environmentVariables,
	}
	payload, err := json.Marshal(environmentVariablesPayload)

	client := &http.Client {Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	var environmentVariablesResponse CreateEnvironmentVariablesResponse

	if err != nil {
		fmt.Println(err)
		return environmentVariablesResponse.EnvironmentVariables, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return environmentVariablesResponse.EnvironmentVariables, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&environmentVariablesResponse)

	if(environmentVariablesResponse.Error) {
		return environmentVariablesResponse.EnvironmentVariables, errors.New(environmentVariablesResponse.Message)
	}

	return environmentVariablesResponse.EnvironmentVariables, nil
}

// put environment variable endpoint
func PutEnvironmentVariable(projectId string, environmentVariableId string, environmentVariable EnvironmentVariables) (UpdatedEnvironmentVariable, error) {
	url := SPHERON_BASE_URL + "/v1/project/" + projectId + "/environment-variables/" + environmentVariableId
	method := "PUT"

	payload, err := json.Marshal(environmentVariable)

	client := &http.Client {Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	
	var environmentVariableResponse UpdateEnvironmentVariablesResponse

	if err != nil {
		fmt.Println(err)
		return environmentVariableResponse.Updated, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return environmentVariableResponse.Updated, err
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&environmentVariableResponse)

	if(environmentVariableResponse.Error) {
		return environmentVariableResponse.Updated, errors.New(environmentVariableResponse.Message)
	}

	return environmentVariableResponse.Updated, nil
}

// delete environment variable endpoint
func DeleteEnvironmentVariable(projectId string, environmentVariableId string) (bool, error) {
	url := SPHERON_BASE_URL + "/v1/project/" + projectId + "/environment-variables/" + environmentVariableId
	method := "DELETE"

	client := &http.Client {Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, url, nil)

	var environmentVariableResponse DeleteEnvironmentVariablesResponse

	if err != nil {
		fmt.Println(err)
		return environmentVariableResponse.Success, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return environmentVariableResponse.Success, err
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&environmentVariableResponse)

	if(environmentVariableResponse.Success) {
		return environmentVariableResponse.Success, nil
	} else {
		return environmentVariableResponse.Success, errors.New(environmentVariableResponse.Message)
	}

	return environmentVariableResponse.Success, nil
}

// get deployment environment variables endpoint
func GetDeploymentEnvironmentVariables(projectId string) ([] DeploymentEnvironment, error) {
	url := SPHERON_BASE_URL + "/v1/project/" + projectId + "/deployment-environments"
	method := "GET"

	client := &http.Client {Timeout: 10 * time.Second}
	req, err := http.NewRequest(method, url, nil)

	var deploymentEnvironments [] DeploymentEnvironment

	var deploymentEnvironmentResponse DeploymentEnvironmentsResponse
	
	if err != nil {
		fmt.Println(err)
		return deploymentEnvironments, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return deploymentEnvironments, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&deploymentEnvironmentResponse)

	if(deploymentEnvironmentResponse.Error) {
		return deploymentEnvironments, errors.New(deploymentEnvironmentResponse.Message)
	}

	for _, deploymentEnvironment := range deploymentEnvironmentResponse.Result {
		createdDeploymentEnvironment := DeploymentEnvironment {
			ID: deploymentEnvironment.ID,
			Name: deploymentEnvironment.Name,
			Branches: deploymentEnvironment.Branches,
			Status: deploymentEnvironment.Status,
			Protocol: deploymentEnvironment.Protocol,
			CreatedAt: deploymentEnvironment.CreatedAt,
			UpdatedAt: deploymentEnvironment.UpdatedAt,
		}
		deploymentEnvironments = append(deploymentEnvironments, createdDeploymentEnvironment)
	}

	return deploymentEnvironments, nil
}

// post deployment environment variable endpoint
func PostDeploymentEnvironmentVariable(projectId string, deploymentEnvironment DeploymentEnvironment) (DeploymentEnvironment, error) {
	url := SPHERON_BASE_URL + "/v1/project/" + projectId + "/deployment-environments"
	method := "POST"

	payload, err := json.Marshal(deploymentEnvironment)

	client := &http.Client {Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	var createdDeploymentEnvironment DeploymentEnvironment

	var deploymentEnvironmentVariablesResponse CreateDeploymentEnvironmentResponse

	if err != nil {
		fmt.Println(err)
		return createdDeploymentEnvironment, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return createdDeploymentEnvironment, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&deploymentEnvironmentVariablesResponse)

	if(deploymentEnvironmentVariablesResponse.Error) {
		return createdDeploymentEnvironment, errors.New(deploymentEnvironmentVariablesResponse.Message)
	}

	createdDeploymentEnvironment = DeploymentEnvironment {
		ID: deploymentEnvironmentVariablesResponse.NewEnvironment.ID,
		Name: deploymentEnvironmentVariablesResponse.NewEnvironment.Name,
		Branches: deploymentEnvironmentVariablesResponse.NewEnvironment.Branches,
		Status: deploymentEnvironmentVariablesResponse.NewEnvironment.Status,
		Protocol: deploymentEnvironmentVariablesResponse.NewEnvironment.Protocol,
		CreatedAt: deploymentEnvironmentVariablesResponse.NewEnvironment.CreatedAt,
		UpdatedAt: deploymentEnvironmentVariablesResponse.NewEnvironment.UpdatedAt,
	}

	return createdDeploymentEnvironment, nil
}

// put deployment environment variable endpoint
func PutDeploymentEnvironmentVariable(projectId string, deploymentEnvironmentId string, deploymentEnvironment DeploymentEnvironment) (DeploymentEnvironment, error) {
	// TODO: evaluate merging deployment environment with deployment environments (same struct, plural vs singular)
	url := SPHERON_BASE_URL + "/v1/project/" + projectId + "/deployment-environments/" + deploymentEnvironmentId
	method := "PUT"

	payload, err := json.Marshal(deploymentEnvironment)

	client := &http.Client {Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	var deploymentEnvironmentVariablesResponse UpdateDeploymentEnvironmentResponse

	if err != nil {
		fmt.Println(err)
		return deploymentEnvironmentVariablesResponse.DeploymentEnvironment, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return deploymentEnvironmentVariablesResponse.DeploymentEnvironment, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&deploymentEnvironmentVariablesResponse)	

	if(deploymentEnvironmentVariablesResponse.Error) {
		return deploymentEnvironmentVariablesResponse.DeploymentEnvironment, errors.New(deploymentEnvironmentVariablesResponse.Message)
	}

	return deploymentEnvironmentVariablesResponse.DeploymentEnvironment, nil
}

// delete deployment environment variable endpoint
func DeleteDeploymentEnvironmentVariable(projectId string, deploymentEnvironmentId string) (bool, error) {
	url := SPHERON_BASE_URL + "/v1/project/" + projectId + "/deployment-environments/" + deploymentEnvironmentId
	method := "DELETE"

	client := &http.Client {Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, url, nil)

	// assume we aren't able to delete it by default, if we dont get an error, we can assume it was deleted
	successfulDelete := false

	var deploymentEnvironmentResponse DeleteDeploymentEnvironmentResponse

	if err != nil {
		fmt.Println(err)
		return successfulDelete, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return successfulDelete, err
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&deploymentEnvironmentResponse)

	if(deploymentEnvironmentResponse.Error) {
		return successfulDelete, errors.New(deploymentEnvironmentResponse.Message)
	} else {
		successfulDelete = true
	}

	return successfulDelete, nil
}

// patch deactivate deployment environment variable endpoint
func DeactivateDeploymentEnvironmentVariable(projectId string, deploymentEnvironmentId string) (DeploymentEnvironment, error) {
	// TODO: evaluate merging deployment environment with deployment environments (same struct, plural vs singular)
	url := SPHERON_BASE_URL + "/v1/project/" + projectId + "/deployment-environments/" + deploymentEnvironmentId + "/deactivate"
	method := "PATCH"

	client := &http.Client {Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, url, nil)

	var deploymentEnvironmentResponse PatchDeploymentEnvironmentResponse

	if err != nil {
		fmt.Println(err)
		return deploymentEnvironmentResponse.DeploymentEnvironment, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return deploymentEnvironmentResponse.DeploymentEnvironment, err
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&deploymentEnvironmentResponse)

	if(deploymentEnvironmentResponse.Error) {
		return deploymentEnvironmentResponse.DeploymentEnvironment, errors.New(deploymentEnvironmentResponse.Message)
	}

	return deploymentEnvironmentResponse.DeploymentEnvironment, nil
}

// TODO: evaluate merging the two patch endpoints into one, they are the same except for the url
// though this removes the "thinness" of the sdk, so maybe not

// patch activate deployment environment variable endpoint
func ActivateDeploymentEnvironmentVariable(projectId string, deploymentEnvironmentId string) (DeploymentEnvironment, error) {
	url := SPHERON_BASE_URL + "/v1/project/" + projectId + "/deployment-environments/" + deploymentEnvironmentId + "/activate"
	method := "PATCH"

	client := &http.Client {Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, url, nil)

	var deploymentEnvironmentResponse PatchDeploymentEnvironmentResponse

	if err != nil {
		fmt.Println(err)
		return deploymentEnvironmentResponse.DeploymentEnvironment, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return deploymentEnvironmentResponse.DeploymentEnvironment, err
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&deploymentEnvironmentResponse)

	if(deploymentEnvironmentResponse.Error) {
		return deploymentEnvironmentResponse.DeploymentEnvironment, errors.New(deploymentEnvironmentResponse.Message)
	}

	return deploymentEnvironmentResponse.DeploymentEnvironment, nil
}

// get domains endpoint
func GetDomains(projectId string) ([]Domains, error) {
	url := SPHERON_BASE_URL + "/v1/project/" + projectId + "/domains"
	method := "GET"

	client := &http.Client {Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, url, nil)

	var domainsResponse DomainsResponse

	if err != nil {
		fmt.Println(err)
		return domainsResponse.Domains, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return domainsResponse.Domains, err
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&domainsResponse)

	if(domainsResponse.Error) {
		return domainsResponse.Domains, errors.New(domainsResponse.Message)
	}

	return domainsResponse.Domains, nil
}

// post domain endpoint
func PostDomain(projectId string, domain CreateDomainPayload) (Domains, error) {
	url := SPHERON_BASE_URL + "/v1/project/" + projectId + "/domains"
	method := "POST"

	client := &http.Client {Timeout: 10 * time.Second}

	jsonValue, _ := json.Marshal(domain)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonValue))

	var domainResponse CreateDomainResponse

	if err != nil {
		fmt.Println(err)
		return domainResponse.Domains, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return domainResponse.Domains, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&domainResponse)	

	if(domainResponse.Error) {
		return domainResponse.Domains, errors.New(domainResponse.Message)
	}

	return domainResponse.Domains, nil
}

// get one domain endpoint 
func GetDomain(projectId string, domainId string) (Domains, error) {
	url := SPHERON_BASE_URL + "/v1/project/" + projectId + "/domains/" + domainId
	method := "GET"

	client := &http.Client {Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, url, nil)

	var domainResponse DomainResponse

	if err != nil {
		fmt.Println(err)
		return domainResponse.Domains, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return domainResponse.Domains, err
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&domainResponse)

	if(domainResponse.Error) {
		return domainResponse.Domains, errors.New(domainResponse.Message)
	}

	return domainResponse.Domains, nil
}

// patch domain endpoint
func PatchDomain(projectId string, domainId string, domain UpdateDomainPayload) (Domains, error) {
	url := SPHERON_BASE_URL + "/v1/project/" + projectId + "/domains/" + domainId
	method := "PATCH"

	client := &http.Client {Timeout: 10 * time.Second}

	jsonValue, _ := json.Marshal(domain)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonValue))

	var domainResponse UpdateDomainResponse

	if err != nil {
		fmt.Println(err)
		return domainResponse.Domains, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return domainResponse.Domains, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&domainResponse)	

	if(domainResponse.Error) {
		return domainResponse.Domains, errors.New(domainResponse.Message)
	}

	return domainResponse.Domains, nil
}

// delete domain endpoint
func DeleteDomain(projectId string, domainId string) (bool, error) {
	url := SPHERON_BASE_URL + "/v1/project/" + projectId + "/domains/" + domainId
	method := "DELETE"

	client := &http.Client {Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, url, nil)

	var domainResponse DeleteDomainResponse

	if err != nil {
		fmt.Println(err)
		return domainResponse.DeleteDomainResult.Success, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return domainResponse.DeleteDomainResult.Success, err
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&domainResponse)

	if(domainResponse.Error) {
		return domainResponse.DeleteDomainResult.Success, errors.New(domainResponse.Message)
	}

	return domainResponse.DeleteDomainResult.Success, nil
}

// verify domain endpoint
func VerifyDomain(projectId string, domainId string) (bool, error) {
	url := SPHERON_BASE_URL + "/v1/project/" + projectId + "/domains/" + domainId + "/verify"
	method := "PATCH"

	client := &http.Client {Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, url, nil)

	var domainResponse VerifyDomainResponse

	if err != nil {
		fmt.Println(err)
		return domainResponse.Success, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return domainResponse.Success, err
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&domainResponse)

	if(domainResponse.Error) {
		return domainResponse.Success, errors.New(domainResponse.Message)
	}

	return domainResponse.Success, nil
}

// post deployment endpoint 
func PostDeployment(deployment CreateDeploymentPayload) (DeploymentDomain, error) {
	url := SPHERON_BASE_URL + "/v1/deployment"
	method := "POST"

	client := &http.Client {Timeout: 10 * time.Second}

	jsonValue, _ := json.Marshal(deployment)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonValue))

	var deploymentResponse CreateDeploymentResponse

	if err != nil {
		fmt.Println(err)
		return deploymentResponse.DeploymentDomain, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return deploymentResponse.DeploymentDomain, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&deploymentResponse)	

	if(deploymentResponse.Error) {
		return deploymentResponse.DeploymentDomain, errors.New(deploymentResponse.Message)
	}

	return deploymentResponse.DeploymentDomain, nil
}

// get deployment endpoint
func GetDeployment(deploymentId string) (Deployment, error) {
	url := SPHERON_BASE_URL + "/v1/deployment/" + deploymentId
	method := "GET"

	client := &http.Client {Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, url, nil)

	var deploymentResponse DeploymentResponse

	if err != nil {
		fmt.Println(err)
		return deploymentResponse.Deployment, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return deploymentResponse.Deployment, err
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&deploymentResponse)

	if(deploymentResponse.Error) {
		return deploymentResponse.Deployment, errors.New(deploymentResponse.Message)
	}

	return deploymentResponse.Deployment, nil
}

// post authorize deployment endpoint
func PostAuthorizeDeployment(deploymentId string) (bool, error) {
	url := SPHERON_BASE_URL + "/v1/deployment/" + deploymentId + "/authorize"
	method := "POST"

	client := &http.Client {Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, url, nil)

	var deploymentResponse AuthorizeDeploymentResponse

	if err != nil {
		fmt.Println(err)
		return deploymentResponse.Success, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return deploymentResponse.Success, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&deploymentResponse)	

	if(deploymentResponse.Error) {
		return deploymentResponse.Success, errors.New(deploymentResponse.Message)
	}

	return deploymentResponse.Success, nil
}

// post cancel deployment endpoint
func PostCancelDeployment(deploymentId string) (bool, error) {
	url := SPHERON_BASE_URL + "/v1/deployment/" + deploymentId + "/cancel"
	method := "POST"

	client := &http.Client {Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, url, nil)

	var deploymentResponse CancelDeploymentResponse

	if err != nil {
		fmt.Println(err)
		return deploymentResponse.Canceled, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return deploymentResponse.Canceled, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&deploymentResponse)	

	if(deploymentResponse.Error) {
		return deploymentResponse.Canceled, errors.New(deploymentResponse.Message)
	}

	return deploymentResponse.Canceled, nil
}

// post redeploy deployment endpoint
func PostRedeployDeployment(deploymentId string) (bool, error) {
	url := SPHERON_BASE_URL + "/v1/deployment/" + deploymentId + "/redeploy"
	method := "POST"

	client := &http.Client {Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, url, nil)

	var deploymentResponse RedeployDeploymentResponse

	if err != nil {
		fmt.Println(err)
		return deploymentResponse.Success, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return deploymentResponse.Success, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&deploymentResponse)	

	if(deploymentResponse.Error) {
		return deploymentResponse.Success, errors.New(deploymentResponse.Message)
	}

	return deploymentResponse.Success, nil
}

// get framework suggestions endpoint (Possibly unused in the CLI)
func GetFrameworkSuggestions(owner string, branch string, repo string, providerName string, root string) (string, error) {
	url := SPHERON_BASE_URL + "/v1/framework/suggestions?owner=" + owner + "&branch=" + branch + "&repo=" + repo + "&providerName=" + providerName + "&root=" + root
	method := "GET"

	client := &http.Client {Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, url, nil)

	var frameworkResponse SuggestedFrameworkResponse

	if err != nil {
		fmt.Println(err)
		return frameworkResponse.SuggestedFramework, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return frameworkResponse.SuggestedFramework, err
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&frameworkResponse)

	if(frameworkResponse.Error) {
		return frameworkResponse.SuggestedFramework, errors.New(frameworkResponse.Message)
	}

	return frameworkResponse.SuggestedFramework, nil
}

// upload built files endpoint
func UploadFiles(organizationId string, projectName string, protocol string, files []FileContent) (UploadFilesDeploymentResponse, error) {
	url := SPHERON_BASE_URL + "/v1/deployment/upload?organization=" + organizationId + "&project=" + projectName + "&protocol=" + protocol
	method := "POST"

	var uploadResponse UploadFilesDeploymentResponse

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	for _, file := range files {
		folderParts := strings.Split(file.Fname, string(os.PathSeparator))
		_, reducedFolderParts := folderParts[0], folderParts[1:]
		reducedFileName := strings.Join(reducedFolderParts, string(os.PathSeparator))
		part, err := writer.CreateFormFile(file.Ftype, reducedFileName)
		if err != nil {
			return uploadResponse, err
		}
		part.Write(file.Fcontent)
	}
	err := writer.Close()
	if err != nil {
		return uploadResponse, err
	}

	client := &http.Client {Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return uploadResponse, err
	}

	req.Header.Add("Authorization", "Bearer " + viper.GetString("secret"))
	req.Header.Add("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return uploadResponse, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&uploadResponse)	

	if(uploadResponse.Error) {
		return uploadResponse, errors.New(uploadResponse.Message)
	}

	return uploadResponse, nil

}