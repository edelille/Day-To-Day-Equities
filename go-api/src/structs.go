package main

type QuoteJSON struct {
	Ticker string
	Date   string
}

type QuoteSingular struct {
	C   int
	H   int
	L   int
	N   int
	O   int
	Otc bool
	T   int
	V   int
	Vw  int
}

type QuoteResponse struct {
	Ticker       string
	Adjusted     bool
	QueryCount   int
	RequestId    string
	ResultsCount int
	Status       string
	Results      []QuoteSingular
	NextUrl      string
}
