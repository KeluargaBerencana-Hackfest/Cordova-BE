package firebase

import (
	"context"
	"fmt"
	"log"
	"os"

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

func (f *FirebaseClient) UploadFile(ctx context.Context, file []byte, fileName string) (string, error) {
	firebaseStorage, err := f.Storage(context.Background())
	if err != nil {
		return "", err
	}

	bucket, err := firebaseStorage.Bucket(os.Getenv("FIREBASE_BUCKET"))
	if err != nil {
		return "", err
	}

	wc := bucket.Object(fileName).NewWriter(ctx)
	if _, err = wc.Write(file); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	_, err = bucket.Object(fileName).Attrs(ctx)
	if err != nil {
		return "", err
	}

	link := fmt.Sprintf(os.Getenv("FIREBASE_IMAGE_URL"), fileName)
	return link, nil
}
