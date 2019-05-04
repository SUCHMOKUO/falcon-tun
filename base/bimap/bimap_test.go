package bimap

import (
	"log"
	"sync"
	"testing"
)

type BiMap64To16 struct {
	key uint64
	value uint16
}

var (
	k1 uint64 = 10000
	v1 uint16 = 10001
	k2 uint64 = 20000
	v2 uint16 = 20001

	k3 uint16 = 30000
	v3 uint64 = 30001
	k4 uint16 = 40000
	v4 uint64 = 40001
)

var biMap = New()

func init() {
	err := biMap.Put(k1, v1)
	if err != nil {
		log.Fatalln(err)
	}
	err = biMap.Put(k2, v2)
	if err != nil {
		log.Fatalln(err)
	}
	err = biMap.Put(k3, v3)
	if err != nil {
		log.Fatalln(err)
	}
	err = biMap.Put(k4, v4)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestBiMap(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(4)
	go func() {
		defer wg.Done()

		if v_1, ok := biMap.Get(k1); !ok {
			t.Errorf("k: %d\n", k1)
		} else {
			log.Printf("k: %d -> v: %d\n", k1, v_1)
		}

		if k_1, ok := biMap.Get(v1); !ok {
			t.Errorf("v: %d\n", v1)
		} else {
			log.Printf("v: %d -> k: %d\n", v1, k_1)
		}
	}()

	go func() {
		defer wg.Done()

		if v_2, ok := biMap.Get(k2); !ok {
			t.Errorf("k: %d\n", k2)
		} else {
			log.Printf("k: %d -> v: %d\n", k2, v_2)
		}

		if k_2, ok := biMap.Get(v2); !ok {
			t.Errorf("v: %d\n", v2)
		} else {
			log.Printf("v: %d -> k: %d\n", v2, k_2)
		}
	}()

	go func() {
		defer wg.Done()

		if v_3, ok := biMap.Get(k3); !ok {
			t.Errorf("k: %d\n", k3)
		} else {
			log.Printf("k: %d -> v: %d\n", k3, v_3)
		}

		if k_3, ok := biMap.Get(v3); !ok {
			t.Errorf("v: %d\n", v3)
		} else {
			log.Printf("v: %d -> k: %d\n", v3, k_3)
		}
	}()

	go func() {
		defer wg.Done()

		if v_4, ok := biMap.Get(k4); !ok {
			t.Errorf("k: %d\n", k4)
		} else {
			log.Printf("k: %d -> v: %d\n", k4, v_4)
		}

		if k_4, ok := biMap.Get(v4); !ok {
			t.Errorf("v: %d\n", v4)
		} else {
			log.Printf("v: %d -> k: %d\n", v4, k_4)
		}
	}()

	wg.Wait()
}

func TestBiMap_Del(t *testing.T) {
	if v_1, ok := biMap.Get(k1); ok {
		log.Println("get v1:", v_1)
	} else {
		t.Error("should get v1")
	}

	biMap.Del(k1)

	if v_1, ok := biMap.Get(k1); ok {
		t.Error("should not get v1:", v_1)
	}

	if k_1, ok := biMap.Get(v1); ok {
		t.Error("should not get k1:", k_1)
	}
}