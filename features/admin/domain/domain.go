package domain

import (
	"mime/multipart"
	"time"
)

type UserCore struct {
	ID          uint
	FullName    string
	Email       string
	Password    string
	Role        string
	Phone       string
	Address     string
	Dob         string
	Gender      string
	UserPicture string
}

type BookingCore struct {
	ID              uint
	IdUser          uint
	DateStart       time.Time
	DateEnd         time.Time
	Entrance        string
	Ticket          int
	IdRanger        uint
	GrossAmount     int
	Token           string
	OrderId         string
	Link            string
	StatusBooking   string
	StatusPendakian string
	FullName        string
	Phone           string
}

type ProductCore struct {
	ID             uint
	ProductName    string
	RentPrice      int
	Detail         string
	Note           string
	ProductPicture string
}

type RangerCore struct {
	ID          uint
	UserID      uint
	Docs        string
	Price       uint
	Detail      string
	Status      string
	StatusApply string
	User        UserCore
}

type ClimberCore struct {
	ID            uint
	IsClimber     int
	MaleClimber   int
	FemaleClimber int
}

type Code struct {
	ID           uint
	Code         string
	AccessToken  string
	TokenType    string
	RefreshToken string
	Expiry       time.Time
}

type PesanCore struct {
	ID       uint
	IdRanger uint
	Email    string
	Status   string
}

// DOMAIN REPOSITORY
type Repository interface {
	GetPendaki() ([]BookingCore, ClimberCore, error)
	InsertClimber(data ClimberCore) (ClimberCore, error)
	GetProduct(page int) ([]ProductCore, int, int, error)
	InsertProduct(newProduct ProductCore) (ProductCore, error)
	UpdateProduct(newProduct ProductCore) (ProductCore, error)
	DeleteProduct(id int) error
	GetAllRanger() ([]RangerCore, []RangerCore, error)
	EditRanger(data RangerCore, datas UserCore, id uint) (RangerCore, UserCore, PesanCore, error)
	DeleteRanger(id int) error
	GetCode() (Code, error)
}

// DOMAIN SERVICES
type Services interface {
	GetPendaki() ([]BookingCore, ClimberCore, error)
	AddClimber(data ClimberCore) (ClimberCore, error)
	GetProduct(page int) ([]ProductCore, int, int, error)
	AddProduct(newProduct ProductCore, file multipart.File, fileheader *multipart.FileHeader) (ProductCore, error)
	EditProduct(newProduct ProductCore, file multipart.File, fileheader *multipart.FileHeader) (ProductCore, error)
	RemoveProduct(id int) error
	ShowAllRanger() ([]RangerCore, []RangerCore, error)
	UpdateRanger(data RangerCore, datas UserCore, id uint) (RangerCore, UserCore, PesanCore, error)
	RemoveRanger(id int) error
	GetCode() (Code, error)
}
