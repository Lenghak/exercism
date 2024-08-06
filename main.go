package main

import (
	parsinglogfiles "exercism/parsing_log_files"
	"fmt"
)

func main() {

	fmt.Println(parsinglogfiles.TagWithUserName([]string{
		"[WRN] User James123 has exceeded storage space.\n",
		"[WRN] Host down. User   Michelle4 lost connection.\n",
		"[INF] Users can login again after 23:00. \n",
		"[DBG] We need to check that user names are at least 6 chars long. \n",
	}))
}
