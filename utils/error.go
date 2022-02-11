package utils

import (
	"github.com/fatih/color"
)

func HandleError(err error) bool {
	if err != nil {
		color.HiRed(err.Error())
		return true
	}
	return false
}
