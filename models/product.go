package models

// Product mewakili informasi produk yang diterima dari Product Service
type Product struct {
	ID    uint   `json:"id"  `   // ID produk
	Name  string `json:"name"`   // Nama produk (opsional, bisa digunakan untuk logging)
	Stock int    ` json:"stock"` // Stok awal produk
}
