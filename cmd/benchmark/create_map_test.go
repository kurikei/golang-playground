package main

import (
	"testing"
)

func BenchmarkCreateMapEmptyString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateMapEmptyString()
	}
}

func BenchmarkCreateMapEmptyBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateMapEmptyBool()
	}
}

func BenchmarkCreateMapEmptyStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateMapEmptyStruct()
	}
}
