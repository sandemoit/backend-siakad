// @title           API SIAKAD PONPES
// @version         1.0
// @description     API Kebutuhan untuk SIAKAD PONPES
// @host            localhost:3000
// @BasePath        /api/v1
package main

import (
	"os"
	"siakad/config"
	"siakad/middleware"
	"siakad/routes"
	"siakad/utils"
)

func main() {
	// laod env
	utils.LoadEnv()

	// Koneksi ke database
	db := config.InitDB()
	defer config.CloseDB(db)

	// Inisialisasi Fiber
	app := middleware.CreateApp()

	// Setup routes
	routes.SetupRoutes(app)

	// Jalankan server
	app.Listen(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}
