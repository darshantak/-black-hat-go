package main

import (
	"encoding/json"
	"os"
)

func (kv *KeyValueStore) SaveToFile(filename string) error {
	kv.mutex.Lock()

	defer kv.mutex.Unlock()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(kv.store)
}

func (kv *KeyValueStore) LoadFromFile(filename string) error {
	kv.mutex.Lock()

	defer kv.mutex.Unlock()

	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	return decoder.Decode(&kv.store)
}
