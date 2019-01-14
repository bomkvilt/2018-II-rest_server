package models

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//go:generate easyjson .

// easyjson:json
type Error struct {
	Message string `json:"message"`
}

// easyjson:json
type Status struct {
	Forum  int `json:"forum"`
	Post   int `json:"post"`
	Thread int `json:"thread"`
	User   int `json:"user"`
}

// go tool pprof pprof_1.exe cpu_out.txt
// go tool pprof -http=:8081 pprof_1.exe cpu_out.txt
