package auth

type CheckResultResponse struct {
	Kind    string            `json:"kind"`
	Actions map[string]string `json:"actions"`
}

type CheckResponse struct {
	Results []CheckResultResponse `json:"results"`
}
