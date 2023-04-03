package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./jajanku-cb527-firebase-adminsdk-88zdp-175d3509ca.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		fmt.Printf("error initializing app: %v\n", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		fmt.Printf("error getting Auth client: %v\n", err)
	}

	fmt.Printf("this client: %v\n", client)

	userIdToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6Ijg3YzFlN2Y4MDAzNGJiYzgxYjhmMmRiODM3OTIxZjRiZDI4N2YxZGYiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vamFqYW5rdS1jYjUyNyIsImF1ZCI6ImphamFua3UtY2I1MjciLCJhdXRoX3RpbWUiOjE2ODA1NDcxMDMsInVzZXJfaWQiOiJiNEdNcFg0YU5OVEw2MUZHQnNLM3k1Q2FnbmQyIiwic3ViIjoiYjRHTXBYNGFOTlRMNjFGR0JzSzN5NUNhZ25kMiIsImlhdCI6MTY4MDU0NzEwMywiZXhwIjoxNjgwNTUwNzAzLCJlbWFpbCI6ImdpbGx5QGV4YW1wbGUuY29tIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7ImVtYWlsIjpbImdpbGx5QGV4YW1wbGUuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoicGFzc3dvcmQifX0.mdhRVCTBtUilzF39bc8aSYZLgAdFLls3nhMFQW-5dVPflZkgH1P5hC2l0IjYQWNLAmjvtqA9ymZc547m53s9lGee-7GNUBOU2GaoUK7MUsmDT1_7bIpUUZzIVB7tRlmhwo_XKfB5mVf0nFd5RuH15czLP7SinNvOGp5B_GN8UJDIjsqbSSA4RK-Ig4a05n5dAqeWgSjdJU_zDsqjs8jvkxRTPIl5iWzd23szpkEbPpWddms8qxd3mYQBJWQNUZ5Oy1gC-2VBgA8_MM7lzpaxaVwi33FNFXIV7aUjn7g0U4aBu70QYtOW6ZTZ-Y-1kRJZXMFZYjcOk0az2eFKcP1ojA"

	user, err := client.VerifyIDToken(ctx, userIdToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	uid := user.UID
	fmt.Printf("this uid: %v\n", uid)

	_user, err := client.GetUser(ctx, uid)
	fmt.Printf("this email: %v\n", _user.Email)

	// Use the client to verify user tokens and perform other authentication tasks
}
