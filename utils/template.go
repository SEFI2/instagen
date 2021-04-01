package utils

import (
	"github.com/SEFI2/instagen/image"
	img "image"
)

func CreateInstagramPost(title, content, author string) (img.Image, error) {
	// Randomly choose background image
	// seed := rand.NewSource(time.Now().UnixNano())
	// random := rand.New(seed)
	// backgroundPath := fmt.Sprintf("BackgroundImage/Square/land%d.jpg", random.Int() % 40)

	// Get instagram post size
	width, height := image.InstagramSquarePostSize()
	image := image.NewImage(width, height)

	background, err := RandomSquareImage("")
	if err != nil {
		return nil, err
	}

	if err := image.LoadBackground(background); err != nil {
		return nil, err
	}

	if err := image.GenerateOverlay(20); err != nil {
		return nil, err
	}

	if err := image.AddTitleText(title); err != nil {
		return nil, err
	}

	if err := image.AddContentText(content); err != nil {
		return nil, err
	}

	if err := image.AddAuthorText(author); err != nil {
		return nil, err
	}


	// if err := image.Save(outputPath); err != nil {
	// 	return err
	//}

	return image.GetImage(), nil
}
