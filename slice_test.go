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

	a1 := NewSet[int64](1, 23, 4, 5, 5, 6, 9)
	a2 := NewSet[int64](1, 23, 19)
	a3 := a1.UnionAll(a2)
	a1.Sorted()
	a3.Sorted()
	t.Log(a3, a1)

}
