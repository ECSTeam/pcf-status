package models

type cfData struct {
	Name     string `json:"name"`
	FileName string `json:"filename"`
	GUID     string `json:"guid"`
	Created  string `json:"created_at"`
	Updated  string `json:"updated_at"`
}

type cfContainer struct {
	Resources []struct {
		Entry struct {
			Name     string `json:"name"`
			FileName string `json:"filename"`
		} `json:"entity"`
		Metadata struct {
			GUID    string `json:"guid"`
			Created string `json:"created_at"`
			Updated string `json:"updated_at"`
		} `json:"metadata"`
	} `json:"resources"`
}

func (container *cfContainer) Dump(err error) interface{} {
	var data []cfData
	if err == nil {
		for _, resource := range container.Resources {
			data = append(data, cfData{
				Name:     resource.Entry.Name,
				FileName: resource.Entry.FileName,
				GUID:     resource.Metadata.GUID,
				Created:  resource.Metadata.Created,
				Updated:  resource.Metadata.Updated,
			})
		}
	}

	return data
}
