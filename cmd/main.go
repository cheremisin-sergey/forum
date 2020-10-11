package cmd

import (
	"fmt"
	"github.com/cheremisin-sergey/forum/config"
)

func main() {

	config := config.NewConfig()
	fmt.Println(config.AppName)

}
