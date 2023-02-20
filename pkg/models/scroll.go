package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/nezaYSR/go-mux-sql/pkg/config"
)

var db *gorm.DB

type Scroll struct {
	gorm.Model
	Title        string `json:"title"`
	MagicianName string `json:"magician_name"`
	Element      string `json:"element" sql:"type:VARCHAR(10) CHECK (element IN ('fire', 'air', 'water', 'thunder'))"`
	Rarity       string `json:"rarity" sql:"type:VARCHAR(10) CHECK (rarity IN ('junk', 'common', 'uncommon', 'rare', 'mythical', 'godly'))"`
	Price        int32  `json:"price"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Scroll{})
}

func (s *Scroll) CreateScroll() (*Scroll, error) {
	db.NewRecord(s)
	db.Create(&s)
	if err := s.BeforeCreate(db); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Scroll) BeforeCreate(tx *gorm.DB) (err error) {
	if !s.isElementValid() {
		return fmt.Errorf("invalid element value: %s", s.Element)
	}
	if !s.isRarityValid() {
		return fmt.Errorf("invalid rarity value: %s", s.Rarity)
	}
	return nil
}

func (s *Scroll) isElementValid() bool {
	switch s.Element {
	case "fire", "air", "water", "thunder":
		return true
	default:
		return false
	}
}

func (s *Scroll) isRarityValid() bool {
	switch s.Rarity {
	case "junk", "uncommon", "common", "rare", "mythical", "godly":
		return true
	default:
		return false
	}
}

func GetAllScrolls() []Scroll {
	var Scrolls []Scroll
	db.Find(&Scrolls)
	return Scrolls
}

func GetScrollById(id int64) (*Scroll, *gorm.DB) {
	var getScroll Scroll
	db := db.Where("ID=?", id).Find(&getScroll)
	return &getScroll, db
}

func DeleteScroll(id int64) Scroll {
	var scroll Scroll
	db.Where("ID=?", id).Delete(scroll)
	return scroll
}
