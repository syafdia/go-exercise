package entity

type TargetType = string

const (
	TargetTypeLegalEntity TargetType = "legal_entity"
	TargetTypeIndustry    TargetType = "industry"
)

type TargetAction struct {
	ID         int64
	ResourceID int64
	Name       string
	TargetID   int64
	TargetType TargetType
}
