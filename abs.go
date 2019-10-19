package calc

func Abs(a int64) int64 {
	// http://cavaliercoder.com/blog/optimized-abs-for-int64-in-go.html
	// https://books.google.com.au/books?id=VicPJYM0I5QC&lpg=PA18&ots=2o-SROAuXq&dq=hackers%20delight%20absolute&pg=PA18#v=onepage&q=hackers%20delight%20absolute&f=false
	x := a >> 63
	// return (a ^ x) - x
	return (a + x) ^ x
}

func Nabs(a int64) int64 {
	x := a >> 63
	// return x -  (a ^ x)
	return (x - a) ^ x
}
