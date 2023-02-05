package orm

type (
	preloadOpt struct {
		query string
		args  []interface{}
	}

	preloadOpts []preloadOpt
)

func SetPreload(query string, args ...interface{}) preloadOpt {
	return preloadOpt{
		query: query,
		args:  args,
	}
}
