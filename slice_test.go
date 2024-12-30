/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package lambda

import (
	"fmt"
	"testing"
)

type personal struct {
	age  int
	name string
}

func (p *personal) String() string {
	return fmt.Sprintf("%d-%s", p.age, p.name)
}

func (p *personal) HashCode() int {
	return p.age
}

func TestSlice_Sort(t *testing.T) {

	a1 := NewSet[int64](1, 23, 4, 5, 5, 6, 9, 1, 2, 1, 1)
	a2 := NewSet[int64](1, 23, 19)
	a3 := a1.UnionAll(a2)

	a3.Sorted()
	t.Log(a3, a1)
	var a4 []int64 = a1.Slice().Values()
	t.Log(a4)
	values := []int{1, 2, 3, 3}
	sm := Stream[float64](values)
	data := sm.Map(func(i any) float64 {
		return float64(i.(int)) / 0.3
	}).Slice()
	t.Log(data)
	var out []int64
	a1.Slice().Foreach(func(i int64) {
		out = append(out, i)
	})
	t.Log(out)
}
