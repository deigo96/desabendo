package router

import (
	"desa-sragen/api/controller"
	"desa-sragen/api/middleware"
	"desa-sragen/bootrstrap"
	"desa-sragen/repository"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootrstrap.Env, db bootrstrap.Databases, router *gin.RouterGroup) {
	repo := repository.NewRepository(db.Db)
	cx := &controller.Controller{
		Config: env,
		Repo:   repo,
	}

	// adminRoute := router.Group("admin")
	publicRoute := router.Group("")
	protectedRoute := router.Group("").Use(
		middleware.ValidateJwt(env.SecretKey),
	)

	tokenRoute := router.Group("").Use(
		middleware.ValidateToken(env.SecretKey),
	)

	userRoute := router.Group("").Use(
		middleware.ValidateJwt(env.SecretKey),
		middleware.ValidateUser(),
	)

	adminRoute := router.Group("").Use(
		middleware.ValidateJwt(env.SecretKey),
		middleware.ValidateAdmin(),
	)

	superRoute := router.Group("").Use(
		middleware.ValidateJwt(env.SecretKey),
		middleware.ValidateSuper(),
	)

	userRoute.GET("validate-token", cx.ValidateToken)
	adminRoute.GET("validate-token-admin", cx.ValidateTokenAdmin)
	superRoute.GET("validate-token-super", cx.ValidateTokenSuper)
	publicRoute.GET("/", cx.Dashboard)
	publicRoute.GET("/login", cx.Login)
	publicRoute.GET("/register", cx.Register)
	publicRoute.POST("/login-handler", cx.LoginHandler)
	publicRoute.POST("/register-handler", cx.RegisterHandler)
	tokenRoute.GET("/success-login", cx.SuccessLogin)

	publicRoute.GET("/dashboard", cx.AdminDashboard)
	publicRoute.GET("/daftar-admin", cx.PageAdmin)
	publicRoute.GET("/daftar-user", cx.ListUsers)
	publicRoute.GET("/ganti-password", cx.PagePassword)
	publicRoute.GET("/change-password", cx.PagePasswordUser)
	adminRoute.POST("/change-admin-password", cx.ChangePassword)
	userRoute.POST("/change-user-password", cx.ChangePassword)
	publicRoute.GET("/detail-pengajuan/:id", cx.DetailPengajuan)
	protectedRoute.GET("redirect/:type", cx.RedirectUrl)
	adminRoute.POST("/add-teams", cx.AddTeams)
	publicRoute.GET("/teams", cx.GetTeams)

	publicRoute.GET("/pengajuan-surat-domisili", cx.PengajuanDomisili)
	publicRoute.GET("/pengajuan-surat-kehilangan", cx.PengajuanKehilangan)
	publicRoute.GET("/pengajuan-surat-kelahiran", cx.PengajuanKelahiran)
	publicRoute.GET("/pengajuan-surat-kematian", cx.PengajuanKematian)
	publicRoute.GET("/pengajuan-kartu-keluarga", cx.PengajuanKk)
	publicRoute.GET("/pengajuan-kartu-tanda-penduduk", cx.PengajuanKtp)
	publicRoute.GET("/pengajuan-surat-usaha", cx.PengajuanSku)
	protectedRoute.GET("/list-dokumen", cx.ListDokumen)
	protectedRoute.POST("/kirim-pengajuan", cx.KirimPengajuan)

	// protected
	protectedRoute.GET("asd", cx.GetAuth)

	// super
	protectedRoute.GET("get-admin", cx.GetAllAdmin)
	superRoute.POST("tambah-admin", cx.TambahAdmin)
	superRoute.POST("aksi-dokumen", cx.AksiDokumen)
	superRoute.POST("in-active-admin/:id", cx.InActiveAdmin)
	superRoute.POST("activation-admin/:id", cx.ActivationAdmin)
	publicRoute.GET("/surat-domisili", cx.SuratDomisili)
	publicRoute.GET("/surat-kartu-keluarga", cx.SuratKartuKeluarga)
	publicRoute.GET("/surat-keterangan-usaha", cx.SuratKeteranganUsaha)
	publicRoute.GET("/surat-kelahiran", cx.SuratKelahiran)
	publicRoute.GET("/surat-kematian", cx.SuratKematian)
	publicRoute.GET("/surat-kehilangan", cx.SuratKehilangan)
	publicRoute.GET("/surat-permohonan-ktp", cx.SuratKtp)


	publicRoute.POST("/send-email", cx.SendEmail)
}
