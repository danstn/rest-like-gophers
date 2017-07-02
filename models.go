package main

// Address example type
type Address struct {
	Line1    string
	Line2    string
	Postcode uint8
	Country  string
}

// Location example type
type Location struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

// Locations wrapper
type Locations []Location
