package main

import (
	"heimdall/memtable"
	"heimdall/memtable/simpletable"
)

type Database struct {
	memtable memtable.MemTable
}

func NewDatabase() *Database {
	return &Database{
		memtable: &simpletable.SimpleTable{},
	}
}
