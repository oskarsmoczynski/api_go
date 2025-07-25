package models

import (
	"time"
)

type User struct {
	UserId    uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"size:100;not null"`
	Email     string    `gorm:"size:100;unique;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type Order struct {
	OrderId     uint      `gorm:"primaryKey;autoIncrement"`
	UserID      uint      `gorm:"not null"`
	OrderDate   time.Time `gorm:"autoCreateTime"`
	TotalAmount float64   `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type Product struct {
	ProductId  uint     `gorm:"primaryKey;autoIncrement"`
	Name       string   `gorm:"size:100;not null"`
	Price      float64  `gorm:"not null"`
	Stock      int      `gorm:"default:0"`
	CategoryID uint     `gorm:"not null"`
	Category   Category `gorm:"foreignKey:CategoryID;references:CategoryId"`
}

type Category struct {
	CategoryId uint   `gorm:"primaryKey;autoIncrement"`
	Name       string `gorm:"size:50;unique;not null"`
}

type OrderItem struct {
	OrderItemId uint    `gorm:"primaryKey;autoIncrement"`
	OrderID     uint    `gorm:"not null"`
	ProductID   uint    `gorm:"not null"`
	Quantity    int     `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	Order       Order   `gorm:"foreignKey:OrderID;references:OrderId"`
	Product     Product `gorm:"foreignKey:ProductID;references:ProductId"`
}

type Review struct {
	ReviewId  uint      `gorm:"primaryKey;autoIncrement"`
	ProductID uint      `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	Rating    int       `gorm:"check:rating >= 1 AND rating <= 5"`
	Comment   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Product   Product   `gorm:"foreignKey:ProductID;references:ProductId"`
	User      User      `gorm:"foreignKey:UserID;references:UserId"`
}
