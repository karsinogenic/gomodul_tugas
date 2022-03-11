package modulasi

//cobalagi

import "fmt"

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
