package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	// check.CheckImportFunction(5)
	// check.InitDb()
	// routes.Get()
	FetchJsonData()
}
type Post struct {
    UserID int    `json:"userId"`
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Body   string `json:"body"`
}


func FetchJsonData(){
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
        if err != nil {
			panic(err)
        }
var posts []Post
        if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
			panic(err)
        }
        for _, post := range posts {
    jsonPost, _ := json.Marshal(post)
    fmt.Println(string(jsonPost))
            }
        defer resp.Body.Close()
}