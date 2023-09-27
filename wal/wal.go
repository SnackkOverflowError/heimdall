package wal

type WAL struct {
	filename string
}

type walAction int

const (
	WRITE walAction = iota
	FLUSH
)

type walElement struct {
	action walAction
	key    string
	value  string
}

func (w *WAL) RunWALService(actionQueue chan walElement) {

	// open the wal file

}
