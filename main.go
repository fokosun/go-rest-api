package main

import (
    "github.com/fokosun/go-rest-api/routes"
)

func main() {
    router := routes.SetupRouter()
    router.Run(":8080")
}
