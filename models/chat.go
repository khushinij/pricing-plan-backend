package models

type ChatMessage struct {
	ID         uint `gorm:"primaryKey"`
	MerchantID string
	Timestamp  int64 `gorm:"index;sort:desc"` // Ensure index for sorting in descending order
	Text       string
	IsBot      bool
}

//CREATE TABLE chat_messages (
//ID SERIAL PRIMARY KEY,
//Merchant_ID VARCHAR(255) UNIQUE,
//Timestamp INT,
//Text VARCHAR(255),
//IsBot BOOLEAN
//);
