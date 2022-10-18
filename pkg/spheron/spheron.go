package spheron

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var ConfigName string = "spheron"
var ConfigType string = "json"
var ConfigDir string
var ConfigPath string

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