```bash
$ go test -bench Benchmark -benchmem
BenchmarkCreateMapEmptyString-8   	      50	  26146700 ns/op	11200486 B/op	  106112 allocs/op
BenchmarkCreateMapEmptyBool-8     	     100	  21971030 ns/op	 6485665 B/op	  106130 allocs/op
BenchmarkCreateMapEmptyStruct-8   	     100	  21979310 ns/op	 6121120 B/op	  106119 allocs/op
```
