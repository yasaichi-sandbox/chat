package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func uploaderHandler(w http.ResponseWriter, req *http.Request) {
	file, header, err := req.FormFile("avatarFile")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	userID := req.FormValue("userID")
	filename := filepath.Join("avatars", userID+filepath.Ext(header.Filename))
	ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, "成功")
}
