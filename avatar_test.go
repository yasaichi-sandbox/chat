package main

import (
	"testing"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar

	client := new(client)
	url, err := authAvatar.URL(client)
	if err != ErrNoAvatarURL {
		t.Error("値が存在しない場合、AuthAvatar.URL()はErrNoAvatarURLを返すべきです")
	}

	testURL := "http://url-to-avatar"
	// NOTE: The zero value of a map is nil.
	client.userData = map[string]interface{}{"avatar_url": testURL}
	url, err = authAvatar.URL(client)
	if err != nil {
		t.Error("値が存在する場合、AuthAvatar.URL()はエラーを返すべきではありません")
	} else if url != testURL {
		t.Error("AuthAvatar.URL()は正しいURLを返すべきです")
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar

	client := new(client)
	client.userData = map[string]interface{}{"email": "MyEmailAddress@example.com"}

	url, err := gravatarAvatar.URL(client)
	if err != nil {
		t.Error("GravatarAvatar.URL()はエラーを返すべきではありません")
	} else if url != "//www.gravatar.com/avatar/0bc83cb571cd1c50ba6f3e8a78ef1346" {
		t.Errorf("GravatarAvatar.URL()が%sという誤った値を返しました", url)
	}
}
