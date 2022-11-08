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

func TestSlice_Sort(t *testing.T) {
	a1 := Slice[int64]{1, 23, 4, 5, 6, 5, 9}
	a2 := Slice[int64]{8, 9, 5, 1, 0, 1}
	a3 := a1.RemoveAll(&a2)
	a3.Sorted()
	t.Log(a3)
	a3.Remove(4)
	t.Log(a3)

	ap := Slice[personal]{{age: 10, name: "10"}, {age: 10, name: "10"}, {age: 11, name: "11"}}
	ap.Sorted()
	t.Log(ap)
	var i complex128 = complex(3, 7)
	j := complex(3, 17)

	t.Log(Equal(i, j))
}
