package main

import (
	components "DanuhPutra/MiniProject3/component"
	"DanuhPutra/MiniProject3/config"
	"fmt"
	"os"

	"github.com/MasterDimmy/go-cls"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func Init(){
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Env not found, using system environment")
	}
	config.OpenDB()
}
func main(){
	cls.CLS()
	var PilihanAksi int
	Init()

	fmt.Println("===========================================")
	fmt.Println("Aplikasi Manajemen Daftar Buku Perpustakaan")
	fmt.Println("===========================================")
	fmt.Println("silahkan pilih menu : ")
	fmt.Println("1. Menambahkan Buku Baru Perpustakaan")
	fmt.Println("2. Menampilkan Buku Perpustakaan")
	fmt.Println("3. Hapus Buku Perpustakaan")
	fmt.Println("4. Edit Buku Perpustakaan")
	fmt.Println("5. Import .csv file to database")
	fmt.Println("6. Keluar dari Program")
	fmt.Println("===========================================")
	fmt.Print("masukan pilihan : ")
	_, err := fmt.Scanln(&PilihanAksi)
	if err != nil {
		fmt.Println("Ups, Terjadi error pada aksi yang kamu pilih!", err)
	}

	switch PilihanAksi{
		case 1 :
			components.TambahBukuBaru(&gorm.DB{})
		case 2 :
			components.TampilkanListBuku(&gorm.DB{})
		case 3 :
			components.HapusDataBukuPerpustakaan(&gorm.DB{})
		case 4 :
			components.UpdateDataBukuPerpustakaan(&gorm.DB{})
		case 5 : 
			components.ImportFile()
		case 6 : 
			os.Exit(0)
	}

	main()
}