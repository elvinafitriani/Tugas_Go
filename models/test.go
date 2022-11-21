package models

type Barang struct {
	Id_barang         int    `gorm:"primaryKey;autoIncrement;" json:"id_barang"`
	Nama_barang      string `json:"nama_barang"`
	Harga      int    `json:"harga"`
	KategoriID int    `json:"kategori_id"`
}

type Jual struct {
	Id_jual      int    `gorm:"primaryKey;autoIncrement;" json:"id_jual"`
	Barang_Id int    `json:"barang_id"`
	Barang    Barang `gorm:"foreignKey:Barang_Id;references:Id_barang;" json:"barang"`
}

type Kategori struct {
	Id_kategori     int      `gorm:"primaryKey;autoIncrement;" json:"id_kategori"`
	Nama_kategori   string   `json:"nama_kategori"`
	Barang []Barang `gorm:"foreignKey:KategoriID;" json:"barang"`
}

//task:
//get barang : -> id brg, nama brg, harga,kategori-> id,nama
//jual : -> id jual,
