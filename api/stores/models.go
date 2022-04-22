package stores

type Store struct {
	id         uint
	name       string
	address    string
	coordinate Coordinate
}

type Coordinate struct {
	lat float32
	lon float32
}
