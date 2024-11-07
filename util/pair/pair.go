package pair

type Pair[L, R any] struct {
	Left  L
	Right R
}

func New[L, R any](right R, left L) Pair[L, R] {
	return Pair[L, R]{
		Left:  left,
		Right: right,
	}
}
