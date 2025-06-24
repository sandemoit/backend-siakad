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
