package main

import (
	"github.com/gorilla/sessions"
)

// Create a new cookie store for storing session data securely.
var store = sessions.NewCookieStore([]byte("secret-key"))
