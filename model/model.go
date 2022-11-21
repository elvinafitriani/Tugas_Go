package model

type Barang struct {
	IdBarang int    `gorm:"column:id_bar;primaryKey;autoIncrement" json:"id_bar"`
	Nama     string `gorm:"column:nama_barang" json:"nama"`
	Harga    int    `gorm:"column:harga" json:"harga"`
	// KategoriID int
	KategoriID int `json:"kategori_id"`
}

type Jual struct {
	Id       int    `gorm:"column:id_jual;primaryKey;autoIncrement;" json:"id_jual"`
	BarangId int    `gorm:"column:id_barang" json:"barang_id"`
	Barang   Barang `gorm:"foreignkey:id_barang;references:id_bar;"`
}

type Kategori struct {
	IdKat   int      `gorm:"column:id_kategori;primaryKey;autoIncrement" json:"id_kat"`
	NamaKat string   `json:"nama_kategori"`
	Barang  []Barang `gorm:"foreignkey:kategori_id;"`
}

type Select struct {
	IdBar      int    `json:"id_bar"`
	NamaBarang string `json:"nama_barang"`
	Harga      int    `json:"harga"`
	KategoriID int    `json:"kategori_id"`
	NamaKat    string `json:"nama_kat"`
	IdJual     int    `json:"id_jual"`
}
