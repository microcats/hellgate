package main

import "github.com/gin-gonic/gin"
//import "time"

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        //time.Sleep(2000 * time.Millisecond)
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.GET("/ping/pong", func(c *gin.Context) {
        //time.Sleep(2000 * time.Millisecond)
        c.JSON(200, gin.H{
            "message": "pong",
            "aaa": c.Query("aa"),
        })
    })


    r.POST("/ping", func(c *gin.Context) {
        //time.Sleep(2000 * time.Millisecond)
        c.JSON(200, gin.H{
            "message": "pong",
            "aaa": c.PostForm("aaaa"),
        })
    })
    r.Run(":9090") // listen and server on 0.0.0.0:8080
}
