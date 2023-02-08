package logger

type (
	Attribute interface {
		Key() string
		Value() interface{}
	}
)
