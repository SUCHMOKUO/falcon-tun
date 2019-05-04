package bimap

import (
	"errors"
	"sync"
)

var (
	errorType = errors.New("invalid type for BiMap")
)

// BiMap offer a bidirectional mapping
// between uint16 and uint64.
type BiMap struct {
	lock sync.Mutex
	uint64ToUint16 map[uint64]uint16
	uint16ToUint64 map[uint16]uint64
}

// New return a new instance of BiMap.
func New() *BiMap {
	biMap := new(BiMap)
	biMap.uint16ToUint64 = make(map[uint16]uint64, 128)
	biMap.uint64ToUint16 = make(map[uint64]uint16, 128)
	return biMap
}

// Put create a mapping of values offered.
// will return error if type mismatched.
func (bm *BiMap) Put(key, value interface{}) error {
	bm.lock.Lock()
	defer bm.lock.Unlock()

	switch k := key.(type) {
	case uint64:
		// key is uint64.
		if v, ok := value.(uint16); ok {
			bm.uint64ToUint16[k] = v
			bm.uint16ToUint64[v] = k
		} else {
			return errorType
		}
	case uint16:
		// key is uint16.
		if v, ok := value.(uint64); ok {
			bm.uint16ToUint64[k] = v
			bm.uint64ToUint16[v] = k
		} else {
			return errorType
		}
	default:
		return errorType
	}
	return nil
}

// Get return the value of the key.
// if key does not exists or the type of key is wrong,
// it will return false.
func (bm *BiMap) Get(key interface{}) (interface{}, bool) {
	bm.lock.Lock()
	defer bm.lock.Unlock()

	var v interface{}
	var ok bool

	switch k := key.(type) {
	case uint64:
		// key is uint64.
		v, ok = bm.uint64ToUint16[k]
	case uint16:
		// key is uint16.
		v, ok = bm.uint16ToUint64[k]
	}

	return v, ok
}

// Del will delete a bidirectional mapping of key offered.
func (bm *BiMap) Del(key interface{}) {
	v, ok := bm.Get(key)
	if !ok {
		return
	}

	bm.lock.Lock()
	defer bm.lock.Unlock()

	switch k := key.(type) {
	case uint16:
		// key is uint16.
		delete(bm.uint16ToUint64, k)
		delete(bm.uint64ToUint16, v.(uint64))
	case uint64:
		// key is uint64.
		delete(bm.uint64ToUint16, k)
		delete(bm.uint16ToUint64, v.(uint16))
	}
}