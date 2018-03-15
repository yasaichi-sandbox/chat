package main

import (
	gomniauthtest "github.com/stretchr/gomniauth/test"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar

	testUser := &gomniauthtest.TestUser{}
	testUser.On("AvatarURL").Return("", ErrNoAvatarURL)
	testChatUser := &chatUser{User: testUser}
	url, err := authAvatar.URL(testChatUser)

	if err != ErrNoAvatarURL {
		t.Error("値が存在しない場合、AuthAvatar.URL()はErrNoAvatarURLを返すべきです")
	}

	testURL := "http://url-to-avatar"
	testUser = &gomniauthtest.TestUser{}
	testUser.On("AvatarURL").Return(testURL, nil)
	testChatUser.User = testUser
	url, err = authAvatar.URL(testChatUser)

	if err != nil {
		t.Error("値が存在する場合、AuthAvatar.URL()はエラーを返すべきではありません")
	} else if url != testURL {
		t.Error("AuthAvatar.URL()は正しいURLを返すべきです")
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar

	user := &chatUser{uniqueID: "abc"}
	url, err := gravatarAvatar.URL(user)

	if err != nil {
		t.Error("GravatarAvatar.URL()はエラーを返すべきではありません")
	} else if url != "//www.gravatar.com/avatar/abc" {
		t.Errorf("GravatarAvatar.URL()が%sという誤った値を返しました", url)
	}
}

func TestFileSystemAvatar(t *testing.T) {
	filename := filepath.Join("avatars", "abc.jpg")
	ioutil.WriteFile(filename, []byte{}, 0644)
	defer func() { os.Remove(filename) }()

	var fileSystemAvatar FileSystemAvatar

	user := &chatUser{uniqueID: "abc"}
	url, err := fileSystemAvatar.URL(user)

	if err != nil {
		t.Error("FileSystemAvatar.URL()はエラーを返すべきではありません")
	} else if url != "/avatars/abc.jpg" {
		t.Errorf("FileSystemAvatar.URL()が%sという誤った値を返しました", url)
	}
}
