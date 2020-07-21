package gochimp3

const (
	templateFoldersPath = "/template-folders"
	// single folder endpoint not implemented
)

type TemplateFolderQueryParams struct {
	ExtendedQueryParams
}

type ListOfTemplateFolders struct {
	baseList
	Folders []TemplateFolder `json:"folders"`
}

type TemplateFolder struct {
	withLinks

	Count uint   `json:"count"`
	ID    string `json:"id"`
	Name  string `json:"name"`

	api *API
}

type TemplateFolderCreationRequest struct {
	Name string `json:"name"`
}

func (api API) GetTemplateFolders(params *TemplateFolderQueryParams) (*ListOfTemplateFolders, error) {
	response := new(ListOfTemplateFolders)

	err := api.Request("GET", templateFoldersPath, params, nil, response)
	if err != nil {
		return nil, err
	}

	for _, l := range response.Folders {
		l.api = &api
	}

	return response, nil
}

func (api API) CreateTemplateFolder(body *TemplateFolderCreationRequest) (*TemplateFolder, error) {
	response := new(TemplateFolder)
	response.api = &api
	return response, api.Request("POST", templateFoldersPath, nil, body, response)
}
