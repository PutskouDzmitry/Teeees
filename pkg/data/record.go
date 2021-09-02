package data

import (
	"fmt"
	"gorm.io/gorm"
)

// Visit represents database table "visits_scud".
type Visit struct {
	Id              int    `gorm:"id" json:"id"`
	CharId          string `gorm:"_id" json:"_id"`
	EmployeeId      string `gorm:"employeeId" json:"employeeId,omitempty"`
	TimePACS        string `gorm:"timePACS" json:"timePACS,omitempty"`
	TimeUTC         string `gorm:"timeUTC" json:"timeUTC,omitempty"`
	Access          string `gorm:"access" json:"access,omitempty"`
	CodeDevice      string `gorm:"codeDevice" json:"codeDevice,omitempty"`
	Device          string `gorm:"device" json:"device,omitempty"`
	Card            string `gorm:"card" json:"card,omitempty"`
	Personal        string `gorm:"personal" json:"personal,omitempty"`
	VisitorType     int    `gorm:"visitorType" json:"visitorType,omitempty"`
	Position        string `gorm:"position" json:"position,,omitempty"`
	SpaceId         string `gorm:"spaceId" json:"spaceId,omitempty"`
	EventType       string `gorm:"eventType" json:"eventType,omitempty"`
	OsmUnitId       string `gorm:"osmUnitId" json:"osmUnitId,omitempty"`
	TimeDeviceLocal string `gorm:"timeDeviceLocal" json:"timeDeviceLocal,omitempty"`
}

// VisitData allows performing database operations.
type VisitData struct {
	db *gorm.DB
}

// NewVisitData creates new VisitData instance.
func NewVisitData(db *gorm.DB) *VisitData {
	return &VisitData{db: db}
}

// Read gets list of visits by passing search parameters from database.
func (d VisitData) Read(from, to, spaceId string) ([]Visit, error) {
	var records []Visit
	result := d.db.Raw(`SELECT * FROM visits_scud WHERE "timeUTC" > ? AND "timeUTC" < ? AND "spaceId" = ?;`, from, to, spaceId).Scan(&records)
	if result.Error != nil {
		return nil, result.Error
	}
	return records, nil
}

func (d VisitData) GetCount(field string) ([]int, error) {
	var id []int
	query := fmt.Sprintf("SELECT COUNT(%v) FROM visits_scud", field)
	result := d.db.Raw(query).Scan(&id)
	if result.Error != nil {
		return nil, result.Error
	}
	return id, nil

}

func (d VisitData) Update(field string, value string, toId string, fromId string) error {
	result := d.db.Raw(`UPDATE visits_scud SET ? = ? WHERE id < ? AND id > ?`, field, value, toId, fromId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d VisitData) Delete(id string) error {
	result := d.db.Raw(`DELETE FROM visits_scud WHERE id=?`, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
