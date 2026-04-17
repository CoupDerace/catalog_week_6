package config

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// FirebaseApp adalah instance Firebase global yang dipakai di seluruh aplikasi
var FirebaseAuth *auth.Client

func InitFirebase() {
	credPath : os.Getenv("FIREBASE_CREDENTIALS_PATH")

	// Inisialisasi Firebase App dengan kredensial
	opt := option.WithCredentialsFile(credPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Gagal inisialisasi Firebase App: %v", err)
	}

	// Dapatkan Firebase Auth Client
	FirebaseAuth, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Gagal mendapatkan Firebase Auth Client: %v", err)
	}
	log.Println("Firebase Auth Client berhasil diinisialisasi")
}