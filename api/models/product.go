package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Product struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name        string    `gorm:"size:255;not null" json:"name"`
	Stock       uint64    `gorm:"not null" json:"stock"`
	Description string    `gorm:"size:255;not null" json:"description"`
	Price       float64   `gorm:"not null" json:"price"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Product) prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.Stock = 0
	p.Description = html.EscapeString(strings.TrimSpace(p.Description))
	p.Price = 0
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.New("Required Product Name")
	}
	if p.Description == "" {
		return errors.New("Required Description")
	}
	return nil
}

func (p *Product) SaveProduct(db *gorm.DB) (*Product, error) {
	var err error
	err = db.Debug().Model(&Product{}).Create(&p).Error
	if err != nil {
		return &Product{}, err
	}
	if p.ID != 0 {

	}
	return p, nil
}

func (p *Product) FindAllProduct(db *gorm.DB) (*[]Product, error) {
	var err error
	product := []Product{}
	err = db.Debug().Model(&Product{}).Limit(100).Find(&product).Error
	if err != nil {
		return &[]Product{}, err
	}
	return &product, nil
}

func (p *Product) FindProductByID(db *gorm.DB, pid uint64) (*Product, error) {
	var err error
	err = db.Debug().Model(&Product{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func (p *Product) UpdateAProduct(db *gorm.DB) (*Product, error) {
	var err error
	err = db.Debug().Model(&Product{}).Where("id = ?", p.ID).Updates(Product{Name: p.Name, Stock: p.Stock, Description: p.Description, Price: p.Price, CreatedAt: time.Now()}).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func (p *Product) DeleteAProduct(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&Product{}).Where("id = ?", pid).Take(&Product{}).Delete(&Product{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Product not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
