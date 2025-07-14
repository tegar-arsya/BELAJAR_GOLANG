package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"portfolio-backend/config"
	"portfolio-backend/internal/handler"
	"portfolio-backend/internal/middleware"
	"portfolio-backend/internal/repository"
	"portfolio-backend/internal/routes"
	"portfolio-backend/internal/service"
	"portfolio-backend/migrations"

	"strings"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 100 << 20

	// ✅ Middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://dashboard.tegararsyadani.my.id", "https://tegararsyadani.my.id"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Static file cache headers (untuk CDN)
	r.Use(func(c *gin.Context) {
		if c.Request.Method == "GET" &&
			(strings.HasPrefix(c.Request.RequestURI, "/porto/") ||
				strings.HasPrefix(c.Request.RequestURI, "/image/") ||
				strings.HasPrefix(c.Request.RequestURI, "/journey/") ||
                strings.HasPrefix(c.Request.RequestURI, "/about/") ||
                strings.HasPrefix(c.Request.RequestURI, "/file/") ||
				strings.HasPrefix(c.Request.RequestURI, "/public/")) {
			c.Header("Cache-Control", "public, max-age=86400")
		}
		c.Next()
	})

	// ✅ DB dan Migrate
	config.ConnectDatabase()
	migrations.Migrate()

    r.Static("/image", "./public/image")
    r.Static("/file", "./public/file")
    r.Static("/serti", "./public/serti")
    r.Static("/journey", "./public/journey")
    r.Static("/porto", "./public/porto")
    r.Static("/me", "./public/me")
    r.Static("/article", "./public/article")
r.Static("/article-content", "./public/article-content")

	// ✅ Inisialisasi langsung tanpa NewFunction
	aboutHandler := &handler.AboutHandler{
		Service: &service.AboutService{
			Repo: &repository.AboutRepository{DB: config.DB},
		},
	}
	cvHandler := &handler.CvHandler{
		Service: &service.CvService{
			Repo: &repository.CvRepository{DB: config.DB},
		},
	}
	portfolioHandler := &handler.PortofolioHandler{
		Service: &service.PortofolioService{
			Repo: &repository.PortofolioRepository{DB: config.DB},
		},
	}
	sertifikatHandler := &handler.SertifikatHandler{
		Service: &service.SertifikatService{
			Repo: &repository.SertifikatRepository{DB: config.DB},
		},
	}
	pengalamanHandler := &handler.PengalamanHandler{
		Service: &service.PengalamanService{
			Repo: &repository.PengalamanRepository{DB: config.DB},
		},
	}
	authHandler := &handler.AuthHandler{
		Service: &service.AuthService{
			Repo: &repository.AuthRepository{DB: config.DB},
            JwtKey: config.JwtKey,
		},
	}

    articleHandler := &handler.ArticleHandler{
    Service: &service.ArticleService{
        Repo: &repository.ArticleRepository{DB: config.DB},
    },
}


	// ✅ Auth routes
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.POST("/logout", middleware.AuthMiddleware(), authHandler.Logout)

	// ✅ Public dan Admin routes
	public := r.Group("/api/public")
	admin := r.Group("/api/admin", middleware.AuthMiddleware())

	routes.RegisterAboutRoutes(public, admin, aboutHandler)
	routes.RegisterCvRoutes(public, admin, cvHandler)
	routes.RegisterPortfolioRoutes(public, admin, portfolioHandler)
	routes.RegisterSertifikatRoutes(public, admin, sertifikatHandler)
	routes.RegisterPengalamanRoutes(public, admin, pengalamanHandler)
routes.RegisterArticleRoutes(public, admin, articleHandler)
	// ✅ Jalankan server
	r.Run("0.0.0.0:8181")
//    r.Run("localhost:9191")        // ✅ listen on 127.0.0.1:9191

}
