package datatypes

type OpenGraphTag struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type OpenGraph struct {
	Tags []OpenGraphTag `json:"tags"`
}
