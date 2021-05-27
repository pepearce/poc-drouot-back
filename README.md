# poc-drouot-back
Backend golang app for mini drouot

# Installation

1) Install Go  
2) Create a mysql database for the app  
3) Configure the session key in "session.env"
4) Configure the database access credentials in "/models/db.env"  
5) In a terminal cd to source folder and run : 
```
go get ./...
go build main.go && main
```
>You can also run :
>```
>go run main.go
>```
>However, under Windows, this will prompt a firewall authorization window on each launch.  

The API should be running  

Once this is done, you can start the [front end app](https://github.com/pepearce/poc-drouot-front)  
