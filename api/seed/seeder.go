package seed

import (
	"log"

	"github.com/AlexSwiss/prentice/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Firstname: "Alexander",
		Lastname:  "Ikeh",
		Email:     "aleks@gmail.com",
		Phone:     "08182432388",
		Gender:    "M",
		Country:   "Nigeria",
		Bio:       "Fullstack developer with enthusis=asm to solve problems",
		Password:  "password",
	},
	models.User{
		Firstname: "Jon",
		Lastname:  "Snow",
		Email:     "snow@gmail.com",
		Phone:     "27002983999",
		Gender:    "F",
		Country:   "Nigerian",
		Bio:       "Front end developer with an eye for good design",
		Password:  "password",
	},
}

var shops = []models.Shop{
	models.Shop{
		AdminName:   "Chris",
		AdminEmail:  "chris@gmail.com",
		ShopName:    "09039947634",
		Stack:       "React, Node.js, AWS, MongoDB",
		TeamMembers: "3",
		Bio:         "Startup digging deep into the edu space",
		Password:    "password",
	},
	models.Shop{
		AdminName:   "Swiss",
		AdminEmail:  "swiss@gmail.com",
		ShopName:    "090399485834",
		Stack:       "React, Node.js, AWS",
		TeamMembers: "5",
		Bio:         "Startup digging deep into the fitness space",
		Password:    "password",
	},
}

// Load data
func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}, &models.Shop{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Shop{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	for i, _ := range shops {
		err = db.Debug().Model(&models.Shop{}).Create(&shops[i]).Error
		if err != nil {
			log.Fatalf("cannot seed shop table: %v", err)
		}
	}
}
