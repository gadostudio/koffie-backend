package settings

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"os"
)

func SetupFirebase(ctx context.Context) *auth.Client {

	app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsJSON([]byte(os.Getenv("JSON_CREDS"))))

	if err != nil {
		panic("Firebase load error")
	}

	client, err := app.Auth(ctx)
	if err != nil {
		panic("Firebase load error")
	}
	return client
}
