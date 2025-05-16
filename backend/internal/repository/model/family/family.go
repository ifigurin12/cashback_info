package family

type Family struct {
	ID       string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title    string `gorm:"type:varchar(50);not null" json:"title"`
	LeaderID string `gorm:"type:uuid;not null" json:"leader_id"`
}
