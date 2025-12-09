// main_test.go
package main

import "testing"

// run with go test -bench=Benchmark -benchmem -benchtime 10s
// to understand the output check https://dev-to-uploads.s3.amazonaws.com/uploads/articles/484ap11qw8d81b43gg0v.png
// original article at https://www.freecodecamp.org/news/how-to-write-benchmark-tests-for-your-golang-functions/
// if its down check:
// https://web.archive.org/web/20250606195415im_/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/484ap11qw8d81b43gg0v.png
// and https://web.archive.org/web/20250806061109/https://www.freecodecamp.org/news/how-to-write-benchmark-tests-for-your-golang-functions/
func BenchmarkPart2(b *testing.B) {
	for b.Loop() {
		part2("input.txt")
	}
}

func BenchmarkPart2opt2(b *testing.B) {
	for b.Loop() {
		part2opt2("input.txt")
	}
}
