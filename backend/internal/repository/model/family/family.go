package family

import (
	"cashback_info/internal/repository/model/user"

	"github.com/google/uuid"
)

type FamilyDB struct {
	ID       uuid.UUID     `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title    string        `gorm:"type:varchar(50);not null" json:"title"`
	LeaderID uuid.UUID     `gorm:"type:uuid;not null" json:"leader_id"`
	Leader   user.UserDB   `gorm:"foreignKey:leader_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"leader"`
	Members  []user.UserDB `gorm:"many2many:families_users;foreignKey:ID;joinForeignKey:family_id;References:ID;joinReferences:user_id" json:"members"`
}

func (FamilyDB) TableName() string {
	return "families"
}
