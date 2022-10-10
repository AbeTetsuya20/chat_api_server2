package main

import (
	"context"
	"fmt"
	"os"
	"server/server/driver"
)

func main() {
	ctx := context.Background()
	sql, err := driver.InitDatabase("localhost", 3306, "server", "root", "tmp")
	if err != nil {
		fmt.Printf("failed to connect database: %s\n", err)
		os.Exit(2)
	}

	serviceDriver, err := driver.InitDriver(sql)
	if err != nil {
		fmt.Printf("failed to create driver: %s\n", err)
		os.Exit(2)
	}
	serviceDriver.ChatService(ctx)

	fmt.Println("終了")
}
