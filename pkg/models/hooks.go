package models

import (
	"gorm.io/gorm"
)

func (p *Pupil) BeforeUpdates(db *gorm.DB) (err error) {
	// finish making logic
	//r.Database.SavePoint()
	//r.Database.Commit()
	//r.Database.RollbackTo()
	// если роль изменилась
	//if tx.Statement.Changed("Role") {
	//	return errors.New("role not allowed to change")
	//}
	//
	//if tx.Statement.Changed("Name", "Admin") { если имя или роль изменились
	//	tx.Statement.SetColumn("Age", 18)
	//}
	//
	//// если любое поле изменилось
	//if tx.Statement.Changed() {
	//	tx.Statement.SetColumn("RefreshedAt", time.Now())
	//}
	//return nil

	return nil
}
