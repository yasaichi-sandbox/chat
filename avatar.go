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

type TryAvatars []Avatar

func (a TryAvatars) URL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.URL(u); err == nil {
			return url, nil
		}
	}

	return "", ErrNoAvatarURL
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

func (AuthAvatar) URL(u ChatUser) (string, error) {
	if url := u.AvatarURL(); url != "" {
		return url, nil
	}

	return "", ErrNoAvatarURL
}

type GravatarAvatar struct{}

var UseGravatar GravatarAvatar

func (GravatarAvatar) URL(u ChatUser) (string, error) {
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}

type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

func (FileSystemAvatar) URL(u ChatUser) (string, error) {
	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}

			if match, _ := filepath.Match(u.UniqueID()+".*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}

	return "", ErrNoAvatarURL
}
