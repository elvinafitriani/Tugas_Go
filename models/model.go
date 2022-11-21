package models

type Barang struct {
	Id_brg     int    `gorm:"primaryKey;autoIncrement;" json:"id_brg"`
	Nama_brg   string `json:"nama_brg"`
	Harga      int    `json:"harga"`
	KategoriID int    `json:"kategori_id"`
}

type Jual struct {
	Id_jual   int    `gorm:"primaryKey:autoIncrement;" json:"id_jual"`
	Barang_Id int    `json:"barang_id"`
	Barang    Barang `gorm:"foreignKey:Barang_Id;references:Id_brg;" json:"barang"`
}

type Kategori struct {
	Id_ktg   int      `gorm:"primaryKey;autoIncrement;" json:"id_ktg"`
	Nama_ktg string   `json:"nama_ktg"`
	Barang   []Barang `gorm:"foreignKey:KategoriID;" json:"barang"`
}

type Join struct {
	Id_brg     int    `json:"id_brg"`
	Nama_brg   string `json:"nama_brg"`
	Harga      int    `json:"harga"`
	KategoriID int    `json:"kategori_id"`
	Nama_ktg   string `json:"nama_ktg"`
	Id_jual    int    `json:"id_jual"`
}

//task//

//barang = 	id barang
//jual +	nama
//			harga
//			kategori - id,nama
//			jual - idjual

//select join
