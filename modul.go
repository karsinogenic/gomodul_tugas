package modulasi

//cobalagi

import "fmt"

type item struct {
	nama     string
	stock    int
	harga    int
	nBelanja int
	nTotal   int
}

func Kasir(uang, bayar int) int {
	var kembali int
	kembali = uang - bayar

	return kembali
}

func Troli(harga, banyak int) (int, bool) {
	var bayar int
	var choice bool
	var choice1 string
	bayar = (harga * banyak) + bayar
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

func (minyak item) Belanja1() (int, int) {
	var banyak int
	var ntotal int
ulang1:
	fmt.Printf("Anda memilih minyak, berapa banyak yang ingin anda beli : ")
	//fmt.Printf("1.%s stock:%d harga:%s || 2.%s stock:%d harga:%s\n//////////////////////////////////////////////////////\n", minyak.nama, minyak.stock, minyak.harga, terigu.nama, terigu.stock, terigu.harga)
	//fmt.Print("Pilihan anda (1-2) : ")
	fmt.Scanln(&banyak)
	if banyak > minyak.stock {
		fmt.Println("Stock tidak mencukupi, jumlah stock : ", minyak.stock)
		goto ulang1
	} else {
		minyak.nBelanja = minyak.nBelanja + banyak
		fmt.Println("Total minyak yang anda beli : ", minyak.nBelanja)
		minyak.nTotal = minyak.nTotal + (minyak.nBelanja * minyak.harga)
		ntotal = minyak.nTotal
		fmt.Println("Total harga minyak yang anda beli : ", minyak.nTotal)
		minyak.stock = minyak.stock - minyak.nBelanja
		fmt.Println("Sisa stock: ", minyak.stock)
	}
	return banyak, ntotal
}
