package main

import (
	"github.com/ymichael/sessions"
	"github.com/zenazn/goji/web"
)

type sessionData struct {
	Username string

	// File uploads persist across requests.
	UploadedName string
	UploadedSize string
}

func initSessions(secret string) *sessions.SessionOptions {
	return sessions.NewSessionOptions(secret, &sessions.MemoryStore{})
}

func getSessionData(c web.C) sessionData {
	sd := sessionManager.GetSessionObject(&c)["s"]
	if sd == nil {
		return sessionData{}
	}

	return sd.(sessionData)
}
