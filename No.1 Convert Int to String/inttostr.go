package main

import (
	"fmt"
)

var NumberToWord = map[int]string{
	1:  "satu",
	2:  "dua",
	3:  "tiga",
	4:  "empat",
	5:  "lima",
	6:  "enam",
	7:  "tujuh",
	8:  "delapan",
	9:  "sembilan",
	10: "sepuluh",
	11: "sebelas",
	12: "dua belas",
	13: "tiga belas",
	14: "empat belas",
	15: "lima belas",
	16: "enam belas",
	17: "tujuh belas",
	18: "delapan belas",
	19: "sembilan belas",
	20: "dua puluh",
	30: "tiga puluh",
	40: "empat puluh",
	50: "lima puluh",
	60: "enam puluh",
	70: "tujuh puluh",
	80: "Delapan puluh",
	90: "sembilan puluh",
}

func convert1to99(n int) (w string) {
	if n < 20 {
		w = NumberToWord[n]
		return
	}

	r := n % 10
	if r == 0 {
		w = NumberToWord[n]
	} else {
		w = NumberToWord[n-r] + " " + NumberToWord[r]
	}
	return
}

func convert100to999(n int) (w string) {
	q := n / 100
	r := n % 100
	w = NumberToWord[q]
	if n == 100 {
		w = "seratus"
		return
	}
	if r == 0 {
		return
	} else {
		w = w + " " + convert1to99(r)
	}
	return
}

func convert1000to9999(n int) (w string) {
	q := n / 1000
	r := n % 1000
	w = NumberToWord[q] + " ribu"
	if n == 1000 {
		w = "seribu"
		return
	}
	if r == 0 {
		return
	} else {
		w = w + " " + convert100to999(r)
	}
	return
}

func convert10000to99999(n int) (w string) {
	q := n / 10000
	r := n % 10000
	w = NumberToWord[q]
	if n == 10000 {
		w = "sepuluh ribu"
		return
	}
	if r == 0 {
		return
	} else {
		w = w + " " + convert1000to9999(r)
	}
	return
}

func convert100000to999999(n int) (w string) {
	q := n / 10000
	r := n % 10000
	w = NumberToWord[q] + " puluh"
	if n == 100000 {
		w = "seratus ribu"
		return
	}
	if r == 0 {
		return
	} else {
		w = w + " " + convert10000to99999(r)
	}
	return
}
func convert1000000to9999999(n int) (w string) {
	q := n / 100000
	r := n % 100000
	w = NumberToWord[q] + " ratus ribu "
	if n == 1000000 {
		w = "satu juta"
		return
	}
	if r == 0 {
		return
	} else {
		w = w + " " + convert100000to999999(r)
	}
	return
}

func Convert1to1000000(n int) (w string) {
	if n > 1000000 || n < 1 {
		panic("func Convert1to1000000000: n > 1000000 or n < 1")
	}
	if n < 100 {
		w = convert1to99(n)
		return
	}
	if n < 1000 {
		w = convert100to999(n)
		return
	}
	if n < 10000 {
		w = convert1000to9999(n)
		return
	}
	if n < 100000 {
		w = convert10000to99999(n)
		return
	}
	w = convert1000000to9999999(n)
	return
}

func main() {
	fmt.Println(Convert1to1000000(12))
	fmt.Println(Convert1to1000000(2012))
	fmt.Println(Convert1to1000000(999999))

}
