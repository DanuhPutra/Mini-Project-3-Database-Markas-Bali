package model

import (
	"time"

	"gorm.io/gorm"
)

type DataBuku struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ISBN      string         `json:"isbn"`
	Penulis   string         `json:"penulis"`
	Tahun     uint           `json:"tahun"`
	Judul     string         `json:"judul"`
	Gambar    string         `json:"gambar"`
	Stok      uint           `json:"stok"`
}