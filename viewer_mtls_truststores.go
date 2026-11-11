package ioriver

type ViewerMTLSTrustStore struct {
	Id      string `json:"id,omitempty"`
	Account string `json:"account,omitempty"`
	Name    string `json:"name"`
	Bundle  string `json:"bundle"`
}

const viewerMTLSTrustStoresBasePath = "viewer-mtls-truststore/"

func (client *IORiverClient) GetViewerMTLSTrustStore(id string) (*ViewerMTLSTrustStore, error) {
	return Get[ViewerMTLSTrustStore](client, viewerMTLSTrustStoresBasePath+id+"/")
}

func (client *IORiverClient) ListViewerMTLSTrustStores() ([]ViewerMTLSTrustStore, error) {
	return List[ViewerMTLSTrustStore](client, viewerMTLSTrustStoresBasePath)
}

func (client *IORiverClient) CreateViewerMTLSTrustStore(trustStore ViewerMTLSTrustStore) (*ViewerMTLSTrustStore, error) {
	return Create[ViewerMTLSTrustStore](client, viewerMTLSTrustStoresBasePath, trustStore)
}

func (client *IORiverClient) DeleteViewerMTLSTrustStore(id string) error {
	return Delete(client, viewerMTLSTrustStoresBasePath+id+"/")
}
