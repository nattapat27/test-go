package loop

func While() {
	i := 0
	for i < 10 {
		print(i)
		if i == 9 {
			println()
		} else {
			print(",")
		}

		i++
	}
}
