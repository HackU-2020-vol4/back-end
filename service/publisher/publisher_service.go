package publisher

import (
	"github.com/HackU-2020-vol4/back-end/db"
	"github.com/HackU-2020-vol4/back-end/entity"
	"github.com/HackU-2020-vol4/back-end/helpers"
	"github.com/gin-gonic/gin"
)

type Service struct{}

type Publisher entity.Publisher

func (s Service) CreateModel(c *gin.Context) (Publisher, error) {
	db := db.GetDB()
	var p Publisher
	var roomid string

	for {
		roomid = helpers.RandomString(20)
		// RoomIDがUniqueかチェック
		isRecord := RecordCheck(roomid)
		// TrueならOK
		if isRecord {
			break;
		}
	}
	p.RoomID = roomid
	if err := db.Create(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}

func RecordCheck(roomid string) bool {
	db := db.GetDB()
	var p Publisher
	var recordNotFound bool
	// RecordNotFound()はレコードに値がない時Trueを返す
	recordNotFound = db.Where("room_id = ?", roomid).First(&p).RecordNotFound()
	return recordNotFound
}