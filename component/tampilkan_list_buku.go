package component

import (
	"DanuhPutra/MiniProject3/config"
	"DanuhPutra/MiniProject3/model"
	"bufio"
	"fmt"
	"os"
	"sort"
	"sync"

	cls "github.com/MasterDimmy/go-cls"
	"gorm.io/gorm"
)

func LihatBuku(ch <-chan model.DataBuku, db *gorm.DB, chPesanan chan model.DataBuku, wg *sync.WaitGroup) {
	for dataSemuaBuku := range ch {
		_, err := dataSemuaBuku.GetSemuaListBukuDatabase(config.Mysql.DB)
		if err != nil {
			fmt.Println("Terjadi error saat membaca file dari database!", err)
			continue
		}
		chPesanan <- dataSemuaBuku
	}
	wg.Done()
}

func TampilkanListBuku(db *gorm.DB){
	cls.CLS()
	fmt.Println("======================================")
	fmt.Println("List Buku yang ada di Perpustakaan ini")
	fmt.Printf("======================================\n")

	var listBook []model.DataBuku
	daftarBuku := &model.DataBuku{}

	books, err := daftarBuku.GetSemuaListBukuDatabase(config.Mysql.DB)
	if err != nil {
		fmt.Printf("Gagal memuat data buku: %v\n", err)
		return
	}
	listBook = books
	wg := sync.WaitGroup{}
	chPesanan := make(chan model.DataBuku, len(listBook))
	ch := make(chan model.DataBuku)

	wg.Add(len(listBook))

	for _, book := range listBook {
		go LihatBuku(ch, db, chPesanan, &wg)
		ch <- book
	}
	close(ch)

	wg.Wait()

	close(chPesanan)

	var orderedBooks []model.DataBuku
	for dataPesanan := range chPesanan {
		orderedBooks = append(orderedBooks, dataPesanan)
	}

	// Urutkan buku berdasarkan waktu pembuatan
	sort.Slice(orderedBooks, func(i, j int) bool {
		return orderedBooks[i].CreatedAt.Before(orderedBooks[j].CreatedAt)
	})

	// Tampilkan buku yang telah diurutkan
	for urutan, book := range orderedBooks {
		fmt.Printf("%d. ID Buku : %d, ISBN : %s, Penulis : %s, Tahun : %d, Judul : %s, Gambar : %s, Stok : %d\n",
			urutan+1,
			book.ID,
			book.ISBN,
			book.Penulis,
			book.Tahun,
			book.Judul,
			book.Gambar,
			book.Stok,
		)
	}

	fmt.Println("\n======================================")
	fmt.Println("Tekan 'Enter' untuk melanjutkan...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

}

