package websockets

import (
	"drouotBack/models"
	"drouotBack/services"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	BidWebSocketPools = make(map[string]bidWebSocketPool)
	bidService        = services.NewBidService()
)

type BidWebSocketPool interface {
	WSAdd(*gin.Context)
	WSReader(conn *websocket.Conn)
	WSWriter(int, []byte, *websocket.Conn)
	WSBroadcast(models.Bid)
}

type bidWebSocketPool struct {
	ArticleId string
	Clients   map[*websocket.Conn]string
}

func GetBidWebSocketPool(articleId string) BidWebSocketPool {
	val, ok := BidWebSocketPools[articleId]
	if ok {
		return val
	}
	var bidPool bidWebSocketPool
	bidPool.ArticleId = articleId
	bidPool.Clients = make(map[*websocket.Conn]string)
	BidWebSocketPools[articleId] = bidPool
	return bidPool
}

func (bwsp bidWebSocketPool) WSAdd(c *gin.Context) {
	// manage cors issues (quick and dirty for dev purposes)
	// !! Must correct before production !!
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	session := sessions.Default(c)
	sessionId := session.Get("localSession")
	if sessionId == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not signed in"})
		log.Println("session-id is nil")
	} else {
		// Open ws connection
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
		}

		// Check if connection is not already registered as subscriber and if not add connection to list of subscribers
		if _, ok := bwsp.Clients[ws]; !ok {
			if bwsp.Clients[ws] != sessionId {
				bwsp.Clients[ws] = sessionId.(string)
				log.Println("connection added to pool")
			}
		}
		// else log
		log.Println("Client Already Connected to : ", bwsp.ArticleId)
		bid, err := services.NewArticleService().GetHighestBidArticle(bwsp.ArticleId)
		if err != nil {
			log.Println(err)
		}
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
