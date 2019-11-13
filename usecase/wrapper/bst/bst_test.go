package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBst(t *testing.T) {
	a := assert.New(t)

	tree := New(nil)

	a.False(tree.Has(12), "doesn't have a value not from constructor")
	a.NotPanics(func() { tree.Del(12) }, "ok for empty constructor")
	tree.Put(12)
	a.True(tree.Has(12), "got inserted value")

	tree = New([]int{45, 12, 78})

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
