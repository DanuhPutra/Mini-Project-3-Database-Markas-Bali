package component

import (
	"DanuhPutra/MiniProject3/config"
	"DanuhPutra/MiniProject3/model"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func UpdateDataBukuPerpustakaan(db *gorm.DB) {
	fmt.Println("=================================")
	fmt.Println("Edit Buku")
	fmt.Println("=================================")
	TampilkanListBuku(db)

	var buku model.DataBuku
	var isbnBaru, penulisBaru, judulBaru, gambarBaru string
	var stokBaru, tahunBaru, idbaru uint
	var err error
	for {
		fmt.Print("Masukan ID Buku yang Ingin Dihapus : ")
		_, err := fmt.Scanln(&IDPerubahan)
		if err != nil {
			fmt.Println("Terjadi error:", err)
			return
		}

		var buku model.DataBuku
		buku, err = buku.GetByID(config.Mysql.DB, IDPerubahan)
		fmt.Println("ini test", buku)
		if err != nil {
			fmt.Println("Buku dengan ID", IDPerubahan, "tidak ditemukan")
			return
		}
		break
	}

	fmt.Print("Jumlah Id Baru :")
	_, err = fmt.Fscanf(inputUser, "%d\n", &idbaru)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Print("ISBN Baru: ")
	isbnBaru, err = inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	isbnBaru = strings.TrimSpace(isbnBaru)

	fmt.Print("Penulis Baru: ")
	penulisBaru, err = inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	penulisBaru = strings.TrimSpace(penulisBaru)

	fmt.Print("Tahun Baru :")
	_, err = fmt.Fscanf(inputUser, "%d\n", &tahunBaru)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Print("Judul Baru: ")
	judulBaru, err = inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	judulBaru = strings.TrimSpace(judulBaru)

	fmt.Print("Gambar Baru: ")
	gambarBaru, err = inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	gambarBaru = strings.TrimSpace(gambarBaru)

	fmt.Print("Jumlah Stok Baru :")
	_, err = fmt.Fscanf(inputUser, "%d\n", &stokBaru)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	buku.ID = idbaru
	buku.ISBN = judulBaru
	buku.Judul = judulBaru
	buku.Penulis = penulisBaru
	buku.Tahun = tahunBaru
	buku.Judul = judulBaru
	buku.Gambar = gambarBaru
	buku.Stok = stokBaru

	if err := buku.UpdateOne(config.Mysql.DB); err != nil {
		fmt.Println("Terjadi error saat mengubah data buku:", err)
		return
	}
	fmt.Println("Data buku berhasil diperbarui.")

}