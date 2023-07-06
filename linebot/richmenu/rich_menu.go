// // Copyright 2016 LINE Corporation
// //
// // LINE Corporation licenses this file to you under the Apache License,
// // version 2.0 (the "License"); you may not use this file except in compliance
// // with the License. You may obtain a copy of the License at:
// //
// //   http://www.apache.org/licenses/LICENSE-2.0
// //
// // Unless required by applicable law or agreed to in writing, software
// // distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// // WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// // License for the specific language governing permissions and limitations
// // under the License.

package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	flag.Parse()
	bot, err := linebot.New(
		"6e45b0810ef87b495b0521a325c82a32",
		"sKrOayAXwLTFlzAqMGwqw0XXgQDvhcEUBPztcBE4fHNi5ELaB+SJV0jJaU4tgEChhihoNxpUHIA4ztYknKKomjiOoAeKMR+srT9D6zgdlBbDw0BbeYcUU3HulJZS4zVO5/vbxNnxSqN8Grq6V3LQrwdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		log.Fatal(err)
	}

// 		// ---------------------- rich menu
	// rich menu 1 (Register)  [richmenu-9f73a96f2bf5a8601f4a0d73d61ff2f8]
	// rich menu 2 (detail/logout) [richmenu-b63c48e1bd729f967755f9f0ff3f2bd7]

	// teatime rich menu 1 (Register) [richmenu-5c1705a09661a32a9dfe0072b4b6a3ce]
	// teatime rich menu 2 (login) [richmenu-ccda289b84596e0be0d5d248ae5d4a86] -----> [richmenu-89ab940234b7c94d2b20fbff9af45e50]
	// teatime rich menu 3 (edit) [richmenu-3db64483ee4b980b8823c93250b7c934]
	var (
		mode = flag.String("mode", "upload", "mode of richmenu helper [list|create|link|unlink|bulklink|bulkunlink|get|delete|upload|download|alias_create|alias_get|alias_update|alias_delete|alias_list]")
		aid      = flag.String("aid", "", "alias id")
		uid      = flag.String("uid", "U7e34782d5ddff5b68b00a2ca9e5fbc3d", "user id")
		rid      = flag.String("rid", "richmenu-ccda289b84596e0be0d5d248ae5d4a86", "richmenu id")
		filePath = flag.String("image.path", "/Users/teerapat/Documents/COOP/Blockfint/example/ent-example/linebot/richmenu/TeaTime_richmenu_login.jpeg", "path to image, used in upload/download mode")
	)
	log.Println("-------------- RICH MENU ----------------")

	switch *mode {
	case "upload":
		if _, err = bot.UploadRichMenuImage(*rid, *filePath).Do(); err != nil {
			log.Fatal(err)
		}
	case "download":
		res, err := bot.DownloadRichMenuImage(*rid).Do()
		if err != nil {
			log.Fatal(err)
		}
		defer res.Content.Close()
		f, err := os.OpenFile(*filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(f, res.Content)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Image is written to %s", *filePath)
	case "alias_create":
		if _, err = bot.CreateRichMenuAlias(*aid, *rid).Do(); err != nil {
			log.Fatal(err)
		}
	case "alias_get":
		if res, err := bot.GetRichMenuAlias(*aid).Do(); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("%v\n", res)
		}
	case "alias_update":
		if _, err = bot.UpdateRichMenuAlias(*aid, *rid).Do(); err != nil {
			log.Fatal(err)
		}
	case "alias_delete":
		if _, err = bot.DeleteRichMenuAlias(*aid).Do(); err != nil {
			log.Fatal(err)
		}
	case "alias_list":
		res, err := bot.GetRichMenuAliasList().Do()
		if err != nil {
			log.Fatal(err)
		}
		for _, alias := range res {
			log.Printf("%v\n", alias)
		}
	case "link":
		if _, err = bot.LinkUserRichMenu(*uid, *rid).Do(); err != nil {
			log.Fatal(err)
		}
	case "unlink":
		if _, err = bot.UnlinkUserRichMenu(*uid).Do(); err != nil {
			log.Fatal(err)
		}
	case "bulklink":
		if _, err = bot.BulkLinkRichMenu(*rid, *uid).Do(); err != nil {
			log.Fatal(err)
		}
	case "bulkunlink":
		if _, err = bot.BulkUnlinkRichMenu(*uid).Do(); err != nil {
			log.Fatal(err)
		}
	case "list":
		log.Println("-----Get List Rich-------")
		res, err := bot.GetRichMenuList().Do()
		if err != nil {
			log.Fatal(err)
		}
		for _, richmenu := range res {
			log.Printf("%v\n", richmenu)
		}
	case "get_default":
		res, err := bot.GetDefaultRichMenu().Do()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%v\n", res)
	case "set_default":
		if _, err = bot.SetDefaultRichMenu(*rid).Do(); err != nil {
			log.Fatal(err)
		}
	case "cancel_default":
		if _, err = bot.CancelDefaultRichMenu().Do(); err != nil {
			log.Fatal(err)
		}
	case "create":
		richMenu := linebot.RichMenu{
			Size:        linebot.RichMenuSize{Width: 2500, Height: 1686},
			Selected:    false,
			Name:        "Menu1",
			ChatBarText: "ChatText",
			Areas: []linebot.AreaDetail{
				{
					Bounds: linebot.RichMenuBounds{X: 0, Y: 0, Width: 2500, Height: 1686},
					Action: linebot.RichMenuAction{
						Type:	linebot.RichMenuActionTypeMessage,
						Text: 	"Register success",
					},
				},
				// {
				// 	Bounds: linebot.RichMenuBounds{X:0, Y: 0, Width: 2444, Height: 222},
				// 	Action: linebot.RichMenuAction{
				// 		Type: linebot.RichMenuActionTypeRichMenuSwitch,
				// 		RichMenuAliasID: "A-To-B",
				// 		Data: "A-To-B",
				// 	},
				// }
				// {
				// 	Bounds: linebot.RichMenuBounds{X: 1250, Y: 0, Width: 1250, Height: 212},
				// 	Action: linebot.RichMenuAction{
				// 		Type:            linebot.RichMenuActionTypeRichMenuSwitch,
				// 		RichMenuAliasID: "richmenu-alias-b",
				// 		Data:            "action=richmenu-changed-to-b",
				// 	},
				// },
				// {
				// 	Bounds: linebot.RichMenuBounds{X: 0, Y: 212, Width: 1250, Height: 737},
				// 	Action: linebot.RichMenuAction{
				// 		Type: linebot.RichMenuActionTypePostback,
				// 		Data: "action=buy&itemid=123",
				// 	},
				// },
				// {
				// 	Bounds: linebot.RichMenuBounds{X: 1250, Y: 212, Width: 1250, Height: 737},
				// 	Action: linebot.RichMenuAction{
				// 		Type: linebot.RichMenuActionTypeURI,
				// 		URI:  "https://developers.line.me/",
				// 		Text: "click me",
				// 	},
				// },
				// {
				// 	Bounds: linebot.RichMenuBounds{X: 0, Y: 949, Width: 1250, Height: 737},
				// 	Action: linebot.RichMenuAction{
				// 		Type: linebot.RichMenuActionTypeMessage,
				// 		Text: "hello world!",
				// 	},
				// },
				// {
				// 	Bounds: linebot.RichMenuBounds{X: 1250, Y: 949, Width: 1250, Height: 737},
				// 	Action: linebot.RichMenuAction{
				// 		Type: linebot.RichMenuActionTypeDatetimePicker,
				// 		Data: "datetime picker!",
				// 		Mode: "datetime",
				// 	},
				// },
			},
		}
		res, err := bot.CreateRichMenu(richMenu).Do()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res.RichMenuID)
	case "create2" :
		richMenu := linebot.RichMenu{
			Size:        linebot.RichMenuSize{Width: 2500, Height: 1686},
			Selected:    true,
			Name:        "Menu1",
			ChatBarText: "ChatText",
			Areas: []linebot.AreaDetail{
				{
					Bounds: linebot.RichMenuBounds{X: 0, Y: 4, Width: 1250, Height: 1686},
					Action: linebot.RichMenuAction{
						Type:	linebot.RichMenuActionTypePostback,
						Data: 	"getDetail",
					},
				},
				{
					Bounds: linebot.RichMenuBounds{X: 1246, Y: 0, Width: 1250, Height: 1686},
					Action: linebot.RichMenuAction{
						Type: linebot.RichMenuActionTypeMessage,
						Text: "logout",
					},
				},
			},
		}
		res, err := bot.CreateRichMenu(richMenu).Do()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res.RichMenuID)
	case "delete": 
		if _, err := bot.DeleteRichMenu(*rid).Do() ; err != nil {
			log.Fatal(err)
		}

	case "create_TeaTime_register" :
		richMenu := linebot.RichMenu{
			Size:        linebot.RichMenuSize{Width: 2500, Height: 1686},
			Selected:    true,
			Name:        "TimeTime_Register",
			ChatBarText: "Register",
			Areas: []linebot.AreaDetail{
				{
					Bounds: linebot.RichMenuBounds{X: 4, Y: 837, Width: 1251, Height: 849},
					Action: linebot.RichMenuAction{
						Type:	linebot.RichMenuActionTypeMessage,
						Text:   "Register",
					},
				},
				{
					Bounds: linebot.RichMenuBounds{X: 1250, Y: 837, Width: 1246, Height: 849},
					Action: linebot.RichMenuAction{
						Type: linebot.RichMenuActionTypeMessage,
						Text: "สินเชื่อ",
					},
				},
			},
		}
		res, err := bot.CreateRichMenu(richMenu).Do() 
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res.RichMenuID)
	case "create_TeaTime_login" :
		richMenu := linebot.RichMenu{
			Size:        linebot.RichMenuSize{Width: 2500, Height: 1686},
			Selected:    true,
			Name:        "TimeTime_Login",
			ChatBarText: "My Account",
			Areas: []linebot.AreaDetail{
				{
					Bounds: linebot.RichMenuBounds{X: 4, Y: 520, Width: 829, Height: 834},
					Action: linebot.RichMenuAction{
						Type:	linebot.RichMenuActionTypePostback,
						Data: 	"getDetail",
					},
				},
				{
					Bounds: linebot.RichMenuBounds{X: 833, Y: 524, Width: 842, Height: 833},
					Action: linebot.RichMenuAction{
						Type: linebot.RichMenuActionTypeMessage,
						Text: "เปลี่ยนวันที่",
					},
				},
				{
					Bounds: linebot.RichMenuBounds{X: 1675, Y: 525, Width: 825, Height: 832},
					Action: linebot.RichMenuAction{
						Type: linebot.RichMenuActionTypeMessage,
						Text: "logout",
					},
				},
			},
		}
		res, err := bot.CreateRichMenu(richMenu).Do() 
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res.RichMenuID)
	default:
		log.Fatal("implement me")
	}
}
