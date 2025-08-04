// internal/model/entities.go
package model

import "time"

type Consumer struct {
	ID            string    `gorm:"primaryKey" json:"id"`
	NIK           string    `json:"nik"`
	FullName      string    `json:"full_name"`
	LegalName     string    `json:"legal_name"`
	TempatLahir   string    `json:"tempat_lahir"`
	TanggalLahir  time.Time `json:"tanggal_lahir"`
	Gaji          float64   `json:"gaji"`
	FotoKTPURL    string    `json:"foto_ktp_url"`
	FotoSelfieURL string    `json:"foto_selfie_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Limit struct {
	ID         string
	ConsumerID string
	TenorMonth int
	TotalLimit float64
	UsedLimit  float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Transaction struct {
	ID             string
	ContractNumber string
	ConsumerID     string
	TenorMonth     int
	JumlahOTR      float64
	AdminFee       float64
	JumlahCicilan  float64
	JumlahBunga    float64
	NamaAsset      string
	CreatedAt      time.Time
}
