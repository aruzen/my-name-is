package repository

type rowScanner interface {
	Scan(dest ...interface{}) error
}
