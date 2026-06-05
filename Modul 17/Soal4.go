package main

import "fmt"

func main() {

	var jumlah float64
	var si, sip1 float64
	var i int

	for i = 0; ; i++ {

		if i%2 == 0 {
			si = 1.0 / float64(2*i+1)
		} else {
			si = -1.0 / float64(2*i+1)
		}

		jumlah += si

		if (i+1)%2 == 0 {
			sip1 = 1.0 / float64(2*(i+1)+1)
		} else {
			sip1 = -1.0 / float64(2*(i+1)+1)
		}

		selisih := si - sip1

		if selisih < 0 {
			selisih = -selisih
		}

		if selisih <= 0.00001 {

			fmt.Println("Hasil PI :", 4*jumlah)
			fmt.Println("Pada i ke:", i+1)

			break
		}
	}
}