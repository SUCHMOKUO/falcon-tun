package queue

import (
	"log"
	"testing"
)

var (
	q = New()
	data = []uint16{1,2,3,4,5,6,7,8,9}
)

func TestQueue(t *testing.T) {
	for _, n := range data {
		q.Put(n)
	}
	for _ = range data {
		log.Println(q.Poll())
	}
}
