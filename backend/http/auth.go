package http

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

// registerAuthRoutes is a helper function to register auth routes to a router.
func (s *Server) registerAuthRoutes(r *mux.Router) {
	r.HandleFunc("/login", s.handleLogin).Methods("GET")
}

// handleLogin handles the "GET /login" route. It simply renders an HTML login form.
func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	// Load and parse the HTML template.
	tmpl := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Login with Gmail</title>
			<script src="https://apis.google.com/js/platform.js" async defer></script>
			<meta name="google-signin-client_id" content="YOUR_CLIENT_ID.apps.googleusercontent.com">
		</head>
		<body>
			<h1>Login with Gmail</h1>
			<div class="g-signin2" data-onsuccess="onSignIn"></div>
			<script>
				function onSignIn(googleUser) {
					// Handle the signed-in user here (e.g., send user info to your server).
					var profile = googleUser.getBasicProfile();
					console.log("ID: " + profile.getId()); // Don't send this directly to your server!
					console.log("Name: " + profile.getName());
					console.log("Email: " + profile.getEmail()); // This should be sent to your server.
				}
			</script>
		</body>
		</html>
	`
	t, err := template.New("login").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template and write it to the response.
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
