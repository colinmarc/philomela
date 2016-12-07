package main

import (
	"html/template"
	"net/http"
	"fmt"
	"log"
	"path/filepath"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/mrjones/oauth"
)

var (
	indexTemplate *template.Template = loadTemplate("index.html")
	userpageTemplate *template.Template = loadTemplate("userpage.html")
	landingTemplate *template.Template = loadTemplate("landing.html")

	secrets = loadSecrets()
	sessionManager = initSessions(secrets.SessionKey)
	oauthConsumer = oauth.NewConsumer(
		secrets.TwitterAuth.Key,
		secrets.TwitterAuth.Secret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		},
	)
)

func loadTemplate(name string) *template.Template {
	t, err := template.ParseFiles("templates/base.html", filepath.Join("templates", name))
	if err != nil {
		log.Fatal(err)
	}

	return t
}

func main() {
	goji.Use(sessionManager.Middleware())

	static := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	goji.Get("/static/*", static)
	goji.Get("/favicon.ico", func(c web.C, w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/favicon.ico")
	})

	goji.Get("/auth/twitter", twitterAuth)
	goji.Get("/auth/twitter/callback", twitterAuthCallback)
	goji.Get("/logout", logout)

	goji.Get("/", index)
	goji.Post("/upload", upload)
	goji.Post("/publish", publish)

	goji.Get("/:user", userpage)
	goji.Get("/:user/:slug", landing)

	goji.Serve()
}


func index(c web.C, w http.ResponseWriter, r *http.Request) {
	data := struct {
		Session sessionData
		Error string
	}{
		getSessionData(c),
		"",
	}

	err := indexTemplate.Execute(w, data)
	if err != nil {
		log.Println("Error rendering template:", err)
		w.WriteHeader(500)
	}
}

func userpage(c web.C, w http.ResponseWriter, r *http.Request) {
}

func landing(c web.C, w http.ResponseWriter, r *http.Request) {
}

func upload(c web.C, w http.ResponseWriter, r *http.Request) {
}

func publish(c web.C, w http.ResponseWriter, r *http.Request) {
}

func twitterAuth(c web.C, w http.ResponseWriter, r *http.Request) {
	log.Fatal("secret", secrets.TwitterAuth.Key)

	tokenUrl := fmt.Sprintf("%s://%s/twitter/auth/callback", r.URL.Scheme, r.Host)
	_, requestUrl, err := oauthConsumer.GetRequestTokenAndUrl(tokenUrl)
	if err != nil {
		log.Println("Error starting oauth:", err)
		w.WriteHeader(500)
		return
	}

	http.Redirect(w, r, requestUrl, http.StatusTemporaryRedirect)
}

func twitterAuthCallback(c web.C, w http.ResponseWriter, r *http.Request) {
	log.Println(r)
}

func logout(c web.C, w http.ResponseWriter, r *http.Request) {
}
