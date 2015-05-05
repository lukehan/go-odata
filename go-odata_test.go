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

var serviceRoot = "/odata/tests"

func TestAverage(t *testing.T) {
	port := "8080"

	restful.DefaultContainer.Router(restful.CurlyRouter{})

	ws := new(restful.WebService)
	ws.Path(serviceRoot)

	ws.Route(ws.GET("/").To(ReadTheServiceRoot).
	Doc("Read the service root").
	Operation("ReadTheServiceRoot").
	Writes(true).
	Produces(restful.MIME_JSON))
	log.Println("Publishing \"01. Read the service root\" method at " + serviceRoot + "/")

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