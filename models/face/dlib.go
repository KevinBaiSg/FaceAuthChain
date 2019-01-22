package face

import (
	"errors"
	"github.com/Kagami/go-face"
	"log"
)

type DLib struct {
}

func (d *DLib) Auth(descriptor interface{}, images []byte) error {
	rec, err := face.NewRecognizer("./dlib-models")
	if err != nil {
		log.Fatalf("Can't init face recognizer: %v", err)
		return err
	}
	defer rec.Close()

	des, ok := descriptor.(face.Descriptor)
	if !ok {
		return errors.New("the descriptor is not type of face.Descriptor")
	}

	if images == nil {
		return errors.New("the images is must not nil")
	}

	recoginzeFace, err := rec.RecognizeSingle(images)
	if err != nil {
		log.Printf("Can't recognize: %v\n", err )
		return err
	}

	if recoginzeFace == nil {
		log.Println("not a single face on the images")
		return errors.New("not a single face on the images")
	}

	var samples []face.Descriptor
	var cats []int32
	samples = append(samples, des)
	cats = append(cats, 0)
	rec.SetSamples(samples, cats)

	catID := rec.Classify(recoginzeFace.Descriptor)
	if catID != 0 {
		log.Println("can't classify")
		return errors.New("can't classify")
	}

	return nil
}

