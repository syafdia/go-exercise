package auth

type FindResourcesByComplianceRequest struct {
	IndustryID    int64
	LegalEntityID int64
	Resources     []string
}
