package main

import "fmt"

func main() {
	var ch int
	runProgram := true

	for {
		fmt.Println("")
		fmt.Println("*--------------------------------------------------------------------*")
		fmt.Println("1.Penjumlahan  2.Pengurangan  3.Perkalian  4.Transpose  8.Keluar")
		fmt.Println("*--------------------------------------------------------------------*")
		fmt.Print("Input menu : ")
		
		fmt.Scanln(&ch)

		switch ch {
			case 1:
				m1, m2 := scanTwoMatrix()
				mOutput, err := m1.Plus(m2)

				if err != nil {
					fmt.Println(err)
					break
				}

				fmt.Println("Hasil Matriks 1 + Matriks 2 = ")
				fmt.Println(mOutput.String())
				break

			case 2:
				m1, m2 := scanTwoMatrix()
				mOutput, err := m1.Minus(m2)

				if err != nil {
					fmt.Println(err)
					break
				}
				fmt.Println("Hasil Matriks 1 - Matriks 2 = ")
				fmt.Println(mOutput.String())
				break

			case 3:
				var scalar float64
				fmt.Print("Input Skalar = ")
				fmt.Scanln(&scalar)
				
				m := ScanMatrix()

				mOutput := m.ScalarMultiply(scalar)

				fmt.Println("Hasil Skalar", scalar, "* Matriks = ")
				fmt.Println(mOutput.String())
				break

			case 4:
				m := ScanMatrix()

				mOutput := m.Transpose()

				fmt.Println("Hasil Transpose Matriks = ")
				fmt.Println(mOutput.String())
				break

			default :
				runProgram = false
				break
		}

		// m6 := m.ScalarMultiply(5)
		// fmt.Println("|m| * |5| = ", m6)

		// fmt.Println("m.Transpose() = ", m.Transpose())
		

		if !runProgram {
			break
		}
	}

	return	
}