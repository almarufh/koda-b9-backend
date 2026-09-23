package main

import (
	"fmt"
	"log"
	"os"

	"github.com/almarufh/koda-b9-backend/internal/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Println("ENV :", os.Getenv("PORT"))
	fmt.Println("FAJAR? :", os.Getenv("CEK"))

	r := gin.Default()
	routers.Routers(r)
	port := os.Getenv("PORT")
	log.Println("\033[32m" + fmt.Sprintf("Server berhasil berjalan di port %s", port) + "\033[0m")
	r.Run(fmt.Sprintf(":%s", port))
}

// package main

// import (
// 	"log"
// 	"os"
// )

// // Definisi kode warna ANSI
// const (
// 	colorReset  = "\033[0m"
// 	colorRed    = "\033[31m"
// 	colorGreen  = "\033[32m"
// 	colorYellow = "\033[33m"
// 	colorCyan   = "\033[36m"
// )

// func main() {
// 	// 1. Mewarnai pesan langsung secara manual
// 	log.Println(colorGreen + "Server berhasil berjalan di port 8080" + colorReset)
// 	log.Printf("%sPeringatan:%s Kuota memori tersisa %d%%%s\n", colorYellow, colorReset, 15, colorReset)

// 	// 2. Membuat Logger kustom dengan Prefix berwarna
// 	infoLog := log.New(os.Stdout, colorCyan+"[INFO] "+colorReset, log.Ldate|log.Ltime)
// 	warnLog := log.New(os.Stdout, colorYellow+"[WARN] "+colorReset, log.Ldate|log.Ltime)
// 	errorLog := log.New(os.Stderr, colorRed+"[ERROR] "+colorReset, log.Ldate|log.Ltime|log.Lshortfile)

// 	infoLog.Println("Menerima request GET /users")
// 	warnLog.Println("Response time lambat: 1200ms")
// 	errorLog.Println("Gagal terhubung ke database")
// }
