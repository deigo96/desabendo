package controller

import (
	"desa-sragen/domain"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

const maxFileSize = 5 * 1024 * 1024 // 5 MB

func (cx *Controller) PengajuanDomisili(c *gin.Context) {

	c.HTML(200, "pengajuan-surat-domisili.html", gin.H{
		"url": cx.Env(),
	})
}

func (cx *Controller) PengajuanKtp(c *gin.Context) {

	c.HTML(200, "pengajuan-ktp.html", gin.H{
		"url": cx.Env(),
	})
}

func (cx *Controller) PengajuanKehilangan(c *gin.Context) {

	c.HTML(200, "pengajuan-surat-kehilangan.html", gin.H{
		"url": cx.Env(),
	})
}

func (cx *Controller) PengajuanKelahiran(c *gin.Context) {

	c.HTML(200, "pengajuan-surat-kelahiran.html", gin.H{
		"url": cx.Env(),
	})
}

func (cx *Controller) PengajuanKematian(c *gin.Context) {

	c.HTML(200, "pengajuan-surat-kematian.html", gin.H{
		"url": cx.Env(),
	})
}

func (cx *Controller) PengajuanKk(c *gin.Context) {

	c.HTML(200, "pengajuan-surat-kk.html", gin.H{
		"url": cx.Env(),
	})
}

func (cx *Controller) PengajuanSku(c *gin.Context) {

	c.HTML(200, "pengajuan-surat-ku.html", gin.H{
		"url": cx.Env(),
	})
}

func (cx *Controller) ListDokumen(c *gin.Context) {
	auth_id := c.GetInt("userId")
	if auth_id == 0 {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{})
		return
	}

	documents, err := cx.Repo.GetDocumentsUser(auth_id)
	if err != nil {
		c.HTML(500, "page-404.html", gin.H{
			"url": cx.Env(),
		})
		return
	}

	c.HTML(200, "list-document.html", gin.H{
		"url":  cx.Env(),
		"data": documents,
	})
}

func (cx *Controller) KirimPengajuan(c *gin.Context) {

	auth_id := c.GetInt("userId")
	nik := c.PostForm("nik")
	nama := c.PostForm("nama")
	phone := c.PostForm("phone")
	jenis_kelamin := c.PostForm("jenis_kelamin")
	agama := c.PostForm("agama")
	pekerjaan := c.PostForm("pekerjaan")
	tempat_lahir := c.PostForm("tempat_lahir")
	kepala_keluarga := c.PostForm("kepala_keluarga")
	nomor_kk := c.PostForm("nomor_kk")
	tanggal_lahir := c.PostForm("tanggal_lahir")
	alamat := c.PostForm("alamat")
	alasan := c.PostForm("alasan")
	jenis_pengajuan := c.PostForm("jenis_pengajuan")
	keperluan := c.PostForm("keperluan")
	status_perkawinan := c.PostForm("status_perkawinan")
	jenis_usaha := c.PostForm("jenis_usaha")
	nama_anak := c.PostForm("nama_anak")
	nama_ibu := c.PostForm("nama_ibu")
	nama_ayah := c.PostForm("nama_ayah")
	nama_pengaju := c.PostForm("nama_pengaju")
	jenis_kelamin_anak := c.PostForm("jenis_kelamin_anak")
	hubungan_pengaju := c.PostForm("hubungan_pengaju")
	hari_meninggal := c.PostForm("hari_meninggal")
	benda_hilang := c.PostForm("benda_hilang")
	umur := c.PostForm("umur")

	var extFoto string
	var extKtp string
	var extTtd string

	if jenis_pengajuan == "1" {
		fileFoto, err := c.FormFile("foto")
		if err != nil {
			c.JSON(http.StatusBadGateway, domain.JsonResponse(http.StatusBadGateway, "Gagal upload file", domain.Empty{}))
			return
		}

		fileKtp, err := c.FormFile("ktp")
		if err != nil {
			c.JSON(http.StatusBadGateway, domain.JsonResponse(http.StatusBadGateway, "Gagal upload file", domain.Empty{}))
			return
		}

		if fileFoto.Size > maxFileSize || fileKtp.Size > maxFileSize {
			c.JSON(http.StatusBadRequest, domain.JsonResponse(http.StatusBadRequest, "Ukuran file max 2 MB", domain.Empty{}))
			return
		}

		if !isAllowedExtension(fileFoto.Filename, []string{".jpg", ".jpeg"}) || !isAllowedExtension(fileKtp.Filename, []string{".jpg", ".jpeg"}) {
			c.JSON(http.StatusBadRequest, domain.JsonResponse(http.StatusBadRequest, "Format file harus jpg atau jpeg", domain.Empty{}))
			return
		}

		extFoto = domain.GenerateUuid() + filepath.Ext(fileFoto.Filename)
		extKtp = domain.GenerateUuid() + filepath.Ext(fileKtp.Filename)

		// Save the uploaded file to a specific location (e.g., ./uploads)
		destination := filepath.Join("./uploads/users", extFoto)
		if err := c.SaveUploadedFile(fileFoto, destination); err != nil {
			c.JSON(http.StatusInternalServerError, domain.JsonResponse(http.StatusInternalServerError, "Failed to save file", domain.Empty{}))
			return
		}

		destinationKtp := filepath.Join("./uploads/users", extKtp)
		if err := c.SaveUploadedFile(fileKtp, destinationKtp); err != nil {
			c.JSON(http.StatusInternalServerError, domain.JsonResponse(http.StatusInternalServerError, "Failed to save file", domain.Empty{}))
			return
		}
	}

	if jenis_pengajuan == "2" || jenis_pengajuan == "6" || jenis_pengajuan == "5" || jenis_pengajuan == "7" {
		fileKtp, err := c.FormFile("ktp")
		if err != nil {
			c.JSON(http.StatusBadGateway, domain.JsonResponse(http.StatusBadGateway, "Gagal upload file", domain.Empty{}))
			return
		}

		if fileKtp.Size > maxFileSize {
			c.JSON(http.StatusBadRequest, domain.JsonResponse(http.StatusBadRequest, "Ukuran file max 2 MB", domain.Empty{}))
			return
		}

		if !isAllowedExtension(fileKtp.Filename, []string{".jpg", ".jpeg"}) {
			c.JSON(http.StatusBadRequest, domain.JsonResponse(http.StatusBadRequest, "Format file harus jpg atau jpeg", domain.Empty{}))
			return
		}

		extKtp = domain.GenerateUuid() + filepath.Ext(fileKtp.Filename)

		// Save the uploaded file to a specific location (e.g., ./uploads)
		destination := filepath.Join("./uploads/users", extKtp)
		if err := c.SaveUploadedFile(fileKtp, destination); err != nil {
			c.JSON(http.StatusInternalServerError, domain.JsonResponse(http.StatusInternalServerError, "Failed to save file", domain.Empty{}))
			return
		}
	}

	if jenis_pengajuan == "3" || jenis_pengajuan == "4" {
		fileFoto, err := c.FormFile("fileTtd")
		log.Println(err)
		if err != nil {
			c.JSON(http.StatusBadGateway, domain.JsonResponse(http.StatusBadGateway, "Gagal upload file", domain.Empty{}))
			return
		}

		if fileFoto.Size > maxFileSize {
			c.JSON(http.StatusBadRequest, domain.JsonResponse(http.StatusBadRequest, "Ukuran file max 2 MB", domain.Empty{}))
			return
		}

		if !isAllowedExtension(fileFoto.Filename, []string{".jpg", ".jpeg"}) {
			c.JSON(http.StatusBadRequest, domain.JsonResponse(http.StatusBadRequest, "Format file harus jpg atau jpeg", domain.Empty{}))
			return
		}

		extTtd = domain.GenerateUuid() + filepath.Ext(fileFoto.Filename)

		// Save the uploaded file to a specific location (e.g., ./uploads)
		destination := filepath.Join("./uploads/users", extTtd)
		if err := c.SaveUploadedFile(fileFoto, destination); err != nil {
			c.JSON(http.StatusInternalServerError, domain.JsonResponse(http.StatusInternalServerError, "Failed to save file", domain.Empty{}))
			return
		}
	}

	umurs := domain.StringToInt(umur)

	jsonData := domain.PengajuanParam{
		AuthId:           int64(auth_id),
		Nama:             nama,
		JenisKelamin:     handleEmpty(jenis_kelamin),
		Nik:              handleEmpty(nik),
		Phone:            handleEmpty(phone),
		Agama:            handleEmpty(agama),
		Pekerjaan:        handleEmpty(pekerjaan),
		TempatLahir:      handleEmpty(tempat_lahir),
		TanggalLahir:     handleEmpty(tanggal_lahir),
		KepalaKeluarga:   handleEmpty(kepala_keluarga),
		NomorKk:          handleEmpty(nomor_kk),
		Alamat:           handleEmpty(alamat),
		Alasan:           handleEmpty(alasan),
		JenisPengajuan:   int8(domain.StringToInt(jenis_pengajuan)),
		Keperluan:        handleEmpty(keperluan),
		StatusPerkawinan: handleEmpty(status_perkawinan),
		JenisUsaha:       handleEmpty(jenis_usaha),
		NamaAnak:         handleEmpty(nama_anak),
		NamaIbu:          handleEmpty(nama_ibu),
		NamaAyah:         handleEmpty(nama_ayah),
		NamaPengaju:      handleEmpty(nama_pengaju),
		JenisKelaminAnak: handleEmpty(jenis_kelamin_anak),
		HubunganPengaju:  handleEmpty(hubungan_pengaju),
		Status:           0,
		HariMeninggal:    handleEmpty(hari_meninggal),
		BendaHilang:      handleEmpty(benda_hilang),
		Umur:             &umurs,
	}

	jsonData.CreatedAt = domain.TimeNow()
	jsonData.Foto = handleEmpty(extFoto)
	jsonData.Ktp = handleEmpty(extKtp)
	jsonData.TandaTangan = handleEmpty(extTtd)

	if err := cx.Repo.CreatePengajuan(jsonData); err != nil {
		c.JSON(500, domain.JsonResponse(500, err.Error(), domain.Empty{}))
		return
	}

	c.JSON(201, domain.JsonResponse(201, "success", domain.Empty{}))
}

func isAllowedExtension(filename string, allowedExts []string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			return true
		}
	}
	return false
}

func handleEmpty(req string) *string {
	if req == "" {
		return nil
	}

	return &req
}
