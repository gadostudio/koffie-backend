package settings

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"path/filepath"
)

func SetupFirebase(ctx context.Context) *auth.Client {
	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		panic("Firebase load error")
	}

	client, err := app.Auth(ctx)
	if err != nil {
		panic("Firebase load error")
	}
	return client
}
