package strategies

import (
	"fmt"
	"openLinksInBrowser/commands"
)

func runStrategy() {

}

func addLinkStrategy() {

}
func showAllLinksStrategy() {

}

func defineStrategy(params []string) {
	var command string = params[0]

	switch command {
	case fmt.Sprint(commands.RUN):
		{
			runStrategy()
		}
	case fmt.Sprint(commands.ADD_LINK):
		{
			addLinkStrategy()
		}
	case fmt.Sprint(commands.SHOW_ALL_LINKS):
		{
			showAllLinksStrategy()
		}

	}
	fmt.Println("", params)
}
