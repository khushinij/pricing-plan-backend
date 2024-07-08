package models

type Parameter struct {
	ID          uint   `gorm:"primaryKey"`
	Volume      string `gorm:"type:varchar(255)"`
	Terminal    string `gorm:"type:varchar(255)"`
	MCC         int    `gorm:"type:number"`
	Competition string `gorm:"type:varchar(255)"`
	Round       string `gorm:"type:varchar(255)"`
	UPI         string `gorm:"type:varchar(255)"`
	NB          string `gorm:"type:varchar(255)"`
	CC          string `gorm:"type:varchar(255)"`
	DC          string `gorm:"type:varchar(255)"`
	Others      string `gorm:"type:varchar(255)"`
}
