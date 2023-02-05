package jwt

import "github.com/naufalfmm/aquafarm-management-service/utils/token"

type JWT struct {
	Encoder token.Encoder
	Decoder token.Decoder
}

func NewJWT(publicKey, alg string) (JWT, error) {
	encoder := NewEncoder(publicKey, alg)
	decoder := NewDecoder(publicKey)

	return JWT{
		Encoder: encoder,
		Decoder: decoder,
	}, nil
}
