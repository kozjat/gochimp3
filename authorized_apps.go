package gochimp3

import "fmt"

const (
	authorizedAppsPath = "/authorized-apps"
	authorizedAppPath  = authorizedAppsPath + "/%s"
)

type ListOfAuthorizedApps struct {
	baseList `json:""`
	Apps     []AuthorizedApp `json:""`
}

type AuthorizedAppRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AuthorizedApp struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Users       []string `json:"users "`
	withLinks
}

type AuthorizedAppCreateResponse struct {
	AccessToken string `json:"access_token"`
	ViewerToken string `json:"viewer_token"`
}

func (api API) GetAuthorizedApps(params *ExtendedQueryParams) (*ListOfAuthorizedApps, error) {
	response := new(ListOfAuthorizedApps)

	err := api.Request("GET", authorizedAppsPath, params, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (api API) CreateAuthorizedApp(body *AuthorizedAppRequest) (*AuthorizedAppCreateResponse, error) {
	response := new(AuthorizedAppCreateResponse)

	err := api.Request("GET", authorizedAppsPath, nil, body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (api API) GetAuthorizedApp(id string, params *BasicQueryParams) (*AuthorizedApp, error) {
	response := new(AuthorizedApp)
	endpoint := fmt.Sprintf(authorizedAppPath, id)

	err := api.Request("GET", endpoint, params, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
