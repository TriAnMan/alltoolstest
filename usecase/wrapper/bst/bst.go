// Package bst provides a tread-safe CRUD wrapper for a Binary Search Tree structure
package bst

import (
	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/sirupsen/logrus"
	"sync"
)

type Bst struct {
	tree *rbt.Tree
	mx   sync.RWMutex
	log  logrus.FieldLogger
}

func New(log logrus.FieldLogger, values []int) *Bst {
	log.WithField("subsystem", "bst").WithField("action", "create").Debug(values)

	t := &Bst{
		tree: rbt.NewWithIntComparator(),
		mx:   sync.RWMutex{},
		log:  log,
	}

	for _, value := range values {
		t.tree.Put(value, struct{}{})
	}

	return t
}

func (t *Bst) Put(value int) {
	t.log.WithField("subsystem", "bst").WithField("action", "put").Debug(value)

	t.mx.Lock()
	defer t.mx.Unlock()

	t.tree.Put(value, struct{}{})
}

func (t *Bst) Del(value int) {
	t.log.WithField("subsystem", "bst").WithField("action", "del").Debug(value)

	t.mx.Lock()
	defer t.mx.Unlock()

	t.tree.Remove(value)
}

func (t *Bst) Has(value int) bool {
	t.log.WithField("subsystem", "bst").WithField("action", "has").Debug(value)

	t.mx.RLock()
	defer t.mx.RUnlock()

	_, ok := t.tree.Get(value)

	return ok
}
