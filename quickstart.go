package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v2"
	"os/exec"
	"runtime"

	"github.com/pandreyn/go-grive/gdrive"
	"strconv"
)

func openBrowser(hostname string, port int) error {

	var err error
	address := "http://"+hostname+":"+strconv.Itoa(port)+"/"

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", address).Start()
	case "windows", "darwin":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", address).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return err
}

func getFilesFromGDrive() {
	ctx := context.Background()

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, drive.DriveMetadataReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := gdrive.GetClient(ctx, config)

	srv, err := drive.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve drive Client %v", err)
	}

	r, err := srv.Files.List().MaxResults(10).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files.", err)
	}

	fmt.Println("Files:")
	if len(r.Items) > 0 {
		for _, i := range r.Items {
			fmt.Printf("%s (%s)\n", i.Title, i.Id)
		}
	} else {
		fmt.Print("No files found.")
	}

}
