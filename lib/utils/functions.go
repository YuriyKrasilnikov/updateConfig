package utils

func Map[A, R any](s []A, f func(A) R) []R {
	result := make([]R, len(s))
	for i := range s {
		result[i] = f(s[i])
	}
	return result
}

func Reduce[A, R any](s []A, f func(R, A) R, init R) R {
	result := init
	for _, v := range s {
		result = f(result, v)
	}
	return result
}

func ForEach[A any](s []A, f func(A)) {
	for _, v := range s {
		f(v)
	}
}
