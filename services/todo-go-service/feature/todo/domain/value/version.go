package value

type Version int

func (v Version) Value() int { return int(v) }
