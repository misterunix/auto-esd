package main

import (
	"fmt"
	"html/template"
	"os"
)

// smallImageTemplate is the template for the small image generation script
func sit(sd *stable) error {

	passOne := "smallImage.tpl"

	tmpl, err := template.New(passOne).Parse(smallImageTemplate)
	if err != nil {
		fmt.Fprintln(os.Stderr, "smallImageTemplate Parse:", err)
		return err
	}

	small, err := os.OpenFile(sd.SmallImagePython, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "smallImageTemplate OpenFile:", err)
		return err
	}

	//CheckFatal(err)

	defer small.Close()

	err = tmpl.Execute(small, sd)
	if err != nil {
		fmt.Fprintln(os.Stderr, "sd.SmallImagePython Execute:", err)
		return err
	}

	//Cmd := exec.Command("./installer_files/env/bin/python", sd.SmallImagePython)
	//err = Cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "sd.SmallImagePython Run:", err)
		return err
	}

	return nil
}

// largeImageTemplate is the template for the large image generation script
func lit(sd *stable) error {

	passOne := "largeImage.tpl"

	tmpl, err := template.New(passOne).Parse(largeImageTemplate)
	if err != nil {
		fmt.Fprintln(os.Stderr, "largeImageTemplate Parse:", err)
		return err
	}

	small, err := os.OpenFile(sd.LargeImagePython, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "largeImageTemplate OpenFile:", err)
		return err
	}

	//CheckFatal(err)

	defer small.Close()

	err = tmpl.Execute(small, sd)
	if err != nil {
		fmt.Fprintln(os.Stderr, "sd.LargeImagePython Execute:", err)
		return err
	}

	//Cmd := exec.Command("./installer_files/env/bin/python", sd.SmallImagePython)
	//err = Cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "sd.LargeImagePython Run:", err)
		return err
	}

	return nil
}
