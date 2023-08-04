package domain

import "time"

type Auth struct {
	AuthId    int
	Username  string
	Email     string
	Password  string
	RoleId    int8
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ClaimedToken struct {
	UserId   int
	Username string
	Role     int
}

type Documents struct {
	IdPengajuan      uint      `json:"id_pengajuan"`
	AuthId           int64     `json:"auth_id"`
	Nama             string    `json:"nama"`
	Nik              string    `json:"nik"`
	Phone            string    `json:"phone"`
	JenisKelamin     string    `json:"jenis_kelamin"`
	Agama            string    `json:"agama"`
	Pekerjaan        string    `json:"pekerjaan"`
	TempatLahir      string    `json:"tempat_lahir"`
	KepalaKeluarga   string    `json:"kepala_keluarga"`
	NomorKk          string    `json:"nomor_kk"`
	TanggalLahir     time.Time `form:"tanggal_lahir" time_format:"2006-01-02" json:"tanggal_lahir"`
	Alamat           string    `json:"alamat"`
	Foto             string    `json:"foto"`
	Ktp              string    `json:"ktp"`
	TandaTangan      string    `json:"tanda_tangan"`
	Alasan           string    `json:"alasan"`
	Status           int       `json:"status"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	UpdatedBy        int
	JenisPengajuan   int8    `json:"jenis_pengajuan"`
	Keperluan        string  `json:"keperluan"`
	StatusPerkawinan string  `json:"status_perkawinan"`
	JenisUsaha       string  `json:"jenis_usaha"`
	NamaAnak         string  `json:"nama_anak"`
	NamaIbu          string  `json:"nama_ibu"`
	NamaAyah         string  `json:"nama_ayah"`
	NamaPengaju      string  `json:"nama_pengaju"`
	JenisKelaminAnak string  `json:"jenis_kelamin_anak"`
	HubunganPengaju  string  `json:"hubungan_pengaju"`
	NamaPengajuan    string  `json:"nama_pengajuan"`
	AlasanPenolakan  string  `json:"alasan_penolakan"`
	Dokumen          string  `json:"dokumen"`
	HariMeninggal    *string `json:"hari_meninggal"`
	BendaHilang      *string `json:"benda_hilang"`
	Umur             *int    `json:"umur"`
}
