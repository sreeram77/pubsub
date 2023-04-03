package storage

type Cache interface {
	Get(string) []any
	Set(string, any)
}
