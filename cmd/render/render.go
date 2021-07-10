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

	views := map[string]gin.H{
		"index":    nil,
		"settings": nil,
		"commands": {
			"commands": config.Conf.Commands,
		},
		"tags": nil,
		"panels": nil,
		"premium": nil,
		"claiming": nil,
		"privacy": nil,
	}

	for view, data := range views {
		f, err := os.Create(fmt.Sprintf("./build/%s.html", view))
		utils.Must(err)

		utils.Must(render.CreateTemplate(view).ExecuteTemplate(f, "layout", data))
	}
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
