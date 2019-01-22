package face

type Face interface {
	Auth(descriptor interface{}, images []byte) error
}
