package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func SetupFirebase() *firebase.App {
	opt := option.WithCredentialsFile("./firebaseServiceKey.json")
	fb, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln("Error initializing Firebase App:", err)
	}
	return fb
}
