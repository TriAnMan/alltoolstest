package domain

type Collection interface {
	Has(int) bool
	Put(int)
	Del(int)
}
