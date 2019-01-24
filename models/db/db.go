package db

type db interface {
	Init() error
	enroll(descriptor []byte) (id string, err error)
	enrollBy(id string, descriptor []byte) error
	query(id string) (descriptor []byte, err error)
}
