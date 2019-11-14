package bst

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

type devNull struct{}

func (devNull) Write(p []byte) (int, error) {
	return len(p), nil
}

func TestBst(t *testing.T) {
	a := assert.New(t)

	log := logrus.StandardLogger()
	log.SetOutput(devNull{})

	tree := New(log, nil)

	a.False(tree.Has(12), "doesn't have a value not from constructor")
	a.NotPanics(func() { tree.Del(12) }, "ok for empty constructor")
	tree.Put(12)
	a.True(tree.Has(12), "got inserted value")

	tree = New(log, []int{45, 12, 78})

	a.True(tree.Has(12), "has an value from constructor")
	tree.Del(12)
	a.False(tree.Has(12), "dont have a removed value")
	a.NotPanics(func() { tree.Del(12) }, "ok for double deletion")
	a.False(tree.Has(12), "dont have a removed value")

	a.False(tree.Has(13), "doesn't have a value not from constructor")
	tree.Put(13)
	a.True(tree.Has(13), "got inserted value")
	a.NotPanics(func() { tree.Put(13) }, "ok for double insertion")
	a.True(tree.Has(13), "got inserted value")
}

func TestTreadSafety(t *testing.T) {
	if !isRace {
		t.Skip("no race detector found")
	}

	log := logrus.StandardLogger()
	log.SetOutput(devNull{})

	tree := New(log, []int{45, 12, 78})

	go tree.Has(13)
	go tree.Put(13)
	go tree.Del(13)

	go tree.Has(13)
	go tree.Del(13)
	go tree.Put(13)

	go tree.Del(13)
	go tree.Put(13)
	go tree.Has(13)

	go tree.Del(13)
	go tree.Has(13)
	go tree.Put(13)

	go tree.Put(13)
	go tree.Del(13)
	go tree.Has(13)

	go tree.Put(13)
	go tree.Has(13)
	go tree.Del(13)
}
