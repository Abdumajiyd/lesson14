package main

import (
    "encoding/json"
    "fmt"
    "os"
    "time"
)

type Info struct {
    ID                 int       `json:"id"`
    Title              string    `json:"title"`
    Price              int       `json:"price"`
    Quantity           int       `json:"quantity"`
    Total              int       `json:"total"`
    DiscountPercentage float64   `json:"discountPercentage"`
    DiscountedPrice    int       `json:"discountedPrice"`
    Thumbnail          string    `json:"thumbnail"`
    CreatedAt          time.Time `json:"created_at"`
}

func main() {
    f, err := os.Open("file.json")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    var info Info
    err = json.NewDecoder(f).Decode(&info)
    if err != nil {
        panic(err)
    }

    fmt.Println("ID:", info.ID)
    fmt.Println("Title:", info.Title)
    fmt.Println("Price:", info.Price)
    fmt.Println("Quantity:", info.Quantity)
    fmt.Println("Total:", info.Total)
    fmt.Println("DiscountPercentage:", info.DiscountPercentage)
    fmt.Println("DiscountedPrice:", info.DiscountedPrice)
    fmt.Println("Thumbnail:", info.Thumbnail)
    fmt.Println("CreatedAt:", info.CreatedAt.Format("02/01 15:04:05 2006"))
}
