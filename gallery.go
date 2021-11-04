package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/widget"
	"io/ioutil"
    "log"
	//"fmt"
	"strings"
	"fyne.io/fyne/v2/canvas"
)

func showGalleryApp() {
	w:= myApp.NewWindow("Gallery");
	w.Resize(fyne.NewSize(1000, 500)) 
	root := "C:\\Users\\HP\\OneDrive\\Desktop\\Photos"
	//var imagearr[] string
	tabs := container.NewAppTabs()
	files, err := ioutil.ReadDir(root)
    if err != nil {
        log.Fatal(err)
    }
    
    for _, file := range files {
        //fmt.Println(file.Name(), file.IsDir())
		if file.IsDir() == false{
            extension:= strings.Split(file.Name(), ".")[1];
			if extension == "jpg" || extension == "png"{
				//imagearr = append(imagearr, root+"\\"+file.Name())
				image := canvas.NewImageFromFile(root + "\\" + file.Name())
				tabs.Append(container.NewTabItem(file.Name(), image))
				tabs.SetTabLocation(container.TabLocationLeading)
		}
       }
	}
	//image := canvas.NewImageFromFile(imagearr[0])
	
	w.SetContent(container.NewBorder(DeskBtn,nil,nil,nil,tabs),
	)
	w.Show()
}