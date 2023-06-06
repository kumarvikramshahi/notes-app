package schemas

type User struct {
	Email    string `form:"email" json:"email" gorm:"primaryKey;not null;size:50" binding:"required"`
	Password string `form:"password" json:"password,omitempty" gorm:"not null;size:50" binding:"required"`
	Name     string `form:"name" json:"name" gorm:"not null;size:50"`
}

type Notes struct {
	NID   uint32 `form:"nid" json:"nid" gorm:"primaryKey;autoIncrement;not null"`
	Email string `form:"email" json:"email,omitempty" binding:"required"`
	Note  string `form:"note" json:"note,omitempty" gorm:"size:400" binding:"required"`
}

type IncomingNotes struct {
	Note string `form:"note" json:"note,omitempty" gorm:"size:400" binding:"required"`
}
