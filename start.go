package main

import (
	"context"
	"log"
	"os"

	"entdemo/ent"
	"entdemo/linebot/route"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(".env") ; err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	client, err := ent.Open("postgres", os.Getenv("PSQL_DB_CONNECT"))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	defer client.Close()
	ctx := context.Background()

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	r.POST("/callback", route.HandlerReply)
	
	r.Run(":" + os.Getenv("PORT"))
}


