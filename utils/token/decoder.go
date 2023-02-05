package token

type Decoder interface {
	DecodeToken(token string) (Claims, error)
}
