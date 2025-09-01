package main

import (
	"testing"

	"github.com/dracory/hb"
)

// user is a sample struct used for benchmarking.
var user = struct {
	FirstName      string
	FavoriteColors []string
	RawContent     string
	EscapedContent string
}{
	FirstName:      "Bob",
	FavoriteColors: []string{"blue", "green", "mauve"},
	RawContent:     "<div><p>Raw Content to be displayed</p></div>",
	EscapedContent: "<div><div><div>Escaped</div></div></div>",
}

// BenchmarkPageWithHeaderAndUL benchmarks the creation of a page with a header and unordered list.
// To run this benchmark, use the following command:
// go test -bench=^BenchmarkPageWithHeaderAndUL*  -benchtime 5s -count 2 -cpu 1,4
//
// Flags explanation:
// -bench=^BenchmarkPageWithHeaderAndUL*: specifies to run benchmarks matching the name "BenchmarkPageWithHeaderAndUL".
// -benchtime=5s: specifies the minimum running time for each benchmark to be 5 seconds.
// -count=2: specifies that each benchmark will be run 2 times.
// -cpu=1,4: specifies the number of CPUs to use for running benchmarks (1 and 4 in this case).
//
// To see the number of allocations, use the -benchmem flag:
// go test -bench=^BenchmarkPageWithHeaderAndUL* -benchmem
//
// To run profiling, use the following command from the project root:
// task profile
func BenchmarkPageWithHeaderAndUL(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ul := hb.UL()
			for i := 0; i < len(user.FavoriteColors); i++ {
				li := hb.LI()
				li.Text(user.FavoriteColors[i]) // Use Text to reduce allocations
				ul.Child(li)
			}
			page := hb.NewTag("html"). // Create a new HTML tag
							Child(hb.NewTag("body"). // Add a body tag as a child
											Child(hb.Heading1().Text(user.FirstName)).                            // Add an H1 heading with the user's first name, using Text to reduce allocations
											Child(hb.Paragraph().Text("Here's a list of your favorite colors:")). // Add a paragraph, using Text to reduce allocations
											Child(ul))                                                            // Add the unordered list
			page.ToHTML() // Render the HTML page
		}
	})
}
