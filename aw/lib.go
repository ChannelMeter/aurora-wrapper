package aw

import (
  "github.com/logrusorgru/aurora"
  "golang.org/x/crypto/ssh/terminal"
  "os"
)

var isTerminal = terminal.IsTerminal(int(os.Stdout.Fd()))
var hasColors = os.Getenv("cm_no_colors") != "yes"
var x = aurora.NewAurora(isTerminal && hasColors)

var Faint = x.Faint
var Bold = x.Bold
var Underline = x.Underline
var Italic = x.Italic

// green
var Green = x.Green
var BgGreen = x.BgGreen
var BrightGreen = x.BrightGreen
var BgBrightGreen = x.BgBrightGreen

// white
var White = x.White
var BgWhite = x.BgWhite
var BrightWhite = x.BrightWhite
var BgBrightWhite = x.BgBrightWhite;

// magenta
var Magenta = x.Magenta
var BgMagenta = x.BgMagenta
var BrightMagenta = x.BrightMagenta
var BgBrightMagenta = x.BgBrightMagenta

// red
var Red = x.Red
var BrightRed = x.BrightRed
var BgRed = x.BgRed;
var BgBrightRed = x.BgBrightRed

// yellow
var Yellow = x.Yellow
var BgYellow = x.BgYellow
var BrightYellow = x.BrightYellow
var BgBrightYellow = x.BgBrightYellow

// gray
var Gray = x.Gray
var BgGray = x.BgGray

// black
var Black = x.Black
var BrightBlack = x.BrightBlack
var BgBlack = x.BgBlack
var BgBrightBlack = x.BgBrightBlack

// blue
var Blue = x.Blue
var BgBlue = x.BgBlue
var BgBrightBlue = x.BgBrightBlue;
var BrightBlue = x.BrightBlue


// cyan
var Cyan = x.Cyan;
var BgCyan = x.BgCyan;
var BrightCyan = x.BrightCyan
var BgBrightCyan = x.BgBrightCyan

