package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/abbot/go-http-auth"
)

func main() {
	opts, err := ParseArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	secret := func(user, realm string) string {
		if user == opts.User {
			return string(auth.MD5Crypt([]byte(opts.Pass), []byte(RandStringRunes(10)), []byte("$1$")))
		}
		return ""
	}
	handler := http.FileServer(http.Dir(opts.Path))
	authHandler := func(res http.ResponseWriter, req *auth.AuthenticatedRequest) {
		handler.ServeHTTP(res, &req.Request)
	}
	authenticator := auth.NewBasicAuthenticator(opts.Realm, secret)
	log.Fatal(http.ListenAndServe(opts.Listen, authenticator.Wrap(authHandler)))
}
