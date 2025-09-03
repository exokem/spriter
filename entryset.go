package main

import "sync"

type Entry struct {
	name string
	path string
}

type EntrySet struct {
	index map[string]Entry
	mu    sync.Mutex
}

func NewEntrySet() *EntrySet {
	set := &EntrySet{
		make(map[string]Entry),
		sync.Mutex{},
	}

	return set
}

func (set *EntrySet) Store(name string, path string) *Entry {
	set.mu.Lock()
	defer set.mu.Unlock()

	entry := &Entry{name, path}
	set.index[name] = *entry

	return entry
}
