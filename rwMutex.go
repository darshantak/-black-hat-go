package main

import (
	"fmt"
	"sync"
	"time"
)

type KeyValue struct {
	Value   string
	Expires time.Time
}

type kvStore struct {
	store map[string]KeyValue
	mutex sync.RWMutex
}

func (kv *kvStore) Set(key string, value KeyValue) {
	kv.mutex.RLock()
	defer kv.mutex.RUnlock()

	kv.store[key] = value
}

func (kv *kvStore) Get(key string) string {
	kv.mutex.RLock()
	defer kv.mutex.RUnlock()
	keyValue, exists := kv.store[key]
	if exists {
		if time.Now().After(keyValue.Expires) {
			kv.mutex.RUnlock()
			kv.Delete(key)
			return ""
		}
		return keyValue.Value
	}
	return ""
}

func (kv *kvStore) Delete(key string) {
	kv.mutex.Lock()
	defer kv.mutex.Unlock()
	delete(kv.store, key)
}

func runRW() {
	kv := kvStore{store: make(map[string]KeyValue)}
	kv.Set("session", KeyValue{Value: "abceded", Expires: time.Now().Add(5 * time.Second)})

	value := kv.Get("session")
	fmt.Printf("Value: %s", value)

	time.Sleep(7 * time.Second)
	value = kv.Get("session")

	if value == "" {
		fmt.Printf("Key is expired")
	} else {
		fmt.Printf("Value: %s", value)
	}
}
