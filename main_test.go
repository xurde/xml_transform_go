package main

import "testing"

func BenchmarkTransformation(b *testing.B) {
	doTransformation()
}

func BenchmarkDummy(b *testing.B) {
}
