package database

import (
	"context"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

func GetFirestoreConn(projectID string) (*firestore.Client, error) {
	// Use the application default credentials
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	return client, nil
}
