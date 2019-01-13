package models

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Error struct {
	Message string `json:"message"`
}

type Status struct {
	Forum  int `json:"forum"`
	Post   int `json:"post"`
	Thread int `json:"thread"`
	User   int `json:"user"`
}
