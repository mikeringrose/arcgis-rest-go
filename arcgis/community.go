package arcgis

import "fmt"

// CommunityAPI todo
type CommunityAPI api

// Group todo
type Group struct {
	ID               string   `json:"id"`
	Title            string   `json:"title"`
	IsInvitationOnly bool     `json:"isInvitationOnly"`
	Owner            string   `json:"owner"`
	Description      string   `json:"description"`
	Snippet          string   `json:"snippet"`
	Tags             []string `json:"tags"`
	Phone            string   `json:"phone"`
	SortField        string   `json:"sortField"`
	SortOrder        string   `json:"sortOrder"`
}

// Group todo
func (api CommunityAPI) Group(id string) (*Group, error) {
	req, err := api.client.NewRequest("GET", fmt.Sprintf("/community/groups/%s?f=json", id), nil)

	if err != nil {
		return nil, fmt.Errorf("Get: unable to create request %+v", err)
	}

	group := &Group{}
	_, errDo := api.client.Do(req, group)

	if errDo != nil {
		return nil, fmt.Errorf("Get: request failed %+v", errDo)
	}

	return group, nil
}
