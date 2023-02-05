package token

type Claims interface {
	Valid() error
	SetExp(exp int64) Claims
}
