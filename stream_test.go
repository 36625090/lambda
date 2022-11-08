/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package lambda

import (
	"log"
	"testing"
)

type User struct {
	Id   int
	Name string
}

func (u User) HashCode() int {
	return u.Id
}

func TestStream(t *testing.T) {
	users := []User{
		{Id: 1, Name: "a"}, {Id: 5, Name: "b"}, {Id: 3, Name: "b"}, {Id: 3, Name: "x"},
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	x2 := Stream(users).Filter(func(i User) bool { return true })
	//.Slice()
	x2.Slice().Reverse()
	t.Log(x2)

	x2.Map(func(i User) any { return i }).Slice().Reverse()
	t.Log(x2)

}
