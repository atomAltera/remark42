package rest

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"hash/crc64"
	"log"
	"net/http"

	"github.com/umputun/remark/app/store"
)

// ContextKey is a type to match on context
type ContextKey string

// GetUserInfo returns user from request context
func GetUserInfo(r *http.Request) (user store.User, err error) {

	ctx := r.Context()
	if ctx == nil {
		return store.User{}, errors.New("no info about user")
	}

	if u, ok := ctx.Value(ContextKey("user")).(store.User); ok {
		return u, nil
	}

	return store.User{}, errors.New("user can't be parsed")
}

// EncodeID hashes user id to sha1
func EncodeID(id string) string {
	h := sha1.New()
	if _, err := h.Write([]byte(id)); err != nil {
		// fail back to crc64
		log.Printf("[WARN] can't hash id %s, %s", id, err)
		return fmt.Sprintf("%x", crc64.Checksum([]byte(id), crc64.MakeTable(crc64.ECMA)))
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}