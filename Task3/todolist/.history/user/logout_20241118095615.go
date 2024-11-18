package user

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("秘密密钥"))

	session, _ := store.Get(r, "session")
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
