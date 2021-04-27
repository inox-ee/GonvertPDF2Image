package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/cheggaaa/pb"
	"github.com/joho/godotenv"
	"gopkg.in/gographics/imagick.v2/imagick"
)

func main() {
	if err := godotenv.Load(".env.pub"); err != nil {
		log.Fatalf("Error at %v\n", err)
	}

	targetDir := "./resources/"

	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	files, err := ioutil.ReadDir(targetDir)
	if err != nil {
		panic(err)
	}

	for i, file := range files {
		filename := file.Name()
		if filepath.Ext(filename) != ".pdf" {
			log.Fatalln("input files must be PDF.")
		}
		fmt.Printf("%d / %d ... start > ", i+1, len(files))
		err = mw.SetResolution(300, 300)
		handleErr(err)
		var longfilename string
		if strings.HasSuffix(targetDir, "/") {
			longfilename += targetDir + filename
		} else {
			longfilename += targetDir + "/" + filename
		}
		fmt.Println(longfilename)
		err = mw.ReadImage(longfilename)
		handleErr(err)
		mw.SetIteratorIndex(0)
		pagenum := mw.GetNumberImages()
		log.Println("Total Page: ", pagenum)
		err = mw.SetImageFormat("png")
		handleErr(err)

		pb := pb.StartNew(int(pagenum))

		for j := 0; j < int(pagenum); j++ {
			if ret := mw.SetIteratorIndex(j); !ret {
				break
			}
			err = mw.WriteImage(fmt.Sprintf("./dist/%s_%02d.png", getFileNameWithoutExt(filename), j+1))
			handleErr(err)
			pb.Increment()
		}

		pb.Finish()
		mw.Clear()
	}
}

func handleErr(err error, msgs ...string) {
	if err != nil {
		log.Fatalf("Failed at %s\n", err)
	}
	if msgs != nil {
		log.Printf("--- %s ---\n", msgs)
	}
}

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
