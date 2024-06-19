package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/atotto/clipboard"
	"github.com/fatih/color"
)

type Cache struct {
	Files []os.DirEntry
}

func main() {
	cache := Cache{}
	ShowInfoMessage()
	filePath := ""
	err := ShowCurrentDir()
	if err != nil {
		fmt.Println(" 	", color.RedString(err.Error()))
		return
	}
	for {
		dir, err := GetInput()
		if err != nil {
			fmt.Println("	", color.RedString(err.Error()))
			continue
		}
		switch dir {
		case "stop":
			if filePath != "" {
				fmt.Println(color.YellowString(" Path:") + color.GreenString("		"+filePath))
				clipboard.WriteAll(filePath)
			} else {
				fmt.Println(color.GreenString("done"))
			}
			return

		case "st":
			if filePath != "" {
				fmt.Println(color.YellowString(" Path:") + color.GreenString("		"+filePath))
				clipboard.WriteAll(filePath)
			} else {
				fmt.Println(color.GreenString("done"))
			}
			return
		case "fls":
			ShowFiles(cache.Files)
		case "dirs":
			ShowAllDirs(cache.Files)
		default:
			err = ChangeDir(dir)
			if err != nil {
				fmt.Println(color.RedString("	wrong path"))
				SHowCurrentDirPath()
			} else {
				c, err := os.ReadDir(".")
				if err != nil {
					fmt.Println("	", color.RedString(err.Error()))
					continue
				}
				ShowAllDirs(c)
				SHowCurrentDirPath()
				cache.Files = c
				filePath = filepath.Join(filePath, dir)
			}
		}
	}
}
func GetInput() (string, error) {
	var dir string
	_, err := fmt.Scan(&dir)
	if err != nil {
		return "", err
	}

	return dir, nil
}
func ChangeDir(dir string) error {
	err := os.Chdir(dir)
	if err != nil {
		return err
	}
	return nil
}
