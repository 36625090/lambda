/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package lambda

import (
	"sort"
)

type Slice[T interface{}] []T

func (m *Slice[T]) Filter(f func(i T) bool) []T {
	var out []T
	for _, t := range *m {
		if f(t) {
			out = append(out, t)
		}
	}
	return out
}

func (m *Slice[T]) Contains(i T) bool {
	for _, t := range *m {
		if Equal(i, t) {
			return true
		}
	}
	return false
}

func (m *Slice[T]) Append(values ...T) {
	*m = append(*m, values...)
}

func (m *Slice[T]) Insert(values ...T) {
	*m = append(values, *m...)
}

func (m *Slice[T]) Sort(cmp func(i, j T) bool) {
	s := &slice[T]{data: *m}
	s.cmp = cmp
	sort.Sort(s)
	*m = s.data
}

func (m *Slice[T]) Sorted() {
	s := &slice[T]{data: *m}
	sort.Sort(s)
	*m = s.data
}

func (m *Slice[T]) Reverse() {
	s := &slice[T]{data: *m}
	sort.Sort(sort.Reverse(s))
	*m = s.data
}

func (m *Slice[T]) Remove(i T) bool {
	for i, v := range *m {
		if Equal(v, i) {
			if i+1 == len(*m) {
				*m = (*m)[:i]
			} else {
				copy((*m)[i:], (*m)[i+1:])
				*m = (*m)[:len(*m)-1]
			}
			return true
		}
	}
	return false
}

// RetailAll 求交集
func (m *Slice[T]) RetailAll(i *Slice[T]) Slice[T] {
	var tmp Slice[T]
	for _, v := range *m {
		if i.Contains(v) {
			tmp.Append(v)
		}
	}
	if len(tmp) == 0 {
		return tmp
	}
	var result Slice[T]
	for _, v := range *i {
		if tmp.Contains(v) && !result.Contains(v) {
			result.Append(v)
		}
	}
	return result
}

// UnionAll 求并集
func (m *Slice[T]) UnionAll(i *Slice[T]) Slice[T] {

	sl := make(Slice[T], len(*m)+len(*i))
	copy(sl, *m)
	copy(sl[len(*m):], *i)
	var result Slice[T]
	for _, v := range sl {
		if !result.Contains(v) {
			result.Append(v)
		}
	}
	return result
}

// RemoveAll 求差集
func (m *Slice[T]) RemoveAll(i *Slice[T]) Slice[T] {
	if len(*m) == 0 {
		return Slice[T]{}
	}
	if len(*i) == 0 {
		return *m
	}
	var result Slice[T]
	for _, v := range *m {
		if !i.Contains(v) {
			result.Append(v)
		}
	}
	return result
}

func (m *Slice[T]) ContainsAll(i *Slice[T]) Slice[T] {
	if len(*m) == 0 {
		return Slice[T]{}
	}
	if len(*i) == 0 {
		return *m
	}

	var result Slice[T]
	for _, v := range *m {
		if i.Contains(v) {
			result.Append(v)
		}
	}
	return result
}

type slice[T any] struct {
	data Slice[T]
	cmp  func(i, j T) bool
}

func (s *slice[T]) Len() int {
	return len(s.data)
}

func (s *slice[T]) Less(i, j int) bool {
	if s.cmp == nil {
		return Less(&s.data[i], &s.data[j])
	}
	return s.cmp(s.data[i], s.data[j])
}

func (s *slice[T]) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}
