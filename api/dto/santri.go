package dto

import "time"

// SantriRequest - DTO untuk create dan update santri
type SantriRequest struct {
	UserID        uint       `json:"user_id" validate:"required"`
	NIS           string     `json:"nis" validate:"required,min=3,max=20"`
	NamaLengkap   string     `json:"nama_lengkap" validate:"required,min=2,max=100"`
	JenisKelamin  string     `json:"jenis_kelamin" validate:"required,oneof=L P"`
	TanggalLahir  *time.Time `json:"tanggal_lahir" validate:"required"`
	TempatLahir   string     `json:"tempat_lahir" validate:"required,min=2,max=50"`
	Alamat        string     `json:"alamat" validate:"required,min=10"`
	Telepon       string     `json:"telepon" validate:"omitempty,min=10,max=15"`
	NamaAyah      string     `json:"nama_ayah" validate:"required,min=2,max=100"`
	NamaIbu       string     `json:"nama_ibu" validate:"required,min=2,max=100"`
	PekerjaanAyah string     `json:"pekerjaan_ayah" validate:"omitempty,max=50"`
	PekerjaanIbu  string     `json:"pekerjaan_ibu" validate:"omitempty,max=50"`
	TeleponOrtu   string     `json:"telepon_ortu" validate:"required,min=10,max=15"`
	TanggalMasuk  *time.Time `json:"tanggal_masuk" validate:"required"`
	Status        string     `json:"status" validate:"required,oneof=aktif non_aktif lulus keluar"`
	KelasID       uint       `json:"kelas_id" validate:"required"`
	AsramaID      uint       `json:"asrama_id" validate:"omitempty"`
	Foto          string     `json:"foto" validate:"omitempty,url"`
	SekolahID     uint       `json:"sekolah_id"`
	IsActive      bool       `json:"is_active"`
}

// SantriResponse - DTO untuk response santri
type SantriResponse struct {
	ID            uint             `json:"id"`
	UserID        uint             `json:"user_id"`
	NIS           string           `json:"nis"`
	NamaLengkap   string           `json:"nama_lengkap"`
	JenisKelamin  string           `json:"jenis_kelamin"`
	TanggalLahir  *time.Time       `json:"tanggal_lahir"`
	TempatLahir   string           `json:"tempat_lahir"`
	Alamat        string           `json:"alamat"`
	Telepon       string           `json:"telepon"`
	NamaAyah      string           `json:"nama_ayah"`
	NamaIbu       string           `json:"nama_ibu"`
	PekerjaanAyah string           `json:"pekerjaan_ayah"`
	PekerjaanIbu  string           `json:"pekerjaan_ibu"`
	TeleponOrtu   string           `json:"telepon_ortu"`
	TanggalMasuk  *time.Time       `json:"tanggal_masuk"`
	Status        string           `json:"status"`
	KelasID       uint             `json:"kelas_id"`
	AsramaID      uint             `json:"asrama_id"`
	Foto          string           `json:"foto"`
	SekolahID     uint             `json:"sekolah_id"`
	IsActive      bool             `json:"is_active"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
	Sekolah       *SekolahResponse `json:"sekolah,omitempty"`
	User          *UserResponse    `json:"user,omitempty"`
	Kelas         *KelasResponse   `json:"kelas,omitempty"`
	Asrama        *AsramaResponse  `json:"asrama,omitempty"`
}

// SantriListResponse - DTO untuk response list santri dengan pagination
type SantriListResponse struct {
	Data       []SantriResponse `json:"data"`
	Total      int64            `json:"total"`
	Page       int              `json:"page"`
	Limit      int              `json:"limit"`
	TotalPages int              `json:"total_pages"`
}

// SantriSearchRequest - DTO untuk pencarian santri
type SantriSearchRequest struct {
	Name      string `json:"name" validate:"omitempty,min=2"`
	NIS       string `json:"nis" validate:"omitempty,min=3"`
	KelasID   uint   `json:"kelas_id" validate:"omitempty"`
	AsramaID  uint   `json:"asrama_id" validate:"omitempty"`
	Status    string `json:"status" validate:"omitempty,oneof=aktif non_aktif lulus keluar"`
	IsActive  *bool  `json:"is_active" validate:"omitempty"`
	Page      int    `json:"page" validate:"min=1"`
	Limit     int    `json:"limit" validate:"min=1,max=100"`
	SortBy    string `json:"sort_by" validate:"omitempty,oneof=nama_lengkap nis created_at updated_at"`
	SortOrder string `json:"sort_order" validate:"omitempty,oneof=asc desc"`
}

// SantriBulkDeleteRequest - DTO untuk bulk delete
type SantriBulkDeleteRequest struct {
	IDs []uint `json:"ids" validate:"required,min=1,max=100"`
}

// SantriBulkUpdateRequest - DTO untuk bulk update
type SantriBulkUpdateRequest struct {
	IDs      []uint `json:"ids" validate:"required,min=1,max=100"`
	KelasID  *uint  `json:"kelas_id" validate:"omitempty"`
	AsramaID *uint  `json:"asrama_id" validate:"omitempty"`
	Status   string `json:"status" validate:"omitempty,oneof=aktif non_aktif lulus keluar"`
	IsActive *bool  `json:"is_active" validate:"omitempty"`
}

// SantriStatsResponse - DTO untuk statistik santri
type SantriStatsResponse struct {
	TotalSantri     int64 `json:"total_santri"`
	SantriAktif     int64 `json:"santri_aktif"`
	SantriNonAktif  int64 `json:"santri_non_aktif"`
	SantriLaki      int64 `json:"santri_laki"`
	SantriPerempuan int64 `json:"santri_perempuan"`
	SantriLulus     int64 `json:"santri_lulus"`
	SantriKeluar    int64 `json:"santri_keluar"`
}

// Related DTOs for nested responses

// SekolahResponse - DTO untuk response sekolah (nested)
type SekolahResponse struct {
	ID          uint   `json:"id"`
	NamaSekolah string `json:"nama_sekolah"`
	Alamat      string `json:"alamat"`
	Telepon     string `json:"telepon"`
}

// UserResponse - DTO untuk response user (nested)
type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// KelasResponse - DTO untuk response kelas (nested)
type KelasResponse struct {
	ID        uint   `json:"id"`
	NamaKelas string `json:"nama_kelas"`
	Tingkat   string `json:"tingkat"`
}

// AsramaResponse - DTO untuk response asrama (nested)
type AsramaResponse struct {
	ID         uint   `json:"id"`
	NamaAsrama string `json:"nama_asrama"`
	Kapasitas  int    `json:"kapasitas"`
}
