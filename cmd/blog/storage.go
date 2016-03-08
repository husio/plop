package main

import (
	"sync"
	"time"

	"github.com/husio/plop/randkey"
	"golang.org/x/net/context"
)

type Database struct {
	mu  sync.Mutex
	mem map[string]Entry
}

func NewDatabase() *Database {
	return &Database{mem: make(map[string]Entry)}
}

type Entry struct {
	ID      string
	UserID  string
	Title   string
	Content string
	Created time.Time
}

func (db *Database) Create(e Entry) Entry {
	e.ID = randkey.New()
	e.Created = time.Now()

	db.mu.Lock()
	defer db.mu.Unlock()

	db.mem[e.ID] = e
	return e
}

func (db *Database) List() []Entry {
	db.mu.Lock()
	defer db.mu.Unlock()

	res := make([]Entry, 0, len(db.mem))
	for _, e := range db.mem {
		res = append(res, e)
	}
	return res
}

func WithDB(ctx context.Context, db *Database) context.Context {
	return context.WithValue(ctx, "db", db)
}

func DB(ctx context.Context) *Database {
	db := ctx.Value("db")
	if db == nil {
		panic("database not present in context")
	}
	return db.(*Database)
}
