package queue

const size = 65535 - 49152 + 1

// Queue is a concurrent safe queue.
type Queue chan uint16

// New return a new instance of Queue.
func New() Queue {
	c := make(chan uint16, size)
	return Queue(c)
}

// Put add an element into queue.
func (q Queue) Put(v uint16) {
	q <- v
}

// Poll remove an element from the head of queue.
// if the queue is empty, it will return false,
// otherwise it will return true.
func (q Queue) Poll() uint16 {
	return <-q
}
