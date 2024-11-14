package value

type TodoStatus bool

const (
	TodoStatusIncomplete TodoStatus = false
	TodoStatusComplete   TodoStatus = true
)

func (t TodoStatus) Value() bool { return bool(t) }
