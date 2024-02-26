// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type EditTKegSurvei struct {
	ID             string     `json:"id"`
	SurveiKd       string     `json:"survei_kd"`
	KegKd          string     `json:"keg_kd"`
	Status         int        `json:"status"`
	TglBuka        *time.Time `json:"tgl_buka,omitempty"`
	TglRekMulai    *time.Time `json:"tgl_rek_mulai,omitempty"`
	TglRekSelesai  *time.Time `json:"tgl_rek_selesai,omitempty"`
	TglMulai       *time.Time `json:"tgl_mulai,omitempty"`
	TglSelesai     *time.Time `json:"tgl_selesai,omitempty"`
	IsRekrutmen    *int       `json:"is_rekrutmen,omitempty"`
	IsMulti        *int       `json:"is_multi,omitempty"`
	IsConfirm      bool       `json:"is_confirm"`
	IsAddIndicator bool       `json:"is_add_indicator"`
	UpdatedBy      string     `json:"updated_by"`
}

type Kegiatan struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
	User   *User  `json:"user"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Mutation struct {
}

type NewKegiatan struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

type NewTKegSurvei struct {
	SurveiKd       string     `json:"survei_kd"`
	KegKd          string     `json:"keg_kd"`
	Status         int        `json:"status"`
	TglBuka        *time.Time `json:"tgl_buka,omitempty"`
	TglRekMulai    *time.Time `json:"tgl_rek_mulai,omitempty"`
	TglRekSelesai  *time.Time `json:"tgl_rek_selesai,omitempty"`
	TglMulai       *time.Time `json:"tgl_mulai,omitempty"`
	TglSelesai     *time.Time `json:"tgl_selesai,omitempty"`
	IsRekrutmen    *int       `json:"is_rekrutmen,omitempty"`
	IsMulti        *int       `json:"is_multi,omitempty"`
	IsConfirm      bool       `json:"is_confirm"`
	IsAddIndicator bool       `json:"is_add_indicator"`
	CreatedBy      string     `json:"created_by"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Query struct {
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TKegSurvei struct {
	ID             string     `json:"id"`
	SurveiKd       string     `json:"survei_kd"`
	KegKd          string     `json:"keg_kd"`
	Status         int        `json:"status"`
	TglBuka        *time.Time `json:"tgl_buka,omitempty"`
	TglRekMulai    *time.Time `json:"tgl_rek_mulai,omitempty"`
	TglRekSelesai  *time.Time `json:"tgl_rek_selesai,omitempty"`
	TglMulai       *time.Time `json:"tgl_mulai,omitempty"`
	TglSelesai     *time.Time `json:"tgl_selesai,omitempty"`
	IsRekrutmen    *int       `json:"is_rekrutmen,omitempty"`
	IsMulti        *int       `json:"is_multi,omitempty"`
	IsConfirm      bool       `json:"is_confirm"`
	IsAddIndicator bool       `json:"is_add_indicator"`
	CreatedBy      string     `json:"created_by"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedBy      *string    `json:"updated_by,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}
