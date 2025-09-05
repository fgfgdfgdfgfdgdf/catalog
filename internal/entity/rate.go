package entity

import "github.com/jackc/pgx/v5/pgtype"

type Rate struct {
	ID         int64              `gorm:"type:bigserial; autoIncrement; not null"`
	UsdPerTon  pgtype.Numeric     `gorm:"type:numeric(12, 4); not null"`
	UsdPerStar pgtype.Numeric     `gorm:"type:numeric(12, 6); not null"`
	IsActive   bool               `gorm:"type:boolean; index:unique_true_value; default:true; not null"`
	CreatedAt  pgtype.Timestamptz `gorm:"type:timestamptz; default:now(); not null"`
}

type RateBody struct {
	UsdPerTon  string `form:"usdPerTon"`
	UsdPerStar string `form:"usdPerStar"`
}
