package models

import "gorm.io/gorm"

// gorm.Model: เป็นโครงสร้างที่มีแท็กของ GORM เพื่อรองรับการจัดการข้อมูลที่เกี่ยวข้องกับฐานข้อมูล เช่น ID, CreatedAt, UpdatedAt, DeletedAt ซึ่งจะถูกเพิ่มโดยอัตโนมัติในฐานข้อมูล
type Fact struct {
	gorm.Model
	Question string `json:"question" gorm:"text;not null;default:null"`
	Answer   string `json:"answer" gorm:"text;not null;default:null"`
}
