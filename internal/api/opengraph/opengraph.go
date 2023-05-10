package opengraph

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func GetOpengraph(imageUrl string, width int, height int, title string, description string, link string, plugin string) (result datatypes.OpenGraph, err error) {
	_, err = url.Parse(link)
	if err != nil {
		return result, err
	}

	result = datatypes.OpenGraph{
		Tags: []datatypes.OpenGraphTag{
			{Type: "og:image", Value: imageUrl},
			{Type: "og:image:width", Value: fmt.Sprintf("%d", width)},
			{Type: "og:image:height", Value: fmt.Sprintf("%d", height)},
			{Type: "og:title", Value: title},
			{Type: "og:description", Value: description},
			{Type: "og:url", Value: link},
			{Type: "og:site_name", Value: plugin},
		},
	}
	return result, nil
}

func OpengraphTag(c *gin.Context) {
	if !define.ENABLE_PLUGIN {
		c.JSON(http.StatusForbidden, "PLIGIN DISABLED")
		return
	}

	target := c.Query("url")
	if target == "" {
		c.JSON(http.StatusNotFound, "url is required")
		return
	}

	imageUrl := "https://m.media-amazon.com/images/M/MV5BMjIwNTM0Mzc5MV5BMl5BanBnXkFtZTgwMDk5NDU1ODE@._V1_.jpg"
	width := 1200
	height := 630
	title := "title"
	description := "Card description, some words are omitted here...................."
	url := target
	plugin := "plugin name"

	result, err := GetOpengraph(
		imageUrl, width, height,
		title, description, url,
		plugin,
	)

	if err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
