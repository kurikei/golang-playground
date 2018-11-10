package main

import (
	"strconv"
	"testing"
)

const mapSize = 100000

func BenchmarkCreateMapEmptyString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createMapEmptyString()
	}
}

func BenchmarkCreateMapEmptyBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createMapEmptyBool()
	}
}

func BenchmarkCreateMapEmptyStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createMapEmptyStruct()
	}
}

func createMapEmptyString() {
	emptyStringMap := map[string]string{}
	for i := 0; i < mapSize; i++ {
		emptyStringMap[strconv.Itoa(i)] = ""
	}
}

func createMapEmptyBool() {
	emptyStringMap := map[string]bool{}
	for i := 0; i < mapSize; i++ {
		emptyStringMap[strconv.Itoa(i)] = true
	}
}

func createMapEmptyStruct() {
	emptyStringMap := map[string]struct{}{}
	for i := 0; i < mapSize; i++ {
		emptyStringMap[strconv.Itoa(i)] = struct{}{}
	}
}
