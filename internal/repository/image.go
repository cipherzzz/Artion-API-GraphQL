package repository

import (
	"artion-api-graphql/internal/types"
	"bytes"
	"fmt"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

const thumbnailMaxHeight = 500
const thumbnailMaxWidth = 500

// createThumbnail resize the Image
func createThumbnail(input types.Media) (output types.Media, err error) {
	// skip thumbnail of SVG, WebP and empty files
	if input.Type == types.MediaTypeSvg || input.Type == types.MediaTypeWebp || len(input.Data) == 0 {
		return input, nil
	}

	// create a thumbnail of a video file
	if input.Type == types.MediaTypeMp4 {
		input, err = createVideoThumbnail(input)
		if err != nil {
			return types.Media{}, err
		}
		return input, nil
	}

	// simple image frame nail
	return createImageThumbnail(input)
}

// createImageThumbnail creates a smaller banner for the given image.
func createImageThumbnail(input types.Media) (output types.Media, err error) {
	reader := bytes.NewReader(input.Data)

	img, err := imaging.Decode(reader, imaging.AutoOrientation(true))
	if err != nil {
		return types.Media{}, fmt.Errorf("decoding failed (type %d); %s", input.Type, err)
	}

	small := imaging.Fit(img, thumbnailMaxWidth, thumbnailMaxHeight, imaging.Linear)

	var writer bytes.Buffer
	if input.Type == types.MediaTypeJpeg {
		err = imaging.Encode(&writer, small, imaging.JPEG, imaging.JPEGQuality(80))
	} else {
		err = imaging.Encode(&writer, small, imaging.PNG) // also for GIFs
		input.Type = types.MediaTypePng
	}

	if err != nil {
		return types.Media{}, err
	}
	return types.Media{
		Data: writer.Bytes(),
		Type: input.Type,
	}, nil
}

func createVideoThumbnail(input types.Media) (output types.Media, err error) {
	inputReader := bytes.NewReader(input.Data)
	writer := bytes.NewBuffer(nil)
	frameNum := 3
	err = ffmpeg.
		Input("pipe:").
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithInput(inputReader).
		WithOutput(writer).
		Run()
	if err != nil {
		return types.Media{}, err
	}
	return types.Media{
		Data: writer.Bytes(),
		Type: types.MediaTypeJpeg,
	}, nil
}
