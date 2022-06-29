package base

import (
	"time"
)

type Entity struct {
	Id      int
	created time.Time
}

func New(id int) Entity {
	return Entity{Id: id, created: time.Now()}
}
