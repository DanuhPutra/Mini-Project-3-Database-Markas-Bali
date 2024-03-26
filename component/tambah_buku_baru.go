package component

import (
	"DanuhPutra/MiniProject3/config"
	"DanuhPutra/MiniProject3/model"
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	cls "github.com/MasterDimmy/go-cls"
	"gorm.io/gorm"
)
var inputUser = bufio.NewReader(os.Stdin)

func TambahBukuBaru(db *gorm.DB){
	cls.CLS()

	// deklarasi variabel
	var id, tahun, stok uint
	judulBukuBaru := bufio.NewReader(os.Stdin)
	isbnBukuBaru := bufio.NewReader(os.Stdin)
	penulisBukuBaru := bufio.NewReader(os.Stdin)
	gambarBukuBaru := bufio.NewReader(os.Stdin)

	fmt.Println("==============================")
	fmt.Println("Menambahkan Buku Baru")
	fmt.Printf("==============================\n")
	simpanBuku := []model.DataBuku{}

	for {
		// kode buku
			fmt.Print("Kode Buku Baru : ")
			_, err := fmt.Scanln(&id)
			if err != nil {
				fmt.Println("Ups, Terjadi error pada Kode Buku!", err)
				return
			}
			// id = strings.TrimSpace(id)

			// if kodeBukuExists(id) {
			// 	fmt.Println("Kode buku sudah digunakan. Masukkan kode buku yang berbeda.")
			// } else {
			// 	break
			// }

		// isbn buku
		fmt.Print("Isbn Buku Baru : ")
		isbnBukuTambah, err := isbnBukuBaru.ReadString('\n')
		if err != nil {
			fmt.Println("Ups, Terjadi error pada Judul Buku!", err)
			return
		}
		isbnBukuTambah = strings.TrimSpace(isbnBukuTambah)

		// judul buku
		fmt.Print("Judul Buku Baru : ")
		JudulBukuTambah, err := judulBukuBaru.ReadString('\n')
		if err != nil {
			fmt.Println("Ups, Terjadi error pada Judul Buku!", err)
			return
		}
		JudulBukuTambah = strings.TrimSpace(JudulBukuTambah)

		// penulis buku
		fmt.Print("Penulis Buku Baru : ")
		PengarangBukuTambah, err := penulisBukuBaru.ReadString('\n')
		if err != nil {
			fmt.Println("Ups, Terjadi error pada Penulis Buku!", err)
			return
		}
		PengarangBukuTambah = strings.TrimSpace(PengarangBukuTambah)

		// gambar buku
		fmt.Print("Gambar Buku Baru : ")
		GambarBukuTambah, err := gambarBukuBaru.ReadString('\n')
		if err != nil {
			fmt.Println("Ups, Terjadi error pada Penulis Buku!", err)
			return
		}
		GambarBukuTambah = strings.TrimSpace(GambarBukuTambah)

		// tahun buku
		fmt.Print("silahkan masukan Tahun Terbit pada Buku Baru :")
		_, err = fmt.Scanln(&tahun)
		if err != nil {
			fmt.Println("Ups, Terjadi error pada Tahun Terbit Buku! :", err)
			return
		}

		// tahun terbit buku
		fmt.Print("silahkan masukan Stok Buku :")
		_, err = fmt.Scanln(&stok)
		if err != nil {
			fmt.Println("Ups, Terjadi error pada Stok Buku! :", err)
			return
		}

		simpanBuku = append(simpanBuku, model.DataBuku{
			ID : id,
			ISBN : isbnBukuTambah,
			Penulis : PengarangBukuTambah,
			Tahun : tahun,
			Judul : JudulBukuTambah,
			Gambar : GambarBukuTambah,
			Stok : stok,
		})

		var pilihanMenuBuku = 0
		fmt.Println("Ketik 1 untuk menambah buku lagi, Ketik 0 untuk kembali")
		_, err = fmt.Scanln(&pilihanMenuBuku)
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}

		if pilihanMenuBuku == 0 {
			break
		}

	}

	fmt.Println("Menambahkan Buku Kedalam Perpustakaan...")
	_ = os.Mkdir("books", 0755)
	ch := make(chan model.DataBuku)
	wg := sync.WaitGroup{}
	jumlahStafBuku := 5

	for i := 0; i < jumlahStafBuku; i++ {
		wg.Add(1)
		go simpanBukuTambahan(ch, db, &wg, i)
	}

	for _, kodeBuku := range simpanBuku {
		ch <- kodeBuku
	}

	close(ch)
	wg.Wait()

	fmt.Println("berhasil menambahkan buku baru kedalam perpustakaan!")
	fmt.Println("\n======================================")
	fmt.Println("Tekan 'Enter' untuk melanjutkan...")		
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func simpanBukuTambahan(ch <-chan model.DataBuku, db *gorm.DB, wg *sync.WaitGroup, noStaff int){
	for buku := range ch {
		err := buku.BuatBuku(config.Mysql.DB)
		if err != nil {
			fmt.Println("terjadi error!", err)
			continue
		}

		fmt.Printf("staff No %d Memproses buku baru dengan KodeBuku : %d!\n", noStaff, buku.ID)

	}
	wg.Done()
}