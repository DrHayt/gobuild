package messaging

import (
	"fmt"
	"os"

	"github.com/CrowBits/gobuild/Godeps/_workspace/src/github.com/fatih/color"
)

// Print Size Satandards
const (
	ScreenWidth = 80
	StatMsgSize = 12
)

// Color Constants
var (
	ColHiGreen   = color.New(color.FgHiGreen)
	ColHiYellow  = color.New(color.FgHiYellow)
	ColHiRed     = color.New(color.FgHiRed)
	ColHiCyan    = color.New(color.FgHiCyan)
	ColHiMegenta = color.New(color.FgHiMagenta)
)

// Msg Constants
var (
	TxtNo      = "No"
	TxtYes     = "Yes"
	TxtFail    = "Fail"
	TxtSuccess = "Success"
	TxtError   = "Error"
)

// SetNoColor turns color off
func SetNoColor() {
	color.NoColor = true
}

// SectionHeader will print the section splitter
func SectionHeader(title string) {
	boxClr := color.New(color.FgHiBlue).SprintFunc()
	boxTTL := color.New(color.FgHiGreen).SprintFunc()

	fmt.Println("")
	fmt.Println(boxTTL(fmt.Sprint(" ", title)))
	fmt.Printf("%s\n", boxClr(padLeft("", "=", ScreenWidth)))
}

// KeyVal will print a key and val on one lind
func KeyVal(key, val string) {
	keyCLR := ColHiMegenta.SprintFunc()
	valCLR := ColHiCyan.SprintFunc()
	fmt.Println(keyCLR(key), valCLR(val))
}

// Action will print the description
func Action(act, padc string) {
	fmt.Print(padRight(fmt.Sprint(act, " "), ".", ScreenWidth-StatMsgSize))
}

// Col4Bar will print a table header with 4 columns
func Col4Bar(col1, col2, col3, col4 string) {

	txtCol := color.New(color.FgHiMagenta).SprintFunc()

	col1pat := padRight("-[%s]-", "-", (ScreenWidth+2)-(len(col1)+(StatMsgSize*3)))
	col2pat := padLeft("-[%s]-", "-", (StatMsgSize+2)-len(col2))
	col3pat := padLeft("-[%s]-", "-", (StatMsgSize+2)-len(col3))
	col4pat := padLeft("-[%s]-", "-", (StatMsgSize+2)-len(col4))

	fmt.Printf("%s%s%s%s\n",
		fmt.Sprintf(col1pat, txtCol(col1)),
		fmt.Sprintf(col2pat, txtCol(col2)),
		fmt.Sprintf(col3pat, txtCol(col3)),
		fmt.Sprintf(col4pat, txtCol(col4)))
}

// Col4Text print the first column of the table row
func Col4Text(msgTxt string) {
	fmt.Print(padRight(fmt.Sprint(msgTxt, " "), " ", ScreenWidth-(StatMsgSize*3)))
}

// ErrorMsg Will print a standardized msg with ERROR and exit if code provided
func ErrorMsg(err error, code int) {
	msgColor := color.New(color.FgHiRed)
	MsgPrint(msgColor, TxtError, err, code)
}

// StatOut will print a status msg
func StatOut(col color.Color, myMsg, padChr string, rtrn bool) {
	// Prety Up the Msg
	myMsg = fmt.Sprint("[", myMsg, "]")

	// Build Prefix / Padding
	preFx := padLeft("", padChr, StatMsgSize-len(myMsg))

	// Get Color Function
	msgColp := col.SprintFunc()

	// Print The Stat
	fmt.Print(preFx, msgColp(myMsg), getRtrn(rtrn))
}

// =============================================================================[ Generic Msging]

// MsgPrint will print a one line msg with prefix in a passed a color
func MsgPrint(preFxColor *color.Color, preFxMsg string, msgError error, exitCode int) {

	// If we have a error get to work
	if msgError != nil {
		// Make it the function we want
		preFxColorP := preFxColor.SprintFunc()

		// Stype the way we want
		preFxMsg = fmt.Sprintf("[%s]", preFxMsg)

		// Print the msg
		fmt.Printf("%s %s\n", preFxColorP(preFxMsg), msgError.Error())

		// Check ExitCode to exit
		if exitCode > 0 {
			os.Exit(exitCode)
		}
	}
}

// =============================================================================[ Msging Functions ]

func padRight(sIn string, padChr string, lenOut int) string {
	for len(sIn) < lenOut {
		sIn = fmt.Sprint(sIn, padChr)
	}
	return sIn
}

func padLeft(sIn, padChr string, lenOut int) string {
	for len(sIn) < lenOut {
		sIn = fmt.Sprint(padChr, sIn)
	}
	return sIn
}

func getRtrn(rtrn bool) (strOut string) {
	if rtrn {
		strOut = "\n"
	}
	return
}
