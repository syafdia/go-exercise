package auth

type CheckResourceRequest struct {
	Kind string `json:"kind"`
}

type CheckPrincipalRequest struct {
	Kind     string            `json:"kind"`
	ID       string            `json:"id"`
	Roles    []string          `json:"roles"`
	Metadata map[string]string `json:"metadata"`
}

type CheckRequest struct {
	Actions   []string              `json:"actions"`
	Resource  CheckResourceRequest  `json:"resource"`
	Principal CheckPrincipalRequest `json:"principal"`
}
