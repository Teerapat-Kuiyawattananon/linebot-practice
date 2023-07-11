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

	if err := CreateCars(client, ctx) ; err != nil {
		log.Fatal(err)
	}

	r.POST("/callback", route.HandlerReply)

	r.Run(":" + os.Getenv("PORT"))
}

func CreateCars(client *ent.Client, ctx context.Context) error {
	if client.Car.Query().CountX(ctx) > 0 {
		return nil
	}
	if err := client.Car.
				Create().
				SetModel("Tesla").
				SetPrice(1759000).
				SetImagePath("https://static-assets.tesla.com/configurator/compositor?&bkba_opt=1&view=STUD_3QTR&size=1400&model=m3&options=$APBS,$DV2W,$IBB1,$PPSW,$PRM30,$SC04,$MDL3,$W40B,$MT322,$CPF0,$DRRH,$RSF1,$CW03&crop=1400,850,300,130&").
				Exec(ctx) ; err != nil {
					return err
				}
		
	if err := client.Car.
				Create().
				SetModel("Ford ranger").
				SetPrice(1264000).
				SetImagePath("https://carsguide-res.cloudinary.com/image/upload/f_auto,fl_lossy,q_auto,t_default/v1/editorial/vhs/Ford-Ranger-dual-cab_1.png").
				Exec(ctx) ; err != nil {
					return err
				}

	if err := client.Car.
				Create().
				SetModel("Toyota yaris").
				SetPrice(1264000).
				SetImagePath(`https://www.hongtonggas.com/wp-content/uploads/2022/08/%E0%B8%A3%E0%B8%B5%E0%B8%A7%E0%B8%B4%E0%B8%A7-Toyota-Yaris-Ativ-1.2-%E0%B8%95%E0%B8%B4%E0%B8%94%E0%B9%81%E0%B8%81%E0%B9%8A%E0%B8%AA-Prins-%E0%B8%AB%E0%B8%87%E0%B8%A9%E0%B9%8C%E0%B8%97%E0%B8%AD%E0%B8%87%E0%B9%81%E0%B8%81%E0%B9%8A%E0%B8%AA.jpg`).
				Exec(ctx) ; err != nil {
					return err
				}
	
	return nil
}
