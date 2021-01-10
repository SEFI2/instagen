package utils

import (
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"time"
)

// ImageFromUrl
func ImageFromUrl(url string) (image.Image, error) {
	img, _ := os.Create("image.jpg")
	defer img.Close()

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	im, _, err := image.Decode(resp.Body)
	if err := resp.Body.Close(); err != nil {
		return nil, err
	}

	return im, err
}

// RandomSquareImage
func RandomSquareImage(keyword string) (image.Image, error) {
	url, err := GenerateRandomSquareImageUrl(keyword)
	if err != nil {
		return nil, err
	}
	return ImageFromUrl(url)
}

// GenerateRandomSquareImageUrl
func GenerateRandomSquareImageUrl(keyword string) (string, error) {
	accessKey := "w-skuXJfa51crHs_LV53C0rCjWboXD-G0b734w3EoRE"
	orientation := "squarish"
	query := keyword
	url := fmt.Sprintf("https://api.unsplash.com/photos/random?client_id=%s&orientation=%s&query=%s", accessKey, orientation, query)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var data struct {
		Urls struct {
			Regular string
		}
	}
	if err = decoder.Decode(&data); err != nil {
		return "", err
	}
	return data.Urls.Regular, nil
}

// GenerateRandomImageUrl
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

