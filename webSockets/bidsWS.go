package websockets

import (
	"drouotBack/models"
	"drouotBack/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// This is an attempt at a web socket connection manager. This is not functionning as is an requires redesigning.
var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	// A BidWebSocketPool is a map of ws that listen for a specific article.
	// This map contains all existing bidwebsocketpools that exist in order to keep track of them.
	BidWebSocketPools = make(map[string]bidWebSocketPool)
	// Bid service instance
	bidService = services.NewBidService()
	// Article service instance
	articleService = services.NewArticleService()
)

// Each websocket pool has the following methods :
type BidWebSocketPool interface {
	// Add a new ws to the pool
	WSAdd(*gin.Context)
	// Read incoming messages
	WSReader(conn *websocket.Conn)
	// Write a response to specific ws connection
	WSWriter(int, []byte, *websocket.Conn)
	// Broadcast to all existing websocket connections in the pool
	WSBroadcast(models.Bid)
}

// A bid websocket pool has two properties :
type bidWebSocketPool struct {
	// The id of the article concerned by the bids
	ArticleId string
	// A map of all active web socket connections currently listening (for broadcasting)
	Clients map[*websocket.Conn]string
}

// This function returns a bidwebsocketPool for the article specified
func GetBidWebSocketPool(articleId string) BidWebSocketPool {
	// If one already exists, return it
	val, ok := BidWebSocketPools[articleId]
	if ok {
		return val
	}
	// If one doesn't exist yet, create and configure a new one :
	var bidPool bidWebSocketPool
	// Set article id to specified value
	bidPool.ArticleId = articleId
	// Initialize the web socket connection map
	bidPool.Clients = make(map[*websocket.Conn]string)
	// Add the newly created bid ws pool to the map of existing pools
	BidWebSocketPools[articleId] = bidPool
	// Return the bid ws pool
	return bidPool
}

func (bwsp bidWebSocketPool) WSAdd(c *gin.Context) {
	// manage cors issues (quick and dirty for dev purposes)
	// !! Must correct before production !!
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// // Attempt at using sessions in order to avoid duplicate ws connections (not working...)
	// session := sessions.Default(c)
	// sessionId := session.Get("localSession")

	// In order to keep track of the different users subsribed to the pool, the session is will be used to identify each subscriber.
	// At the time of writing, sessionId is not implemented...
	sessionId, err := c.Cookie("sessionId")
	if err != nil || sessionId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not signed in"})
		log.Println("session id is nil")
	} else {
		// Open ws connection
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
		}

		// Check if connection is not already registered as subscriber and if not add connection to list of subscribers
		if _, ok := bwsp.Clients[ws]; !ok {
			bwsp.Clients[ws] = sessionId
			log.Println("connection added to pool")
		} else {
			// else log
			log.Println("Client Already Connected to : ", bwsp.ArticleId)
		}

		// Get current highest bid
		bid, err := articleService.GetHighestBidArticle(bwsp.ArticleId)
		if err != nil {
			log.Println(err)
		}
		// Send value to the new web socket connection
		ws.WriteJSON(bid)

		// For dev only
		log.Println("Client Successfully Connected to : ", bwsp.ArticleId)
		for key, value := range bwsp.Clients {
			log.Println(&key, &value)
		}
	}
}

func (bwsp bidWebSocketPool) WSReader(client *websocket.Conn) {
	// Infinite loop to listen for incoming messages
	for {
		var bid models.Bid
		if err := client.ReadJSON(&bid); err != nil {
			log.Println(err)
			bwsp.WSWriter(1, []byte(err.Error()), client)
			return
		}
		newBid, err := bidService.CreateBid(bid)
		if err != nil {
			log.Println(err)
			bwsp.WSWriter(1, []byte(err.Error()), client)
			return
		}
		bwsp.WSBroadcast(newBid)
	}
}

func (bwsp bidWebSocketPool) WSWriter(mt int, message []byte, client *websocket.Conn) {
	if err := client.WriteMessage(mt, message); err != nil {
		log.Println(err)
	}
}

func (bwsp bidWebSocketPool) WSBroadcast(bid models.Bid) {

	for subscriber := range bwsp.Clients {
		err := subscriber.WriteJSON(bid)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
