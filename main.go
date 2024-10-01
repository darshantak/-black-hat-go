package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type KeyValueStore struct {
	store map[string]string
	mutex sync.RWMutex
}

func (kv *KeyValueStore) SetHandler(w http.ResponseWriter, r *http.Request) {
	kv.mutex.Lock()
	defer kv.mutex.Unlock()

	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	if key == "" || value == "" {
		http.Error(w, "Missing k or v", http.StatusBadRequest)
		return
	}

	kv.store[key] = value
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Set key : %s and value : %s", key, value)
}

func (kv *KeyValueStore) GetHandler(w http.ResponseWriter, r *http.Request) {
	kv.mutex.Lock()
	defer kv.mutex.Unlock()

	key := r.URL.Query().Get("key")
	value, exists := kv.store[key]
	if !exists {
		http.Error(w, "Key Not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Value for key : %s is %s", key, value)
}

func (kv *KeyValueStore) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	kv.mutex.Lock()
	defer kv.mutex.Unlock()

	key := r.URL.Query().Get("key")
	_, exists := kv.store[key]
	if !exists {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	delete(kv.store, key)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted key: %s", key)
}

func main() {
	kv := KeyValueStore{store: make(map[string]string)}

	http.HandleFunc("/set", kv.SetHandler)
	http.HandleFunc("/get", kv.GetHandler)
	http.HandleFunc("/delete", kv.DeleteHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

