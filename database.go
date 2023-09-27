package main

import (
	"errors"
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

var databases map[string]*Database

func (db *Database) Insert(key, value string) error {
	return errors.ErrUnsupported
}

func (db *Database) Get(key string) (string, error) {
	return "", errors.ErrUnsupported
}
