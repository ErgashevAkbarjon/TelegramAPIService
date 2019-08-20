package main

/*
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/", mainPageHandler)
	http.ListenAndServe(":4343", nil)
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY")

	if err != nil {
		fmt.Println(err.Error())
	}
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Fprint(w, string(body))
	fmt.Println("Endpoint hit: home page")
}
*/

import (
	"github.com/zelenin/go-tdlib/client"
	"log"
	"path/filepath"
)

func WithLogs() client.Option {
	return func(tdlibClient *client.Client) {
		tdlibClient.SetLogVerbosityLevel(&client.SetLogVerbosityLevelRequest{
			NewVerbosityLevel: 1,
		})
	}
}

func main() {

	const (
		apiID   = 954242
		apiHash = "ed33628ede24f596c53edbecfd40ca0b"
	)

	authorizer := client.ClientAuthorizer()
	go client.CliInteractor(authorizer)

	authorizer.TdlibParameters <- &client.TdlibParameters{
		UseTestDc:              false,
		DatabaseDirectory:      filepath.Join(".tdlib", "database"),
		FilesDirectory:         filepath.Join(".tdlib", "files"),
		UseFileDatabase:        true,
		UseChatInfoDatabase:    true,
		UseMessageDatabase:     true,
		UseSecretChats:         false,
		ApiId:                  apiID,
		ApiHash:                apiHash,
		SystemLanguageCode:     "en",
		DeviceModel:            "Server",
		SystemVersion:          "1.0.0",
		ApplicationVersion:     "1.0.0",
		EnableStorageOptimizer: true,
		IgnoreFileNames:        false,
	}

	tdlibClient, err := client.NewClient(authorizer, WithLogs())
	if err != nil {
		log.Fatalf("NewClient error: %s", err)
	}

	listener := tdlibClient.GetListener()
	defer listener.Close()

	for update := range listener.Updates {

		if update.GetClass() == client.ClassUpdate {

			statusUpdate, _ := update.(*client.UpdateUserStatus)

			log.Printf("%v", statusUpdate.Type)

		}

	}

	// log.Printf("Me: %s %s [%s]", me.FirstName, me.LastName, me.Username)
	// log.Println(me)
}
