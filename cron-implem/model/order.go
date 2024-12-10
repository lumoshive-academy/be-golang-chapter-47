package model

// Order adalah struktur data untuk merepresentasikan data order
type Order struct {
	OrderID   string
	Customer  string
	Amount    float64
	OrderDate string
}
