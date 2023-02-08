package token

type Decoder interface {
	DecodeToken(t string) (Claims, error)
}
