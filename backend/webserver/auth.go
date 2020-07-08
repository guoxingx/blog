package webserver

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// Authorizer ...
type Authorizer struct {
	raws    map[string]string    // map[raw]cookie
	cookies map[string]string    // map[cookie]raw
	expires map[string]time.Time // map[cookie]expire
	mu      sync.Mutex
	salt    string
	expire  time.Duration
	cname   string
	format  string
	users   map[string]string
}

type cookie struct {
	c string
	t time.Time
}

// NewAuthorizer ...
func NewAuthorizer(salt string, users map[string]string) (*Authorizer, error) {
	a := Authorizer{
		raws:    make(map[string]string, 0),
		cookies: make(map[string]string, 0),
		expires: make(map[string]time.Time, 0),
		salt:    salt,
		expire:  time.Hour * 48,
		cname:   "darktemplar",
		format:  "%smiss%slester%s",
		users:   users,
	}
	return &a, nil
}

// Check ...
func (a *Authorizer) Check(req *http.Request) bool {
	c, err := req.Cookie(a.cname)
	if err != nil {
		return false
	}

	cookie := c.String()
	cookie = cookie[12:]
	raw, ok := a.cookies[cookie]
	if !ok {
		a.mu.Lock()
		defer a.mu.Unlock()
		delete(a.raws, raw)
		return false
	}

	t, ok := a.expires[cookie]
	if !ok {
		a.mu.Lock()
		defer a.mu.Unlock()
		delete(a.raws, raw)
		delete(a.cookies, cookie)
		return false
	}
	if time.Now().After(t) {
		a.mu.Lock()
		defer a.mu.Unlock()
		delete(a.raws, raw)
		delete(a.cookies, cookie)
		delete(a.expires, cookie)
		return false
	}

	return true
}

// CheckAndResponse ...
func (a *Authorizer) CheckAndResponse(w http.ResponseWriter, req *http.Request) (bool, int, interface{}) {
	ok := a.Check(req)
	if !ok {
		return false, 403, []byte("cookie invalid or expired")
	}
	return true, 0, nil
}

// GenerateAndResponse ...
func (a *Authorizer) GenerateAndResponse(raw string, w http.ResponseWriter) string {
	cookie := a.generate(raw)
	// fmt.Println(raw, cookie)

	a.mu.Lock()
	defer a.mu.Unlock()

	oldCookie, ok := a.raws[raw]
	if ok {
		delete(a.cookies, oldCookie)
		delete(a.expires, oldCookie)
	}
	a.raws[raw] = cookie
	a.cookies[cookie] = raw
	a.expires[cookie] = time.Now().Add(a.expire)

	w.Header().Add("Set-Cookie", cookie)
	return cookie
}

// Remove ...
func (a *Authorizer) Remove(req *http.Request) {
	c, err := req.Cookie(a.cname)
	if err != nil {
		return
	}

	cookie := c.String()
	raw, ok := a.cookies[cookie]
	if ok {
		a.mu.Lock()
		defer a.mu.Unlock()
		delete(a.raws, raw)
		delete(a.cookies, cookie)
		delete(a.expires, cookie)
	}
}

func (a *Authorizer) generate(raw string) string {
	w := md5.New()
	rand := fmt.Sprintf("%d", time.Now().Second()%(time.Now().Minute()+1))
	io.WriteString(w, fmt.Sprintf(a.format, raw, a.salt, rand))
	return fmt.Sprintf("%x", w.Sum(nil))
}

func signPassword(username, password string) string {
	w := md5.New()
	io.WriteString(w, fmt.Sprintf("%s{%s}", password, username))
	return fmt.Sprintf("%x", w.Sum(nil))
}
