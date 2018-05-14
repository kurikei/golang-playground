package main

import "strconv"

const mapSize = 100000

func CreateMapEmptyString() {
	emptyStringMap := map[string]string{}
	for i := 0; i < mapSize; i++ {
		emptyStringMap[strconv.Itoa(i)] = ""
	}
}

func CreateMapEmptyBool() {
	emptyStringMap := map[string]bool{}
	for i := 0; i < mapSize; i++ {
		emptyStringMap[strconv.Itoa(i)] = true
	}
}

func CreateMapEmptyStruct() {
	emptyStringMap := map[string]struct{}{}
	for i := 0; i < mapSize; i++ {
		emptyStringMap[strconv.Itoa(i)] = struct{}{}
	}
}
