package models

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Error struct {
	Message string
}
