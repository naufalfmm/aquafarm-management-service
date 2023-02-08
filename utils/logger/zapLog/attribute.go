package zapLog

type (
	attribute struct {
		key   string
		value interface{}
	}
)

func (a attribute) Key() string        { return a.key }
func (a attribute) Value() interface{} { return a.value }

func SetAttribute(key string, val interface{}) attribute {
	return attribute{key, val}
}
