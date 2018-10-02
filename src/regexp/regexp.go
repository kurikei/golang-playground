package main

import (
	"fmt"
	"regexp"
)

const src =  "google.com"
var (
	regMatch    = regexp.MustCompile("o+")
	regNotMatch = regexp.MustCompile("hogefugapiyo")
)

func main() {
	// FindString
	fmt.Println("=== FindString ===")
	p("FindString:match", regMatch.FindString(src))       // => "oo"
	p("FindString:notmatch", regNotMatch.FindString(src)) // => ""

	// FindStringIndex
	fmt.Println("=== FindStringIndex ===")
	p("FindStringIndex:match", regMatch.FindStringIndex(src))       // => []int{1, 3}
	p("FindStringIndex:notmatch", regNotMatch.FindStringIndex(src)) // => nil

	// FindStringSubmatch
	fmt.Println("=== FindStringSubmatch ===")
	p("FindStringSubmatch:match", regMatch.FindStringSubmatch(src))       // => []string{"oo"}
	p("FindStringSubmatch:notmatch", regNotMatch.FindStringSubmatch(src)) // => nil

	// FindStringSubmatchIndex
	fmt.Println("=== FindStringSubmatchIndex ===")
	p("FindStringSubmatchIndex:match", regMatch.FindStringSubmatchIndex(src))       // => []int{1, 3}
	p("FindStringSubmatchIndex:notmatch", regNotMatch.FindStringSubmatchIndex(src)) // => nil

	// FindAllString
	fmt.Println("=== FindAllString ===")
	p("FindAllString:match:n=0", regMatch.FindAllString(src, 0))       // => nil
	p("FindAllString:match:n=1", regMatch.FindAllString(src, 1))       // => []string{"oo"}
	p("FindAllString:match:n=2", regMatch.FindAllString(src, 2))       // => []string{"oo", "o"}
	p("FindAllString:notmatch:n=1", regNotMatch.FindAllString(src, 1)) // => nil

	// FindAllStringIndex
	fmt.Println("=== FindAllStringIndex ===")
	p("FindAllStringIndex:match:n=0", regMatch.FindAllStringIndex(src, 0))       // => nil
	p("FindAllStringIndex:match:n=1", regMatch.FindAllStringIndex(src, 1))       // => [][]int{[]int{1, 3}}
	p("FindAllStringIndex:match:n=2", regMatch.FindAllStringIndex(src, 2))       // => [][]int{[]int{1, 3}, []int{8, 9}}
	p("FindAllStringIndex:notmatch:n=1", regNotMatch.FindAllStringIndex(src, 1)) // => nil

	// FindAllStringSubmatch
	fmt.Println("=== FindAllStringSubmatch ===")
	p("FindAllStringSubmatch:match:n=0", regMatch.FindAllStringSubmatch(src, 0))       // => nil
	p("FindAllStringSubmatch:match:n=1", regMatch.FindAllStringSubmatch(src, 1))       // => [][]string{[]string{"oo"}}
	p("FindAllStringSubmatch:match:n=2", regMatch.FindAllStringSubmatch(src, 2))       // => [][]string{[]string{"oo"}, []string{"o"}}
	p("FindAllStringSubmatch:notmatch:n=1", regNotMatch.FindAllStringSubmatch(src, 1)) // => nil

	// FindAllStringSubmatchIndex
	fmt.Println("=== FindAllStringSubmatchIndex ===")
	p("FindAllStringSubmatchIndex:match:n=0", regMatch.FindAllStringSubmatchIndex(src, 0))       // => nil
	p("FindAllStringSubmatchIndex:match:n=1", regMatch.FindAllStringSubmatchIndex(src, 1))       // => [][]int{[]int{1, 3}}
	p("FindAllStringSubmatchIndex:match:n=2", regMatch.FindAllStringSubmatchIndex(src, 2))       // => [][]int{[]int{1, 3}, []int{8, 9}}
	p("FindAllStringSubmatchIndex:notmatch:n=1", regNotMatch.FindAllStringSubmatchIndex(src, 1)) // => nil
}

func p(message string, content interface{}) {
	format := message + ": %#v\n"
	fmt.Printf(format, content)
}
