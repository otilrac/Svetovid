package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/akamensky/argparse"
	"github.com/atotto/clipboard"
	"github.com/fatih/color"
	. "github.com/redcode-labs/Coldfire"
)

var green = color.New(color.FgGreen).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var cyan = color.New(color.FgBlue).SprintFunc()

func StartServer(port, dir string) error {
	file_server := http.FileServer(http.Dir(dir))
	return http.ListenAndServe(":"+port, file_server)
}

func main() {
	fmt.Println()
	fmt.Println("\t\t/// Svetovid ///")
	fmt.Println()
	fmt.Println("\tCreated by: redcodelabs.io", red("<*>"))
	fmt.Println()
	parser := argparse.NewParser("svetovid", "")
	var port *int = parser.Int("p", "port", &argparse.Options{Default: 8080, Help: "Local port to start server on"})
	var tunnel *bool = parser.Flag("t", "tunnel", &argparse.Options{Required: false, Help: "Enable Ngrok tunelling"})
	var dir *string = parser.String("d", "dir", &argparse.Options{Required: false, Default: ".", Help: "Directory to serve"})
	var clip *bool = parser.Flag("c", "clip", &argparse.Options{Required: false, Help: "Copy server's URL to clipboard"})
	var global_ip *bool = parser.Flag("g", "global", &argparse.Options{Required: false, Help: "Use global IP address instead of local"})
	err := parser.Parse(os.Args)
	ExitOnError(err)
	c2_addr := GetLocalIp() + ":" + IntToStr(*port)
	if *global_ip {
		c2_addr = GetGlobalIp() + ":" + IntToStr(*port)
	}
	err = StartServer(IntToStr(*port), *dir)
	if err != nil {
		PrintInfo(F("Started server on port %s", cyan(*port)))
	}
	if *tunnel {
		go StartNgrokTCP(*port)
		c2_addr, err = GetNgrokURL()
		PrintInfo(F("Launched tunnel (%s)", green(c2_addr)))
	}
	if *clip {
		clipboard.WriteAll(c2_addr)
		PrintInfo("Copied server address to clipboard")
	}
}
