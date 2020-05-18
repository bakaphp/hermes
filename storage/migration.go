package storage

import (
	"src/feeds/models"
	"src/feeds/services"
)

// MigrateDB : Table migrations
func MigrateDB() {

	db := services.MysqlConnect()

	usersFollows := models.UsersFollows{}
	usersMessages := models.UserMessages{}
	groupMessages := models.GroupMessages{}

	db.Debug().CreateTable(&usersFollows)
	db.Debug().CreateTable(&usersMessages)
	db.Debug().CreateTable(&groupMessages)
}

// SeedDB database
func SeedDB() {

	db := services.MysqlConnect()

	usersFollowers := models.UsersFollows{
		UsersID: 1, EntityID: 4, EntityNamespace: "Canvas\\Models\\Users", IsDeleted: 0,
	}
	usersMessages := models.UserMessages{
		MessageID: 1, UsersID: 1, IsDeleted: 0,
	}
	groupMessages := models.GroupMessages{
		MessageID: 1, GroupID: 1, IsDeleted: 0,
	}

	// for _, usersFollower := range usersFollowers {
	db.Debug().Create(&usersFollowers)
	// }

	// for _, usersMessage := range usersMessages {
	db.Debug().Create(&usersMessages)
	// }

	// for _, groupMessage := range groupMessages {
	db.Debug().Create(&groupMessages)
	// }

}
