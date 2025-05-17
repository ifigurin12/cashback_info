package bank

type BankDB struct {
	ID   int32  `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(50);unique;not null" json:"name"`
}

func (b *BankDB) TableName() string {
	return "banks"
}
