package model
type DummyModel struct {
	Id uint
	UId  int `gorm:"primaryKey"`
	Name string
	Price int
}