package gochimp3

const (
	// single folder endpoint not implemented
	campaignFoldersPath = "/campaign-folders"
)

type CampaignFolderQueryParams struct {
	ExtendedQueryParams
}

type ListOfCampaignFolders struct {
	baseList
	Folders []CampaignFolder `json:"folders"`
}

type CampaignFolder struct {
	withLinks

	ID    string `json:"id"`
	Count uint   `json:"count"`
	Name  string `json:"name"`

	api *API
}

type CampaignFolderCreationRequest struct {
	Name string `json:"name"`
}

func (api API) GetCampaignFolders(params *CampaignFolderQueryParams) (*ListOfCampaignFolders, error) {
	response := new(ListOfCampaignFolders)

	err := api.Request("GET", campaignFoldersPath, params, nil, response)
	if err != nil {
		return nil, err
	}

	for _, l := range response.Folders {
		l.api = &api
	}

	return response, nil
}

func (api API) CreateCampaignFolder(body *CampaignFolderCreationRequest) (*CampaignFolder, error) {
	response := new(CampaignFolder)
	response.api = &api
	return response, api.Request("POST", campaignFoldersPath, nil, body, response)
}
