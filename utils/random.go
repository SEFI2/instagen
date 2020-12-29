package utils

import (
	"fmt"
	"image"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"time"
)

func ImageFromUrl(url string) (image.Image, error) {
	img, _ := os.Create("image.jpg")
	defer img.Close()

	resp, _ := http.Get(url)
	defer resp.Body.Close()
	im, _, err := image.Decode(resp.Body)
	return im, err
}

func GenerateRandomImageUrl(keyword string) (string, error) {
	url := fmt.Sprintf("https://www.google.com/search?q=%s&tbm=isch", keyword)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile("src=\"(http[^\"]+)\"")
	matches := re.FindAllStringSubmatch(string(body), -1)

	potatoes := make([]string, len(matches))

	for index, match := range matches {
		potatoes[index] = match[1]
	}
	rand.Seed(time.Now().UnixNano())

	//get random image url and print to stdout
	randomUrl := potatoes[rand.Intn(len(potatoes))]
	return randomUrl, nil
}

