/*
 *	Site: 	Hacker Rank
 *	Link: 	https://www.hackerrank.com/challenges/counting-valleys
 *  Level:	Easy
 */

package main

func countingValleys(n int32, s string) int32 {

	data := []byte(s)
	level := 0
	valley := false
	valleys := 0

	if data[0] == 'D' {
		valley = true
		level--
	} else {
		level++
	}

	for i := 1; i < len(data); i++ {
		if data[i] == 'D' {
			level--
		}
		if data[i] == 'U' {
			level++
		}

		if level == 0 && data[i] == 'U' {
			if valley {
				valleys++
			}
		}

		if level < 0 && data[i] == 'D' {
			valley = true
		}
		if level > 0 && data[i] == 'U' {
			valley = false
		}
	}

	return int32(valleys)
}

func main() {
	r := countingValleys(1, "UDUUUDUDDD")
	println(r)
}
