package main

import (
	"fmt"
	"log"
	"io/ioutil"
    "github.com/ledongthuc/pdf"
	
	htgotts "github.com/hegedustibor/htgo-tts"
	handlers "github.com/hegedustibor/htgo-tts/handlers"
	voices "github.com/hegedustibor/htgo-tts/voices"
)

func main() {
	/*content, err := ioutil.ReadFile("test.txt")
    if err != nil {
        log.Fatal(err)
    }

    // Convert []byte to string and print to screen
    text := string(content)
	speech := htgotts.Speech{Folder: "audio", Language: voices.Spanish, Handler: &handlers.MPlayer{}}
	speech.Speak(text)*/

	/*content, err := readPdf("test.pdf")
    if err != nil {
        panic(err)
    }*/

	content, err := ioutil.ReadFile("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(content)

	speech := htgotts.Speech{Folder: "audio", Language: voices.Spanish, Handler: &handlers.MPlayer{}}
	myString := string(content[:])
	speech.Speak(myString)
	return
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		return "", err
	}
	totalPage := r.NumPage()

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		rows, _ := p.GetTextByRow()
		for _, row := range rows {
		    //println(">>>> row: ", row.Position)
		    for _, word := range row.Content {
		        fmt.Println(word.S)
		    }
		}
	}
	return "", nil
}