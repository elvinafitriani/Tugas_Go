package models

type Barang struct {
	Id_Barang   int    `gorm:"primaryKey;autoIncreament;" json:"id_barang"`
	Nama_Barang string `json:"nama_barang"`
	Harga       int    `json:"harga"`
	Kategori_Id int    `json:"kategori_id"`
}

type Jual struct {
	Id_Jual   int    `gorm:"primaryKey;autoIncreament;" json:"id_jual"`
	Barang_Id int    `json:"barang_id"`
	Barang    Barang `gorm:"foreignKey:Barang_Id;references:Id_Barang" json:"barang"`
}

type Kategori struct {
	Id_Ktg   int      `gorm:"primaryKey;autoIncreament;" json:"id_ktg"`
	Nama_Ktg string   `json:"nama_ktg"`
	Barang   []Barang `gorm:"foreignKey:Kategori_Id;references:Id_Ktg" json:"barang"`
}

type Join struct {
	Id_Barang   int    `jaon:"id_barang"`
	Nama_Barang string `json:"nama_barang"`
	Harga       int    `json:"harga"`
	Id_Ktg      int    `json:"id_ktg"`
	Nama_Ktg    string `json:"nama_ktg"`
	Id_Jual     int    `json:"id_jual"`
}
