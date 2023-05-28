package product

type Product struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Price    float32 `json:"price"`
	Quantity uint16  `json:"quantity"` // I believe one day we can migrate to uint32 type ;)
}
