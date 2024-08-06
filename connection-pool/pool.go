package connectionpool

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type conn struct {
	DB *sql.DB
}

type CPool interface {
	Get() *conn
	Put(*conn)
	Close() error
}

type pool struct {
	channel chan *conn
	cap     int
}

func (p *pool) Put(c *conn) {
	p.channel <- c
}

func (p *pool) Get() *conn {
	return <-p.channel
}

func (p *pool) Close() error {
	close(p.channel)

	for conn := range p.channel {
		err := conn.DB.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func New(cap int) *pool {
	channel := make(chan *conn, cap)
	for i := 0; i < cap; i++ {
		conn := newConn()
		channel <- conn
	}

	return &pool{
		channel: channel,
		cap:     cap,
	}
}

func newConn() *conn {
	db, err := sql.Open("mysql", "root:example@(localhost:3307)/")
	if err != nil {
		panic(fmt.Sprintf("could not create connection, err: %+v", err))
	}
	return &conn{
		DB: db,
	}
}
