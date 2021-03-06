package main

import (
	"Perekoter/core"
	"Perekoter/models"
	"Perekoter/utility"
	"fmt"
	"os"
)

func main() {
	models.Init()
	fmt.Println("Connection to the database was created!")
	os.Mkdir("./covers", 0777)
	utility.Config.Read()
	config := utility.Config.Get()

	utility.CurrentUsercode.PasscodeAuth()

	utility.StartInterval()

	router := core.GetRouter()

	router.Run(":" + config.Port)
}
