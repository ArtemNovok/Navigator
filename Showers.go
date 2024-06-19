package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func ShowAllDirs(c []os.DirEntry) {
	ch := false
	for _, entry := range c {
		if entry.IsDir() {
			ch = true
			fmt.Println(color.CyanString("	Dir: ") + color.WhiteString(entry.Name()))
		}
	}
	if !ch {
		fmt.Println(color.YellowString("	there is no dirs in current directoire"))
	}
}

func ShowCurrentDir() error {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("	", color.RedString(err.Error()))
		return err
	}
	c, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("	", color.RedString(err.Error()))
		return err
	}
	ShowAllDirs(c)
	fmt.Println("	Current dir:", color.GreenString(dir))
	return nil
}
func SHowCurrentDirPath() error {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("	", color.RedString(err.Error()))
		return err
	}
	fmt.Println("	Current dir:", color.GreenString(dir))
	return nil
}
func ShowInfoMessage() {
	fmt.Println("  ", color.GreenString("use it like cd command, but to stop app enter 'stop' or 'st' "))
}

func ShowFiles(c []os.DirEntry) {
	ch := false
	for ind, entry := range c {
		if !entry.IsDir() {
			ch = true
			fmt.Println(color.MagentaString(fmt.Sprintf("	File %v: ", ind)) + color.WhiteString(entry.Name()))
		}
	}
	if !ch {
		fmt.Println(color.YellowString("	there is no files in current directoire"))
	}
}
