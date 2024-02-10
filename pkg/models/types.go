package models

import (
	"gorm.io/gorm"
)

type BasicAuth struct {
	Login, Password string
}

type Answer struct {
	Answer string `json:"answer"`
}

func (p *Pupil) BeforeUpdates(db *gorm.DB) (err error) {
	// finish making logic
	//r.Database.SavePoint()
	//r.Database.Commit()
	//r.Database.RollbackTo()
	return nil
}
