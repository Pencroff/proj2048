package resources

import (
	"embed"
	log "github.com/go-pkgz/lgr"
	fontConverter "github.com/tdewolff/canvas/font"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed assets/*
var assets embed.FS

var CeilOneCharFF font.Face
var CeilTwoCharFF font.Face
var CeilThreeCharFF font.Face
var CeilFourCharFF font.Face
var CeilFiveCharFF font.Face

var HeaderFontSize float64
var HeaderFF font.Face
var TitleFontSize float64
var TitleFF font.Face
var TextFontSize float64
var TextFF font.Face

func InitFonts() {

	//log.Printf("Content: %v", content)
	//file, err := content.Open("assets/ClearSans-Bold-webfont.woff")
	//if err != nil {
	//	log.Printf("[ERROR], Can't open embed file: %v", err)
	//}
	//stat, err := file.Stat()
	//if err != nil {
	//	log.Printf("[ERROR], Stat err: %v", err)
	//}
	//log.Printf("Stat: %v, %v", stat.Name(), stat.Size())

	var regularFont = LoadFont("assets/ClearSans-Bold-webfont.woff")
	var boldFont = LoadFont("assets/ClearSans-Bold-webfont.woff")

	CeilOneCharFF = MakeScreenFontFace(boldFont, 55)
	CeilTwoCharFF = MakeScreenFontFace(boldFont, 55)
	CeilThreeCharFF = MakeScreenFontFace(boldFont, 45)
	CeilFourCharFF = MakeScreenFontFace(boldFont, 35)
	CeilFiveCharFF = MakeScreenFontFace(boldFont, 30)

	HeaderFontSize = 24
	HeaderFF = MakeScreenFontFace(boldFont, HeaderFontSize)
	TitleFontSize = 20
	TitleFF = MakeScreenFontFace(boldFont, TitleFontSize)
	TextFontSize = 16
	TextFF = MakeScreenFontFace(regularFont, TextFontSize)
}

func MakeScreenFontFace(fnt *opentype.Font, size float64) font.Face {
	return MakeFontFace(fnt, size, 72)
}

func MakeFontFace(fnt *opentype.Font, size float64, dpi float64) font.Face {
	if dpi == 0 {
		dpi = 72
	}

	fontFace, err := opentype.NewFace(fnt, &opentype.FaceOptions{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Printf("[ERROR] Can't make font face: %v", err)
	}
	return fontFace
}

func LoadFont(fontPath string) *opentype.Font {
	//currentFolder, err := os.Getwd()
	//if err != nil {
	//	log.Printf("[ERROR] Can't resolve os.Getwd: %v", err)
	//}
	//filePath := strings.Replace(path.Join(currentFolder, fontPath), "/", "\\", -1)

	rawData, err := assets.ReadFile(fontPath)
	if err != nil {
		log.Printf("[ERROR] Can't find file: %v", err)
	}
	fntData, err := fontConverter.ToSFNT(rawData)
	if err != nil {
		log.Printf("[ERROR] Can't transform woff: %v", err)
	}
	fnt, err := opentype.Parse(fntData)
	if err != nil {
		log.Printf("[ERROR] Can't parse font: %v", err)
	}
	return fnt
}
