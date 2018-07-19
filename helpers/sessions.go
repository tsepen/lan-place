package helpers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("FUIylxerKLwnZ5MiTE1boAAhBRSJI8Qg")
	store = sessions.NewCookieStore(key)
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateSession(w http.ResponseWriter, r *http.Request, id int) {
	session, _ := store.Get(r, "lan-place-session")
	fmt.Println(id)
	session.Values["authenticated"] = true
	session.Values["id"] = id
	session.Save(r, w)
}

func GetSession(w http.ResponseWriter, r *http.Request) *User {
	session, _ := store.Get(r, "lan-place-session")
	u := new(User)

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return u
	}

	db := Db
	id := session.Values["id"]

	row := db.QueryRow("SELECT * FROM users WHERE id=$1;", id)
	err := row.Scan(&u.ID, &u.Email, &u.Name, &u.Password)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
	}

	return u
}

func DestroySession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "lan-place-session")

	session.Values["authenticated"] = false
	session.Save(r, w)
}
