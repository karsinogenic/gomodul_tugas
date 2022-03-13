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

func Troli(troli_item []string, troli_bnyk []int, total int) bool {
	var choice string
	var ulang3 bool
	fmt.Print("isi troli anda : ")
	for i := 0; i < len(troli_bnyk); i++ {
		fmt.Printf("%s :%d buah ", troli_item[i], troli_bnyk[i])
	}
ulang4:
	fmt.Println("Ingin belanja lagi (y/n) ?")
	fmt.Scanln(&choice)
	if choice == "y" {
		ulang3 = true
	} else if choice == "n" {
		fmt.Printf("Total belanja anda adalah sebesar : %d", total)
	} else {
		fmt.Println("input salah")
		goto ulang4
	}
	return ulang3
}

// var minyak = item{"minyak", 10, 30000, 0, 0}

func (item Item) Belanja1() (int, int, int) {
	var banyak int
	// var Ntotal int
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
		item.Stock = item.Stock - item.Nbelanja
		//fmt.Println("Total minyak yang anda beli : ", item.Nbelanja)
		item.Ntotal = item.Ntotal + (item.Nbelanja * item.Harga)
		//Ntotal = item.Ntotal
		//fmt.Println("Total Harga minyak yang anda beli : ", item.Ntotal)
		//fmt.Println("Sisa Stock: ", item.Stock)
	}
	return item.Stock, item.Nbelanja, item.Ntotal
}
