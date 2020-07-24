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
	ID               string    `json:"id"`
	Username         string    `json:"usernanme"`
	FullName         string    `json:"fullName"`
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	Email            string    `json:"email"`
	Description      string    `json:"description"`
	IDPUsername      string    `json:"idpUsername"`
	FavGroupID       string    `json:"favGroupId"`
	PreferredView    string    `json:"preferredView"`
	LastLogin        EpochTime `json:"lastLogin"`
	AvailableCredits float64   `json:"availableCredits"`
	AssignedCredits  float64   `json:"assignedCredits"`
}

// Self todo
type Self struct {
	User User `json:"user"`
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
