package data

import (
	"fyne.io/fyne"
	"io/ioutil"
	"log"
)

func Logo() fyne.Resource {

	b, err := ioutil.ReadFile("client/img/logo.png")
	if err != nil {
		log.Println(err)
	}
	return fyne.NewStaticResource("logo", b)
}
