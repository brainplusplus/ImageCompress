// Copyright [2016] [hoenir]
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	// import all other side packages

	"github.com/hoenirvili/ImageCompress/apis/imgur"
	"github.com/hoenirvili/ImageCompress/apis/tinypng"
)

const (
	clientID     = "40aea5f08c0f717"
	clientSecret = "a72c35d27b38d27114b4503e5b9acc835861ed8c"
)

// imgurToTiny just send image to tinyPNG
func imgurToTiny() {
	// alloc
	imgur := imgur.NewImgur()
	tiny := tinypng.NewTiny()
	// set
	imgur.SetClientID(clientID)
	imgur.SetClientSecret(clientSecret)
	imgur.SetBody(new(bytes.Buffer))
	// get request JSON response
	v, err := imgur.ImageJSON("https://api.imgur.com/3/gallery/image/i0xn0Dx")
	if err != nil {
		log.Fatal(err)
	}

	if v.Data.Type == "image/png" || v.Data.Type == "image/jpeg" {
		byteImage, err := imgur.ImageByte(v.Data.Link, v.Data.Type)
		if err != nil {
			log.Fatal(err)
		}
		// prepare tiny for post Request
		// setting body with the image downloaded from imgur
		tiny.SetBody(byteImage)
		tiny.PostTry(v.Data.Type)
		// t, err := tiny.PostGetJSON(v.Data.Type)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		//
		// // get pic from tiny api
		// pic, err := tiny.Get(t.Output.URL)
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "%s", err)
		// }
		// fmt.Println(pic)
		// // save image pic
		// err = tiny.SaveImage(pic, "~/Work/Go/src/github.com/hoenirvili/ImageCompress/newImageCompressed", v.Data.Type)
		// if err != nil {
		// 	log.Fatal(err)
		// }
	} else {
		fmt.Fprintf(os.Stderr, "Error: image is not PNG/JPG type\n")
		fmt.Fprintf(os.Stderr, "Please enter a valid PNG/JPG file type\n")
		os.Exit(1)
	}
}

// func shackToTiny() {
// 	// alloc
// 	shack := imageshack.NewImageShack()
// 	tiny := tinypng.NewTiny()
// 	v := shack.ImageJSON("https://api.imageshack.com/v2/images/pbzPCsEij")
// 	url, err := utils.Concat("https://", v.Result.Direct_link)
//
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	if strings.Contains(url, ".jpg") {
// 		byteImage := shack.ImageByte(url)
// 		tiny.SetBody(byteImage)
// 		tiny.Post("image/jpeg")
// 	} else if strings.Contains(url, ".png") {
// 		byteImage := shack.ImageByte(url)
// 		tiny.SetBody(byteImage)
// 		tiny.Post("image/png")
// 	} else {
// 		fmt.Fprintf(os.Stderr, "Error: image is not PNG/JPG type\n")
// 		fmt.Fprintf(os.Stderr, "Please enter a valid PNG/JPG file type\n")
// 		os.Exit(1)
// 	}
// }

func main() {

	imgurToTiny()
	//	shackToTiny()

}
