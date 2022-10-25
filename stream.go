/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package lambda

import (
	"encoding/json"
	"fmt"
)

type Lambda[T Any] struct {
	data []T
}

func Stream[T Any](in []T) *Lambda[T] {
	return &Lambda[T]{
		data: in,
	}
}

func (l *Lambda[T]) List(f func(i T) any) []any {
	return []any{f(l.data[0])}
}

func (l *Lambda[T]) IntList() []int {
	var result []int
	for _, t := range l.data {
		var x interface{} = t
		result = append(result, x.(int))
	}
	return result
}

func (l *Lambda[T]) StringList() []string {
	var result []string
	for _, t := range l.data {
		result = append(result, fmt.Sprintf("%v", t))
	}
	return result
}

func (l *Lambda[T]) Group(k func(i T) any, v func(i T) any) map[any][]any {
	result := make(map[any][]any)
	for _, t := range l.data {
		result[k(t)] = append(result[k(t)], v(t))
	}
	return result
}

func (l *Lambda[T]) StringGroup(k func(i T) string, v func(i T) any) map[string][]any {
	result := make(map[string][]any)
	for _, t := range l.data {
		result[k(t)] = append(result[k(t)], v(t))
	}
	return result
}

func (l *Lambda[T]) IntGroup(k func(i T) int, v func(i T) any) map[int][]any {
	result := make(map[int][]any)
	for _, t := range l.data {
		result[k(t)] = append(result[k(t)], v(t))
	}
	return result
}

func (l *Lambda[T]) FlatMap(k func(i T) any, v func(i T) any) map[any]any {
	result := make(map[any]any)
	for _, t := range l.data {
		result[k(t)] = v(t)
	}
	return result
}

func (l *Lambda[T]) FlatStringMap(k func(i T) string, v func(i T) any) map[string]any {
	result := make(map[string]any)
	for _, t := range l.data {
		result[k(t)] = v(t)
	}
	return result
}
func (l *Lambda[T]) FlatIntMap(k func(i T) int, v func(i T) any) map[int]any {
	result := make(map[int]any)
	for _, t := range l.data {
		result[k(t)] = v(t)
	}
	return result
}

func (l *Lambda[T]) String() string {
	data, _ := json.Marshal(l.data)
	return string(data)
}
