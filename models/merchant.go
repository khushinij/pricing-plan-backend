package models

type Merchant struct {
	ID              uint    `gorm:"primaryKey"`
	MerchantID      string  `gorm:"uniqueIndex"`
	TerminalGateway string  `gorm:"type:varchar(255)"`
	MCC             int     `gorm:"type:number"`
	GMV             float64 `gorm:"type:number"`
}
