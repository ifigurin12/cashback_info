package card

type BankType string

const (
	BankTypeVTB     BankType = "vtb"
	BankTypeAlfa    BankType = "alfa"
	BankTypeTinkoff BankType = "tinkoff"
	BankTypePochta  BankType = "pochta"
	BankTypeGazprom BankType = "gazprom"
)

func CreateBankTypeFromString(value string) *BankType {
	var result BankType
	switch value {
	case string(BankTypeVTB):
		result = BankTypeVTB
	case string(BankTypeAlfa):
		result = BankTypeAlfa
	case string(BankTypeTinkoff):
		result = BankTypeTinkoff
	case string(BankTypePochta):
		result = BankTypePochta
	case string(BankTypeGazprom):
		result = BankTypeGazprom
	default:
		return nil
	}

	return &result
}
