package gochimp3

import (
	"errors"
	"fmt"
)

const (
	membersPath      = "/lists/%s/members"
	singleMemberPath = membersPath + "/%s"

	memberActivityPath = singleMemberPath + "/activity"
	memberGoalsPath    = singleMemberPath + "/goals"

	memberNotesPath      = singleMemberPath + "/notes"
	singleMemberNotePath = memberNotesPath + "/%s"
)

type ListOfMembers struct {
	baseList

	ListID  string   `json:"list_id"`
	Members []Member `json:"members"`
}

type Tag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type MemberRequest struct {
	EmailAddress    string                 `json:"email_address"`
	EmailType       string                 `json:"email_type,omitempty"`
	Status          string                 `json:"status"`
	StatusIfNew     string                 `json:"status_if_new,omitempty"`
	MergeFields     map[string]interface{} `json:"merge_fields,omitempty"`
	Interests       map[string]bool        `json:"interests,omitempty"`
	Language        string                 `json:"language"`
	VIP             bool                   `json:"vip"`
	Location        *MemberLocation        `json:"location,omitempty"`
	IPOpt           string                 `json:"ip_opt,omitempty"`
	IPSignup        string                 `json:"ip_signup,omitempty"`
	Tags            []string               `json:"tags,omitempty"`
	TimestampSignup string                 `json:"timestamp_signup,omitempty"`
	TimestampOpt    string                 `json:"timestamp_opt,omitempty"`
}

type Member struct {
	MemberRequest

	ID            string          `json:"id"`
	ListID        string          `json:"list_id"`
	UniqueEmailID string          `json:"unique_email_id"`
	EmailType     string          `json:"email_type"`
	Stats         MemberStats     `json:"stats"`
	MemberRating  int             `json:"member_rating"`
	LastChanged   string          `json:"last_changed"`
	EmailClient   string          `json:"email_client"`
	LastNote      MemberNoteShort `json:"last_note"`
	Tags          []Tag           `json:"tags"`

	api *API
}

func (mem Member) CanMakeRequest() error {
	if mem.ListID == "" {
		return errors.New("No ListID provided")
	}

	if mem.ID == "" {
		return errors.New("No ID provided")
	}

	return nil
}

type MemberStats struct {
	AvgOpenRate  float64 `json:"avg_open_rate"`
	AvgClickRate float64 `json:"avg_click_rate"`
}

type MemberLocation struct {
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	GMTOffset   int     `json:"gmtoff"`
	DSTOffset   int     `json:"dstoff"`
	CountryCode string  `json:"country_code"`
	Timezone    string  `json:"timezone"`
}

type MemberNoteShort struct {
	ID        int    `json:"note_id"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
	Note      string `json:"note"`
}

type MemberTag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (list ListResponse) GetMembers(params *InterestCategoriesQueryParams) (*ListOfMembers, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(membersPath, list.ID)
	response := new(ListOfMembers)

	err := list.api.Request("GET", endpoint, params, nil, response)
	if err != nil {
		return nil, err
	}

	for _, m := range response.Members {
		m.api = list.api
	}

	return response, nil
}

func (list ListResponse) GetMember(id string, params *BasicQueryParams) (*Member, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(singleMemberPath, list.ID, id)
	response := new(Member)
	response.api = list.api

	return response, list.api.Request("GET", endpoint, params, nil, response)
}

func (list ListResponse) CreateMember(body *MemberRequest) (*Member, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(membersPath, list.ID)
	response := new(Member)
	response.api = list.api

	return response, list.api.Request("POST", endpoint, nil, body, response)
}

func (list ListResponse) UpdateMember(id string, body *MemberRequest) (*Member, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(singleMemberPath, list.ID, id)
	response := new(Member)
	response.api = list.api

	return response, list.api.Request("PATCH", endpoint, nil, body, response)
}

func (list ListResponse) AddOrUpdateMember(id string, body *MemberRequest) (*Member, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(singleMemberPath, list.ID, id)
	response := new(Member)
	response.api = list.api

	return response, list.api.Request("PUT", endpoint, nil, body, response)
}

func (list ListResponse) DeleteMember(id string) (bool, error) {
	if err := list.CanMakeRequest(); err != nil {
		return false, err
	}

	endpoint := fmt.Sprintf(singleMemberPath, list.ID, id)
	return list.api.RequestOk("DELETE", endpoint)
}

// ------------------------------------------------------------------------------------------------
// Activity
// ------------------------------------------------------------------------------------------------

type ListOfMemberActivity struct {
	baseList

	EmailID  string     `json:"email_id"`
	ListID   string     `json:"list_id"`
	Activity []Activity `json:"activity"`
}

type MemberActivity struct {
	Action         string `json:"action"`
	Timestamp      string `json:"timestamp"`
	URL            string `json:"url"`
	Type           string `json:"type"`
	CampaignID     string `json:"campaign_id"`
	Title          string `json:"title"`
	ParentCampaign string `json:"parent_campaign"`
}

func (mem Member) GetActivity(params *BasicQueryParams) (*ListOfMemberActivity, error) {
	if err := mem.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(memberActivityPath, mem.ListID, mem.ID)
	response := new(ListOfMemberActivity)

	return response, mem.api.Request("GET", endpoint, params, nil, response)
}

// ------------------------------------------------------------------------------------------------
// Goals
// ------------------------------------------------------------------------------------------------

type ListOfMemberGoals struct {
	baseList

	EmailID string       `json:"email_id"`
	ListID  string       `json:"list_id"`
	Goals   []MemberGoal `json:"goals"`
}

type MemberGoal struct {
	ID            int    `json:"goal_id"`
	Event         string `json:"event"`
	LastVisitedAt string `json:"last_visited_at"`
	Data          string `json:"data"`
}

func (mem Member) GetGoals(params *BasicQueryParams) (*ListOfMemberGoals, error) {
	if err := mem.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(memberGoalsPath, mem.ListID, mem.ID)
	response := new(ListOfMemberGoals)

	return response, mem.api.Request("GET", endpoint, params, nil, response)
}

// ------------------------------------------------------------------------------------------------
// NOTES
// ------------------------------------------------------------------------------------------------

type ListOfMemberNotes struct {
	baseList

	EmailID string           `json:"email_id"`
	ListID  string           `json:"list_id"`
	Notes   []MemberNoteLong `json:"notes"`
}

type MemberNoteLong struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt string `json:"updated_at"`
	Note      string `json:"note"`
	ListID    string `json:"list_id"`
	EmailID   string `json:"email_id"`

	withLinks
}

func (mem Member) GetNotes(params *ExtendedQueryParams) (*ListOfMemberNotes, error) {
	if err := mem.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(memberNotesPath, mem.ListID, mem.ID)
	response := new(ListOfMemberNotes)

	return response, mem.api.Request("GET", endpoint, params, nil, response)
}

func (mem Member) CreateNote(msg string) (*MemberNoteLong, error) {
	if err := mem.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(memberNotesPath, mem.ListID, mem.ID)
	response := new(MemberNoteLong)

	body := struct{ Note string }{
		Note: msg,
	}

	return response, mem.api.Request("POST", endpoint, nil, &body, response)
}

func (mem Member) UpdateNote(id, msg string) (*MemberNoteLong, error) {
	if err := mem.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(singleMemberNotePath, mem.ListID, mem.ID, id)
	response := new(MemberNoteLong)

	body := struct{ Note string }{
		Note: msg,
	}

	return response, mem.api.Request("PATCH", endpoint, nil, &body, response)
}

func (mem Member) GetNote(id string, params *BasicQueryParams) (*MemberNoteLong, error) {
	if err := mem.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(singleMemberNotePath, mem.ListID, mem.ID, id)
	response := new(MemberNoteLong)

	return response, mem.api.Request("GET", endpoint, params, nil, response)
}

func (mem Member) DeleteNote(id string) (bool, error) {
	if err := mem.CanMakeRequest(); err != nil {
		return false, err
	}

	endpoint := fmt.Sprintf(singleMemberNotePath, mem.ListID, mem.ID, id)
	return mem.api.RequestOk("DELETE", endpoint)
}
