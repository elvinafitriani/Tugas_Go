package models

type Kategori struct {
	IdKat  int      `gorm:"primaryKey;autoIncrement;" json:"id"`
	Nama   string   `json:"nama"`
	Barang []Barang `gorm:"foreignKey:KategoriID;" json:"barang"`
}

type Barang struct {
	Id         int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	NamaBrg    string `json:"nama"`
	Harga      int    `json:"harga"`
	KategoriID int    `json:"kategori_id"`
}

type Jual struct {
	IdJual    int    `gorm:"primaryKey;autoIncrement;" json:"id_Jual"`
	Barang_Id int    `json:"barang_id"`
	Barang    Barang `gorm:"foreignKey:Barang_Id;references:Id;" json:"barang"`
}
