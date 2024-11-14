package value

type TodoNote string

func (tn TodoNote) String() string { return string(tn) }
