package aw

import (
  "fmt"
  "github.com/logrusorgru/aurora"
  "os"
)

var hasColors = os.Getenv("cm_no_colors") != "nope"
var x = aurora.NewAurora(hasColors)

var Red = x.Red
var Faint = x.Faint
var Bold = x.Bold
var Green = x.Green
var Blue = x.Blue
var BgBlue = x.BgBlue
var BgGreen = x.BgGreen
var Yellow = x.Yellow
var BgYellow = x.BgYellow
var BgRed = x.BgRed
var Underline = x.Underline
var Italic = x.Italic
var White = x.White
var Magenta = x.Magenta
var BgMagenta = x.BgMagenta
var BrightMagenta = x.BrightMagenta
var BgBrightMagenta = x.BgBrightMagenta
var BrightRed = x.BrightRed
var Black = x.Black
var BrightBlack = x.BrightBlack
var BgBrightCyan = x.BgBrightCyan
var BgCyan = x.BgCyan;
var Cyan = x.Cyan;
var BrightCyan = x.BrightCyan

func init(){
  fmt.Println(Bold("colors on?:"), Cyan(hasColors))
}
