/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package lambda

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
)

type stream[T any] struct {
	data Slice[any]
	cmp  func(i, j any) bool
}

func Stream[T any](in any) *stream[T] {
	val := reflect.ValueOf(in)
	if val.Kind() != reflect.Array && val.Kind() != reflect.Slice {
		_ = fmt.Errorf("input is not an array or slice")
		return nil
	}

	// 创建一个切片来存储拆解后的元素
	result := make([]any, val.Len())

	// 将数组或切片的每个元素添加到结果切片中
	for i := 0; i < val.Len(); i++ {
		result[i] = val.Index(i).Interface()
	}
	return &stream[T]{
		data: result,
	}
}

func (l *stream[T]) Len() int {
	return len(l.data)
}

func (l *stream[T]) Less(i, j int) bool {
	return l.cmp(l.data[i], l.data[j])
}

func (l *stream[T]) Swap(i, j int) {
	l.data[i], l.data[j] = l.data[j], l.data[i]
}

func (l *stream[T]) Sort(cmp func(i, j any) bool) {
	l.cmp = cmp
	sort.Sort(l)
}

func (l *stream[T]) Foreach(w func(i any)) {
	for _, t := range l.data {
		w(t)
	}
}

func (l *stream[T]) Map(c func(i any) T) *stream[T] {
	var out []any
	for _, t := range l.data {
		out = append(out, c(t))
	}
	return &stream[T]{
		data: out,
	}
}

func (l *stream[T]) Filter(f func(i any) bool) *stream[any] {
	return &stream[any]{data: *l.data.Filter(f)}
}

func (l *stream[T]) Slice() *Slice[T] {
	var out Slice[T]
	for _, datum := range l.data {
		out = append(out, datum.(T))
	}
	return &out
}

func (l *stream[T]) Group(k func(i any) any, v func(i any) T) map[any]Slice[T] {
	result := make(map[any]Slice[T])
	for _, t := range l.data {
		result[k(t)] = append(result[k(t)], v(t))
	}
	return result
}

func (l *stream[T]) FlatMap(k func(i any) any, v func(i any) T) map[any]T {
	result := make(map[any]T)
	for _, t := range l.data {
		result[k(t)] = v(t)
	}
	return result
}

func (l *stream[T]) String() string {
	data, _ := json.Marshal(l.data)
	return string(data)
}

func (l *stream[T]) SumInt(value func(i any) int) int {
	i := 0
	for _, data := range l.data {
		i += value(data)
	}
	return i
}

func (l *stream[T]) SumInt64(value func(i any) int64) int64 {
	var i int64 = 0
	for _, data := range l.data {
		i += value(data)
	}
	return i
}
func (l *stream[T]) SumFloat(value func(i any) float64) float64 {
	i := 0.0
	for _, data := range l.data {
		i += value(data)
	}
	return i
}
