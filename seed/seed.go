package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/CoupDerace/catalog_week_6/config"
	"github.com/CoupDerace/catalog_week_6/models"
)

func main() {
	godotenv.Load()
	config.InitDatabase()

	products := []models.Product{
		{
			Name:        "Helm Bogo Polos",
			Price:       35000,
			Category:    "Helm",
			Stock:       25,
			Description: "Helm retro klasik dengan gaya unik",
			ImageURL:    "https://i.ibb.co.com/Q7mky5XG/Helm-Carglos-Polos-Retro-Bogo-Dewasa-Pria-Wanita-SNI-kekinian-Terbaru.jpg",
		},
		{
			Name:        "Helm Half Face",
			Price:       500000,
			Category:    "Helm",
			Stock:       40,
			Description: "Helm half face nyaman untuk penggunaan harian",
			ImageURL:    "https://i.ibb.co.com/bRBCQyG3/ORIGINAL-HELM-NJS-KAIROZ-HALF-FACE-FREE-SPOILER.jpg",
		},
		{
			Name:        "Helm Full Face",
			Price:       1000000,
			Category:    "Helm",
			Stock:       30,
			Description: "Helm full face dengan desain sporty dan aman",
			ImageURL:    "https://i.ibb.co.com/fdFH7DDx/Casque-int-gral-KYT-R2-R-PLAIN-Gris-Gris.jpg",
		},
		{
			Name:        "Helm Cross Pro Taper",
			Price:       450000,
			Category:    "Helm",
			Stock:       20,
			Description: "Helm modular fleksibel bisa dibuka tutup",
			ImageURL:    "https://i.ibb.co.com/KpXdKBT9/FXR-Motocross-Helm-6-D-ATR-3-FIM-Wei-Schwarz-Maciag-Offroad.jpg",
		},
	}

	for _, p := range products {
		config.DB.Create(&p)
	}

	log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}