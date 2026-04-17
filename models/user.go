package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirebaseUID  string `gorm:"uniqueIndex;size:128;not null"
json: "firebase_uid"`
	Email string `gorm:"uniqueIndex;size:255;not null"`
	Name  string `gorm:"size:100"`
	Role  string `gorm:"size:20;default:'user'"` // UID dari Firebase Auth
	EmailVerified bool   `gorm:"default:false"
json: "email_verified"`
	LastLoginAt   int64  `gorm:"index"
json: "last_login_at, omitempty"`
}

// gorm.Model memberikan fields: 
// ID         uint      (primary key, auto increment) 
// CreatedAt  time.Time (auto fill saat insert) 
// UpdatedAt  time.Time (auto fill saat update) 
// DeletedAt  gorm.DeletedAt (soft delete — row tidak benar-benar dihapus) 
// Struct tag "gorm" mengontrol perilaku GORM: 
// uniqueIndex = buat unique index di kolom ini 
// size:128    = varchar(128) 
// not null    = kolom tidak boleh NULL 
// default:user = nilai default kolom 
// Struct tag "json" mengontrol serialisasi JSON: 
// json:"firebase_uid" = nama key di JSON response 
// json:"-"            = field ini tidak dimasukkan ke JSON 
// json:"...,omitempty" = skip jika nilnya zero/nil 