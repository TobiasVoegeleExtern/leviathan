package models

import "time"

type Haushaltsausgaben struct {
	ID              int       `gorm:"primaryKey"`
	Description     string    `gorm:"type:text"`
	ValueTotal      float64   `gorm:"column:valuetotal;type:numeric"`
	ValueRate       float64   `gorm:"column:valuerate;type:numeric"`
	CreditStart     time.Time `gorm:"column:creditstart"`
	CreditEnd       time.Time `gorm:"column:creditend"`
	Type            string    `gorm:"type:varchar"`
	UserID          int       `gorm:"column:userid;index"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	ChangedAt       time.Time `gorm:"autoUpdateTime"`
	Faelligkeitstag string    `gorm:"column:faelligkeitstag;type:varchar"`
	Zahldatum       time.Time `gorm:"type:timestamp"`
	Receipt         []byte    `gorm:"column:receipt;type:bytea"`
}

func (Haushaltsausgaben) TableName() string {
	return "haushaltsausgaben"
}
