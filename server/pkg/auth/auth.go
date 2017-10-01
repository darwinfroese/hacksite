package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/database"
)

const sessionTokenSize = 64

// SessionMaxAge is the cookie length in seconds
const SessionMaxAge = 900

// SessionCookieName is the name of the cookie
const SessionCookieName = "HacksiteSession"

// SaltPassword generates a salt, puts it into the password and returns
// the salt and the new password or an error
func SaltPassword(password string) (string, string, error) {
	salt := make([]byte, sha256.Size)

	n, err := rand.Read(salt)

	if err != nil {
		return "", "", err
	}

	if n != sha256.Size {
		return "", "", err
	}

	saltedVal := append([]byte(password), salt...)
	encrypted := sha256.Sum256(saltedVal)

	hashStr := base64.StdEncoding.EncodeToString(encrypted[:sha256.Size])
	saltStr := base64.StdEncoding.EncodeToString(salt)

	return saltStr, hashStr, nil
}

// GetSaltedPassword returns the salted version of the password using the
// account's salt
func GetSaltedPassword(password string, salt string) (string, error) {
	s, err := base64.StdEncoding.DecodeString(salt)

	if err != nil {
		return "", err
	}

	saltedVal := append([]byte(password), s...)
	encrypted := sha256.Sum256(saltedVal)

	hashStr := base64.StdEncoding.EncodeToString(encrypted[:sha256.Size])

	return hashStr, nil
}

// CreateSession returns the session token to store in the cookie
func CreateSession(id int) models.Session {
	session := CreateSessionToken()

	return models.Session{
		Token:      session,
		UserID:     id,
		Expiration: time.Now().Add(time.Second * time.Duration(SessionMaxAge)),
	}
}

// CreateSessionToken generates a random session token
func CreateSessionToken() string {
	sesh := make([]byte, sessionTokenSize)

	_, _ = rand.Read(sesh)

	return base64.StdEncoding.EncodeToString(sesh)
}

// SetCookie creates an http cookie and sets it in the response
func SetCookie(w http.ResponseWriter, name, token string) {
	// TODO: Implement remember me functionality (MaxAge: 0)
	http.SetCookie(w, &http.Cookie{
		Name:   name,
		Value:  token,
		MaxAge: SessionMaxAge,
		// TODO: set secure when supporting HTTPS
	})
}

// GetCurrentSession reads the session cookie and grabs the session associated
func GetCurrentSession(r *http.Request, db database.Database) (models.Session, error) {
	cookie, err := r.Cookie(SessionCookieName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return models.Session{}, err
	}

	return db.GetSession(cookie.Value)
}

// GetCurrentUser will use the session model passed in to find the signed in users value
func GetCurrentUser(db database.Database, session models.Session) (models.Account, error) {
	return db.GetAccountByID(session.UserID)
}
