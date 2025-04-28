package main
import "fmt"

const NMAX int = 100
func main(){
	var A[NMAX]int
	var n, i, masukan, jumlah, j  int
	var rerata float64
	fmt.Scan(&n)
	for i =0; i<n;i++{
		fmt.Scan(&masukan)
		if masukan >= 0{
			A[i] = masukan
			jumlah += A[i]
			j++
		}
	}
	rerata = float64(jumlah)/float64(j)
	fmt.Print(rerata)
	
}


