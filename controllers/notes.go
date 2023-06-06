package controllers

import (
	"log"
	"net/http"
	"notesApp/configs"
	"notesApp/schemas"
	"notesApp/validators"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FetchNotes(ctx *gin.Context) {
	// get and validate session id
	incomingUser := &schemas.User{}
	// if err := ctx.ShouldBindJSON(incomingUser); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }
	// TO DO

	requestUser, _ := ctx.Get("user")

	// fetch all notes
	getNotes := []schemas.Notes{}
	result := configs.DefaultDB.Where("email = ?", requestUser.Email).Find(&getNotes)
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
	notes := &schemas.Notes{}
	incomingNotes := &schemas.IncomingNotes{}
	// binding notes json to struct
	if err := ctx.ShouldBindJSON(incomingNotes); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// get and validate session id
	email := validators.EmailExtractor(incomingNotes.Sid)

	// create notes
	notes.Email = email
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
	sid := ctx.Param("sid")
	nid, err := strconv.ParseUint(ctx.Param("nid"), 10, 32)
	if err != nil {
		// Handle error if the parameter cannot be converted to a uint32
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid notes id(nid)",
		})
		return
	}

	notes := &schemas.Notes{}
	notes.Email = email
	notes.NID = uint32(nid)
	result := configs.DefaultDB.Delete(&notes)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	log.Println("deleted rows", result.RowsAffected)
	ctx.JSON(http.StatusOK, gin.H{
		"message": notes,
	})

}
