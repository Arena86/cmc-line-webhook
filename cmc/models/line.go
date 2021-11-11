package models

import "gorm.io/gorm"

type LineEvent struct {
	ID    int    `gorm:"column:Id;primaryKey;autoIncrement;not null" json:"Id"`
	Event string `gorm:"column:Event" json:"Event"`
}

// TableName func
func (LineEvent) TableName() string {
	return "line_event"
}

// BeforeCreate func
func (object *LineEvent) BeforeCreate(db *gorm.DB) (err error) {
	return
}

// BeforeUpdate func
func (object *LineEvent) BeforeUpdate(db *gorm.DB) (err error) {
	return
}
