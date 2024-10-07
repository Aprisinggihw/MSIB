package entities

type Product struct {
	Id          uint `json:"id"`
	Nama        string `json:"nama"`
	Kategori    string `json:"kategori"`
	Stok       int64 `json:"stok"`
	Deskripsi string `json:"deskripsi"`
}
