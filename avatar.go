package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ErrNoAvatarURL is an error raised when Avatar instance can't return the avatar URL.
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません")

// Avatar is a type which represents a profile image of users.
type Avatar interface {
	// URL() returns the avatar URL of the specified client
	// It returns an error when something goes wrong.
	// Especially it does ErrNoAvatarURL when it can't fetch the avatar URL.
	URL(ChatUser) (string, error)
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
	if userID, ok := c.userData["user_id"]; ok {
		if userIDStr, ok := userID.(string); ok {
			return "//www.gravatar.com/avatar/" + userIDStr, nil
		}
	}

	return "", ErrNoAvatarURL
}

type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

func (FileSystemAvatar) URL(c *client) (string, error) {
	if userID, ok := c.userData["user_id"]; ok {
		if userIDStr, ok := userID.(string); ok {
			if files, err := ioutil.ReadDir("avatars"); err == nil {
				for _, file := range files {
					if file.IsDir() {
						continue
					}

					if match, _ := filepath.Match(userIDStr+".*", file.Name()); match {
						return "/avatars/" + file.Name(), nil
					}
				}
			}

		}
	}

	return "", ErrNoAvatarURL
}
