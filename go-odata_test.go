package odata
import (
	"testing"
	"github.com/emicklei/go-restful"
	"github.com/hydrogen18/stoppableListener"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"sync"
)

var serviceRoot = "/odata/examples"

func TestAverage(t *testing.T) {
	port := "8080"

	restful.DefaultContainer.Router(restful.CurlyRouter{})

	ws := new(restful.WebService)
	ws.Path(serviceRoot)

	ws.Route(ws.GET("/").To(ReadTheServiceRoot).
	Doc("101.1 Read the service root").
	Operation("ReadTheServiceRoot").
	Writes("v4.ResponseResourceList").
	Produces(restful.MIME_JSON))
	log.Println("Publishing \"101-1. Read the service root\" method at " + serviceRoot + "/")

	ws.Route(ws.GET("/People").To(ReadAnEntitySet).
	Doc("101.3 Read an entity set").
	Operation("ReadAnEntitySet").
	Writes("v4.ResponseEntrySet").
	Produces(restful.MIME_JSON))
	log.Println("Publishing \"101-3. Read an entity set\" method at " + serviceRoot + "/People")

	ws.Route(ws.GET("/{resource:People\\('.+'\\)}").To(GetASingleEntityFromACollection).
	Doc("101.4 Get a single entity from a collection").
	Operation("ReadAnEntitySet").
	Param(ws.PathParameter("resource", "Resource").
	DataType("string").
	Required(true)).
	Writes("Man").
	Produces(restful.MIME_JSON))
	log.Println("Publishing \"101-4. Get a single entity from a collection\" method at " + serviceRoot + "/People('{id}')")

	restful.Add(ws)

	originalListener, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatal(err)
	}
	sl, err := stoppableListener.New(originalListener)
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{}
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT)
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		defer wg.Done()
		server.Serve(sl)
	}()
	log.Println("Started listening on localhost:", port)

	select {
	case signal := <-stop:
		log.Println("Got signal: ", signal)
	}

	log.Println("Stopping listener")
	sl.Stop()

	log.Println("Waiting on server")
	wg.Wait()
}