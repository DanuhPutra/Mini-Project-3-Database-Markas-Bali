package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (cr *DataBuku) BuatBuku(db *gorm.DB) error {
	err := db.Model(DataBuku{}).Create(&cr).Error
	if err != nil {
		return err
	}
	return nil
}
func (cr *DataBuku) GetByID(db *gorm.DB, id uint) (DataBuku, error) {
	respon := DataBuku{}
	err := db.Model(DataBuku{}).Where("id = ?", id).Take(&respon).Error
	if err != nil {
		return DataBuku{}, err 
	}
	return respon, nil 
}

func (cr *DataBuku) GetSemuaListBukuDatabase(db *gorm.DB) ([]DataBuku, error) {
	respon := []DataBuku{}
	err := db.Model(DataBuku{}).Find(&respon).Error
	if err != nil {
		return []DataBuku{}, err
	}

	return respon, nil
}

func (cr *DataBuku) UpdateOne(db *gorm.DB) error {
	err := db.Model(DataBuku{}).Where("id = ?", cr.ID).Updates(map[string]interface{}{
		"isbn":    cr.ISBN,
		"penulis": cr.Penulis,
		"tahun":   cr.Tahun,
		"judul":   cr.Judul,
		"gambar":  cr.Gambar,
		"stok":    cr.Stok,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (cr *DataBuku) DeleteById(db *gorm.DB) error {
	err := db.Model(DataBuku{}).Where("id = ?", cr.ID).Delete(&cr).Error
	if err != nil {
		return err
	}
	return nil
}
func (book *DataBuku) UpsertBuku(db *gorm.DB) error {
	result := db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface {
		}{
			"isbn": book.ISBN, "penulis": book.Penulis,
			"tahun": book.Tahun, "judul": book.Judul, "gambar": book.Gambar,
			"stok": book.Stok}),
	}).Create(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}