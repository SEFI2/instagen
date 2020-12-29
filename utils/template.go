package utils

import "github.com/SEFI2/instagen/image"

func CreateInstagramPost(title, content, author,  backgroundPath, outputPath string) error {
	width, height := image.InstagramSquarePostSize()
	image := image.NewImage(width, height)

	if err := image.LoadBackgroundFromPath(backgroundPath); err != nil {
		return err
	}

	if err := image.GenerateOverlay(20); err != nil {
		return err
	}

	if err := image.AddTitleText(title); err != nil {
		return err
	}

	if err := image.AddContentText(content); err != nil {
		return err
	}

	if err := image.AddAuthorText(author); err != nil {
		return err
	}


	if err := image.Save(outputPath); err != nil {
		return err
	}

	return nil
}
