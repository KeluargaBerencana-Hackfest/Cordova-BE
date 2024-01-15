package firebase

import (
	"context"
	"encoding/base64"
	"encoding/json"
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
	serviceAccountCredentials := map[string]interface{}{
		"type":                        os.Getenv("FIREBASE_TYPE"),
		"project_id":                  os.Getenv("FIREBASE_PROJECT_ID"),
		"private_key_id":              os.Getenv("FIREBASE_PRIVATE_KEY_ID"),
		"private_key":                 base64.StdEncoding.EncodeToString([]byte(os.Getenv("FIREBASE_PRIVATE_KEY"))),
		"client_email":                os.Getenv("FIREBASE_CLIENT_EMAIL"),
		"client_id":                   os.Getenv("FIREBASE_CLIENT_ID"),
		"auth_uri":                    os.Getenv("FIREBASE_AUTH_URI"),
		"token_uri":                   os.Getenv("FIREBASE_TOKEN_URI"),
		"auth_provider_x509_cert_url": os.Getenv("FIREBASE_AUTH_PROVIDER_x509_CERT_URL"),
		"client_x509_cert_url":        os.Getenv("FIREBASE_CLIENT_x509_CERT_URL"),
		"universe_domain":             os.Getenv("FIREBASE_UNIVERSE_DOMAIN"),
	}

	credsJson, err := json.Marshal(serviceAccountCredentials)
	if err != nil {
		return nil, err
	}

	opt := option.WithCredentialsJSON(credsJson)
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
