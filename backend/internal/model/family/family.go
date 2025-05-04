package family

import "cashback_info/internal/model/user"

type Family struct {
	ID     string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title  string    `gorm:"type:varchar(50);not null" json:"title"`
	Leader user.User `gorm:"foreignKey:leader_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"leader"`
}
