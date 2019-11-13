// Package bst provides a tread-safe CRUD wrapper for a Binary Search Tree structure
package bst

import (
	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"sync"
)

type bst struct {
	tree *rbt.Tree
	mx   sync.RWMutex
}

func New(values []int) *bst {
	t := &bst{
		tree: rbt.NewWithIntComparator(),
		mx:   sync.RWMutex{},
	}

	for _, value := range values {
		t.tree.Put(value, struct{}{})
	}

	return t
}

func (t *bst) Put(value int) {
	t.mx.Lock()
	defer t.mx.Unlock()

	t.tree.Put(value, struct{}{})
}

func (t *bst) Del(value int) {
	t.mx.Lock()
	defer t.mx.Unlock()

	t.tree.Remove(value)
}

func (t *bst) Has(value int) bool {
	t.mx.RLock()
	defer t.mx.RUnlock()

	_, ok := t.tree.Get(value)

	return ok
}
