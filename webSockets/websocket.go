package websockets

import (
	"drouotBack/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// This is the new version of the websocket controller

var (
	upgrader2     = websocket.Upgrader{}
	websocketPool = make(map[*websocket.Conn]bool)
)

// Open a new websocket
func OpenSocket(c *gin.Context) {
	upgrader2.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader2.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatalln("Connection upgrade failed")
	}
	log.Println("Connection successfully upgraded")
	AddSocketToPool(ws)
	WSReader(ws)
}

// Remove a socket from the websocket pool
func RemoveSocket(ws *websocket.Conn) {
	delete(websocketPool, ws)
	log.Println("Socket removed from pool")
	for key := range websocketPool {
		log.Println(&key)
	}
}

// Add a socket to the websocket pool
func AddSocketToPool(ws *websocket.Conn) {

	// If socket is already present in pool, skip add, else add to map
	if _, ok := websocketPool[ws]; ok {
		log.Println("Socket already in pool")
		return
	} else {
		log.Println("Adding socket to pool")
		websocketPool[ws] = true
		log.Println("Socket added to pool")
		for key := range websocketPool {
			log.Println(&key)
		}
	}

}

// Start listening to a websocket connection
func WSReader(ws *websocket.Conn) {
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			// if the error is due to client closing tab
			if err.Error() == "websocket: close 1001 (going away)" {
				log.Println("Websocket closed due to client going away")
				RemoveSocket(ws)
			} else {
				log.Println("Error occured : ", err.Error())
			}
			return
		}

		type bidActionWS struct {
			Message string     `json:"message"`
			Id      int        `json:"id"`
			Bid     models.Bid `json:"bid"`
		}
		var action bidActionWS
		// Unmarshal the message to get access to the action and the payload
		json.Unmarshal([]byte(message), &action)
		log.Println("message type : ", mt, "   message: ", action.Message)
		// if err := ws.ReadJSON(&action); err != nil {
		// 	if err == err.(*websocket.CloseError) {
		// 		ws.Close()
		// 	}
		// 	log.Println("error : ", err.Error())
		// }

		// If action is GETBID, get bid with article id specified in payload
		if action.Message == GETBID {
			bid, err := articleService.GetHighestBidArticle(fmt.Sprint(action.Id))
			log.Println("bid: ", fmt.Sprint(bid))
			if err != nil {
				log.Println("Error : ", err.Error())
			}
			if bid.ID == 0 {
				ws.WriteJSON(0)
			} else {
				ws.WriteJSON(bid)
			}

			log.Println("message : ", action.Message, "  id : ", action.Id)
		}

		// If action is SETBID, create a new bid with the payload
		if action.Message == SETBID {
			log.Println("message : ", action.Message, "  bid : ", action.Bid)
			bid, err := bidService.CreateBid(action.Bid)
			if err != nil {
				log.Println(err.Error())
			}
			WSBroadcaster(bid)
		}
	}

}

// Broad cast new bid to the websocket connection pool
func WSBroadcaster(bid models.Bid) {
	for ws := range websocketPool {
		log.Println("Broadcasting : ", bid, " to :", &ws)
		ws.WriteJSON(&bid)
	}
}
