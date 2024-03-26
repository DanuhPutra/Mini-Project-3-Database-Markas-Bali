package component

import (
	"DanuhPutra/MiniProject3/config"
	"DanuhPutra/MiniProject3/model"
	"bufio"
	"fmt"
	"os"

	cls "github.com/MasterDimmy/go-cls"
	"gorm.io/gorm"
)

var IDPerubahan uint

func HapusDataBukuPerpustakaan(db *gorm.DB) {
	cls.CLS()
	fmt.Println("==============================")
	fmt.Println("Hapus Buku dari Perpustakaan")
	fmt.Println("==============================")

	TampilkanListBuku(db)

	fmt.Print("Masukkan ID buku yang ingin dihapus: ")
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

	err = buku.DeleteById(config.Mysql.DB)
	if err != nil {
		fmt.Printf("Terjadi error saat menghapus buku: %v\n", err)
		return
	}
	fmt.Println("Buku dengan ID", buku.ID, " dengan judul : ", buku.Judul, " Berhasil Dihapus.")
	fmt.Println("Buku berhasil dihapus dari perpustakaan.")
	fmt.Println("Tekan 'Enter' untuk kembali...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}