package db

type db interface {
	Init() error
	set(key string, value []byte) error
	get(key string) ([]byte, error)
}
