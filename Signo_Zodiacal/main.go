package main

import "fmt"

func main() {
	var mes int
	var dia int
	var valmes int = 1
	var valdia int = 1
	for valmes <= 12 && valdia <= 31 {
		fmt.Scan(&dia)
		fmt.Scan(&mes)
		switch {
		case mes == 1:
			if dia < 21 {
				fmt.Println("Capricornio")
			} else if dia <= 31 {
				fmt.Println("Acuario")
			}
		case mes == 2:
			if dia < 21 {
				fmt.Println("Acuario")
			} else if dia <= 29 {
				fmt.Println("Piscis")
			}
		case mes == 3:
			if dia < 21 {
				fmt.Println("Piscis")
			} else if dia <= 31 {
				fmt.Println("Aries")
			}
		case mes == 4:
			if dia < 21 {
				fmt.Println("Aries")
			} else if dia <= 31 {
				fmt.Println("Tauro")
			}
		case mes == 5:
			if dia < 22 {
				fmt.Println("Tauro")
			} else if dia <= 31 {
				fmt.Println("Geminis")
			}
		case mes == 6:
			if dia < 22 {
				fmt.Println("Geminis")
			} else if dia <= 31 {
				fmt.Println("Cancer")
			}
		case mes == 7:
			if dia < 23 {
				fmt.Println("Cancer")
			} else if dia <= 31 {
				fmt.Println("Leo")
			}
		case mes == 8:
			if dia < 23 {
				fmt.Println("Leo")
			} else if dia <= 31 {
				fmt.Println("virgo")
			}
		case mes == 9:
			if dia < 24 {
				fmt.Println("virgo")
			} else if dia <= 31 {
				fmt.Println("Libra")
			}
		case mes == 10:
			if dia < 25 {
				fmt.Println("Libra")
			} else if dia <= 31 {
				fmt.Println("Escorpio")
			}
		case mes == 11:
			if dia < 23 {
				fmt.Println("Escorpio")
			} else if dia <= 31 {
				fmt.Println("Sagitario")
			}
		case mes == 12:
			if dia < 22 {
				fmt.Println("Sagitario")
			} else if dia <= 31 {
				fmt.Println("Capricornio")
			}
		default:

		}
		valdia = dia
		valmes = mes
	}

}
