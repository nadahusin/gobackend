package models

import "time"

type History struct {
	IdHistory uint      `gorm:"primaryKey" json:"id_history"`
	Username  string    `json:"username"`
	Address   string    `json:"address"`
	GoodName  string    `json:"good_name"`
	BookedAt  string    `json:"booked_at"`
	SubTotal  string    `json:"sub_total"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Histories []History
