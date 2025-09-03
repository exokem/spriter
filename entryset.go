package main

import "sync"

type Entry struct {
	Name string
	Path string
}

type EntrySet struct {
	Index map[string]Entry
	mu    sync.Mutex
}

func NewEntrySet() *EntrySet {
	set := &EntrySet{
		make(map[string]Entry),
		sync.Mutex{},
	}

	return set
}

func (set *EntrySet) Store(name string, path string) Entry {
	set.mu.Lock()
	defer set.mu.Unlock()

	entry := Entry{name, path}
	set.Index[name] = entry
	return entry
}

func (set *EntrySet) Absorb(other *EntrySet) *EntrySet {

	for key, value := range other.Index {
		set.Index[key] = value
	}

	return set
}
