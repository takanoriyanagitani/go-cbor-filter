package util

func ErrFn[T, U any](f func(T) U) func(T) (U, error) {
	return func(t T) (U, error) {
		return f(t), nil
	}
}
