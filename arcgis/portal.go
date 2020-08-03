package arcgis

import (
	"fmt"
	"strconv"
	"time"
)

// PortalAPI todo
type PortalAPI api

// EpochTime todo
type EpochTime time.Time

// UnmarshalJSON todo
func (t *EpochTime) UnmarshalJSON(b []byte) (err error) {
	r := string(b)
	q, err := strconv.ParseInt(r, 10, 64)

	if err != nil {
		return err
	}

	*t = EpochTime(time.Unix(q, 0))
	return nil
}

// User todo
type User struct {
	ID                string    `json:"id"`
	Username          string    `json:"usernanme"`
	FullName          string    `json:"fullName"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Email             string    `json:"email"`
	Description       string    `json:"description"`
	IDPUsername       string    `json:"idpUsername"`
	FavGroupID        string    `json:"favGroupId"`
	PreferredView     string    `json:"preferredView"`
	LastLogin         EpochTime `json:"lastLogin"`
	AvailableCredits  float64   `json:"availableCredits"`
	AssignedCredits   float64   `json:"assignedCredits"`
	MFAEnabled        bool      `json:"mfaEnabled"`
	Access            string    `json:"access"`
	StorageUsage      int       `json:"storageUsage"`
	StorageQuota      int       `json:"storageQuota"`
	OrgID             string    `json:"orgId"`
	Role              string    `json:"role"`
	Privileges        []string  `json:"privileges"`
	Disabled          bool      `json:"disabled"`
	UserLicenseTypeID string    `json:"userLicenseTypeId"`
	Units             string    `json:"units"`
	Tags              []string  `json:"tags"`
	Culture           string    `json:"culture"`
	Region            string    `json:"region"`
	Thumbnail         string    `json:"thumbnail"`
	Created           EpochTime `json:"created"`
	Modified          EpochTime `json:"modified"`
	Provider          string    `json:"provider"`
}

// AppInfo todo
type AppInfo struct {
	AppID    string `json:"appId"`
	ItemID   string `json:"itemId"`
	AppOwner string `json:"appOwner"`
	OrgID    string `json:"orgId"`
	AppTitle string `json:"appTitle"`
}

// Layer todo
type Layer struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	ItemID      string  `json:"itemId"`
	LayerType   string  `json:"layerType"`
	URL         string  `json:"url"`
	Visibility  bool    `json:"visibility"`
	Opacity     float64 `json:"opacity"`
	Title       string  `json:"title"`
	IsReference bool    `json:"isReference"`
}

// BaseMap todo
type BaseMap struct {
	ID                string  `json:"id"`
	Title             string  `json:"title"`
	BaseMapLayers     []Layer `json:"baseMapLayers"`
	OperationalLayers []Layer `json:"operationalLayers"`
}

// Self todo
type Self struct {
	Access                        string    `json:"access"`
	AllSSL                        bool      `json:"allSSL"`
	AnalysisLayersGroupQuery      string    `json:"analysisLayersGroupQuery"`
	AuthorizedCrossOriginDomains  []string  `json:"authorizedCrossOriginDomains"`
	AllowedRedirectUris           []string  `json:"allowedRedirectUris"`
	AvailableCredits              float64   `json:"availableCredits"`
	BackgroundImage               string    `json:"backgroundImage"`
	BasemapGalleryGroupQuery      string    `json:"basemapGalleryGroupQuery"`
	BingKey                       string    `json:"bingKey"`
	CanListApps                   bool      `json:"canListApps"`
	CanListData                   bool      `json:"canListData"`
	CanListPreProvisionedItems    bool      `json:"canListPreProvisionedItems"`
	CanProvisionDirectPurchase    bool      `json:"canProvisionDirectPurchase"`
	CanSearchPublic               bool      `json:"canSearchPublic"`
	CanShareBingPublic            bool      `json:"canShareBingPublic"`
	CanSharePublic                bool      `json:"canSharePublic"`
	CanSignInArcGIS               bool      `json:"canSignInArcGIS"`
	CanSignInIDP                  bool      `json:"canSignInIDP"`
	ColorSetsGroupQuery           string    `json:"colorSetsGroupQuery"`
	CommentsEnabled               bool      `json:"commentsEnabled"`
	ContentCategorySetsGroupQuery string    `json:"contentCategorySetsGroupQuery"`
	Created                       EpochTime `json:"created"`
	CreditAssignments             string    `json:"creditAssignments"`
	Culture                       string    `json:"culture"`
	CultureFormat                 string    `json:"cultureFormat"`
	CustomBaseURL                 string    `json:"customBaseUrl"`
	MFAEnabled                    bool      `json:"mfaEnabled"`
	DefaultBaseMap                BaseMap   `json:"defaultBasemap"`
	User                          User      `json:"user"`
	AppInfo                       AppInfo   `json:"appInfo"`
}

// Self todo
func (api PortalAPI) Self() (*Self, error) {
	req, err := api.client.NewRequest("GET", "/portals/self?f=json", nil)

	if err != nil {
		return nil, fmt.Errorf("Self: unable to create request %+v", err)
	}

	self := &Self{}
	_, errDo := api.client.Do(req, self)

	if errDo != nil {
		return nil, fmt.Errorf("Self: request failed %+v", errDo)
	}

	return self, nil
}
