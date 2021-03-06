package jewerly

import (
	"errors"
	"gopkg.in/guregu/null.v3"
)

// Inputs
type CreateProductInput struct {
	Titles       MultiLanguageInput `json:"titles" binding:"required"`
	Descriptions MultiLanguageInput `json:"descriptions" binding:"required"`
	Material     MultiLanguageInput `json:"materials" binding:"required"`
	Price        float32            `json:"price" binding:"required,min=0"`
	Code         string             `json:"code" binding:"required"`
	ImageIds     []int              `json:"image_ids" binding:"required,min=1"`
	CategoryId   Category           `json:"category_id" binding:"required"`
}

func (i CreateProductInput) Validate() error {
	return i.CategoryId.Validate()
}

type UpdateProductInput struct {
	Titles       *MultiLanguageInput `json:"titles"`
	Descriptions *MultiLanguageInput `json:"descriptions"`
	Material     *MultiLanguageInput `json:"materials"`
	Price        null.Float          `json:"price"`
	Code         null.String         `json:"code"`
	CategoryId   *Category           `json:"category_id"`
	InStock      null.Bool           `json:"in_stock"`
}

func (i UpdateProductInput) Validate() error {
	if (UpdateProductInput{}) == i  {
		return errors.New("empty update product input")
	}

	if i.Price.Valid &&  i.Price.Float64 <= 0 {
		return errors.New("price can't be negative or zero")
	}

	if i.CategoryId != nil {
		return i.CategoryId.Validate()
	}

	return nil
}

type MultiLanguageInput struct {
	English   string `json:"english" binding:"required"`
	Russian   string `json:"russian" binding:"required"`
	Ukrainian string `json:"ukrainian" binding:"required"`
}

type GetAllProductsFilters struct {
	Language   string
	Offset     int
	Limit      int
	CategoryId null.Int
}

// Responses
type ProductResponse struct {
	Id          int         `json:"id" db:"id"`
	Title       string      `json:"title" db:"title"`
	Description string      `json:"description" db:"description"`
	Material    string      `json:"material" db:"material"`
	Price       float32     `json:"price" db:"price"`
	Code        null.String `json:"code" db:"code"`
	Images      []Image     `json:"images"`
	CategoryId  Category    `json:"category_id" db:"category_id"`
	InStock     bool        `json:"in_stock" db:"in_stock"`
}

type Image struct {
	Id      int         `json:"id" db:"id"`
	URL     string      `json:"url" db:"url"`
	AltText null.String `json:"alt_text" db:"alt_text"`
}

type ProductsList struct {
	Products []ProductResponse `json:"data"`
	Total    int               `json:"total"`
}

// Categories

type Category int

func (c Category) Validate() error {
	_, ok := Categories[c]
	if !ok {
		return errors.New("invalid category")
	}

	return nil
}

const (
	CategoryRings = iota + 1
	CategoryBracelets
	CategoryPendants
	CategoryEarring
	CategoryNecklaces
	CategorySets

	English    = "english"
	Ukraininan = "ukrainian"
	Russian    = "russian"
)

var (
	Categories = map[Category]bool{
		CategoryRings:     true,
		CategoryBracelets: true,
		CategoryPendants:  true,
		CategoryEarring:   true,
		CategoryNecklaces: true,
		CategorySets:      true,
	}

	languageQueries = map[string]string{
		"en":        English,
		"ru":        Russian,
		"ua":        Ukraininan,
		"english":   English,
		"russian":   Russian,
		"ukrainian": Ukraininan,
	}
)

func GetLanguageFromQuery(query string) string {
	if val, ok := languageQueries[query]; ok {
		return val
	}

	return English
}