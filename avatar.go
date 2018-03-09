package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"strings"
)

// ErrNoAvatarURL is an error raised when Avatar instance can't return the avatar URL.
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません")

// Avatar is a type which represents a profile image of users.
type Avatar interface {
	// URL() returns the avatar URL of the specified client
	// It returns an error when something goes wrong.
	// Especially it does ErrNoAvatarURL when it can't fetch the avatar URL.
	URL(c *client) (string, error)
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

func (AuthAvatar) URL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}

	return "", ErrNoAvatarURL
}

type GravatarAvatar struct{}

var UseGravatar GravatarAvatar

func (GravatarAvatar) URL(c *client) (string, error) {
	if email, ok := c.userData["email"]; ok {
		if emailStr, ok := email.(string); ok {
			m := md5.New()
			io.WriteString(m, strings.ToLower(emailStr))
			return fmt.Sprintf("//www.gravatar.com/avatar/%x", m.Sum(nil)), nil
		}
	}

	return "", ErrNoAvatarURL
}
