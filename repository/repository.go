package repository

import (
	"desa-sragen/domain"
	"errors"
	"log"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.Repo {
	return &repo{
		db: db,
	}
}

func (r *repo) GetUser(username string) (record *domain.Auth, err error) {
	if err := r.db.Table("auth").Where("username = ?", username).Scan(&record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (r *repo) GetUserById(id int) (record *domain.Auth, err error) {
	if err := r.db.Table("auth").Where("auth_id = ?", id).Scan(&record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (r *repo) CreateUser(req domain.Auth) error {
	if err := r.db.Table("auth").Create(&req).Error; err != nil {
		return errors.New("internal server error")
	}

	return nil
}

func (r *repo) GetAllAdmin() (record []domain.Auth, err error) {
	if err := r.db.Table("auth").Where("role_id = 2").Scan(&record).Error; err != nil {
		return nil, errors.New("internal server error")
	}

	return record, nil
}

func (r *repo) InActiveAuth(id int, isActive bool) error {
	data := map[bool]int8{true: 1, false: 0}[isActive]

	if err := r.db.Debug().Table("auth").Where("auth_id = ?", id).Update("is_active", data).Error; err != nil {
		return errors.New("internal server error")
	}

	return nil
}

func (r *repo) CreatePengajuan(req domain.PengajuanParam) error {
	if err := r.db.Table("pengajuan_dokumen").Create(&req).Error; err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}

	return nil
}

func (r *repo) GetDocumentsUser(authId int) (record []domain.Documents, err error) {
	if err := r.db.Table("pengajuan_dokumen pd").Joins("JOIN master_pengajuan mp ON pd.jenis_pengajuan = mp.id_jenis_pengajuan").Where("pd.auth_id = ?", authId).Order("pd.updated_at DESC").Scan(&record).Error; err != nil {
		return nil, errors.New("internal server error")
	}

	return record, nil
}

func (r *repo) GetAllPengajuan(idPengajuan int32) (record []domain.Documents, err error) {
	if err := r.db.Table("pengajuan_dokumen pd").Joins("JOIN master_pengajuan mp ON pd.jenis_pengajuan = mp.id_jenis_pengajuan").Where("pd.jenis_pengajuan = ?", idPengajuan).Order("pd.updated_at DESC").Scan(&record).Error; err != nil {
		return nil, errors.New("internal server error")
	}

	return record, nil
}

func (r *repo) GetPengajuanById(id int, idJenis int8) (record *domain.Documents, err error) {
	if err := r.db.Table("pengajuan_dokumen pd").Joins("JOIN master_pengajuan mp ON pd.jenis_pengajuan = mp.id_jenis_pengajuan").Where("pd.id_pengajuan = ? AND pd.jenis_pengajuan = ?", id, idJenis).Scan(&record).Error; err != nil {
		return nil, errors.New("internal server error")
	}

	return record, nil
}

func (r *repo) UpdateStatus(id int, req domain.Aksi) error {
	if err := r.db.Table("pengajuan_dokumen").Where("id_pengajuan = ? AND status = 0", id).Updates(&req).Error; err != nil {
		return errors.New("internal server error")
	}

	return nil
}

func (r *repo) GetAllUsers() (record []domain.Auth, err error) {
	if err := r.db.Table("auth").Where("role_id = 3").Scan(&record).Error; err != nil {
		return nil, errors.New("internal server error")
	}

	return record, nil
}

func (r *repo) UpdatePassword(id int, pass string) error {
	if err := r.db.Table("auth").Where("auth_id = ?", id).Update("password", pass).Error; err != nil {
		return err
	}

	return nil
}

func (r *repo) StoreTeams(req domain.Teams) error {
	err := r.db.Table("teams").Create(&req)

	return err.Error
}

func (r *repo) GetTeams() (record []domain.Teams, err error) {
	if err := r.db.Table("teams").Scan(&record).Error; err != nil {
		return nil, errors.New("internal server error")
	}

	return record, nil
}

func (r *repo) CountPengajuan() int {
	var count int64
	r.db.Table("pengajuan_dokumen").Count(&count)
	return int(count)
}

func (r *repo) CountBelumDiproses() int {
	var count int64
	r.db.Table("pengajuan_dokumen").Where("status = 2").Count(&count)
	return int(count)
}

func (r *repo) CountUsers() int {
	var count int64
	r.db.Table("auth").Where("role_id = 3").Count(&count)
	return int(count)
}

func (r *repo) LastDokumen() (record []domain.Dashboard, err error) {
	res := r.db.Table("pengajuan_dokumen pd").Joins("join auth a on pd.auth_id = a.auth_id").Joins("join master_pengajuan mp on pd.jenis_pengajuan = mp.id_jenis_pengajuan").Order("pd.created_at DESC").Limit(10).Scan(&record)

	if res.Error != nil {
		log.Println(res.Error)
	}
	
	return record, nil
}