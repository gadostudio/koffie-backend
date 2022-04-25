package stores

type Store struct {
	Id      uint    `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Lat     float32 `json:"lat"`
	Lon     float32 `json:"lon"`
}

type StoreResponse struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`

	Coordinate Coordinate `json:"coordinate"`
}

type Coordinate struct {
	Lat float32
	Lon float32
}
