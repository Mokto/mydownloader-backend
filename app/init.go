package app

import (
	// "time"
	"github.com/revel/revel"
	"net/http"
	"golang.org/x/net/websocket"
	"fmt"
	"localserver/app/models"
	"localserver/app/services/websockets"
	"localserver/app/utils/links"
	"localserver/app/utils/debrid"
	"github.com/satori/go.uuid"
	"time"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.BeforeAfterFilter,       // Call the before and after filter functions
		revel.ActionInvoker,           // Invoke the action.
	}

	// Register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	// revel.OnAppStart(ExampleStartupScript)
	// revel.OnAppStart(InitDB)
	revel.OnAppStart(WatchAllDebrid)
	revel.OnAppStart(StartWebsockets)
	revel.OnAppStart(CheckLinksToDebrid)
}

// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	c.Response.Out.Header().Add("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
	c.Response.Out.Header().Add("Access-Control-Allow-Origin", "*")
	c.Response.Out.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Response.Out.Header().Add("Content-Type", "application/json; charset=UTF-8")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}

func StartWebsockets() {
	http.Handle("/", websocket.Handler(func (ws *websocket.Conn) {
		defer ws.Close()
		fmt.Println("Client Connected")

		connId := uuid.Must(uuid.NewV4()).String()
	
		websockets.AddConnection(ws, connId)
		go links.ListAndSend()
	
	
		for {
			var message string
			err := websocket.Message.Receive(ws, &message)
			if err != nil {
				websockets.RemoveConnection(connId)
				break
			}
		}
	}))

	fmt.Println("Websocket server is listening to : 9001")
	go http.ListenAndServe(":9001", nil)
}


func WatchAllDebrid() {
	go func() {
		for {
			links := links.GetAll()
			for _, link := range links {
				if (link.AllDebridID != 0 && (link.TorrentState != models.TORRENT_DONE)) {
					var debridInstance debrid.Debrid
					debridInstance = &debrid.AllDebrid{}
					fmt.Println("Updating statuses")
					debridInstance.UpdateStatuses(links)
					break
				}
			}
			time.Sleep(time.Second);
		}
	}()
}


func CheckLinksToDebrid() {
	go func() {
		for {
			linksToDebrid := links.GetAll()
			for  i, link := range linksToDebrid {
				if link.DownloadState == models.DOWNLOAD_NOT_READY && link.TorrentState == models.TORRENT_DONE {
					linksToDebrid[i].DownloadState = models.DOWNLOAD_DEBRIDING
					links.Save(linksToDebrid)
					for  j, textLink := range link.Links {
						fmt.Println(textLink)
						var debridInstance debrid.Debrid
						debridInstance = &debrid.AllDebrid{}
						linkDownloadable := debridInstance.GetDownloadableLink(textLink)
						
						linksToDebrid[i].Links[j] = linkDownloadable
					}
					linksToDebrid[i].DownloadState = models.DOWNLOAD_DOWNLOADING
				}
			}
			links.Save(linksToDebrid)
			links.ListAndSend()
			time.Sleep(10 * time.Second);
		}
	}()
}


	// for  i, link := range linksToCheck {
	// 	if link.DownloadState == models.DOWNLOAD_NOT_DEBRIDED {
	// 		for  j, textLink := range link.Links {

	// 		}
	// 	}
	// }