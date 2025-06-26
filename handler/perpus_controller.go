package handler

import (
	"fmt"
	"perpustakaan/contract"
	"perpustakaan/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type perpusController struct {
	service contract.PerpusService
}

func (c *perpusController) getPrefix() string {
	return "/perpus"
}

func (c *perpusController) initService(service *contract.Service) {
	c.service = service.Perpus
}

func (c *perpusController) initRoute(app *gin.RouterGroup) {
	app.GET("/:id", c.GetPerpus)
	app.POST("/create", c.CreatePerpus)
	app.PUT("/:id", c.UpdatePerpus)
	app.PATCH("/:id", c.UpdatePerpus)
	app.DELETE("/:id", c.DeletePerpus)
	app.GET("/:id/pinjam", c.PinjamBuku)
	app.POST("/:id/kembali", c.KembalikanBuku)
	app.GET("/", c.CariJudul)
}

func (c *perpusController) GetPerpus(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	response, err := c.service.GetPerpus(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *perpusController) CreatePerpus(ctx *gin.Context) {
	var payload dto.PerpusRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.service.CreatePerpus(&payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *perpusController) UpdatePerpus(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var payload dto.PerpusRequest
	err = ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.service.UpdatePerpus(intID, &payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (c *perpusController) DeletePerpus(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	response, err := c.service.DeletePerpus(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *perpusController) PinjamBuku(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	response, err := c.service.PinjamBuku(intID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *perpusController) KembalikanBuku(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File tidak ditemukan"})
		return
	}

	dst := fmt.Sprintf("returned/buku_%d_%s", intID, file.Filename)
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
		return
	}

	err = c.service.KembalikanBuku(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":     "Buku berhasil dikembalikan",
		"file_stored": dst,
	})
}

func (c *perpusController) CariJudul(ctx *gin.Context) {
	judul := ctx.Query("judul")

	var result *dto.PerpusListResponse
	var err error

	if judul != "" {
		result, err = c.service.SearchPerpusByJudul(judul)
	} else {
		result, err = c.service.CariJudul()
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}