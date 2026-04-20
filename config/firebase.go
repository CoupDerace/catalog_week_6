package config

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var FirebaseAuth *auth.Client

func InitFirebase() {
	credPath := os.Getenv("FIREBASE_CREDENTIALS_PATH")

	if credPath == "" {
		log.Fatal("FIREBASE_CREDENTIALS_PATH belum di-set")
	}

	opt := option.WithCredentialsFile(credPath)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Gagal init Firebase App: %v", err)
	}

	FirebaseAuth, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Gagal Firebase Auth Client: %v", err)
	}

	log.Println("Firebase berhasil diinisialisasi")
}