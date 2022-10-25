/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package lambda

func (l *Lambda[T]) Filter(f func(i T) bool) *Lambda[T] {
	var out []T
	for _, t := range l.data {
		if f(t) {
			out = append(out, t)
		}
	}
	return &Lambda[T]{data: out}
}
