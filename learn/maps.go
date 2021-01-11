package learn

import "fmt"

func Maps() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	m["k3"] = 13

	fmt.Println("map:", m)

	fmt.Println("len:", len(m))

	_, prs := m["k4"]
	fmt.Println("prs:", prs)
}
