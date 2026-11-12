package ioriver

type Tag struct {
	Id            string `json:"id,omitempty"`
	Account       string `json:"account,omitempty"`
	Name          string `json:"name"`
	Color         string `json:"color,omitempty"`
	ServicesCount int    `json:"services_count,omitempty"`
}

const tagsBasePath = "tags/"

func (client *IORiverClient) GetTag(id string) (*Tag, error) {
	return Get[Tag](client, tagsBasePath+id+"/")
}

func (client *IORiverClient) ListTags() ([]Tag, error) {
	return List[Tag](client, tagsBasePath)
}

func (client *IORiverClient) CreateTag(tag Tag) (*Tag, error) {
	return Create[Tag](client, tagsBasePath, tag)
}

func (client *IORiverClient) UpdateTag(tag Tag) (*Tag, error) {
	return Update[Tag](client, tagsBasePath+tag.Id+"/", tag)
}

func (client *IORiverClient) DeleteTag(id string) error {
	return Delete(client, tagsBasePath+id+"/")
}
