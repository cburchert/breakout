//go:generate file2byteslice -package=data -input=./mplus-1p-regular.ttf -output=./generated/mplus-1p-regular.go -var=MPlus1PRegular_ttf
//go:generate file2byteslice -package=data -input=./heart.png -output=./generated/heart.go -var=Heart_png

package resources

import (
	// Dummy imports for go.mod for some Go files with 'ignore' tags. For example, `go mod tidy` does not
	// recognize Go files with 'ignore' build tag.
	//
	// Note that this affects only importing this package, but not 'file2byteslice' commands in //go:generate.
	_ "github.com/hajimehoshi/file2byteslice"
)
