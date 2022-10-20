package learn

func Struct() {
	type Saiyan struct {
		Name  string
		Power int
	}

	goku := Saiyan{
		Name:  "Goku",
		Power: 9000,
	}

	println(goku.Name)
}
