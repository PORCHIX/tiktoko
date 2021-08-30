package tiktok

import (
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
)

const (
	link = "https://vm.tiktok.com/ZSJsSRafc/"
)

func DownloadTikTokVideo(url string) (string, error) {
	if !isValidTikTokLink(url) {
		return "", errors.New("invalid TikTok link")
	}
	cmd := exec.Command("tiktok-scraper", "video", url, "-d", "-w")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New("\"Error running " + cmd.String() + ": " + err.Error())
	}
	filename, err := extractFileName(string(out))
	if err != nil {
		return "", err
	}
	return string(filename), nil
}

func isValidTikTokLink(link string) bool {
	resp, err := http.Get(link)
	if err != nil {
		return false
	}
	if resp.StatusCode != 200 {
		return false
	}
	u, err := url.Parse(link)
	if err != nil {
		return false
	}
	logrus.Printf(u.Host)
	if u.Host != "vm.tiktok.com" {
		return false
	}
	return true
}

func extractFileName(s string) (string, error) {
	strs := strings.Split(s, " ")
	logrus.Print(len(strs))
	if len(strs) != 3 {
		return "", errors.New("unable to download tik tokvideo")
	}
	filepath := strs[2]
	strs = strings.Split(filepath, "/")
	filename := strs[len(strs)-1]
	return strings.TrimSpace(filename), nil
}