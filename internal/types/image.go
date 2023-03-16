package types

import (
	"fmt"
	"net/http"
	"strings"

	svg "github.com/h2non/go-is-svg"
)

// Media represents image of NFT downloaded from specified URI
type Media struct {
	Data []byte
	Type MediaType
}

type MediaType int8

const (
	MediaTypeUnknown MediaType = iota
	MediaTypeSvg
	MediaTypeGif
	MediaTypeJpeg
	MediaTypePng
	MediaTypeWebp
	MediaTypeMp4
	MediaTypeMp3
	MediaTypeGlb
)

func (i MediaType) Mimetype() string {
	switch i {
	case MediaTypeSvg:
		return "image/svg+xml"
	case MediaTypeGif:
		return "image/gif"
	case MediaTypeJpeg:
		return "image/jpeg"
	case MediaTypePng:
		return "image/png"
	case MediaTypeWebp:
		return "image/webp"
	case MediaTypeMp4:
		return "video/mp4"
	case MediaTypeMp3:
		return "audio/mpeg"
	}
	return ""
}

func (i MediaType) Extension() string {
	switch i {
	case MediaTypeSvg:
		return ".svg"
	case MediaTypeGif:
		return ".gif"
	case MediaTypeJpeg:
		return ".jpg"
	case MediaTypePng:
		return ".png"
	case MediaTypeWebp:
		return ".webp"
	case MediaTypeMp4:
		return ".mp4"
	case MediaTypeMp3:
		return ".mp3"
	}
	return ""
}

func MediaTypeFromMimetype(data []byte) (MediaType, error) {
	mimetype := http.DetectContentType(data)
	switch mimetype {
	case "image/svg+xml":
		return MediaTypeSvg, nil
	case "image/gif":
		return MediaTypeGif, nil
	case "image/jpeg":
		return MediaTypeJpeg, nil
	case "image/png":
		return MediaTypePng, nil
	case "image/webp":
		return MediaTypeWebp, nil
	case "video/mp4":
		return MediaTypeMp4, nil
	case "audio/mpeg":
		return MediaTypeMp3, nil
	}
	if strings.HasPrefix(mimetype, "text/xml") || strings.HasPrefix(mimetype, "text/plain") {
		if svg.Is(data) {
			return MediaTypeSvg, nil
		}
	}

	return MediaTypeUnknown, fmt.Errorf("unrecognized image type %s", mimetype)
}

func MediaTypeFromExtension(uri string) (mimetype MediaType) {
	uri = strings.ToLower(uri)
	if strings.HasSuffix(uri, ".svg") {
		return MediaTypeSvg
	}
	if strings.HasSuffix(uri, ".gif") {
		return MediaTypeGif
	}
	if strings.HasSuffix(uri, ".jpg") || strings.HasSuffix(uri, ".jpeg") {
		return MediaTypeJpeg
	}
	if strings.HasSuffix(uri, ".png") {
		return MediaTypePng
	}
	if strings.HasSuffix(uri, ".webp") {
		return MediaTypeWebp
	}
	if strings.HasSuffix(uri, ".mp4") {
		return MediaTypeMp4
	}
	if strings.HasSuffix(uri, ".mp3") {
		return MediaTypeMp3
	}
	if strings.HasSuffix(uri, ".glb") {
		return MediaTypeGlb
	}
	return MediaTypeUnknown
}
