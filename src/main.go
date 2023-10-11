package main

import (
	"fmt"
	"github.com/nkocsea/skylib_skyutl/config"
	"github.com/nkocsea/skylib_skyutl/skyutl"
)

type PT struct {
	Id int64
	A  *string
	B  string
	C  int64
	D  int64
	E  int32
	F  int32
	G  bool
}

type A struct {
	Id int64
	A  string
	B  *string
	C  int64
	D  *int64
	E  int32
	F  *int32
	G  bool `json:"g,omitempty"`
}

func main() {
	// data, _, err := skyutl.ReadAsBytes("/data/online/test.png")
	
	// fmt.Println(err)

	// imgAsBytes, err := skyutl.ResizeImageWithBytes(data, 0, 300)
	// fmt.Println(err)
	// img, _, err := skyutl.ByteArrayToImage(imgAsBytes)
	// fmt.Println(err)
	
	// skyutl.WriteImage(img, "/data/online/test2.png")
	var test config.AppConfig
	fmt.Println(test)
	fmt.Print(skyutl.StringRightPaddingList([]string{"aaabbbcc", "aaabbbcc"}, []int{20, 20}))
	// fmt.Println(skyutl.DecodeToken("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzg5MTg3MzcsInVzZXJJZCI6IjEiLCJwYXJ0bmVySWQiOiIwIiwidXNlcm5hbWUiOiJyb290IiwiZnVsbE5hbWUiOiIiLCJkZXZpY2VJZCI6IjAiLCJhY2NvdW50VHlwZSI6IjAifQ.Ara6rZR2a22H8jTkXNrfM-G6ZSFKImwSQ4-Zm3bWJBqedSe8zd7DTFz4Xet7KGWTGgrvJQVW3cMriVzav_c3V0yVuclxMVX04q3ekEuN6idK23Yzt8mDgYT0FpvbiKhjMLXfqvekOWK2LliefdGySpngI3mISHGm2gRMw2cl798"))
}

