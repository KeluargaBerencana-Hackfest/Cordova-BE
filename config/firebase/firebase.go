package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type FirebaseClient struct {
	*firebase.App
}

func InitFirebase() (*FirebaseClient, error) {
	opt := option.WithCredentialsFile("./firebase-account-key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return &FirebaseClient{app}, nil
}

func (f *FirebaseClient) AuthFirebase() *auth.Client {
	client, err := f.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	return client
}
