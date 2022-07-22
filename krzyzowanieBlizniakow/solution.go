package main

import (
	"fmt"
)

var (
	numbers = map[int]bool{101: true, 103: true,
		107: true, 109: true,
		137: true, 139: true,
		149: true, 151: true,
		179: true, 181: true,
		191: true, 193: true,
		197: true, 199: true,
		227: true, 229: true,
		239: true, 241: true,
		269: true, 271: true,
		281: true, 283: true,
		311: true, 313: true,
		347: true, 349: true,
		419: true, 421: true,
		431: true, 433: true,
		461: true, 463: true,
		521: true, 523: true,
		569: true, 571: true,
		599: true, 601: true,
		617: true, 619: true,
		641: true, 643: true,
		659: true, 661: true,
		809: true, 811: true,
		821: true, 823: true,
		827: true, 829: true,
		857: true, 859: true,
		881: true, 883: true}
)

func old_really_bad() {
	count := 1
	for i := 107000001; i < 883999999; i += 2 {
		a3 := i % 1000
		if numbers[a3] {
			a2 := (i / 1000) % 1000
			if numbers[a2] {
				a1 := (i / 1000000) % 1000
				if numbers[a1] {
					b1 := (a1/100)*100 + 10*(a2/100) + a3/100
					if numbers[b1] {
						b2 := 100*((a1/10)%10) + 10*((a2/10)%10) + (a3/10)%10
						if numbers[b2] {
							b3 := 100*(a1%10) + 10*(a2%10) + a3%10
							if numbers[b3] {
								if a1 != a2 && a1 != a3 && a1 != b1 && a1 != b2 && a1 != b3 &&
									a2 != a3 && a2 != b1 && a2 != b2 && a2 != b3 &&
									a3 != b1 && a3 != b2 && a3 != b3 &&
									b1 != b2 && b1 != b3 &&
									b2 != b3 {
									fmt.Printf("--- solution %d ---\n", count)
									fmt.Printf("%d\n%d\n%d\n\n", a1, a2, a3)
									count++
								}
							}
						}
					}
				}
			}
		}
	}
}

func new_better_solution() {
	numbers_i := [54]int{101, 103, 107, 109, 137, 139, 149, 151, 179, 181, 191, 193, 197, 199, 227, 229, 239, 241, 269, 271, 281, 283, 311, 313, 347, 349, 419, 421, 431, 433, 461, 463, 521, 523, 569, 571, 599, 601, 617, 619, 641, 643, 659, 661, 809, 811, 821, 823, 827, 829, 857, 859, 881, 883}

	count := 1
	for i := 0; i < 54; i++ {
		for j := 0; j < 54; j++ {
			if j != i {
				for k := 0; k < 54; k++ {
					if k != i && k != j {
						a1 := numbers_i[i]
						a2 := numbers_i[j]
						a3 := numbers_i[k]
						b1 := (a1/100)*100 + 10*(a2/100) + a3/100
						b2 := 100*((a1/10)%10) + 10*((a2/10)%10) + (a3/10)%10
						b3 := 100*(a1%10) + 10*(a2%10) + a3%10
						if numbers[b1] && numbers[b2] && numbers[b3] {
							if a1 != b1 && a1 != b2 && a1 != b3 &&
								a2 != b1 && a2 != b2 && a2 != b3 &&
								a3 != b1 && a3 != b2 && a3 != b3 &&
								b1 != b2 && b1 != b3 && b2 != b3 {
								fmt.Printf("--- solution %d ---\n", count)
								fmt.Printf("%d\n%d\n%d\n\n", a1, a2, a3)
								count++
							}
						}
					}
				}
			}
		}
	}
}

func main() {
	fmt.Println("Krzyżowanie bliźniaków")
	fmt.Println("https://penszko.blog.polityka.pl/2022/07/16/krzyzowanie-blizniakow/")
	// ~4.84 sec
	//old_really_bad()

	// 0.01s
	new_better_solution()
}
