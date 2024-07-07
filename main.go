package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/a-h/templ"
)

//go:generate templ generate

var serviceWebsiteURL = "https://byvko.dev"
var lastUpdatedDate = fmt.Sprint(time.Now().Format("Jan _2 2006"))

func main() {
	err := os.MkdirAll("./build", os.ModePerm)
	if err != nil {
		panic(err)
	}

	err = renderToFile(PrivacyPolicy(), "privacy-policy", true)
	if err != nil {
		panic(err)
	}
	err = renderToFile(TermsOfService(), "terms-of-service", true)
	if err != nil {
		panic(err)
	}
	err = renderToFile(ReturnPolicy(), "return-policy", true)
	if err != nil {
		panic(err)
	}
	err = renderToFile(Directory(), "index", true)
	if err != nil {
		panic(err)
	}
}

func renderToFile(component templ.Component, name string, withPartial bool) error {
	{
		f, err := os.Create(fmt.Sprintf("./build/%s.html", name))
		if err != nil {
			return err
		}
		defer f.Close()
		err = Page(component).Render(context.Background(), f)
		if err != nil {
			return err
		}
	}
	if withPartial {
		f, err := os.Create(fmt.Sprintf("./build/%s-partial.html", name))
		if err != nil {
			return err
		}
		defer f.Close()
		err = component.Render(context.Background(), f)
		if err != nil {
			return err
		}
	}
	return nil
}
