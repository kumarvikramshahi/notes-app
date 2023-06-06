package controllers

import (
	"fmt"
	"net/http"
	"notesApp/configs"
	"notesApp/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FetchNotes(ctx *gin.Context) {
	// get user email sent by validateJwtToken middleware
	userEmail, isPresent := ctx.Get("userEmail")
	if !isPresent {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// fetch all notes
	getNotes := []schemas.Notes{}
	email := fmt.Sprintf("%v", userEmail)
	result := configs.DefaultDB.Where("email = ?", email).Find(&getNotes)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": getNotes,
	})
}

func CreateNotes(ctx *gin.Context) {
	incomingNotes := &schemas.IncomingNotes{}
	// binding notes json to struct
	if err := ctx.ShouldBindJSON(incomingNotes); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// get user email sent by validateJwtToken middleware
	userEmail, isPresent := ctx.Get("userEmail")
	if !isPresent {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// create notes
	notes := &schemas.Notes{}
	notes.Email = fmt.Sprintf("%v", userEmail)
	notes.Note = incomingNotes.Note
	result := configs.DefaultDB.Create(&notes)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"nid": notes.NID,
	})
}

func DeleteNotes(ctx *gin.Context) {
	nid, err := strconv.ParseUint(ctx.Param("nid"), 10, 32)
	if err != nil {
		// Handle error if the parameter cannot be converted to a uint32
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid notes id(nid)",
		})
		return
	}

	// get user email sent by validateJwtToken middleware
	_, isPresent := ctx.Get("userEmail")
	if !isPresent {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	notes := &schemas.Notes{}
	notes.NID = uint32(nid)
	findNid := configs.DefaultDB.First(&notes)
	if findNid.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": findNid.Error.Error(),
		})
		return
	}
	result := configs.DefaultDB.Delete(&notes)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}
