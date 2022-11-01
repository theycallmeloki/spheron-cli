package spheron

import "time"

// Config struct (stored locally)
type Config struct {
	// Secret Access Token that can be generated from Aqua Spheron Console
	Secret  			string
	// Currently active organization that the User is working on
	Organization	 	string
	// Project
	Project 			string
	// Deployment
	Deployment 			string
}

// Scope struct (possibly deprecated)
type Scope struct {
	Scope []string `json:"scope"`
}

// Scope response 
type ScopeResponse struct {
	User          User            `json:"user"`
	Organizations []Organizations `json:"organizations"`
}

// user struct 
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

// organizations struct
type Organizations struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

// organization struct (possibly deprecated)
type Organization struct {
	ID             string        `json:"_id"`
	Profile        Profile       `json:"profile"`
	Users          []Users       `json:"users"`
	Overdue        bool          `json:"overdue"`
	CreatedAt      time.Time     `json:"createdAt"`
	UpdatedAt      time.Time     `json:"updatedAt"`
	InvitedMembers []interface{} `json:"invitedMembers"`
}

// profile struct
type Profile struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

// provider struct
type Provider struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

// platform profile struct
type PlatformProfile struct {
	IsActive bool   `json:"is_active"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

// platform provider profiles struct
type ProviderProfiles struct {
	ID                string    `json:"id"`
	AvatarURL         string    `json:"avatar_url"`
	URL               string    `json:"url"`
	HTMLURL           string    `json:"html_url"`
	FollowersURL      string    `json:"followers_url"`
	FollowingURL      string    `json:"following_url"`
	GistsURL          string    `json:"gists_url"`
	StarredURL        string    `json:"starred_url"`
	SubscriptionsURL  string    `json:"subscriptions_url"`
	OrganizationsURL  string    `json:"organizations_url"`
	ReposURL          string    `json:"repos_url"`
	EventsURL         string    `json:"events_url"`
	ReceivedEventsURL string    `json:"received_events_url"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	PublicRepos       int       `json:"public_repos"`
	PublicGists       int       `json:"public_gists"`
	Followers         int       `json:"followers"`
	Following         int       `json:"following"`
	Username          string    `json:"username"`
	ProviderName      string    `json:"providerName"`
	CreatedAt         time.Time `json:"createdAt"`
	LastLogin         time.Time `json:"lastLogin"`
}

// users struct (possibly deprecated)
type Users struct {
	ID               string             `json:"_id"`
	Organizations    []string           `json:"organizations"`
	Provider         Provider           `json:"provider"`
	CreatedAt        time.Time          `json:"createdAt"`
	UpdatedAt        time.Time          `json:"updatedAt"`
	PlatformProfile  PlatformProfile    `json:"platformProfile"`
	ProviderProfiles []ProviderProfiles `json:"providerProfiles"`
	NftAccess        string             `json:"nftAccess"`
}

// project count response struct
type ProjectCountResponse struct {
	Count string `json:"count"`
}

// organization projects response struct
type OrganizationProjectsResponse struct {
	Projects []Projects `json:"projects"`
}

// delete member from organization response struct
type DeleteOrganizationMemberResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// environment variables struct 
type EnvironmentVariables struct {
	ID		  			   string   `json:"_id,omitempty"`
	Name                   string   `json:"name"`
	Value                  string   `json:"value"`
	DeploymentEnvironments []string `json:"deploymentEnvironments"`
}

// deployment environments struct
type DeploymentEnvironments struct {
	ID        string   `json:"_id,omitempty"`
	Name      string   `json:"name"`
	Branches  []string `json:"branches"`
	Status    string   `json:"status"`
	Protocol  string   `json:"protocol"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
}

// PR Comment IDs struct
type PrCommentIds struct {
	PrID      string `json:"prId"`
	CommentID string `json:"commentId"`
}

// configuration struct
type Configuration struct {
	BuildCommand   string `json:"buildCommand"`
	InstallCommand string `json:"installCommand"`
	Workspace      string `json:"workspace"`
	PublishDir     string `json:"publishDir"`
	Framework      string `json:"framework"`
	NodeVersion    string `json:"nodeVersion"`
}

// logs struct
type Logs struct {
	Time string `json:"time"`
	Log  string `json:"log"`
}

// env struct - (possibly used ?)
type Env struct {
	Key  string `json:"key"`
	Value string `json:"value"`
}

// screenshot struct
type Screenshot struct {
	ID  string `json:"id"`
	Fee string `json:"fee"`
	URL string `json:"url"`
}

// latest deployment struct
type LatestDeployment struct {
	ID                        string        `json:"_id"`
	SitePreview               string        `json:"sitePreview"`
	CommitID                  string        `json:"commitId"`
	CommitMessage             string        `json:"commitMessage"`
	Logs                      []Logs        `json:"logs"`
	BuildDirectory            []string      `json:"buildDirectory"`
	ContentHash               string        `json:"contentHash"`
	Topic                     string        `json:"topic"`
	Status                    string        `json:"status"`
	PaymentID                 string        `json:"paymentId"`
	BuildTime                 int           `json:"buildTime"`
	MemoryUsed                int           `json:"memoryUsed"`
	Env                       Env           `json:"env"`
	Project                   string        `json:"project"`
	Screenshot                Screenshot    `json:"screenshot"`
	DeploymentInitiator       string        `json:"deploymentInitiator"`
	Branch                    string        `json:"branch"`
	ExternalRepositoryName    string        `json:"externalRepositoryName"`
	Protocol                  string        `json:"protocol"`
	DeploymentEnvironmentName string        `json:"deploymentEnvironmentName"`
	FailedMessage             string        `json:"failedMessage"`
	IsFromRequest             bool          `json:"isFromRequest"`
	Configuration             Configuration `json:"configuration"`
	CreatedAt                 string        `json:"createdAt"`
	UpdatedAt                 string        `json:"updatedAt"`
}

// domains struct
type Domains struct {
	ID                       string   `json:"_id,omitempty"`
	Name                     string   `json:"name"`
	Link                     string   `json:"link"`
	IsLatest                 bool     `json:"isLatest"`
	Type                     string   `json:"type"`
	Verified                 bool     `json:"verified"`
	ProjectID                string   `json:"projectId"`
	DeploymentEnvironmentIds []string `json:"deploymentEnvironmentIds"`
	Version                  string   `json:"version"`
	CreatedAt                string   `json:"createdAt"`
	UpdatedAt                string   `json:"updatedAt"`
}


// projects struct
type Projects struct {
	ID                     string                   `json:"_id"`
	Name                   string                   `json:"name"`
	URL                    string                   `json:"url"`
	HookID                 string                   `json:"hookId"`
	Provider               string                   `json:"provider"`
	CreatedAt              string                   `json:"createdAt"`
	UpdatedAt              string                   `json:"updatedAt"`
	CreatedBy              string                   `json:"createdBy"`
	State                  string                   `json:"state"`
	EnvironmentVariables   []EnvironmentVariables   `json:"environmentVariables"`
	Organization           string                   `json:"organization"`
	DeploymentEnvironments []DeploymentEnvironments `json:"deploymentEnvironments"`
	PrCommentIds           []PrCommentIds           `json:"prCommentIds"`
	Configuration          Configuration            `json:"configuration"`
	LatestDeployment       LatestDeployment         `json:"latestDeployment"`
	Domains                []Domains                `json:"domains"`
}

// organization overdue struct
type OrganizationOverdueResponse struct {
	Overdue         bool            `json:"overdue"`
	Message         string          `json:"message"`
	OverdueResponse OverdueResponse `json:"overdueResponse,omitempty"`
	OverdueReasons  []string        `json:"overdueReasons,omitempty"`
}

// exceeded environments for project struct
type ExceededEnvironmentsForProjects struct {
	Amount  int    `json:"amount"`
	Project string `json:"project"`
}

// overdue response struct
type OverdueResponse struct {
	UsedDomains                     int                             `json:"usedDomains"`
	UsedHnsDomains                  int                             `json:"usedHnsDomains"`
	UsedEnsDomains                  int                             `json:"usedEnsDomains"`
	AllowedDomains                  int                             `json:"allowedDomains"`
	AllowedHnsDomains               int                             `json:"allowedHnsDomains"`
	AllowedEnsDomains               int                             `json:"allowedEnsDomains"`
	AllowedMembers                  int                             `json:"allowedMembers"`
	ExceededDomains                 int                             `json:"exceededDomains"`
	ExceededHnsDomains              int                             `json:"exceededHnsDomains"`
	ExceededEnsDomains              int                             `json:"exceededEnsDomains"`
	ExceededMembers                 int                             `json:"exceededMembers"`
	ExceededEnvironmentsForProjects ExceededEnvironmentsForProjects `json:"exceededEnvironmentsForProjects"`
}

// remove member response struct
type RemoveMemberResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// coupons response struct
type CouponsResponse struct {
	Coupons []Coupons `json:"coupons"`
}

// coupon struct
type Coupons struct {
	ID                  string `json:"id"`
	CouponName          string `json:"couponName"`
	ActivationCode      string `json:"activationCode"`
	ActivatedAt         string `json:"activatedAt"`
	ExpiresAt           string `json:"expiresAt"`
	State               string `json:"state"`
	RegisteredAt        string `json:"registeredAt"`
	TotalDays           int    `json:"totalDays"`
	DaysRemaning        int    `json:"daysRemaning"`
	DaysUntilActivation int    `json:"daysUntilActivation"`
}

// organization member invite struct response
type InvitesResponse struct {
	Invites []Invites `json:"invites"`
}

// invites struct
type Invites struct {
	ID           string `json:"_id"`
	UserEmail    string `json:"userEmail"`
	Status       string `json:"status"`
	Link         string `json:"link"`
	Organization string `json:"organization"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

// create member invitation response struct
type CreateInviteResponse struct {
	Message string `json:"message"`
	Invites  Invites `json:"invite"`
	Error bool `json:"error,omitempty"`
}

// delete pending invitation response struct
type DeleteInviteResponse struct {
	ID           string `json:"_id"`
	UserEmail    string `json:"userEmail"`
	Status       string `json:"status"`
	Link         string `json:"link"`
	Organization string `json:"organization"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

// project response struct
type ProjectResponse struct {
	ID                     string                   `json:"_id"`
	Name                   string                   `json:"name"`
	URL                    string                   `json:"url"`
	HookID                 string                   `json:"hookId"`
	Provider               string                   `json:"provider"`
	CreatedAt              string                   `json:"createdAt"`
	UpdatedAt              string                   `json:"updatedAt"`
	CreatedBy              string                   `json:"createdBy"`
	State                  string                   `json:"state"`
	EnvironmentVariables   []EnvironmentVariables   `json:"environmentVariables"`
	Organization           string                   `json:"organization"`
	DeploymentEnvironments []DeploymentEnvironments `json:"deploymentEnvironments"`
	PrCommentIds           []PrCommentIds           `json:"prCommentIds"`
	Configuration          Configuration            `json:"configuration"`
	LatestDeployment       LatestDeployment         `json:"latestDeployment"`
	Domains                []Domains                `json:"domains"`
}

// project list response struct
type ProjectDeploymentsResponse []struct {
	ID                        string        `json:"_id"`
	SitePreview               string        `json:"sitePreview"`
	CommitID                  string        `json:"commitId"`
	CommitMessage             string        `json:"commitMessage"`
	Logs                      []Logs        `json:"logs"`
	BuildDirectory            []string      `json:"buildDirectory"`
	ContentHash               string        `json:"contentHash"`
	Topic                     string        `json:"topic"`
	Status                    string        `json:"status"`
	PaymentID                 string        `json:"paymentId"`
	BuildTime                 int           `json:"buildTime"`
	MemoryUsed                int           `json:"memoryUsed"`
	Env                       Env           `json:"env"`
	Project                   string        `json:"project"`
	Screenshot                Screenshot    `json:"screenshot"`
	DeploymentInitiator       string        `json:"deploymentInitiator"`
	Branch                    string        `json:"branch"`
	ExternalRepositoryName    string        `json:"externalRepositoryName"`
	Protocol                  string        `json:"protocol"`
	DeploymentEnvironmentName string        `json:"deploymentEnvironmentName"`
	FailedMessage             string        `json:"failedMessage"`
	IsFromRequest             bool          `json:"isFromRequest"`
	Configuration             Configuration `json:"configuration"`
	CreatedAt                 string        `json:"createdAt"`
	UpdatedAt                 string        `json:"updatedAt"`
}

// project count response struct
type ProjectDeploymentCountResponse struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Failed     int `json:"failed"`
	Pending    int `json:"pending"`
}

// project state response struct
type ProjectStateResponse struct {
	Message string `json:"message"`
}

// project configuration response struct
type ProjectConfigurationResponse struct {
	Configuration Configuration `json:"configuration"`
}

// create environment variables payload struct 
type CreateEnvironmentVariablesPayload struct {
	EnvironmentVariables []EnvironmentVariables `json:"environmentVariables"`
}

// create environment variables response struct
type CreateEnvironmentVariablesResponse struct {
	EnvironmentVariables []EnvironmentVariables `json:"environmentVariables"`
}

// update environment variables response struct
type UpdateEnvironmentVariablesResponse struct {
	Updated UpdatedEnvironmentVariable `json:"updated"`
}

// updated struct (part of update environment variables response struct)
type UpdatedEnvironmentVariable struct {
	ID                     string                   `json:"_id"`
	Name                   string                   `json:"name"`
	Value                  string                   `json:"value"`
	DeploymentEnvironments []DeploymentEnvironments `json:"deploymentEnvironments"`
}

// delete environment variables response struct
type DeleteEnvironmentVariablesResponse struct {
	Success bool `json:"success"`
}
	
// deployment environment response struct
type DeploymentEnvironmentsResponse struct {
	Result []Result `json:"result"`
}

// result struct (part of deployment environment response struct)
type Result struct {
	ID        string   `json:"_id"`
	Name      string   `json:"name"`
	Branches  []string `json:"branches"`
	Status    string   `json:"status"`
	Protocol  string   `json:"protocol"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
}

// create deployment environment response struct
type CreateDeploymentEnvironmentResponse struct {
	NewEnvironment NewEnvironment `json:"newEnvironment"`
	Message string `json:"message,omitempty"`
	Error bool `json:"error,omitempty"`
}

// new environment struct (part of create deployment environment response struct)
type NewEnvironment struct {
	ID        string   `json:"_id"`
	Name      string   `json:"name"`
	Branches  []string `json:"branches"`
	Status    string   `json:"status"`
	Protocol  string   `json:"protocol"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
}

// update deployment environment response struct
type UpdateDeploymentEnvironmentResponse struct {
	DeploymentEnvironment DeploymentEnvironment `json:"deploymentEnvironment"`
	Message string `json:"message,omitempty"`
	Error bool `json:"error,omitempty"`
}

// delete deployment environment response struct
type DeleteDeploymentEnvironmentResponse struct {
	Error bool `json:"error,omitempty"`
	Message string `json:"message"`
}

// patch deployment environment response struct
type PatchDeploymentEnvironmentResponse struct {
	DeploymentEnvironment DeploymentEnvironment `json:"deploymentEnvironment"`
	Message string `json:"message,omitempty"`
	Error bool `json:"error,omitempty"`
}

// deployment environment struct
type DeploymentEnvironment struct {
	ID        string   `json:"_id"`
	Name      string   `json:"name"`
	Branches  []string `json:"branches"`
	Status    string   `json:"status"`
	Protocol  string   `json:"protocol"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
}

// get domains response struct
type DomainsResponse struct {
	Domains []Domains `json:"domains"`
	Message string `json:"message,omitempty"`
	Error bool `json:"error,omitempty"`
}

// get one domain response struct (todo - check if this is correct, merge with above if feasible)
type DomainResponse struct {
	Domains Domains `json:"domain"`
	Message string `json:"message,omitempty"`
	Error bool `json:"error,omitempty"`
}

// create domain payload struct
type CreateDomainPayload struct {
	Link string `json:"link"`
	Type string `json:"type"`
	DeploymentEnvironments []string `json:"deploymentEnvironments"`
	IsLatest bool `json:"isLatest"`
	Name string `json:"name"`
}

// create domain response struct
type CreateDomainResponse struct {
	Domains Domains `json:"domain"`
	Message string `json:"message,omitempty"`
	Error bool `json:"error,omitempty"`
}

// update domain payload struct
type UpdateDomainPayload struct {
	Link string `json:"link"`
	DeploymentEnvironments []string `json:"deploymentEnvironments"`
	IsLatest bool `json:"isLatest"`
	Name string `json:"name"`
}

// update domain response struct
type UpdateDomainResponse struct {
	Domains Domains `json:"domain"`
	Message string `json:"message,omitempty"`
	Error bool `json:"error,omitempty"`
}

// delete domain response struct
type DeleteDomainResponse struct {
	DeleteDomainResult DeleteDomainResult `json:"domain"`
	Message string `json:"message,omitempty"`
	Error bool `json:"error,omitempty"`
}

// delete domain struct
type DeleteDomainResult struct {
	Success bool `json:"success"`
}

// verify domain response struct
type VerifyDomainResponse struct {
	Success bool   `json:"success"`
	Domains  Domains `json:"domain"`
	Message string `json:"message,omitempty"`
	Error bool `json:"error,omitempty"`
}

// git provider preferences struct (part of project response struct)
type GitProviderPreferences struct {
	PrComments 			bool `json:"prComments"`
	CommitComments 		bool `json:"commitComments"`
	BuildStatus 		bool `json:"buildStatus"`
	GithubDeployment 	bool `json:"githubDeployment"`
}

// create deployment payload struct 
type CreateDeploymentPayload struct {
	OrganizationID 			string `json:"organizationId"`
	GitURL 		   			string `json:"gitUrl"`
	RepoName	   			string `json:"repoName"`
	UniqueTopicID  			string `json:"uniqueTopicId"`
	Configuration Configuration `json:"configuration"`
	Env 		  			string `json:"env"`
	Protocol	  			string `json:"protocol"`
	CreateDefaultWebhook 	bool `json:"createDefaultWebhook"`
	Provider 				string `json:"provider"`
	Branch 					string `json:"branch"`
	GitProviderPreferences GitProviderPreferences `json:"gitProviderPreferences"`
}


// create deployment response struct
type CreateDeploymentResponse struct {
	DeploymentDomain DeploymentDomain `json:"domain"`
	Message string `json:"message,omitempty"`
	Error bool `json:"error,omitempty"`
}

// logs to capture struct
type LogsToCapture struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// body struct (part of create deployment response struct)
type Body struct {
	DeploymentID        string          `json:"deploymentId"`
	GithubURL           string          `json:"githubUrl"`
	FolderName          string          `json:"folderName"`
	Topic               string          `json:"topic"`
	Framework           string          `json:"framework"`
	Branch              string          `json:"branch"`
	BuildCommand        string          `json:"buildCommand"`
	InstallCommand      string          `json:"installCommand"`
	PublishDirectory    string          `json:"publishDirectory"`
	Protocol            string          `json:"protocol"`
	Workspace           string          `json:"workspace"`
	IsWorkspace         bool            `json:"isWorkspace"`
	LogsToCapture       []LogsToCapture `json:"logsToCapture"`
	Env                 Env             `json:"env"`
	PaidViaSubscription bool            `json:"paidViaSubscription"`
	CommitID            string          `json:"commitId"`
}

// deployment domain struct
type DeploymentDomain struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	Topic        string `json:"topic"`
	DeploymentID string `json:"deploymentId"`
	ProjectID    string `json:"projectId"`
	Body         Body   `json:"body"`
}

// deployment response struct
type DeploymentResponse struct {
	Deployment Deployment `json:"deployment"`
	LiveLogs   LiveLogs   `json:"liveLogs"`
	Message    string `json:"message,omitempty"`
	Error      bool `json:"error,omitempty"`
}

// deployment struct
type Deployment struct {
	ID                        string        `json:"_id"`
	SitePreview               string        `json:"sitePreview"`
	CommitID                  string        `json:"commitId"`
	CommitMessage             string        `json:"commitMessage"`
	Logs                      []Logs        `json:"logs"`
	BuildDirectory            []string      `json:"buildDirectory"`
	ContentHash               string        `json:"contentHash"`
	Topic                     string        `json:"topic"`
	Status                    string        `json:"status"`
	PaymentID                 string        `json:"paymentId"`
	BuildTime                 int           `json:"buildTime"`
	MemoryUsed                int           `json:"memoryUsed"`
	Env                       Env           `json:"env"`
	Project                   Projects      `json:"project"`
	Screenshot                Screenshot    `json:"screenshot"`
	DeploymentInitiator       string        `json:"deploymentInitiator"`
	Branch                    string        `json:"branch"`
	ExternalRepositoryName    string        `json:"externalRepositoryName"`
	Protocol                  string        `json:"protocol"`
	DeploymentEnvironmentName string        `json:"deploymentEnvironmentName"`
	FailedMessage             string        `json:"failedMessage"`
	IsFromRequest             bool          `json:"isFromRequest"`
	Configuration             Configuration `json:"configuration"`
	CreatedAt                 string        `json:"createdAt"`
	UpdatedAt                 string        `json:"updatedAt"`
}

// live logs struct (possibly unused)
type LiveLogs struct {
}

// authorize deployment payload struct
type AuthorizeDeploymentPayload struct {
	Message string `json:"message"`
	Success bool `json:"success"`
	Topic string `json:"topic"`
	DeploymentId string `json:"deploymentId"`
	ProjectId string `json:"projectId"`
}

// authorize deployment response struct
type AuthorizeDeploymentResponse struct {
	Message      string `json:"message"`
	Success      bool   `json:"success"`
	Topic        string `json:"topic"`
	DeploymentID string `json:"deploymentId"`
	ProjectID    string `json:"projectId"`
	Error      bool `json:"error,omitempty"`
}

// cancel deployment response struct
type CancelDeploymentResponse struct {
	Message  string `json:"message"`
	Canceled bool   `json:"canceled"`
	Killing  bool   `json:"killing"`
	Error      bool `json:"error,omitempty"`
}

// reploy deployment response struct
type RedeployDeploymentResponse struct {
	Message      string `json:"message"`
	Success      bool   `json:"success"`
	Topic        string `json:"topic"`
	DeploymentID string `json:"deploymentId"`
	ProjectID    string `json:"projectId"`
	Error      bool `json:"error,omitempty"`
}

// suggested framework response struct
type SuggestedFrameworkResponse struct {
	SuggestedFramework string `json:"suggestedFramework"`
	Message string `json:"message,omitempty"`
	Error bool `json:"error,omitempty"`
}
	
type UploadFilesDeploymentResponse struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	DeploymentID string `json:"deploymentId"`
	ProjectID    string `json:"projectId"`
	SitePreview  string `json:"sitePreview"`
	Error 		 bool   `json:"error,omitempty"`
}


type FileContent struct {
	Fname string 
	Ftype string
	Fcontent []byte
}