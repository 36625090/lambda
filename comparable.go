/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package lambda

import (
	"fmt"
	"reflect"
)

type Comparable interface {
	HashCode() int
}

func Equal(i, j any) bool {
	if isComparable(i, j) {
		return equal(i, j)
	}
	iv := reflect.ValueOf(i)
	jv := reflect.ValueOf(j)
	if iv.Type() != jv.Type() {
		return false
	}
	if iv.Kind() == reflect.Pointer {
		iv = iv.Elem()
	}

	if jv.Kind() == reflect.Pointer {
		jv = jv.Elem()
	}
	if iv.Kind() == reflect.Interface {
		return equal(iv.Elem().Interface(), jv.Elem().Interface())
	}

	if iv.CanInt() {
		return iv.Int() == jv.Int()
	}
	if iv.CanUint() {
		return iv.Uint() == jv.Uint()
	}
	if iv.CanFloat() {
		return iv.Float() == jv.Float()
	}
	if iv.CanComplex() {
		ic := iv.Complex()
		jc := jv.Complex()
		return ic == jc
	}

	return false
}

func Less(i, j any) bool {

	if isComparable(i, j) {
		return less(i, j)
	}

	iv := reflect.ValueOf(i)
	jv := reflect.ValueOf(j)
	if iv.Type() != jv.Type() {
		return false
	}

	if iv.Kind() == reflect.Pointer {
		iv = iv.Elem()
	}
	if jv.Kind() == reflect.Pointer {
		jv = jv.Elem()
	}

	if iv.Kind() == reflect.Interface {
		return less(iv.Elem().Interface(), jv.Elem().Interface())
	}

	if iv.CanInt() {
		return iv.Int() < jv.Int()
	}
	if iv.CanUint() {
		return iv.Uint() < jv.Uint()
	}
	if iv.CanFloat() {
		return iv.Float() < jv.Float()
	}
	if iv.CanComplex() {
		ic := iv.Complex()
		jc := jv.Complex()
		return real(ic) < real(jc) && imag(ic) < imag(jc)
	}

	return false
}

func LessEqual(i, j any) bool {
	if isComparable(i, j) {
		return lessEqual(i, j)
	}

	iv := reflect.ValueOf(i)
	jv := reflect.ValueOf(j)
	if iv.Type() != jv.Type() {
		return false
	}

	if iv.Kind() == reflect.Pointer {
		iv = iv.Elem()
	}
	if jv.Kind() == reflect.Pointer {
		jv = jv.Elem()
	}
	if iv.Kind() == reflect.Interface {
		return lessEqual(iv.Elem().Interface(), jv.Elem().Interface())
	}

	if iv.CanInt() {
		return iv.Int() <= jv.Int()
	}
	if iv.CanUint() {
		return iv.Uint() <= jv.Uint()
	}
	if iv.CanFloat() {
		return iv.Float() <= jv.Float()
	}

	if iv.CanComplex() {
		ic := iv.Complex()
		jc := jv.Complex()
		return real(ic) <= real(jc) && imag(ic) <= imag(jc)
	}
	return false
}

func Greater(i, j any) bool {
	if isComparable(i, j) {
		return greater(i, j)
	}

	iv := reflect.ValueOf(i)
	jv := reflect.ValueOf(j)
	if iv.Type() != jv.Type() {
		return false
	}

	if iv.Kind() == reflect.Pointer {
		iv = iv.Elem()
	}
	if jv.Kind() == reflect.Pointer {
		jv = jv.Elem()
	}
	if iv.Kind() == reflect.Interface {
		return greater(iv.Elem().Interface(), jv.Elem().Interface())
	}
	if iv.CanInt() {
		return iv.Int() > jv.Int()
	}
	if iv.CanUint() {
		return iv.Uint() > jv.Uint()
	}
	if iv.CanFloat() {
		return iv.Float() > jv.Float()
	}

	if iv.CanComplex() {
		ic := iv.Complex()
		jc := jv.Complex()
		return real(ic) > real(jc) && imag(ic) > imag(jc)
	}
	return false
}
func GreaterEqual(i, j any) bool {
	if isComparable(i, j) {
		return greaterEqual(i, j)
	}

	iv := reflect.ValueOf(i)
	jv := reflect.ValueOf(j)
	if iv.Type() != jv.Type() {
		return false
	}

	if iv.Kind() == reflect.Pointer {
		iv = iv.Elem()
	}
	if jv.Kind() == reflect.Pointer {
		jv = jv.Elem()
	}
	if iv.Kind() == reflect.Interface {
		return greaterEqual(iv.Elem().Interface(), jv.Elem().Interface())
	}
	if iv.CanInt() {
		return iv.Int() >= jv.Int()
	}
	if iv.CanUint() {
		return iv.Uint() >= jv.Uint()
	}
	if iv.CanFloat() {
		return iv.Float() >= jv.Float()
	}

	if iv.CanComplex() {
		ic := iv.Complex()
		jc := jv.Complex()
		return real(ic) >= real(jc) && imag(ic) >= imag(jc)
	}

	return false
}

func equal(i, j any) bool {
	switch tp := i.(type) {
	case Comparable:
		switch tp2 := j.(type) {
		case Comparable:
			return tp.HashCode() == tp2.HashCode()
		}

	case fmt.Stringer:
		switch tp2 := j.(type) {
		case fmt.Stringer:
			return tp.String() == tp2.String()
		}

	}
	return false
}

func less(i, j any) bool {
	switch tp := i.(type) {
	case Comparable:
		switch tp2 := j.(type) {
		case Comparable:
			return tp.HashCode() < tp2.HashCode()
		}

	case fmt.Stringer:
		switch tp2 := j.(type) {
		case fmt.Stringer:
			return tp.String() < tp2.String()
		}

	}
	return false
}
func lessEqual(i, j any) bool {
	switch tp := i.(type) {
	case Comparable:
		switch tp2 := j.(type) {
		case Comparable:
			return tp.HashCode() <= tp2.HashCode()
		}
	case fmt.Stringer:
		switch tp2 := j.(type) {
		case fmt.Stringer:
			return tp.String() <= tp2.String()
		}

	}
	return false
}

func greater(i, j any) bool {
	switch tp := i.(type) {
	case Comparable:
		switch tp2 := j.(type) {
		case Comparable:
			return tp.HashCode() > tp2.HashCode()
		}

	case fmt.Stringer:
		switch tp2 := j.(type) {
		case fmt.Stringer:
			return tp.String() > tp2.String()
		}

	}
	return false
}
func greaterEqual(i, j any) bool {
	switch tp := i.(type) {
	case Comparable:
		switch tp2 := j.(type) {
		case Comparable:
			return tp.HashCode() >= tp2.HashCode()
		}
	case fmt.Stringer:
		switch tp2 := j.(type) {
		case fmt.Stringer:
			return tp.String() >= tp2.String()
		}

	}
	return false
}

func isComparable(i, j any) bool {
	switch i.(type) {
	case Comparable:
		switch j.(type) {
		case Comparable:
			return true
		}
	case fmt.Stringer:
		switch j.(type) {
		case fmt.Stringer:
			return true
		}
	}
	return false
}
