package util

import (
	"errors"
	"io"
	"reflect"
)

type Iterator struct {
	typ    reflect.Type
	idx    int
	values []interface{}
}

func NewIterator(values ...interface{}) *Iterator {
	if len(values) <= 0 {
		panic("create Iterator")
	}

	typ := reflect.TypeOf(values[0])

	iterator := &Iterator{
		typ:    typ,
		idx:    0,
		values: make([]interface{}, 0, len(values)),
	}

	for i := 0; i < len(values); i++ {
		if err := iterator.Add(values[i]); err != nil {
			panic("create Iterator")
		}
	}

	return iterator
}

func (i *Iterator) Add(v interface{}) error {
	if i.typ != reflect.TypeOf(v) {
		return errors.New("value not fit")
	}
	i.values = append(i.values, v)
	return nil
}

func (i *Iterator) HasNext() bool {
	return i.idx < len(i.values)
}

func (i *Iterator) Next() (interface{}, error) {
	if i.HasNext() {
		v := i.values[i.idx]
		i.idx++
		return v, nil
	}
	return nil, io.EOF
}

func (i *Iterator) Reset() {
	i.idx = 0
}

func (i *Iterator) Clear() {
	i.idx = 0
	i.values = i.values[0:]
}

func (i *Iterator) Clone() *Iterator {
	values := make([]interface{}, len(i.values))
	copy(values, i.values)

	return &Iterator{
		typ:    i.typ,
		idx:    i.idx,
		values: values,
	}
}
