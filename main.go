package main

import (
    "github.com/fvbock/endless"
    "time"
    "fmt"
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "api_project/routers"
    _ "api_project/models"
    _ "api_project/docs"
    _ "github.com/astaxie/beego"
)


func main() {

    fmt.Println("[Server Starting]...")
    gin.SetMode(gin.ReleaseMode)
    router := routers.InitRouter()
    
    server := endless.NewServer("127.0.0.1:8080", router)
    server.ReadTimeout = 3 * time.Second
    server.WriteTimeout = 3 * time.Second

    server.BeforeBegin = func(add string) {
        log.Printf(add)
    }

    err := server.ListenAndServe()

    if err != nil {
        log.Println(err)
    }

    log.Println("Server stopped")
    os.Exit(0)
    
}
