package main

import (
	"fmt"
	"sync"
)

type kvStore struct {
	store map[string]string
	mutex sync.RWMutex
}

func (kv *kvStore) Set(key, value string) {
	kv.mutex.RLock()
	defer kv.mutex.RUnlock()

	kv.store[key] = value
}

func (kv *kvStore) Get(key string) string {
	kv.mutex.Lock()
	defer kv.mutex.Unlock()

	return kv.store[key]
}

func (kv *kvStore) Delete(key string) {
	kv.mutex.Lock()
	defer kv.mutex.Unlock()
	delete(kv.store, key)
}

func runRW() {
	kv := kvStore{store: make(map[string]string)}
	kv.Set("key", "one")

	value := kv.Get("key")
	fmt.Printf("Value: %s", value)
}
