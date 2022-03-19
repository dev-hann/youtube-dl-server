package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"github.com/youtube-dl-server/ngrok"
	"google.golang.org/api/option"
	"log"
)

type Firebase struct {
	CredentialPath string
	Ctx            context.Context
	client         *firestore.Client
}

func NewFirebase(tokenPath string) *Firebase {
	f := Firebase{
		Ctx:            context.Background(),
		CredentialPath: tokenPath,
	}
	f.init()
	return &f
}

func (f *Firebase) init() {
	opt := option.WithCredentialsFile(f.CredentialPath)
	//var app *firebase.App
	app, err := firebase.NewApp(f.Ctx, nil, opt)
	if err != nil {
		log.Panicln(err)
	}

	f.client, err = app.Firestore(f.Ctx)
	if err != nil {
		log.Panicln(err)
	}
}

func (f *Firebase) UpdateNgrok(data *ngrok.Ngrok) {
	//var res *firestore.WriteResult
	res, err := f.client.Collection("youtube-dl").Doc("server").Set(f.Ctx, data)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(res.UpdateTime.String() + " =>  update on FireStore")
}
