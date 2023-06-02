package main

import "github.com/labstack/echo/v4"

type Daftaranime struct {
	Judul string `json:"judul"`
	Tahun int    `json:"tahun rilis"`
}

type BaseResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data anime"`
}

type CariAnimeRequest struct {
	Judul string `json:"judul"`
	Tahun int    `json:"tahun"`
}

func main() {
	e := echo.New()

	// routing (mengarahkan)
	e.GET("/daftaranime", GetAnimeController)
	e.POST("/carianime", FindAnimeController)
	e.Start(":8000")
}

func FindAnimeController(c echo.Context) error {

	judul := c.FormValue("judul")
	tahun := c.FormValue("tahun")
	var anime = Daftaranime{Judul: judul + " tahun rilis: " + tahun, Tahun: 1999}
	return c.JSON(200, anime)
}

func GetAnimeController(c echo.Context) error {

	var dataAnime []Daftaranime

	var anime Daftaranime
	anime.Judul = "One Piece"
	anime.Tahun = 1999
	dataAnime = append(dataAnime, anime)

	anime.Judul = "Demon Slayer: Kimetsu no Yaiba"
	anime.Tahun = 2019
	dataAnime = append(dataAnime, anime)

	var response BaseResponse
	response.Data = dataAnime
	response.Message = "Berhasil"

	return c.JSON(200, response)
}
