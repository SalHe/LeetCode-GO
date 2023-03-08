package main

func main() {
	x := 50
	for i := 0; i < x; i++ {
		if i == 9 {
			x = 9
		}
		println(i)
	}
}
