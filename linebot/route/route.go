package route

import (
	"context"
	"entdemo/ent"
	"entdemo/ent/car"
	_ "entdemo/ent/car"
	_ "entdemo/ent/creditlater"
	"entdemo/ent/linelog"
	"entdemo/ent/lineuser"
	"entdemo/linebot/richmessage"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/line/line-bot-sdk-go/v7/linebot/httphandler"
)

var client *ent.Client
var err error

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	client, err = ent.Open("postgres", os.Getenv("PSQL_DB_CONNECT"))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	fmt.Println("Connect PostgresDB Success")

	// defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func HandlerReply(c *gin.Context) {
	ctx := context.Background()
	handler, err := httphandler.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	bot, err := handler.NewClient()
	if err != nil {
		log.Print(err)
		return
	}

	events, err := bot.ParseRequest(c.Request)
	// New Flex Message
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				CreateLineLog(ctx, client, event.Source.UserID, "send msg", message.Text)
				var lineUser *ent.LineUser
				if client.LineUser.Query().Where(lineuser.UserId(event.Source.UserID)).FirstIDX(ctx) != 0 {
					lineUser = client.LineUser.
						Query().
						Where(lineuser.UserId(event.Source.UserID)).
						OnlyX(ctx)
				}
				// switch text := message.Text
				if message.Text == "สินเชื่อ" {
					bot.ReplyMessage(event.ReplyToken, richmessage.GetSinTrustInfoFlexMessage()).Do()
				} else if message.Text == "Me" {
					profile, _ := bot.GetProfile(event.Source.UserID).Do()
					bot.ReplyMessage(event.ReplyToken, richmessage.GetRichWithJSON(profile)).Do()
				} else if message.Text == "Tea" {
					bot.ReplyMessage(event.ReplyToken, richmessage.GetInfoTeaTimeFlexMessage()).Do()
				} else if message.Text == "MENU" {
					bot.ReplyMessage(event.ReplyToken, richmessage.GetMenuTeaTimeCarousel()).Do()
				} else if message.Text == "สินเชื่อของฉัน" {
					if client.LineUser.Query().Where(lineuser.UserId(event.Source.UserID)).FirstIDX(ctx) == 0 || lineUser.Active == false {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("กรุณาลงทะเบียนหรือเข้าสู่ระบบ")).Do()
						return
					}
					if lineUser.Active {
						bot.ReplyMessage(event.ReplyToken, richmessage.GetSinTrustFlexMessage(lineUser)).Do()
					}
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("กรุณาลงทะเบียนหรือเข้าสู่ระบบ")).Do()
				} else if message.Text == "Logs" {
					if client.LineUser.Query().Where(lineuser.UserId(event.Source.UserID)).FirstIDX(ctx) == 0 || lineUser.Active == false {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("กรุณาลงทะเบียนหรือเข้าสู่ระบบ")).Do()
						return
					}
					bot.ReplyMessage(event.ReplyToken, richmessage.GetLineLogs(lineUser)).Do()
				} else if message.Text == "แก้ไขวันที่" {
					if client.LineUser.Query().Where(lineuser.UserId(event.Source.UserID)).FirstIDX(ctx) == 0 || lineUser.Active == false {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("กรุณาลงทะเบียนหรือเข้าสู่ระบบ")).Do()
						return
					}
					if lineUser.Active {
						bot.ReplyMessage(event.ReplyToken, richmessage.GetSinTrustChangeDate(lineUser)).Do()
					}
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("กรุณาลงทะเบียนหรือเข้าสู่ระบบ")).Do()
				} else if message.Text == "Groups" {
					bot.ReplyMessage(event.ReplyToken, richmessage.GetListGroups(client)).Do()
				} else if message.Text == "Register" {
					profile, _ := bot.GetProfile(event.Source.UserID).Do()
					if client.LineUser.Query().Where(lineuser.UserId(profile.UserID)).FirstIDX(ctx) == 0 {
						newUser := CreateLineUser(ctx, client, profile)
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ลงทะเบียนสำเร็จ..."),
							linebot.NewTextMessage("เข้าสู่ระบบสำเร็จ...")).Do()
						newUser = newUser.Update().SetActive(true).SaveX(ctx)
						if _, err := bot.LinkUserRichMenu(event.Source.UserID, "richmenu-47e96433196357039bdea13c0ac04682").Do(); err != nil {
							bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("เกิดข้อผิดพลาด: " + err.Error())).Do()
						}
						return
					}
					var newTextMessgae []linebot.Message
					newTextMessgae = append(newTextMessgae, linebot.NewTextMessage("คุณได้ลงทะเบียนไปแล้ว"))
					newTextMessgae = append(newTextMessgae, linebot.NewTextMessage("เข้าสู่ระบบสำเร็จ..."))
					bot.ReplyMessage(event.ReplyToken,
						linebot.NewTextMessage("คุณได้ลงทะเบียนไปแล้ว"),
						linebot.NewTextMessage("เข้าสู่ระบบสำเร็จ...")).Do()
					lineUser = lineUser.Update().SetActive(true).SaveX(ctx)
					if _, err := bot.LinkUserRichMenu(event.Source.UserID, "richmenu-47e96433196357039bdea13c0ac04682").Do(); err != nil {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("เกิดข้อผิดพลาด: " + err.Error())).Do()
						return
					}
				} else if message.Text == "logout" {
					if client.LineUser.Query().Where(lineuser.UserId(event.Source.UserID)).FirstIDX(ctx) == 0 {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("กรุณาลงทะเบียนหรือเข้าสู่ระบบ")).Do()
						return
					}
					lineUser = lineUser.Update().SetActive(false).SaveX(ctx)
					if _, err := bot.LinkUserRichMenu(event.Source.UserID, "richmenu-ad61634174b7900b7636c6f035f265e4").Do(); err != nil {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("เกิดข้อผิดพลาด: " + err.Error())).Do()
						return
					}
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ออกจากระบบแล้ว...")).Do()
				} else if message.Text == "แก้ไขสาขา" {
					if client.LineUser.Query().Where(lineuser.UserId(event.Source.UserID)).FirstIDX(ctx) == 0 || lineUser.Active == false {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("กรุณาลงทะเบียนหรือเข้าสู่ระบบ")).Do()
						return
					}
					if lineUser.Active {
						bot.ReplyMessage(event.ReplyToken, richmessage.GetSinTrustChangeBranch(lineUser)).Do()
					}
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("กรุณาลงทะเบียนหรือเข้าสู่ระบบ")).Do()
				} else if match, _ := regexp.MatchString("แก้ไขสาขาเป็น: [ก-๙a-zA-Z]+", message.Text); match == true {
					if client.LineUser.Query().Where(lineuser.UserId(event.Source.UserID)).FirstIDX(ctx) == 0 || lineUser.Active == false {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("กรุณาลงทะเบียนหรือเข้าสู่ระบบ")).Do()
						return
					}
					if lineUser.Active {
						newBeanch := message.Text[41:]
						credit_later := lineUser.QueryCreditlaters().OnlyX(ctx)
						credit_later = credit_later.Update().SetBranch(newBeanch).SaveX(ctx)
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("แก้ไขสาขาของคุณเป็น "+newBeanch+" สำเร็จ")).Do()
					}
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("กรุณาลงทะเบียนหรือเข้าสู่ระบบ")).Do()
				} else if message.Text == "แก้ไขจำนวนเงิน" {
					if client.LineUser.Query().Where(lineuser.UserId(event.Source.UserID)).FirstIDX(ctx) == 0 || lineUser.Active == false {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("กรุณาลงทะเบียนหรือเข้าสู่ระบบ")).Do()
						return
					}
					if lineUser.Active {
						bot.ReplyMessage(event.ReplyToken, richmessage.GetSinTrustChangeAmount(lineUser)).Do()
					}
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("กรุณาลงทะเบียนหรือเข้าสู่ระบบ")).Do()
				} else if match, _ := regexp.MatchString("แก้ไขจำนวนเงินเป็น: [1-9][0-9]+[.]?[0-9]+", message.Text); match == true {
					if client.LineUser.Query().Where(lineuser.UserId(event.Source.UserID)).FirstIDX(ctx) == 0 || lineUser.Active == false {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("กรุณาลงทะเบียนหรือเข้าสู่ระบบ")).Do()
						return
					}
					if lineUser.Active {
						log.Println(strings.Index(message.Text, " "))
						newAmount := message.Text[56:]
						newAmountInt, _ := strconv.Atoi(newAmount)
						credit_later := lineUser.QueryCreditlaters().OnlyX(ctx)
						credit_later = credit_later.Update().SetAmount(newAmountInt).SaveX(ctx)
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("แก้ไขจำนวนเงินของคุณเป็น "+newAmount+" สำเร็จ")).Do()
					}
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("กรุณาลงทะเบียนหรือเข้าสู่ระบบ")).Do()
				} else if message.Text == "Delete line" {

					// delete LineUser
					err := client.LineUser.DeleteOne(lineUser).Exec(ctx)
					if err != nil {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("เกิดข้อผิดพลาด: " + err.Error())).Do()
						return
					}

					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ลบข้อมูลของคูณสำเร็จ")).Do()
					bot.LinkUserRichMenu(event.Source.UserID, "richmenu-ad61634174b7900b7636c6f035f265e4").Do()

				} else if message.Text == "รถ" {
					if client.LineUser.Query().Where(lineuser.UserId(event.Source.UserID)).FirstIDX(ctx) != 0 && lineUser.Active {
						bot.ReplyMessage(event.ReplyToken, richmessage.GetListCars(client, ctx)).Do()
						return
					}
					bot.ReplyMessage(event.ReplyToken, richmessage.GetInfoCars(client, ctx)).Do()

				} else if match, _ := regexp.MatchString("เพิ่มรถ: [ก-๙a-zA-Z| ]+", message.Text); match {
					carName := message.Text[23:]
					// Get car
					car, err := client.Car.
						Query().
						Where(car.Not(
							car.HasOwner(),
						), car.Model(carName)).
						Only(ctx)
					if err != nil {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("เกิดข้อผิดพลาด: " + err.Error())).Do()
						return
					}

					newCar, err := client.Car.
						Create().
						SetModel(car.Model).
						SetPrice(car.Price).
						SetImagePath(car.ImagePath).
						Save(ctx)
					if err != nil {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("เกิดข้อผิดพลาด: " + err.Error())).Do()
						return
					}

					// Add car to user
					err = lineUser.Update().
						AddCars(newCar).
						Exec(ctx)
					if err != nil {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("เกิดข้อผิดพลาด: " + err.Error())).Do()
						return
					}

					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("เพิ่มรถ "+car.Model+" ให้ "+lineUser.DisplyaName+" สำเร็จ")).Do()

				} else if match, _ := regexp.MatchString("ลบรถ: [ก-๙a-zA-Z| ]+\nID: ", message.Text); match {
					lineStr := strings.Split(message.Text, "\n")
					carIDStr := lineStr[1][4:]
					
					carID, _ := strconv.Atoi(carIDStr)

					car, err := lineUser.QueryCars().
									Where(car.ID(carID)).
									Only(ctx)
					if err != nil {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("เกิดข้อผิดพลาด: " + err.Error())).Do()
						return
					}
					
					// remove car
					err = lineUser.Update().
								RemoveCars(car).
								Exec(ctx)
					if err != nil {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("เกิดข้อผิดพลาด: " + err.Error())).Do()
						return
					}

					var newMessage []linebot.SendingMessage
					newMessage = append(newMessage, linebot.NewTextMessage("ลบรถ: " + car.Model + " ออกจาก " + lineUser.DisplyaName + " สำเร็จ"))
				
					// delete car after remove from user
					err = client.Car.
							DeleteOne(car).
							Exec(ctx)
					if err != nil {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("เกิดข้อผิดพลาด: " + err.Error())).Do()
						return
					}
					
					if lineUser.QueryCars().CountX(ctx) == 0 {
						newMessage = append(newMessage, linebot.NewTextMessage("รถถูกลบทั้งหมดแล้ว"))
						bot.ReplyMessage(event.ReplyToken, newMessage...).Do()
						return
					}
					newMessage = append(newMessage, richmessage.GetMyCars(client, ctx, lineUser))
					bot.ReplyMessage(event.ReplyToken, newMessage...).Do()
					// bot.ReplyMessage(event.ReplyToken, richmessage.GetMyCars(client, ctx, lineUser)).Do()
				} else if message.Text == "รถของฉัน" {
					if lineUser.QueryCars().CountX(ctx) == 0 {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("คุณยังไม่มีรถ")).Do()
						return
					}
					// lineUser.Update().
					// 	AddCarIDs(1).
					// 	ExecX(ctx)
					bot.ReplyMessage(event.ReplyToken, richmessage.GetMyCars(client, ctx, lineUser)).Do()

				} else if match, _ := regexp.MatchString("ดูข้อมูลรถ: [a-zA-Z| ]+\nID: [1-9]+[0-9]*", message.Text); match {
					lineStr := strings.Split(message.Text, "\n")
					// carName := lineStr[0][32:]
					carIDStr := lineStr[1][4:]
					carID, _ := strconv.Atoi(carIDStr)

					car, err := client.Car.Query().
						Where(car.ID(carID)).
						Only(ctx)
					if err != nil {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("เกิดข้อผิดพลาด: " + err.Error())).Do()
						return
					}
					msg := fmt.Sprintf("CarID: %d\nModel: %s\nPrice: %s บาท\nDate: %s", car.ID, car.Model, humanize.Comma(int64(car.Price)), car.RegisteredAt.Format("02/01/2006"))
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg)).Do()

				} else if match, _ := regexp.MatchString(`แก้ไข Log ที่ [1-9]?[0-9]+\nAction: [a-zA-Z| ]+`, message.Text); match {
					msgLine := strings.Split(message.Text, "\n")
					logID, _ := strconv.Atoi(msgLine[0][30:])
					logAction := msgLine[1][8:]

					line_log, err := client.LineLog.
						Query().
						Where(linelog.ID(logID)).
						FirstX(ctx).
						Update().
						SetAction(logAction).
						Save(ctx)
					if err != nil {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("เกิดข้อผิดพลาด: " + err.Error())).Do()
					}

					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(fmt.Sprintf("แก้ไข Action's Log ที่ %d เป็น %s เรียบร้อย", line_log.ID, line_log.Action))).Do()
				} else {
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(`คุณสามารถพิมพ์ "สินเชื่อ" หรือ กดปุ่ม "ดูรายละเอียดสินเชื่อ" ด้านล่าง เพื่อดูรายละเอียดการกู้เงิน`),
						linebot.NewTextMessage(`หรือคุณสามารถสมัครเป็นผู้กู้โดยการ กดที่ปุ่ม "ลงทะเบียน/เข้าสู่ระบบ" ด้านล่าง หรือพิมพ์ "Register"`)).Do()
					// bot.ReplyMessage(event.ReplyToken, richmessage.GetTest()).Do()
				}

			}
		} else if event.Type == linebot.EventTypePostback {
			post_data := event.Postback.Data
			// CreateLineLog(ctx, client, event.Source.UserID, "send postback", post_data)
			profile, _ := bot.GetProfile(event.Source.UserID).Do()
			if post_data == "DATE" {
				dateTime := event.Postback.Params.Datetime
				t, _ := time.Parse("2006-01-02T15:04", dateTime)
				data := profile.DisplayName + " ได้เลือกวันที่ " + t.Format("02/01/2006 15:04")
				// data := profile.DisplayName + " ได้เลือกวันที่ " + dateTime

				// data += fmt.Sprintf("(%v)", date)

				bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(data)).Do()
			} else if post_data == "getDetail" {
				if client.LineUser.Query().Where(lineuser.UserId(event.Source.UserID)).FirstIDX(ctx) == 0 {
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("กรุณาลงทะเบียนหรือเข้าสู่ระบบ")).Do()
					return
				}
				line_user := client.LineUser.
					Query().
					Where(lineuser.UserId(event.Source.UserID)).
					OnlyX(ctx)
				if line_user.Active {
					bot.ReplyMessage(event.ReplyToken, richmessage.GetSinTrustFlexMessage(line_user)).Do()
				}
				bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("กรุณาลงทะเบียนหรือเข้าสู่ระบบ")).Do()
			} else if post_data == "changeDateSinTrust" {
				if client.LineUser.Query().Where(lineuser.UserId(event.Source.UserID)).FirstIDX(ctx) == 0 {
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("กรุณาลงทะเบียนหรือเข้าสู่ระบบ")).Do()
					return
				}
				line_user := client.LineUser.
					Query().
					Where(lineuser.UserId(event.Source.UserID)).
					OnlyX(ctx)
				if line_user.Active {
					dateTimeStr := event.Postback.Params.Datetime
					dateTimeRFC, err := time.Parse("2006-01-02T15:04", dateTimeStr)
					dateTimeFormat := dateTimeRFC.Format("02/01/2006 15:04")

					if err != nil {
						log.Println(err)
					}
					credit_later := client.LineUser.
						Query().
						Where(lineuser.UserId(event.Source.UserID)).
						QueryCreditlaters().
						OnlyX(ctx).
						Update().
						SetDate(dateTimeFormat).
						SaveX(ctx)

					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("เปลี่ยนวันที่เป็น "+credit_later.Date+" น.")).Do()
					return
				}
				bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("กรุณาลงทะเบียนหรือเข้าสู่ระบบ")).Do()
			}
		} else if event.Type == linebot.EventTypeFollow {
			bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(`คุณสามารถพิมพ์ "สินเชื่อ" หรือ กดปุ่ม "ดูรายละเอียดสินเชื่อ" ด้านล่าง เพื่อดูรายละเอียดการกู้เงิน`),
				linebot.NewTextMessage(`หรือคุณสามารถสมัครเป็นผู้กู้โดยการ กดที่ปุ่ม "ลงทะเบียน/เข้าสู่ระบบ" ด้านล่าง หรือพิมพ์ "Register"`)).Do()
		}
	}
}

func CreateLineUser(ctx context.Context, client *ent.Client, profile *linebot.UserProfileResponse) *ent.LineUser {
	var lineUser *ent.LineUser
	if client.LineUser.Query().Where(lineuser.UserId(profile.UserID)).FirstIDX(ctx) == 0 {
		lineUser, err = client.LineUser.Create().
			SetUserId(profile.UserID).
			SetDisplyaName(profile.DisplayName).
			SetRegisteredAt(time.Now()).
			Save(ctx)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		_, err := client.CreditLater.Create().
			SetOwner(lineUser).
			SetTransactionRef(strconv.Itoa(rand.Intn(50))).
			Save(ctx)
		if err != nil {
			log.Fatal(err)
			return nil
		}
	}
	return lineUser
}

func CreateLineLog(ctx context.Context, client *ent.Client, userId string, action string, msg string) {
	result := "Please type Regis"
	if client.LineUser.Query().Where(lineuser.UserId(userId)).FirstIDX(ctx) != 0 {
		lineUser := client.LineUser.Query().Where(lineuser.UserId(userId)).OnlyX(ctx)
		log, err := client.LineLog.Create().
			SetAction(action).
			SetMessage(msg).
			SetCreatedAt(time.Now()).
			SetOwner(lineUser).
			Save(ctx)
		if err != nil {
			result = "Created log failed"
		}
		result = "Created log success" + log.String()
	}

	log.Println(result)

}
