package main

import "testing"
import "time"

func BenchmarkTransformation(b *testing.B) {
	doTransformation()
}

func BenchmarkDummy(b *testing.B) {
}
