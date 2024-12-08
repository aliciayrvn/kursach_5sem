package models

type Car struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Model        string `json:"model" binding:"required"`
	LicensePlate string `json:"license_plate" binding:"required"`
	Available    bool   `json:"available"`
}
