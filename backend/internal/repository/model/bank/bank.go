package bank

type Bank struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(50);unique;not null" json:"name"`
}

func (b *Bank) TableName() string {
	return "banks"
}
