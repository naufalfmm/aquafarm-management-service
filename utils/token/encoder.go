package token

type Encoder interface {
	EncodeToken(claims Claims) (string, error)
}
