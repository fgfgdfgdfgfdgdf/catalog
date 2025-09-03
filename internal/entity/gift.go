package entity

import (
	"errors"
	"strings"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"gorm.io/gorm/clause"
)

var (
	GiftColumns = utils.ParseFieldNames(&Gift{})
)

type Gift struct {
	ID         int64              `gorm:"type:bigserial; autoIncrement; not null"`
	Name       string             `gorm:"type:text; not null"`
	Type       string             `gorm:"type:text; check:type IN ('collectible','unique'); not null"`
	Rarity     string             `gorm:"type:text; check:rarity IN ('common','rare','epic','legendary'); not null"`
	PriceStars int32              `gorm:"type:integer; index:,sort:desc;  not null;"`
	PriceUsd   pgtype.Numeric     `gorm:"type:numeric(12, 2); index:,sort:desc; not null"`
	PriceTon   pgtype.Numeric     `gorm:"type:numeric(12, 2); not null"`
	CreatedAt  pgtype.Timestamptz `gorm:"type:timestamptz; default:now(); not null"`
}

type PaginatedGiftResponse struct {
	Items     []*Gift `json:"items"`
	Page      int64   `json:"page"`
	PerPage   int64   `json:"perPage"`
	Total     int64   `json:"total"`
	PageCount int     `json:"pageCount"`
}

type GiftQuery struct {
	Type        string    `form:"type"`
	Rarity      CSV       `form:"rarity"`
	Sort        SortRules `form:"sort"`
	MinPriceUsd int64     `form:"minPriceUsd"`
	MaxPriceUsd int64     `form:"maxPriceUsd"`
	Search      string    `form:"search"`
	Page        int64     `form:"page"`
	PerPage     int64     `form:"perPage"`
}

const (
	descSort = "desc"
	ascSort  = "asc"
)

type CSV []string

type SortRules struct {
	Items []clause.OrderByColumn
}

func (c *SortRules) UnmarshalParam(param string) error {
	arr := strings.Split(param, ",")
	for _, v := range arr {
		sub := strings.Split(v, ":")

		if len(sub) != 2 {
			return errors.New("invalid sort")
		}

		var desc bool

		if sub[1] == descSort {
			desc = true
		} else if sub[1] != ascSort {
			return errors.New("invalid sort")
		}

		c.Items = append(c.Items, clause.OrderByColumn{
			Column: clause.Column{
				Name: sub[0],
			},
			Desc: desc,
		})
	}

	return nil
}

func (c *CSV) UnmarshalParam(param string) error {
	*c = strings.Split(param, ",")
	return nil
}
