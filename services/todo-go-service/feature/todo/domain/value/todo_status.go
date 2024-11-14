package value

type TodoStatus bool

const (
	TodoStatusIncomplete TodoStatus = false
	TodoStatusComplete   TodoStatus = true
)
