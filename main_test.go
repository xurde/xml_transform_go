package main

import "testing"

func BenchmarkTransformation(b *testing.B) {
	doTransformation()
}

func BenchmarkCachedTransformation(b *testing.B) {
	doCachedTransformation()
}

func BenchmarkDummy(b *testing.B) {
}
