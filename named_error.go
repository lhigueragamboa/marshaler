package marshaler

type NamedError interface {
	error
	Name() string
}
