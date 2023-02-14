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

func TestSlice(t *testing.T) {
	a1 := Stream[int64]([]int64{1, 2, 3, 4}).Map(func(i int64) any { return int(i) }).Slice().Values()
	t.Log(a1.([]int))
	t.Log(len(a1.([]int)), cap(a1.([]int)))
}
