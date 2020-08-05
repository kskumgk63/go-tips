package iota

type bloodType int

// blood types
const (
	A bloodType = iota
	B
	O
	AB
)

func (b bloodType) String() string {
	return [...]string{"A", "B", "O", "AB"}[b]
}
