package models

type Pricing struct {
	CommitmentLevel int `gorm:"type:number"`
	Weightage       int `gorm:"type:number"`
}
