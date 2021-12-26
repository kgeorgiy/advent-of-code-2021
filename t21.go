package main

func solve21one(player1 int, player2 int) int {
	dice := 0
	score1 := 0
	score2 := 0
	for {
		player1 = (player1+roll(&dice)+roll(&dice)+roll(&dice)-1)%10 + 1
		score1 += player1
		//println("1", player1)
		if score1 >= 1000 {
			println(dice, score1, score2)
			return score2 * dice
		}
		player2 = (player2+roll(&dice)+roll(&dice)+roll(&dice)-1)%10 + 1
		score2 += player2
		//println("2", player1)
		if score2 >= 1000 {
			println(dice, score1, score2)
			return score1 * dice
		}
	}
}

func roll(dice *int) int {
	*dice++
	return (*dice-1)%100 + 1
}

func solve21two(player1 int, player2 int) int {
	const maxScore = 21
	const positions = 10
	win := [maxScore + 1][maxScore + 1][positions + 1][positions + 1]int{}
	los := [maxScore + 1][maxScore + 1][positions + 1][positions + 1]int{}
	for ts := 1; ts <= maxScore*2; ts++ {
		for s1 := 1; s1 <= intMin(ts, maxScore); s1++ {
			s2 := ts - s1
			if s2 > maxScore {
				continue
			}
			//println(s1, s2)
			for p1 := 1; p1 <= positions; p1++ {
				for p2 := 1; p2 <= positions; p2++ {
					totalWin := 0
					totalLos := 0
					for d1 := 1; d1 <= 3; d1++ {
						for d2 := 1; d2 <= 3; d2++ {
							for d3 := 1; d3 <= 3; d3++ {
								np1 := (p1+d1+d2+d3-1)%10 + 1
								ns1 := s1 - np1
								if ns1 <= 0 {
									totalWin++
								} else {
									//println(s2, ns1, p2, np1)
									totalWin += los[s2][ns1][p2][np1]
									totalLos += win[s2][ns1][p2][np1]
								}
							}
						}
					}
					win[s1][s2][p1][p2] = totalWin
					los[s1][s2][p1][p2] = totalLos
				}
			}
		}
	}
	println(win[maxScore][maxScore][player1][player2], los[maxScore][maxScore][player1][player2])
	return intMax(win[maxScore][maxScore][player1][player2], los[maxScore][maxScore][player1][player2])
}

func main() {
	println(solve21one(4, 8))
	println(solve21one(9, 3))
	println(solve21two(4, 8))
	println(solve21two(9, 3))
}
