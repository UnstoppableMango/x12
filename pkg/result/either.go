package result

type Either[L, R any] struct {
	left  *L
	right *R
}

func left[R, L any](value L) Either[L, R] {
	return Either[L, R]{left: &value}
}

func right[L, R any](value R) Either[L, R] {
	return Either[L, R]{right: &value}
}

func mapLeft[LA, LB, R any](e Either[LA, R], fn func(LA) LB) Either[LB, R] {
	if e.left != nil {
		return left[R](fn(*e.left))
	}
	return right[LB](*e.right)
}

func bindLeft[LA, LB, R any](e Either[LA, R], fn func(LA) Either[LB, R]) Either[LB, R] {
	if e.left != nil {
		return fn(*e.left)
	}
	return right[LB](*e.right)
}

type Result[T any] = Either[T, error]

func Ok[T any](value T) Result[T] {
	return left[error](value)
}

func Err[T any](err error) Result[T] {
	return right[T](err)
}

func Map[A, B any](res Result[A], fn func(A) B) Result[B] {
	return mapLeft(res, fn)
}

func Bind[A, B any](res Result[A], fn func(A) Result[B]) Result[B] {
	return bindLeft(res, fn)
}

type Maybe[T any] struct {
	value *T
}
