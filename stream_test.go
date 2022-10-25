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

type User struct {
	Id   int
	Name string
}

func TestStream(t *testing.T) {
	users := []User{
		{Id: 1, Name: "a"}, {Id: 1, Name: "b"}, {Id: 3, Name: "b"}, {Id: 4, Name: "x"},
	}

	var x = Stream([]int{1, 2, 3, 4}).Filter(func(i int) bool { return i > 2 }).
		Map(func(i int) Value {
			return i
		}).Filter(func(i Value) bool { return i.(int) > 2 })

	t.Log(x.IntList())
	t.Log(x.StringList())
	t.Log(Dump(x.IntList()))
	t.Log(Dump(x.StringList()))
	t.Log(x.String())

	x2 := Stream(users).Filter(func(i User) bool { return i.Id > 0 }).
		Map(func(i User) Value {
			return i.Id
		})

	x3 := x2.StringGroup(func(u Value) string {
		return fmt.Sprintf("val-%d", u)
	}, func(u Value) any {
		return u
	})

	t.Log(Stream(users).String())
	t.Log(x2.String())
	t.Log(Dump(x3), x3)

	x4 := x2.IntGroup(func(u Value) int { return u.(int) }, func(u Value) any { return u })
	t.Log(Dump(x4))

}
