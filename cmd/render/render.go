package main

import (
	"fmt"
	"github.com/TicketsBot/website/config"
	"github.com/TicketsBot/website/render"
	"github.com/TicketsBot/website/utils"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	config.LoadConfig()

	utils.Must(os.RemoveAll("./build"))
	utils.Must(os.MkdirAll("./build", os.ModePerm))

	// copyFile assets
	utils.Must(filepath.Walk("./public/static", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		return copyFile(path, strings.Replace(path, "public/static", "./build", -1))
	}))

	data := gin.H{
		"commands": config.Conf.Commands,
	}

	utils.Must(filepath.Walk("./public/templates/views", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		lastIdx := strings.LastIndex(path, string(os.PathSeparator))
		name := strings.Replace(path[lastIdx+1:], ".tmpl", "", 1)

		f, err := os.Create(fmt.Sprintf("./build/%s.html", name))
		utils.Must(err)

		utils.Must(render.CreateTemplate(name).ExecuteTemplate(f, "layout", data))
		return nil
	}))
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	lastIdx := strings.LastIndex(dst, string(os.PathSeparator))
	if err := os.MkdirAll(dst[:lastIdx], os.ModePerm); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
