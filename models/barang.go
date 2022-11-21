package models

type Barang struct {
	Id         int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Nama_brg   string `json:"nama_brg"`
	Harga      int    `json:"harga"`
	KategoriID int    `json:"kategori_id"`
}

type Jual struct {
	Id_jual   int    `gorm:"primaryKey;autoIncrement;" json:"id_jual"`
	Barang_ID int    `json:"barang_id"`
	Barang    Barang `gorm:"foreignKey:Barang_ID;references:Id;" json:"barang"`
}

type Kategori struct {
	Id_ktg   int      `gorm:"primaryKey;autoIncrement;" json:"id_ktg"`
	Nama_ktg string   `json:"nama_ktg"`
	Barang   []Barang `gorm:"foreignKey:KategoriID;" json:"barang"`
}

//task
//id barang, nama brg, harga, kategori>id,nama, jual=>idjual
