package seed

import (
	"log"

	"github.com/danurwijayanto/golang-api-with-jwt-and-mysql/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Nickname:     "Steven victor",
		Email:        "steven@gmail.com",
		Password:     "password",
		UserCategory: 1,
	},
	models.User{
		Nickname:     "Martin Luther",
		Email:        "luther@gmail.com",
		Password:     "password",
		UserCategory: 3,
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	models.Post{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

var product = []models.Product{
	models.Product{
		Name:        "Lobster",
		Stock:       5,
		Description: "Lobster Laut",
		Price:       120000,
	},
	models.Product{
		Name:        "Tongkol",
		Stock:       10,
		Description: "Tongkol Laut",
		Price:       50000,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}, &models.Product{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.Product{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}

		err = db.Debug().Model(&models.Product{}).Create(&product[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
