package entity

type Action struct {
	ID         int64  `db:"id"`
	ResourceID int64  `db:"resource_id"`
	Name       string `db:"name"`
}
