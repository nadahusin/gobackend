package models

import "time"

type Vehicle struct {
	DetailId  uint      `gorm:"primaryKey" json:"detail_id"`
	Goods     string    `json:"goods"`
	City      string    `json:"city"`
	Status    string    `json:"status"`
	Capacity  int       `json:"capacity"`
	Type      string    `json:"type"`
	Rent      string    `json:"rent"`
	Price     string    `json:"price"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Image     string    `json:"image"`
}

type Vehicles []Vehicle
