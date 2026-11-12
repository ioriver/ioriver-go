package ioriver

const attachedFilesBasePath = "attached-files/"

type AttachedFile struct {
	ID       string `json:"id,omitempty"`
	Account  string `json:"account,omitempty"`
	FileType string `json:"file_type"`
	Contents string `json:"contents"`
}

// CreateAttachedFile uploads an immutable account-scoped text attachment.
// Identical contents and file types within the account may reuse an existing ID.
func (client *IORiverClient) CreateAttachedFile(fileType, contents string) (*AttachedFile, error) {
	return Create[AttachedFile](client, attachedFilesBasePath, AttachedFile{
		FileType: fileType,
		Contents: contents,
	})
}

// GetAttachedFile retrieves attachment metadata and text contents.
func (client *IORiverClient) GetAttachedFile(id string) (*AttachedFile, error) {
	return Get[AttachedFile](client, attachedFilesBasePath+id+"/")
}

// DeleteAttachedFile fails while any service configuration version references the attachment.
func (client *IORiverClient) DeleteAttachedFile(id string) error {
	return Delete(client, attachedFilesBasePath+id+"/")
}
