package models

import (
	"drouotBack/security"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// Database access configuration
type Config struct {
	sqlUser     string
	sqlPassword string
	host        string
	dbName      string
	port        string
	dbType      string
}

var DB *gorm.DB

func ConnectDataBase() {

	// load credentials from db.env file
	err := godotenv.Load("./models/db.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// Database access configuration
	var config = Config{
		sqlUser:     os.Getenv("DB_USER"),
		sqlPassword: os.Getenv("DB_PASSWORD"),
		host:        os.Getenv("HOST"),
		dbName:      os.Getenv("DB_NAME"),
		port:        os.Getenv("PORT"),
		dbType:      os.Getenv("DB_TYPE"),
	}
	// Set database access parameters
	connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.sqlUser, config.sqlPassword, config.host, config.port, config.dbName)

	//database, err := gorm.Open("mysql", "sqlDev:pa$$word123@tcp(127.0.0.1:3306)/drouotDB")
	// Create a connection to the database
	database, err := gorm.Open(config.dbType, connArgs)

	if err != nil {
		log.Fatal(err)
		panic("Failed to connect to database!")
	}

	DB = database

	// Drop all tables (for development purposes only !)
	dropTables()

	// Create tables in the database corresponding to the models
	createTables()

	// Create foreign keys in tables depending on the relations between entities
	foreignKeys()

	// Add unique constraints
	uniqueConstraints()

	//Create composite primary keys
	// primaryKeys()

	// Populate tables (for development purposes only !)
	initEntries()

}

func createTables() {
	DB.AutoMigrate(
		&Article{},
		&Auction{},
		&User{},
		&Bid{},
	)
}

// func primaryKeys() {
// 	DB.Exec("alter table bids add constraint PK_bids Primary key (bidDate,userID,articleID)")
// }

func uniqueConstraints() {
	DB.Exec("ALTER TABLE users ADD UNIQUE UN_users_email(email)")
}

func foreignKeys() {
	DB.Model(&Article{}).AddForeignKey("AuctionID", "Auctions(ID)", "CASCADE", "CASCADE")
	DB.Model(&Bid{}).AddForeignKey("UserID", "Users(ID)", "CASCADE", "CASCADE")
	DB.Model(&Bid{}).AddForeignKey("ArticleID", "Articles(ID)", "CASCADE", "CASCADE")
	DB.Model(&Auction{}).AddForeignKey("UserID", "Users(ID)", "CASCADE", "CASCADE")

}

func dropTables() {

	// Remove existing foreign keys
	DB.Model(&Article{}).RemoveForeignKey("AuctionID", "Auctions(ID)")
	DB.Model(&Article{}).RemoveForeignKey("HighestBidID", "Bid(ID)")
	DB.Model(&Bid{}).RemoveForeignKey("UserID", "Users(ID)")
	DB.Model(&Bid{}).RemoveForeignKey("ArticleID", "Articles(ID)")
	DB.Model(&Auction{}).RemoveForeignKey("UserID", "Users(ID)")

	// Drop tables if they exist
	// For development purposes only
	DB.DropTableIfExists("users")
	DB.DropTableIfExists("articles")
	DB.DropTableIfExists("auctions")
	DB.DropTableIfExists("bids")

}

func initEntries() {
	// Users
	password1, _ := security.Hash("password123")
	user1 := User{
		FirstName: "Paul",
		LastName:  "Pearce",
		Address:   "70 rue mouffetard",
		Email:     "paul.vpearce@gmail.com",
		Password:  string(password1),
		Role:      "admin",
	}
	password2, _ := security.Hash("password321")
	user2 := User{
		FirstName: "John",
		LastName:  "Doe",
		Address:   "1 rue nimp",
		Email:     "john@doe.com",
		Password:  string(password2),
	}

	DB.Model(&User{}).Create(&user1)
	DB.Model(&User{}).Create(&user2)
	// Auctions
	auction1 := Auction{Title: "Auction 1", UserID: 1, Category: "beaux_arts", StartDate: time.Date(2022, 10, 8, 8, 30, 0, 0, time.Local), EndDate: time.Date(2022, 11, 8, 8, 30, 0, 0, time.Local), PhotoURL: "https://source.unsplash.com/collection/190728/400x550"}
	auction2 := Auction{Title: "Auction 2", UserID: 1, Category: "mobilier_obj_art", StartDate: time.Date(2022, 10, 8, 8, 30, 0, 0, time.Local), EndDate: time.Date(2022, 11, 8, 8, 30, 0, 0, time.Local), PhotoURL: "https://source.unsplash.com/collection/190727/400x550"}
	DB.Model(&Auction{}).Create(&auction1)
	DB.Model(&Auction{}).Create(&auction2)

	// Articles
	article1 := Article{
		Title:           "Article 1",
		Description:     "This is article 1 of auction 1 !!",
		AuctionId:       1,
		Estimation:      200,
		InitialOffering: 30,
		PhotoPath:       "https://source.unsplash.com/collection/190728/400x550"}
	article2 := Article{
		Title:           "Article 2",
		Description:     "This is article 2 of auction 1 !!",
		AuctionId:       1,
		Estimation:      350,
		InitialOffering: 50,
		PhotoPath:       "https://source.unsplash.com/collection/190725/400x550"}
	article3 := Article{
		Title:           "Article 3",
		Description:     "This is article 3 of auction 1 !!",
		AuctionId:       1,
		Estimation:      200,
		InitialOffering: 30,
		PhotoPath:       "https://source.unsplash.com/collection/190727/400x550"}
	article4 := Article{
		Title:           "Article 4",
		Description:     "This is article 4 of auction 1 !!",
		AuctionId:       1,
		Estimation:      350,
		InitialOffering: 50,
		PhotoPath:       "https://source.unsplash.com/collection/190726/400x550"}
	article5 := Article{
		Title:           "Article 5",
		Description:     "This is article 5 of auction 1 !!",
		AuctionId:       1,
		Estimation:      200,
		InitialOffering: 30,
		PhotoPath:       "https://source.unsplash.com/collection/190728/400x550"}
	article6 := Article{
		Title:           "Article 6",
		Description:     "This is article 6 of auction 1 !!",
		AuctionId:       1,
		Estimation:      350,
		InitialOffering: 50,
		PhotoPath:       "https://source.unsplash.com/collection/190725/400x550"}
	article7 := Article{
		Title:           "Article 7",
		Description:     "This is article 7 of auction 1 !!",
		AuctionId:       1,
		Estimation:      200,
		InitialOffering: 30,
		PhotoPath:       "https://source.unsplash.com/collection/190727/400x550"}
	article8 := Article{
		Title:           "Article 8",
		Description:     "This is article 8 of auction 1 !!",
		AuctionId:       1,
		Estimation:      350,
		InitialOffering: 50,
		PhotoPath:       "https://source.unsplash.com/collection/190726/400x550"}

	article9 := Article{
		Title:           "Article 1",
		Description:     "This is article 1 of auction 2!!",
		AuctionId:       2,
		Estimation:      1200,
		InitialOffering: 300,
		PhotoPath:       "https://source.unsplash.com/collection/190727/400x550"}
	article10 := Article{
		Title:           "Article 2",
		Description:     "This is arcticle 2 of auction 2 !!",
		AuctionId:       2,
		Estimation:      800,
		InitialOffering: 230,
		PhotoPath:       "https://source.unsplash.com/collection/190726/400x550"}
	article11 := Article{
		Title:           "Article 3",
		Description:     "This is article 3 of auction 2 !!",
		AuctionId:       2,
		Estimation:      200,
		InitialOffering: 30,
		PhotoPath:       "https://source.unsplash.com/collection/190727/400x550"}
	article12 := Article{
		Title:           "Article 4",
		Description:     "This is article 4 of auction 2 !!",
		AuctionId:       2,
		Estimation:      350,
		InitialOffering: 50,
		PhotoPath:       "https://source.unsplash.com/collection/190726/400x550"}

	DB.Model(&Article{}).Create(&article1)
	DB.Model(&Article{}).Create(&article2)
	DB.Model(&Article{}).Create(&article3)
	DB.Model(&Article{}).Create(&article4)
	DB.Model(&Article{}).Create(&article5)
	DB.Model(&Article{}).Create(&article6)
	DB.Model(&Article{}).Create(&article7)
	DB.Model(&Article{}).Create(&article8)
	DB.Model(&Article{}).Create(&article9)
	DB.Model(&Article{}).Create(&article10)
	DB.Model(&Article{}).Create(&article11)
	DB.Model(&Article{}).Create(&article12)

	// Bids
	bid1 := Bid{
		UserID:    1,
		BidDate:   time.Now(),
		ArticleID: 1,
		BidAmount: 40,
	}
	bid2 := Bid{
		UserID:    2,
		BidDate:   time.Now(),
		ArticleID: 1,
		BidAmount: 50,
	}
	bid3 := Bid{
		UserID:    1,
		BidDate:   time.Now().Add(30000000000),
		ArticleID: 1,
		BidAmount: 60,
	}
	bid4 := Bid{
		UserID:    1,
		BidDate:   time.Now(),
		ArticleID: 4,
		BidAmount: 240,
	}
	DB.Model(&Bid{}).Create(&bid1)
	DB.Model(&Bid{}).Create(&bid2)
	DB.Model(&Bid{}).Create(&bid3)
	DB.Model(&Bid{}).Create(&bid4)
}
