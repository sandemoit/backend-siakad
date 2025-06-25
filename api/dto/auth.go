package dto

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterSekolahRequest struct {
	NamaSekolah  string `json:"nama_sekolah"`
	Provinsi     string `json:"provinsi"`
	Kabupaten    string `json:"kabupaten"`
	Kecamatan    string `json:"kecamatan"`
	Alamat       string `json:"alamat"`
	NPSN         string `json:"npsn"`
	EmailSekolah string `json:"email_sekolah"`
	Telepon      string `json:"telepon"`
	Logo         string `json:"logo,omitempty"`

	NamaAdmin    string `json:"nama_admin"`
	JabatanAdmin string `json:"jabatan_admin"`
	NoWaAdmin    string `json:"no_wa_admin"`
	EmailAdmin   string `json:"email_admin"`
}
