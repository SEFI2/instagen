package image

import (
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"

	"image"
	"image/color"
	"path/filepath"
)

type Image struct {
	context *gg.Context
	margin float64
	background image.Image
	curDir string
}

func InstagramStorySize() (int, int) {
	return 1080, 1920
}

func InstagramSquarePostSize() (int, int) {
	return 1080, 1080
}

func InstagramLandscapePostSize() (int, int) {
	return 1080, 608
}

func InstagramPortraitPostSize() (int, int) {
	return 1080, 1350
}

func NewImage(width, height int) *Image {
	c := gg.NewContext(width, height)
	/*_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
		return nil
	}
	curDir := path.Dir(filename)
	*/
	curDir := ""
	return &Image{context: c, margin: 0, curDir: curDir}
}


func (i *Image) LoadBackground(background image.Image) error {
	i.context.DrawImage(background, 0, 0)
	i.background = background
	return nil
}

func (i *Image) LoadBackgroundFromPath(path string) error {
	background, err := gg.LoadImage(path)
	if err != nil {
		return err
	}
	return i.LoadBackground(background)
}

func (i *Image) GenerateOverlay(margin float64) error {
	i.margin = margin
	x, y := margin, margin
	w := float64(i.context.Width()) - (2.0 * margin)
	h := float64(i.context.Height()) - (2.0 * margin)
	i.context.SetColor(color.RGBA{A: 170})
	i.context.DrawRectangle(x, y, w, h)
	i.context.Fill()
	return nil
}

func (i *Image) AddTitleText(title string) error {
	textColor := color.RGBA{R: 250, G: 150, A: 240}
	fontPath := filepath.Join(i.curDir, "Roboto", "Roboto-Bold.ttf")
	if err := i.context.LoadFontFace(fontPath, 30); err != nil {
		return err
	}
	textRightMargin := 3 * i.margin
	textTopMargin := 3 * i.margin
	x, y := textRightMargin, textTopMargin
	maxWidth := float64(i.context.Width()) - textRightMargin - textRightMargin
	i.context.SetColor(textColor)
	i.context.DrawStringWrapped(title, x, y, 0, 0, maxWidth, 1.5, gg.AlignLeft)
	return nil
}

func (i *Image) AddContentText(content string) error {
	if len(content) > 450 {
		content = content[0:450]
	}
	textColor := color.White
	fontPath := filepath.Join(i.curDir, "Roboto", "Roboto-Bold.ttf")
	if err := i.context.LoadFontFace(fontPath, 45); err != nil {
		return err
	}
	textRightMargin := 6 * i.margin
	textTopMargin := float64(i.context.Height()) / 4.0
	x, y := textRightMargin, textTopMargin
	maxWidth := float64(i.context.Width()) - textRightMargin - textRightMargin
	i.context.SetColor(textColor)
	i.context.DrawStringWrapped(content, x, y, 0, 0, maxWidth, 1.5, gg.AlignLeft)
	return nil
}

func (i *Image) AddAuthorText(author string) error {
	fontPath := filepath.Join(i.curDir, "Roboto", "Roboto-Bold.ttf")
	if err := i.context.LoadFontFace(fontPath, 75); err != nil {
		return err
	}
	i.context.SetColor(color.RGBA{G: 205, A: 200})
	textWidth, textHeight := i.context.MeasureString(author)
	x := float64(i.context.Width()) - textWidth - 2 * i.margin
	y := float64(i.context.Height()) - textHeight - 2 * i.margin
	i.context.DrawString(author, x, y)
	return nil
}

func (i *Image) Save(path string) error {
	imaging.Fill(i.background, i.context.Width(), i.context.Height(), imaging.Center, imaging.Lanczos)
	if err := i.context.SavePNG(path); err != nil {
		return err
	}
	return nil
}

func (i *Image) GetImage() image.Image {
	return i.context.Image()
}