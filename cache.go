package main

import (
	"encoding/json"
	"os"
	"path"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func initData() {
	if _, err := os.Stat("data/cache"); os.IsNotExist(err) {
		check(os.MkdirAll("data/cache", 0755))
	}
}

func save(file string, v any) {
	initData()

	data, err := json.Marshal(v)
	check(err)

	check(os.WriteFile(path.Join("data/cache", file), data, 0644))
}

func load(file string, v any, init any) any {
	initData()

	file = path.Join("data/cache", file)

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return init
	}

	bytes, err := os.ReadFile(file)
	check(err)

	json.Unmarshal(bytes, v)

	return v
}

func loadMeta() *Metadata {
	meta := Metadata{}
	load("meta.json", &meta, meta)
	return &meta
}

func saveEntries(file string, set *EntrySet) {
	save(file, set.Index)
}

func loadEntries(file string) *EntrySet {
	initData()

	file = path.Join("data/cache", file)

	// If the requested set is not cached, create an empty one
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return NewEntrySet()
	}

	bytes, err := os.ReadFile(file)
	check(err)

	set := &EntrySet{}

	json.Unmarshal(bytes, &set.Index)

	return set
}
