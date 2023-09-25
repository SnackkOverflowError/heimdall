package memtable

type MemTable interface {
	Insert(key, value string) error
	Remove(key string) error
	Get(key string) (string, error)
}
