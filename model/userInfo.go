package model

type User struct {
	Phone	string	`gorm:"column:phone", json:"phone"`
	Passoword	string	`gorm:"colmun:password", json:"passoword"`
	Id		int	`gorm:"column:id", gorm:"AUTO_INCREMENT", json:"-"'`
	Level	int	`gorm:"column:level"`
	Name	string	`gorm:"column:name", json:"name"`
	Sex		string	`gorm:"column:sex", json:"sex"`
	Birth	string	`gorm:"column:birth", json:"birth"`
	Native	string	`gorm:"column:native", json:"native"`
	NativePlace	string	`gorm:"column:native_place", json:"native_place"`
	Email	string	`gorm:"column:email", json:"email"`
	IdentityNumber	string	`gorm:"column:identity_number",json:"identity_number"`
	ImageUrl	string	`gorm:"column:image_url", json:"image_url"`
}

type LoginInfo struct {
	Phone	string	`gorm:"column:phone", json:"phone"`
	Passoword	string	`gorm:"colmun:password", json:"passoword"`
}

type RegisterInfo struct {
	Phone	string	`gorm:"column:phone", json:"phone"`
	Vcd		string	`gorm:"column:vcd", json:"vcd"`
	Password	string `gorm:"column:passwprd", json:"password"`
}