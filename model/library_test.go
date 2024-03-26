package model_test

import (
	"DanuhPutra/MiniProject3/config"
	"DanuhPutra/MiniProject3/model"
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("env not found, using global env")
	}
	config.OpenDB()
}

func TestMemasukanDataKedalamDatabase(t *testing.T) {
	Init()
	dataBukuBaru := model.DataBuku{
		ISBN:    "123456",
		Penulis: "danuh",
		Tahun:   11111,
		Judul:   "kancil dikejar buaya",
		Gambar:  "https://dsadsadsadsadasda/",
		Stok:    100,
	}
	err := dataBukuBaru.BuatBuku(config.Mysql.DB)
	assert.Nil(t, err)
}

func TestGetCarByID(t *testing.T) {
	Init()
	bukuID := model.DataBuku{
		ID: 2,
	}
	data, err := bukuID.GetByID(config.Mysql.DB, bukuID.ID)
	assert.Nil(t, err)
	fmt.Println(data)
}

func TestTampilkanSemuaBukuDatabase(t *testing.T) {
	Init()
	dataSemuaBuku := model.DataBuku{
		ISBN:    "test",
		Penulis: "penulis1",
		Tahun:   1010,
		Judul:   "judul",
		Gambar:  "1021",
		Stok:    100,
	}
	err := dataSemuaBuku.BuatBuku(config.Mysql.DB)
	assert.Nil(t, err)

	res, err := dataSemuaBuku.GetSemuaListBukuDatabase(config.Mysql.DB)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(res), 1)

	fmt.Println(res)
}

func TestUpdate(t *testing.T) {
	Init()
	var err error
	carInsert := model.DataBuku{
		ISBN:    "654645654654",
		Penulis: "danuhputra",
		Tahun:   33333,
		Judul:   "kancil dikejar buaya sama kucing",
		Gambar:  "https://dsadsadsa342432dsadsa/",
		Stok:    10230,
	}
	err = carInsert.BuatBuku(config.Mysql.DB)
	assert.Nil(t, err)

	carData := model.DataBuku{
		ID:      73201,
		ISBN:    "test kancil dikejar buaya updated1",
		Penulis: "danuh updated1",
		Tahun:   22222222,
		Judul:   "judul",
		Gambar:  "dsadsadsadsa",
		Stok:    1002321312321,
	}
	err = carData.UpdateOne(config.Mysql.DB)
	assert.Nil(t, err)
}

func TestDeleteById(t *testing.T) {
	Init()
	var err error
	carInsert := model.DataBuku{
		ISBN:    "test",
		Penulis: "penulis1",
		Tahun:   1010,
		Judul:   "judul",
		Gambar:  "1021",
		Stok:    100,
	}
	err = carInsert.BuatBuku(config.Mysql.DB)
	assert.Nil(t, err)
	carData := model.DataBuku{
		ID: 73201,
	}
	err = carData.DeleteById(config.Mysql.DB)
	assert.Nil(t, err)
}