package main

import (
	//"fyne.io/fyne/v2"
	"strconv"

	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func showCalc() {
	
	output:="";
	w:=myApp.NewWindow("Calc");
	w.Resize(fyne.NewSize(500,240))
	isHistory:= false
	var str string = ""
	history := widget.NewLabel(str)
    input := widget.NewLabel(output)
	var historyArr [] string 
	historyBtn := widget.NewButton("history",func(){
		if isHistory{
            str = ""
		}else{
			for i := len(historyArr)-1; i>=0; i-- {
				str= str + historyArr[i]
				str = str + "\n"
			}
		}
		history.SetText(str)
        isHistory = !isHistory
	})
	backBtn := widget.NewButton("back",func(){
		if output=="error"{
			output = ""
			input.SetText(output);
		}
	if len(output)>0{
		output = output[:len(output)-1]
		input.SetText(output);}
	})
	clearBtn := widget.NewButton("clear",func(){
		
		output = ""
		input.SetText(output)
		
	})
    openBtn := widget.NewButton("(",func(){
		if output=="error"{
			output = ""
		}
		output = output + "("
		input.SetText(output)
		
	})
	closeBtn := widget.NewButton(")",func(){
		if output=="error"{
			output = ""
		}
		output = output + ")"
		input.SetText(output)
		
	})
	divideBtn := widget.NewButton("/",func(){
		if output=="error"{
			output = ""
		}
		output = output + "/"
		input.SetText(output)
		
	})
	nineBtn := widget.NewButton("9",func(){
		if output=="error"{
			output = ""
		}
		output = output + "9"
		input.SetText(output)
	})
	eightBtn := widget.NewButton("8",func(){
		if output=="error"{
			output = ""
		}
		
		output = output + "8"
		input.SetText(output)
		
	})
	sevenBtn := widget.NewButton("7",func(){
		if output=="error"{
			output = ""
		}
		
		output = output + "7"
		input.SetText(output)
		
	})
	multiplyBtn := widget.NewButton("*",func(){
		if output=="error"{
			output = ""
		}
		output = output + "*"
		input.SetText(output)
	})
	fourBtn := widget.NewButton("4",func(){
		if output=="error"{
			output = ""
		}
		output = output + "4"
		input.SetText(output)
		
	})
	fiveBtn := widget.NewButton("5",func(){
		if output=="error"{
			output = ""
		}
		output = output + "5"
		input.SetText(output)
	})
	sixBtn := widget.NewButton("6",func(){
		if output=="error"{
			output = ""
		}
		output = output + "6"
		input.SetText(output)
	})
	minusBtn := widget.NewButton("-",func(){
		if output=="error"{
			output = ""
		}
		output = output + "-"
		input.SetText(output)
	})
	oneBtn := widget.NewButton("1",func(){
		if output=="error"{
			output = ""
		}
		output = output + "1"
		input.SetText(output)
	})
	twoBtn := widget.NewButton("2",func(){
		if output=="error"{
			output = ""
		}
		output = output + "2"
		input.SetText(output)
	})
	threeBtn := widget.NewButton("3",func(){
		if output=="error"{
			output = ""
		}
		output = output + "3"
		input.SetText(output)
		
	})
	plusBtn := widget.NewButton("+",func(){
		if output=="error"{
			output = ""
		}
		output = output + "+"
		input.SetText(output)
		
	})
	zeroBtn := widget.NewButton("0",func(){
		if output=="error"{
			output = ""
		}
		output = output + "0"
		input.SetText(output)
	})
	dotBtn := widget.NewButton(".",func(){
		if output=="error"{
			output = ""
		}
		output = output + "."
		input.SetText(output)
	})
	equalBtn := widget.NewButton("=",func(){
		if output=="error"{
			output = "error"
		}else if len(output) > 0{
        expression, err := govaluate.NewEvaluableExpression(output);
		if err==nil{
	        result, err := expression.Evaluate(nil);
			if err==nil{
				ans:= strconv.FormatFloat(result.(float64),'f',-1,64) //result is converted into string
			    strToAppend:= output + "=" + ans
				historyArr = append(historyArr,strToAppend)
				output = ans 
			
			}else{
				output = "error"
			}
		}else{
			output = "error"
		}}else{
			output = "error"
		}
	// result is now set to "true", the bool value.
        input.SetText(output)
	})

	calcContainer:=container.NewVBox(
		input,
		history,
		container.NewGridWithColumns(2,
		historyBtn,
	    backBtn),
		container.NewGridWithColumns(4,
		clearBtn,
	    openBtn,
	    closeBtn,
	    divideBtn),
		container.NewGridWithColumns(4,
			nineBtn,
			eightBtn,
			sevenBtn,
			multiplyBtn,
		),
		container.NewGridWithColumns(4,
		    fourBtn,
		    fiveBtn,
		    sixBtn,
		    minusBtn,
		),
		container.NewGridWithColumns(4,
		    oneBtn,
		    twoBtn,
		    threeBtn,
		    plusBtn),
		container.NewGridWithColumns(2,
				container.NewGridWithColumns(2,
				zeroBtn,
				dotBtn),
	    equalBtn),

		
	)
	
    w.SetContent(
		container.NewBorder(DeskBtn,nil,nil,nil,calcContainer),
		
	)
	w.Show()
}