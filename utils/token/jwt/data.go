package jwt

type UserLogin struct {
	ID         uint64 `json:"id"`
	Email      string `json:"email"`
	IsVerified bool   `json:"isVerified"`
}

func (l *UserLogin) CreatedBy() string {
	return l.Email
}
