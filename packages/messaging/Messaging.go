package messaging

import (
	"fmt"
	"os"

	"github.com/fatih/color"
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
