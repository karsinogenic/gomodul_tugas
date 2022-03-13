package modul

//coba

import "fmt"

type Item struct {
	Nama     string
	Stock    int
	Harga    int
	Nbelanja int
	Ntotal   int
}

func Kasir(uang, bayar int) int {
	var kembali int
	kembali = uang - bayar

	return kembali
}

func Troli(Harga, banyak int) (int, bool) {
	var bayar int
	var choice bool
	var choice1 string
	bayar = (Harga * banyak) + bayar
	fmt.Println("total belanja anda adalah ", bayar)
ulang:
	fmt.Print("ingin terus belanja (y/n)?")
	fmt.Scanln(&choice1)
	if choice1 == "y" {
		choice = true
	} else if choice1 == "n" {
		choice = false
	} else {
		fmt.Print("Input salah?")
		goto ulang
	}
	return bayar, choice
}

// var minyak = item{"minyak", 10, 30000, 0, 0}

func (item Item) Belanja1() (int, int) {
	var banyak int
	var Ntotal int
ulang1:
	fmt.Printf("Anda memilih %s, berapa banyak yang ingin anda beli : ", item.Nama)
	//fmt.Printf("1.%s Stock:%d Harga:%s || 2.%s Stock:%d Harga:%s\n//////////////////////////////////////////////////////\n", minyak.Nama, minyak.Stock, minyak.Harga, terigu.Nama, terigu.Stock, terigu.Harga)
	//fmt.Print("Pilihan anda (1-2) : ")
	fmt.Scanln(&banyak)
	if banyak > item.Stock {
		fmt.Println("Stock tidak mencukupi, jumlah Stock : ", item.Stock)
		goto ulang1
	} else {
		item.Nbelanja = item.Nbelanja + banyak
		fmt.Println("Total minyak yang anda beli : ", item.Nbelanja)
		item.Ntotal = item.Ntotal + (item.Nbelanja * item.Harga)
		Ntotal = item.Ntotal
		fmt.Println("Total Harga minyak yang anda beli : ", item.Ntotal)
		item.Stock = item.Stock - item.Nbelanja
		fmt.Println("Sisa Stock: ", item.Stock)
	}
	return banyak, Ntotal
}
