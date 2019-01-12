package main

import (
	"bytes"
	"encoding/json"
	"github.com/Kagami/go-face"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Response ...
type Response struct {
	Status string `json:"status"`
}

var rec *face.Recognizer = nil
var userNames []string
var DBDirectory = "DB"
var IdentifyFacesDirectory = "cache"

const (
	ROUTER_API_ENROLL   = "VeriLook/face/apiEnroll"
	ROUTER_API_IDENTIFY = "VeriLook/face/apiIdentify"
)

// Identify create a new item
func enroll(w http.ResponseWriter, r *http.Request) {
}

// Identify create a new item
func identify(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		_ = json.NewEncoder(w).Encode(Response{Status: "Parse Form failed\n"})
		return
	}
	file, _, err := r.FormFile("faceImage")
	if err != nil {
		_ = json.NewEncoder(w).Encode(Response{Status: "Identification failed\n"})
		return
	}
	defer func() { _ = file.Close() }()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		_ = json.NewEncoder(w).Encode(Response{Status: "Data error\n"})
		return
	}

	// save image for test
	saveReader := ioutil.NopCloser(bytes.NewBuffer(buf))
	defer saveReader.Close()
	_, _ = SaveFace(saveReader)

	recoginzeFace, err := rec.RecognizeSingle(buf)
	if err != nil {
		log.Printf("Can't recognize: %v\n", err)
		_ = json.NewEncoder(w).Encode(Response{Status: "Data error\n"})
		return
	}

	if recoginzeFace == nil {
		log.Println("Not a single face on the image")
	}
	catID := rec.Classify(recoginzeFace.Descriptor)
	if catID < 0 {
		log.Println("Can't classify")
	}
	if catID < len(userNames) {
		name := userNames[catID]
		status := "Matched with ID: '" + name + "' with score ** \n"
		_ = json.NewEncoder(w).Encode(Response{Status: status})
	} else {
		_ = json.NewEncoder(w).Encode(Response{Status: "Identify Failed"})
	}
}

func InitRecognizerDB() []string {
	var names []string
	var samples []face.Descriptor
	var cats []int32
	files, err := ioutil.ReadDir(DBDirectory)
	if err != nil {
		log.Fatal(err)
	}

	index := 0
	for _, f := range files {
		fileName := f.Name()
		ext := filepath.Ext(fileName)
		if ext != ".jpeg" && ext != ".jpg" {
			continue
		}
		imagePristin := filepath.Join(DBDirectory, fileName)
		faces, err := rec.RecognizeFile(imagePristin)
		if err != nil {
			continue
		}
		if len(faces) != 1 {
			continue
		}
		samples = append(samples, faces[0].Descriptor)
		cats = append(cats, int32(index))
		index++
		names = append(names, fileName[0:len(fileName)-len(ext)])
	}
	rec.SetSamples(samples, cats)
	return names
}

func SaveFace(file io.Reader) (written int64, err error) {
	t := time.Now()
	fileName := t.Format(time.RFC3339) + ".jpg"
	saveFile := filepath.Join(IdentifyFacesDirectory, fileName)
	f, err := os.Create(saveFile)
	if err != nil {
		return 0, err
	}

	return io.Copy(f, file)
}

func main() {
	var err error
	// please download models
	rec, err = face.NewRecognizer("./dlib-models")
	if err != nil {
		log.Fatalf("Can't init face recognizer: %v", err)
	}
	defer rec.Close()

	// TODO: Code refactoring by database
	userNames = InitRecognizerDB()

	router := mux.NewRouter()
	router.HandleFunc(ROUTER_API_ENROLL, enroll).Methods("POST")
	router.HandleFunc(ROUTER_API_IDENTIFY, identify).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
