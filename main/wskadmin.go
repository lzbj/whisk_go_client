package main

import (
	//"../whisk"
	flag "github.com/docker/docker/pkg/mflag"
	"fmt"
	"os"
)

const DB_WHISK_AUTHS = "DB_WHISK_AUTHS"
const DB_PROTOCOL = "DB_PROTOCOL"
const DB_HOST = "DB_HOST"
const DB_PORT = "DB_PORT"
const DB_USERNAME = "DB_USERNAME"
const DB_PASSWORD = "DB_PASSWORD"

const usage = `manage users
create: create a user and show authorization key
delete: delete a user
get:get authorization key for user
whois:identify user from UUID (for convenience you can provide the entire authorization key)
`

var (
	option string
	help   bool
)

func init() {
	flag.Bool([]string{"#hp", "#-help"}, false, "display the help")
	flag.StringVar(&option, []string{"user"}, "create", usage)
	flag.Parse()
}
func main() {
	//fmt.Printf("ARGS: %v\n", flag.Args())
	if len(flag.Args()) == 0{
		fmt.Println("please choose command, see help")
		os.Exit(1)
	}
	if flag.Args()[0] != "user" {
		fmt.Println("not supported command, please choose from {user}, see help")
		os.Exit(1)
	}

	fmt.Println("continue here")
}
