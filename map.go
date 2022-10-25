/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package lambda

func (l *Lambda[T]) Map(f func(i T) Value) *Lambda[Value] {
	var out []Value
	for _, t := range l.data {
		out = append(out, f(t))
	}

	return &Lambda[Value]{
		data: out,
	}
}
