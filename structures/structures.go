package structures

type Todo struct {
	Id   int
	Item string
	Done bool
}

type Data struct {
	Todos []Todo
}
