package main

import (
	"fmt"
	"image/color"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"encoding/json"
	"net/http"
)

func showWeatherApp() {
	w:= myApp.NewWindow("Weather App");
	w.Resize(fyne.NewSize(600,600))


	input := widget.NewEntry()
	input.SetPlaceHolder("Enter City Name")
    
	//api part of the code
	// openweathermap.org  --> API --> create API keys 
	
	text := canvas.NewText("https://api.openweathermap.org/data/2.5/weather?q=delhi&APPID=43a6cd4cffb4ddd6c32017f7a570fffb", color.White)
	text.TextSize = 10
	
	
	res, err := http.Get(text.Text)
    
	if err!=nil{
		fmt.Println(err)
	}
	defer res.Body.Close() // close the http part after we have our response, now we will work on response
    // now we have to work with the html body we have, the api will give us a JSON file
	// we are reading that json file in body variable from response, and error
	body , err := ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Println(err)
	}
    // using unmarshalweather structure we will have weather details in weather variable
	// this is the last part in getting data from API
	weather , err := UnmarshalWelcome(body)

	// UI part of the code
	img := canvas.NewImageFromFile("bg1.png")
	img.FillMode = canvas.ImageFillOriginal
	//img.Resize(fyne.NewSize(500,500))
	label1:= canvas.NewText("Weather Details", color.Black )
	label1.TextStyle = fyne.TextStyle{Bold: true}
	label2:= canvas.NewText(fmt.Sprintf("Country : %s",weather.Sys.Country), color.Black )
	label3:= canvas.NewText(fmt.Sprintf("Temperature : %.2f °C",weather.Main.Temp-273.15), color.Black )
	label4 := canvas.NewText(fmt.Sprintf("Wind Speed : %.2f mph", weather.Wind.Speed), color.Black)
	label5 := canvas.NewText(fmt.Sprintf("City : %s", weather.Name), color.Black)
	label6 := canvas.NewText(fmt.Sprintf("Feels Like : %.2f °C", weather.Main.FeelsLike-273.15), color.Black)
	label7 := canvas.NewText(fmt.Sprintf("Description : %s", weather.Weather[0].Description), color.Black)
	label8 := canvas.NewText(fmt.Sprintf("Humidity : %d %%", weather.Main.Humidity), color.Black)
    

	btn1 := widget.NewButton("SEARCH", func() {

		if input.Text == "" {
			input.Text = "Delhi"
		}

		text.Text = "https://api.openweathermap.org/data/2.5/weather?q=" + input.Text + "&APPID=43a6cd4cffb4ddd6c32017f7a570fffb"
		text.Refresh()

		// fmt.Println(text.Text)

		response, err := http.Get(text.Text)

		if err != nil { // if we encounter error
			fmt.Println("Error, try again")
		}

		// now we have to work with the html body we have, the api will give us a JSON file
		// we are reading that json file in body variable from response, and error
		body, err := ioutil.ReadAll(response.Body)

		if err != nil { // if we encounter error
			fmt.Println("Error, try again")
		}

		// using unmarshalweather structure we will have weather details in weather variable
		// this is the last part in getting data from API
		weather, err := UnmarshalWelcome(body)

		if err != nil { // if we encounter error
			fmt.Println("Error, try again")
		}

		label2.Text = fmt.Sprintf("Country : %s", weather.Sys.Country)
		label4.Text = fmt.Sprintf("Temperature : %.2f °C", weather.Main.Temp-273.15)
		label3.Text = fmt.Sprintf("Wind Speed : %.2f mph", weather.Wind.Speed)
		label5.Text = fmt.Sprintf("City : %s", weather.Name)
		label6.Text = fmt.Sprintf("Feels Like : %.2f °C", weather.Main.FeelsLike-273.15)
		label7.Text = fmt.Sprintf("Description : %s", weather.Weather[0].Description)
		label8.Text = fmt.Sprintf("Humidity : %d %%", weather.Main.Humidity)

		label2.Refresh()
		label3.Refresh()
		label4.Refresh()
		label5.Refresh()
		label6.Refresh()
		label7.Refresh()
		label8.Refresh()
		// fmt.Printf(weather.Sys.Country)

	})
	
		weatherContainer:= container.NewVBox(
			label1,
			img , 
			input,
			btn1,
			label2,
			label3,
			label4,
			label5,
			label6,
			label7,
			label8,
			container.NewGridWithColumns(1,
			),
		)
	
	w.SetContent(container.NewBorder( DeskBtn,nil,nil,nil,weatherContainer))
	w.Show()
}

func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Coord      Coord     `json:"coord"`     
	Weather    []Weather `json:"weather"`   
	Base       string    `json:"base"`      
	Main       Main      `json:"main"`      
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`      
	Clouds     Clouds    `json:"clouds"`    
	Dt         int64     `json:"dt"`        
	Sys        Sys       `json:"sys"`       
	Timezone   int64     `json:"timezone"`  
	ID         int64     `json:"id"`        
	Name       string    `json:"name"`      
	Cod        int64     `json:"cod"`       
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`      
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`  
	TempMax   float64 `json:"temp_max"`  
	Pressure  int64   `json:"pressure"`  
	Humidity  int64   `json:"humidity"`  
}

type Sys struct {
	Type    int64  `json:"type"`   
	ID      int64  `json:"id"`     
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"` 
}

type Weather struct {
	ID          int64  `json:"id"`         
	Main        string `json:"main"`       
	Description string `json:"description"`
	Icon        string `json:"icon"`       
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`  
	Gust  float64 `json:"gust"` 
}
