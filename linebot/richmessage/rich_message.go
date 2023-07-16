package richmessage

import (
	"context"
	"entdemo/ent"
	"entdemo/ent/car"
	"fmt"
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/dustin/go-humanize"
)

func GetRichMessage() *linebot.FlexMessage {
	var contents []linebot.FlexComponent
	text := linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   "Brown Cafe",
		Weight: "bold",
		Size:   linebot.FlexTextSizeTypeXl,
	}
	contents = append(contents, &text)
	// Make Hero
	hero := linebot.ImageComponent{
		Type:        linebot.FlexComponentTypeImage,
		URL:         "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_1_cafe.png",
		Size:        "full",
		AspectRatio: linebot.FlexImageAspectRatioType20to13,
		AspectMode:  linebot.FlexImageAspectModeTypeCover,
		Action:      linebot.NewMessageAction("left", "left clicked"),
	}
	// Make Body
	body := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeVertical,
		Contents: contents,
	}
	// Build Container
	bubble := linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Hero: &hero,
		Body: &body,
	}
	// New Flex Message
	flexMessage := linebot.NewFlexMessage("FlexWithCode", &bubble)

	return flexMessage
}

func GetRichProfile(profile *linebot.UserProfileResponse) *linebot.FlexMessage {
	var contents []linebot.FlexComponent
	text := linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   profile.DisplayName,
		Weight: "bold",
		Size:   linebot.FlexTextSizeTypeXl,
	}
	contents = append(contents, &text)
	// Make Hero
	hero := linebot.ImageComponent{
		Type:        linebot.FlexComponentTypeImage,
		URL:         profile.PictureURL,
		Size:        "full",
		AspectRatio: linebot.FlexImageAspectRatioType20to13,
		AspectMode:  linebot.FlexImageAspectModeTypeCover,
		Action:      linebot.NewMessageAction("left", "left clicked"),
	}
	// Make Body
	body := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeVertical,
		Contents: contents,
	}
	// Build Container
	bubble := linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Hero: &hero,
		Body: &body,
	}
	// New Flex Message
	flexMessage := linebot.NewFlexMessage("FlexWithCode", &bubble)

	return flexMessage
}

func GetRichWithJSON(profile *linebot.UserProfileResponse) *linebot.FlexMessage {
	jsonStr := `{
		"type": "bubble",
		"direction": "ltr",
		"header": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "%s",
			  "align": "center",
			  "contents": []
			}
		  ]
		},
		"hero": {
		  "type": "image",
		  "url": "%s",
		  "size": "full",
		  "aspectMode": "cover"
		},
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "Statue : %s",
			  "align": "center",
			  "contents": []
			}
		  ]
		},
		"footer": {
		  "type": "box",
		  "layout": "horizontal",
		  "contents": [
			{
			  "type": "button",
			  "action": {
				"type": "message",
				"label": "ดูรายละเอียด",
				"text": "love chu~~~~"
			  },
			  "style": "primary"
			}
		  ]
		}
	  }`
	
	jsonStr = fmt.Sprintf(jsonStr, profile.DisplayName, profile.PictureURL, profile.StatusMessage)
	// fmt.Println(jsonStr)
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonStr))
	if err != nil {
		log.Println(err)
	}
	// New Flex Message
	flexMessage := linebot.NewFlexMessage("FlexWithJSON", flexContainer)

	return flexMessage
}

func GetInfoTeaTimeFlexMessage() *linebot.FlexMessage {
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(`{
		"type": "bubble",
		"hero": {
		  "type": "image",
		  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_1_cafe.png",
		  "size": "full",
		  "aspectRatio": "20:13",
		  "aspectMode": "cover",
		  "action": {
			"type": "uri",
			"label": "Action",
			"uri": "https://linecorp.com/"
		  }
		},
		"body": {
		  "type": "box",
		  "layout": "horizontal",
		  "spacing": "md",
		  "contents": [
			{
			  "type": "box",
			  "layout": "vertical",
			  "flex": 2,
			  "contents": [
				{
				  "type": "text",
				  "text": "TeaTime For You",
				  "weight": "bold",
				  "size": "lg",
				  "flex": 1,
				  "align": "start",
				  "gravity": "top",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "~ MikeTea",
				  "size": "md",
				  "flex": 2,
				  "align": "start",
				  "gravity": "top",
				  "margin": "md",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "~ Coffee",
				  "size": "md",
				  "flex": 2,
				  "gravity": "center",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "~ Bakery",
				  "size": "md",
				  "flex": 1,
				  "gravity": "bottom",
				  "contents": []
				}
			  ]
			}
		  ]
		},
		"footer": {
		  "type": "box",
		  "layout": "horizontal",
		  "contents": [
			{
			  "type": "button",
			  "action": {
				"type": "message",
				"label": "ดูรายละเอียด",
				"text": "MENU"
			  },
			  "color": "#43BC38FF"
			}
		  ]
		}
	  }`))
	if err != nil {
		log.Println(err)
	}
	// New Flex Message
	flexMessage := linebot.NewFlexMessage("FlexWithJSON", flexContainer)
	return flexMessage
}

// func GetMenuTeaTimeCarousel() *linebot.TemplateMessage {
// 	// New TemplateAction
// 	var actions []linebot.TemplateAction
// 	// Add Actions
// 	actions = append(actions, linebot.NewMessageAction("left", "left clicked"))
// 	actions = append(actions, linebot.NewMessageAction("right", "right clicked"))
// 	// Image URL For CarouselColumn
// 	imgURL := "https://cdn-image.travelandleisure.com/sites/default/files/styles/1600x1000/public/1539963100/sloth-SLOTH1018.jpg?itok=n6IuFyx_"
// 	// New CarouselColumns
// 	var columns []*linebot.CarouselColumn
// 	// Add CarouselColumn
// 	columns = append(columns, linebot.NewCarouselColumn(imgURL, "Title", "description", actions...))
// 	// New CarouselTemplate
// 	carousel := linebot.NewCarouselTemplate(columns...)
// 	// New TemplateMessage
// 	template := linebot.NewTemplateMessage("Carousel", carousel)
// 	// Reply Message
// 	return template
// }

func GetMenuTeaTimeCarousel() *linebot.FlexMessage {
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(`
	{
		"type": "carousel",
		"contents": [
		  {
			"type": "bubble",
			"hero": {
			  "type": "image",
			  "url": "https://img.wongnai.com/p/1920x0/2019/02/15/3fc49b443e224063aabf0e6028fcabd1.jpg",
			  "size": "full",
			  "aspectRatio": "20:13",
			  "aspectMode": "cover"
			},
			"body": {
			  "type": "box",
			  "layout": "vertical",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "Koi the' Goledn Bubble Mike Tea",
				  "weight": "bold",
				  "size": "xl",
				  "wrap": true,
				  "contents": []
				},
				{
				  "type": "box",
				  "layout": "baseline",
				  "margin": "lg",
				  "contents": [
					{
					  "type": "icon",
					  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
					},
					{
					  "type": "icon",
					  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
					},
					{
					  "type": "icon",
					  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
					},
					{
					  "type": "icon",
					  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
					},
					{
					  "type": "icon",
					  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
					},
					{
					  "type": "text",
					  "text": "5.0",
					  "color": "#999999",
					  "margin": "sm",
					  "contents": []
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "vertical",
				  "contents": [
					{
					  "type": "text",
					  "text": "Price : 80 Baht",
					  "weight": "regular",
					  "size": "lg",
					  "flex": 1,
					  "align": "start",
					  "gravity": "top",
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "Stock : 10",
					  "size": "lg",
					  "contents": []
					}
				  ]
				}
			  ]
			},
			"footer": {
			  "type": "box",
			  "layout": "vertical",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "button",
				  "action": {
					"type": "message",
					"label": "Add to Cart",
					"text": "add_miketea_to_cart"
				  },
				  "color": "#43BC38FF",
				  "style": "primary"
				},
				{
				  "type": "button",
				  "action": {
					"type": "uri",
					"label": "Add to wishlist",
					"uri": "https://linecorp.com"
				  }
				}
			  ]
			}
		  },
		  {
			"type": "bubble",
			"hero": {
			  "type": "image",
			  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_1_cafe.png",
			  "size": "full",
			  "aspectRatio": "20:13",
			  "aspectMode": "cover"
			},
			"body": {
			  "type": "box",
			  "layout": "vertical",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "Coffer Lover",
				  "weight": "bold",
				  "size": "xl",
				  "wrap": true,
				  "contents": []
				},
				{
				  "type": "separator",
				  "margin": "xl",
				  "color": "#FFFFFFFF"
				},
				{
				  "type": "box",
				  "layout": "baseline",
				  "margin": "xxl",
				  "contents": [
					{
					  "type": "icon",
					  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
					},
					{
					  "type": "icon",
					  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
					},
					{
					  "type": "icon",
					  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
					},
					{
					  "type": "icon",
					  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
					},
					{
					  "type": "icon",
					  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gray_star_28.png"
					},
					{
					  "type": "text",
					  "text": "4.0",
					  "color": "#999999",
					  "margin": "sm",
					  "contents": []
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "vertical",
				  "flex": 1,
				  "contents": [
					{
					  "type": "text",
					  "text": "Price : 100 Baht",
					  "size": "lg",
					  "flex": 0,
					  "align": "start",
					  "gravity": "top",
					  "wrap": true,
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "Stock : 11",
					  "size": "lg",
					  "contents": []
					}
				  ]
				}
			  ]
			},
			"footer": {
			  "type": "box",
			  "layout": "vertical",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "button",
				  "action": {
					"type": "message",
					"label": "Add to Cart",
					"text": "add_coffee_to_cart"
				  },
				  "flex": 2,
				  "color": "#43BC38FF",
				  "style": "primary"
				},
				{
				  "type": "button",
				  "action": {
					"type": "uri",
					"label": "Add to wish list",
					"uri": "https://linecorp.com"
				  }
				}
			  ]
			}
		  },
		  {
			"type": "bubble",
			"body": {
			  "type": "box",
			  "layout": "vertical",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "button",
				  "action": {
					"type": "uri",
					"label": "See more",
					"uri": "https://linecorp.com"
				  },
				  "flex": 1,
				  "gravity": "center"
				}
			  ]
			}
		  }
		]
	  }`))
	if err != nil {
		log.Println(err)
	}
	// New Flex Message
	flexMessage := linebot.NewFlexMessage("FlexWithJSON", flexContainer)
	return flexMessage
}

func GetSinTrustFlexMessage(lineuser *ent.LineUser) *linebot.FlexMessage {
	creditlater := lineuser.QueryCreditlaters().OnlyX(context.Background())
	// format Credit later ID
	strTmp := fmt.Sprintf("%012d", creditlater.ID)
	creditlater_number := fmt.Sprintf("%s-%s-%s", strTmp[:4], strTmp[4:7], strTmp[7:])

	//Create FlexContainer
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(fmt.Sprintf(
		`{
			"type": "bubble",
			"direction": "ltr",
			"header": {
			  "type": "box",
			  "layout": "vertical",
			  "contents": [
				{
				  "type": "text",
				  "text": "ใช้วงเงิน",
				  "weight": "bold",
				  "size": "xl",
				  "align": "start",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "Buy Now Pay Later",
				  "weight": "bold",
				  "size": "xl",
				  "align": "start",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "เลขที่สัญญา: %s",
				  "margin": "sm",
				  "contents": []
				},
				{
				  "type": "separator",
				  "margin": "lg"
				},
				{
				  "type": "box",
				  "layout": "baseline",
				  "spacing": "xs",
				  "margin": "md",
				  "contents": [
					{
					  "type": "text",
					  "text": "Transaction Ref.",
					  "size": "sm",
					  "color": "#AAA9A9FF",
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%s",
					  "size": "sm",
					  "align": "start",
					  "contents": []
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "horizontal",
				  "spacing": "xs",
				  "margin": "xs",
				  "contents": [
					{
					  "type": "text",
					  "text": "วันที่และเวลา",
					  "size": "sm",
					  "color": "#AAA9A9FF",
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%s",
					  "size": "sm",
					  "contents": []
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "horizontal",
				  "spacing": "xs",
				  "margin": "xs",
				  "contents": [
					{
					  "type": "text",
					  "text": "สาขา",
					  "size": "sm",
					  "color": "#AAA9A9FF",
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%s",
					  "size": "sm",
					  "align": "start",
					  "wrap": true,
					  "contents": []
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "horizontal",
				  "spacing": "xs",
				  "margin": "xs",
				  "contents": [
					{
					  "type": "text",
					  "text": "จำนวนเงิน",
					  "size": "sm",
					  "color": "#AAA9A9FF",
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "฿ %s",
					  "size": "sm",
					  "color": "#6BB977FF",
					  "contents": []
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "horizontal",
				  "spacing": "xs",
				  "margin": "xs",
				  "contents": [
					{
					  "type": "text",
					  "text": "ผ่อนชำระ",
					  "size": "sm",
					  "color": "#AAA9A9FF",
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%d งวด",
					  "size": "sm",
					  "contents": []
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "horizontal",
				  "spacing": "xs",
				  "margin": "xs",
				  "contents": [
					{
					  "type": "text",
					  "text": "หมายเหตุ",
					  "size": "sm",
					  "color": "#AAA9A9FF",
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%s",
					  "size": "sm",
					  "contents": []
					}
				  ]
				}
			  ]
			}
		  }`, creditlater_number, creditlater.TransactionRef, creditlater.Date + " น.", creditlater.Branch , humanize.Comma(int64(creditlater.Amount)), creditlater.Installment, creditlater.Detail)))
	if err != nil {
		log.Println(err)
	}

	flexMessage := linebot.NewFlexMessage("SinTruch", flexContainer)

	return flexMessage
}

func GetSinTrustInfoFlexMessage() *linebot.FlexMessage {
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(`{
		"type": "bubble",
		"hero": {
		  "type": "image",
		  "url": "https://www.arabianbusiness.com/cloud/2021/11/13/NWPALgnN-digital-payment-1.jpg",
		  "size": "full",
		  "aspectRatio": "20:13",
		  "aspectMode": "cover"
		},
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "เครดิต Buy Now Pay Later",
			  "weight": "bold",
			  "size": "xl",
			  "contents": []
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "contents": [
				{
				  "type": "text",
				  "text": "* วงเงินสูงสุด 100,000 บาท",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "* ระยะเวลาการให้สินเชื่อ 30 วัน",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "* ดอกเบี้ย 0.68% ต่อเดือนเมื่อคิดแบบเงินต้นคงที่ หรือเท่ากับ 1.25% ต่อเดือนเมื่อคิดแบบลดต้นลดดอก ตามระยะเวลาผ่อนจริง",
				  "wrap": true,
				  "contents": []
				}
			  ]
			}
		  ]
		},
		"footer": {
		  "type": "box",
		  "layout": "vertical",
		  "flex": 0,
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "button",
			  "action": {
				"type": "postback",
				"label": "ขอรายละเอียด",
				"data": "getDetail"
			  },
			  "color": "#6BB583FF",
			  "height": "sm",
			  "style": "primary"
			}
		  ]
		}
	  }`))
	if err != nil {
		log.Println(err)
	}

	flexMessage := linebot.NewFlexMessage("SinTruchInfo", flexContainer)
	return flexMessage
}


func GetLineLogs(lineuser *ent.LineUser) *linebot.FlexMessage {
	logs := lineuser.QueryLinelogs().AllX(context.Background())
	jsonStr := 
	`{
		"type": "bubble",
		"direction": "ltr",
		"header": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "History",
			  "weight": "bold",
			  "size": "xl",
			  "align": "center",
			  "contents": []
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "margin": "lg",
			  "contents": [
				{
				  "type": "text",
				  "text": "DisplayName :",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "%s",
				  "contents": []
				}
			  ]
			},
			{
			  "type": "separator",
			  "margin": "md",
			  "color": "#FF0000FF"
			}%s
		  ]
		}
	}`

	logStrJson := 
	`,
	{
	  "type": "box",
	  "layout": "baseline",
	  "margin": "sm",
	  "contents": [
		{
		  "type": "text",
		  "text": "Action :",
		  "contents": []
		},
		{
		  "type": "text",
		  "text": "%s",
		  "contents": []
		}
	  ]
	},
	{
	  "type": "box",
	  "layout": "baseline",
	  "contents": [
		{
		  "type": "text",
		  "text": "Created_at :",
		  "contents": []
		},
		{
		  "type": "text",
		  "text": "%s",
		  "contents": []
		}
	  ]
	},
	{
        "type": "button",
        "action": {
			"type": "postback",
			"label": "แก้ไข",
			"data": "changeLogs",
			"inputOption": "openKeyboard",
			"fillInText": "แก้ไข Log ที่ %d\nAction: %s"
		},
        "color": "#6BB583FF",
        "height": "sm",
        "style": "primary"
      },
	{
	  "type": "separator",
	  "margin": "md"
	}`
	tmpStr := ""
	count := len(logs)
	if count >= 5 {
		count = 5
	}
	for _, log := range logs[:count] {
		tmpStr += fmt.Sprintf(logStrJson, log.Action, log.CreatedAt.Format("02/01/2006"), log.ID, log.Action)
	}
	// testLog := fmt.Sprintf(logStrJson, "send power", "2022-22-22")
	testLogAddJson := fmt.Sprintf(jsonStr, lineuser.DisplyaName, tmpStr)
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(testLogAddJson))
	if err != nil {
		log.Println(err)
	}
	flexMessage := linebot.NewFlexMessage("Logs", flexContainer)
	return flexMessage

}

func GetSinTrustChangeDate(lineuser *ent.LineUser) *linebot.FlexMessage {
	credit_later := lineuser.QueryCreditlaters().OnlyX(context.Background())
	jsonStr := `{
		"type": "bubble",
		"direction": "ltr",
		"header": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "แก้ไขวันที่และเวลา",
			  "weight": "bold",
			  "size": "lg",
			  "align": "start",
			  "contents": []
			},
			{
			  "type": "separator",
			  "margin": "sm",
			  "color": "#00C301FF"
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "margin": "lg",
			  "contents": [
				{
				  "type": "text",
				  "text": "วันที่และเวลาเดิม : ",
				  "size": "sm",
				  "color": "#AAA9A9FF",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "%s",
				  "size": "sm",
				  "contents": []
				}
			  ]
			}
		  ]
		},
		"footer": {
			"type": "box",
			"layout": "horizontal",
			"contents": [
			  {
				"type": "button",
				"action": {
				  "type": "datetimepicker",
				  "label": "เลือกวันที่และเวลา",
				  "data": "changeDateSinTrust",
				  "mode": "datetime",
				  "initial": "2023-06-29T11:57",
				  "max": "2024-06-29T11:57",
				  "min": "2022-06-29T11:57"
				},
				"color": "#6BB583FF",
				"style": "primary"
			  }
			]
		  }
		}`
	
	jsonStr = fmt.Sprintf(jsonStr, credit_later.Date + " น.")
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonStr))
	if err != nil {
		log.Println(err)
	}
	flexMessage := linebot.NewFlexMessage("ChangeDate", flexContainer)

	return flexMessage
	
}

func GetTest() *linebot.FlexMessage {
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(`
	{
		"type": "bubble",
		"direction": "ltr",
		"footer": {
		  "type": "box",
		  "layout": "horizontal",
		  "contents": [
			{
			  "type": "button",
			  "action": {
				"type": "postback",
				"label": "Buy",
				"data": "action=buy&itemid=111",
				"displayText": "Buy",
				"inputOption": "openKeyboard",
				"fillInText": "---\nName: \nPhone: \nBirthday: \n---"
			  }
			}
		  ]
		}
	  }
	  `))
	if err != nil {
		log.Println(err)
	}
	flexMessage := linebot.NewFlexMessage("test", flexContainer)
	return flexMessage
	
}

func GetListGroups(client *ent.Client) *linebot.FlexMessage {
	groups := client.Group.
		Query().
		AllX(context.Background())
	
	jsonStr := `
	{
		"type": "bubble",
		"direction": "ltr",
		"header": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "List of Groups",
			  "weight": "bold",
			  "size": "lg",
			  "align": "start",
			  "contents": []
			},
			{
			  "type": "separator",
			  "margin": "md",
			  "color": "#00BF89FF"
			}%s
		  ]
		}
	  }`

	tmpStr := ""
	for _, group := range groups {
		nameStr := ""
		for _, u := range group.QueryLineusers().AllX(context.Background()) {
			nameStr += u.DisplyaName + " "
		}
		tmpStr += fmt.Sprintf(`,
		{
		  "type": "box",
		  "layout": "baseline",
		  "margin": "lg",
		  "contents": [
			{
			  "type": "text",
			  "text": "%s :",
			  "contents": []
			},
			{
			  "type": "text",
			  "text": "%s",
			  "wrap": true,
			  "contents": []
			}
		  ]
		},
		{
		  "type": "separator",
		  "margin": "md",
		  "color": "#00BF89FF"
		}`, group.Name, nameStr)
	}
	flexStr := fmt.Sprintf(jsonStr, tmpStr)
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(flexStr))
	if err != nil {
		log.Println(err)
	}
	flexMessage := linebot.NewFlexMessage("ListGroups", flexContainer)

	return flexMessage
}

func GetSinTrustChangeBranch(lineuser *ent.LineUser) *linebot.FlexMessage {
	credit_later := lineuser.QueryCreditlaters().OnlyX(context.Background())
	jsonStr := `{
		"type": "bubble",
		"direction": "ltr",
		"header": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "แก้ไขสาขา",
			  "weight": "bold",
			  "size": "lg",
			  "align": "start",
			  "contents": []
			},
			{
			  "type": "separator",
			  "margin": "sm",
			  "color": "#00C301FF"
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "margin": "lg",
			  "contents": [
				{
				  "type": "text",
				  "text": "สาขาเดิม : ",
				  "size": "sm",
				  "color": "#AAA9A9FF",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "%s",
				  "size": "sm",
				  "contents": []
				}
			  ]
			}
		  ]
		},
		"footer": {
			"type": "box",
			"layout": "horizontal",
			"contents": [
			  {
				"type": "button",
				"action": {
				  "type": "postback",
				  "label": "แก้ไขสาขา",
				  "data": "changeBranchSinTrust",
				  "inputOption": "openKeyboard",
				  "fillInText": "แก้ไขสาขาเป็น: "
				},
				"color": "#6BB583FF",
				"style": "primary"
			  }
			]
		  }
		}`
	
	jsonStr = fmt.Sprintf(jsonStr, credit_later.Branch)
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonStr))
	if err != nil {
		log.Println(err)
	}
	flexMessage := linebot.NewFlexMessage("ChangeBranch", flexContainer)

	return flexMessage
	
}

func GetSinTrustChangeAmount(lineuser *ent.LineUser) *linebot.FlexMessage {
	credit_later := lineuser.QueryCreditlaters().OnlyX(context.Background())
	jsonStr := `{
		"type": "bubble",
		"direction": "ltr",
		"header": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "แก้ไขจำนวนเงิน",
			  "weight": "bold",
			  "size": "lg",
			  "align": "start",
			  "contents": []
			},
			{
			  "type": "separator",
			  "margin": "sm",
			  "color": "#00C301FF"
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "margin": "lg",
			  "contents": [
				{
				  "type": "text",
				  "text": "จำนวนเงินเดิม : ",
				  "size": "sm",
				  "color": "#AAA9A9FF",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "฿ %d",
				  "size": "sm",
				  "contents": []
				}
			  ]
			}
		  ]
		},
		"footer": {
			"type": "box",
			"layout": "horizontal",
			"contents": [
			  {
				"type": "button",
				"action": {
				  "type": "postback",
				  "label": "แก้ไขจำนวนเงิน",
				  "data": "changeAmountSinTrust",
				  "inputOption": "openKeyboard",
				  "fillInText": "แก้ไขจำนวนเงินเป็น: "
				},
				"color": "#6BB583FF",
				"style": "primary"
			  }
			]
		  }
		}`
	
	jsonStr = fmt.Sprintf(jsonStr, credit_later.Amount)
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonStr))
	if err != nil {
		log.Println(err)
	}
	flexMessage := linebot.NewFlexMessage("ChangeAmount", flexContainer)

	return flexMessage
	
}

func GetListCars(client *ent.Client, ctx context.Context) *linebot.FlexMessage {
	jsonStr := `{
		"type": "carousel",
		"contents": [%s]
	  }`
	
	cars, err := client.Car.
					Query().
					Where(car.Not(
						car.HasOwner(),
					),).
					All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	jsonCars := `,{
		"type": "bubble",
		"hero": {
		  "type": "image",
		  "url": "%s",
		  "size": "full",
		  "aspectRatio": "20:13",
		  "aspectMode": "cover"
		},
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "text",
			  "text": "%s",
			  "weight": "bold",
			  "size": "xl",
			  "wrap": true,
			  "contents": []
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "contents": [
				{
				  "type": "text",
				  "text": "%s บาท",
				  "weight": "bold",
				  "size": "xl",
				  "flex": 0,
				  "wrap": true,
				  "contents": []
				}
			  ]
			}
		  ]
		},
		"footer": {
			"type": "box",
			"layout": "horizontal",
			"spacing": "sm",
			"contents": [
			  {
				  "type": "button",
				  "action": {
					"type": "message",
					"label": "ดูข้อมูลรถ",
					"text": "ดูข้อมูลรถ: %s\nID: %d"
				  },
				  "color": "#AAAAAA",
				  "style": "primary"
			  },
			  {
				  "type": "button",
				  "action": {
					"type": "message",
					"label": "Add Car",
					"text": "เพิ่มรถ: %s"
				  },
				  "color": "#6BB583FF",
				  "style": "primary"
			  }
			]
		  }
	  }`
	jsonTmp := ""
	for i, car := range cars {
		if i == 0 {
			jsonTmp += fmt.Sprintf(jsonCars[1:], car.ImagePath, car.Model, humanize.Comma(int64(car.Price)), car.Model, car.ID, car.Model)
		} else {
			jsonTmp += fmt.Sprintf(jsonCars, car.ImagePath, car.Model, humanize.Comma(int64(car.Price)), car.Model, car.ID, car.Model)
		}
		
	}

	jsonStr = fmt.Sprintf(jsonStr, jsonTmp)
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	flexMessage := linebot.NewFlexMessage("Cars", flexContainer)

	return flexMessage
}

func GetMyCars(client *ent.Client, ctx context.Context, lineUser *ent.LineUser) *linebot.FlexMessage {
	// find Cars's user
	cars, err := lineUser.
					QueryCars().
					All(ctx)
	if err != nil {
		log.Fatal(err)
	}


	jsonStr := `{
		"type": "carousel",
		"contents": [%s]
	  }`
	
	
	jsonCars := `,{
		"type": "bubble",
		"hero": {
		  "type": "image",
		  "url": "%s",
		  "size": "full",
		  "aspectRatio": "20:13",
		  "aspectMode": "cover"
		},
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "text",
			  "text": "%s",
			  "weight": "bold",
			  "size": "xl",
			  "wrap": true,
			  "contents": []
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "contents": [
				{
				  "type": "text",
				  "text": "%s บาท",
				  "weight": "bold",
				  "size": "xl",
				  "flex": 0,
				  "wrap": true,
				  "contents": []
				}
			  ]
			}
		  ]
		},
		"footer": {
		  "type": "box",
		  "layout": "horizontal",
		  "spacing": "sm",
		  "contents": [
			{
				"type": "button",
				"action": {
				  "type": "message",
				  "label": "ดูข้อมูลรถ",
				  "text": "ดูข้อมูลรถ: %s\nID: %d"
				},
				"color": "#6BB583FF",
				"style": "primary"
			},
			{
				"type": "button",
				"action": {
				  "type": "message",
				  "label": "ลบ",
				  "text": "ลบรถ: %s\nID: %d"
				},
				"color": "#F77575FF",
				"style": "primary"
			}
		  ]
		}
	  }`
	jsonTmp := ""
	for i, car := range cars {
		if i == 0 {
			jsonTmp += fmt.Sprintf(jsonCars[1:], car.ImagePath, car.Model, humanize.Comma(int64(car.Price)), car.Model, car.ID, car.Model, car.ID)
		} else {
			jsonTmp += fmt.Sprintf(jsonCars, car.ImagePath, car.Model, humanize.Comma(int64(car.Price)), car.Model, car.ID, car.Model, car.ID)
		}
		
	}

	jsonStr = fmt.Sprintf(jsonStr, jsonTmp)
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	flexMessage := linebot.NewFlexMessage("Cars", flexContainer)

	return flexMessage
}

func GetInfoCars(client *ent.Client, ctx context.Context) *linebot.FlexMessage {
	jsonStr := `{
		"type": "carousel",
		"contents": [%s]
	  }`
	
	cars, err := client.Car.
					Query().
					All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	jsonCars := `,{
		"type": "bubble",
		"hero": {
		  "type": "image",
		  "url": "%s",
		  "size": "full",
		  "aspectRatio": "20:13",
		  "aspectMode": "cover"
		},
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "text",
			  "text": "%s",
			  "weight": "bold",
			  "size": "xl",
			  "wrap": true,
			  "contents": []
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "contents": [
				{
				  "type": "text",
				  "text": "%s บาท",
				  "weight": "bold",
				  "size": "xl",
				  "flex": 0,
				  "wrap": true,
				  "contents": []
				}
			  ]
			}
		  ]
		},
		"footer": {
		  "type": "box",
		  "layout": "vertical",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "button",
			  "action": {
				"type": "message",
				"label": "ดูข้อมูฃ",
				"text": "ดูข้อมูลรถ: %s\nID: %d"
			  },
			  "color": "#6BB583FF",
			  "style": "primary"
			}
		  ]
		}
	  }`
	jsonTmp := ""
	for i, car := range cars {
		if i == 0 {
			jsonTmp += fmt.Sprintf(jsonCars[1:], car.ImagePath, car.Model, humanize.Comma(int64(car.Price)), car.Model, car.ID)
		} else {
			jsonTmp += fmt.Sprintf(jsonCars, car.ImagePath, car.Model, humanize.Comma(int64(car.Price)), car.Model, car.ID)
		}
		
	}

	jsonStr = fmt.Sprintf(jsonStr, jsonTmp)
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	flexMessage := linebot.NewFlexMessage("Cars", flexContainer)

	return flexMessage
}

func GetCarDetail(car *ent.Car) *linebot.FlexMessage {
	jsonStr := `{
		"type": "bubble",
		"direction": "ltr",
		"header": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "%s",
			  "weight": "bold",
			  "size": "lg",
			  "align": "center",
			  "contents": []
			},
			{
			  "type": "separator",
			  "margin": "sm",
			  "color": "#00C301FF"
			}
		  ]
		},
		"hero": {
		  "type": "image",
		  "url": "%s",
		  "size": "full",
		  "aspectRatio": "1.51:1",
		  "aspectMode": "cover"
		},
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "separator",
			  "margin": "none",
			  "color": "#00C301FF"
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "margin": "lg",
			  "contents": [
				{
				  "type": "text",
				  "text": "หมายเลขรถยนต์ :",
				  "size": "sm",
				  "color": "#AAA9A9FF",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "%d",
				  "size": "sm",
				  "wrap": true,
				  "contents": []
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "margin": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "Model :",
				  "size": "sm",
				  "color": "#AAA9A9FF",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "%s",
				  "size": "sm",
				  "wrap": true,
				  "contents": []
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "margin": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "ราคา :",
				  "size": "sm",
				  "color": "#AAA9A9FF",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "%s บาท",
				  "size": "sm",
				  "wrap": true,
				  "contents": []
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "margin": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "วันที่ผลิต :",
				  "size": "sm",
				  "color": "#AAA9A9FF",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "%s",
				  "size": "sm",
				  "contents": []
				}
			  ]
			}
		  ]
		}
	  }`
	json := fmt.Sprintf(jsonStr, car.Model, car.ImagePath, car.ID, car.Model, humanize.Comma(int64(car.Price)), car.RegisteredAt.Format("02/01/2006"))

	// Creater Flex Container
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(json))
	if err != nil {
		log.Panicln(err)
	}

	flexMessage := linebot.NewFlexMessage("CarDetail", flexContainer)

	return flexMessage
}

func GetGroups(client *ent.Client, ctx context.Context) *linebot.FlexMessage {
	
	jsonStr := `{
		"type": "carousel",
		"contents": [%s]
	  }`
	
	groups, err := client.Group.
						Query().
						All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	jsonGroups := `,{
		"type": "bubble",
		"header": {
		  "type": "box",
		  "layout": "vertical",
		  "flex": 0,
		  "contents": [
			{
			  "type": "text",
			  "text": "Group : %s",
			  "weight": "bold",
			  "size": "lg",
			  "align": "center",
			  "contents": []
			}
		  ]
		},
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "box",
			  "layout": "baseline",
			  "contents": [
				{
				  "type": "text",
				  "text": "จำนวนสมาชิก",
				  "color": "#AAA9A9FF",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "%d",
				  "contents": []
				}
			  ]
			}
		  ]
		},
		"footer": {
		  "type": "box",
		  "layout": "horizontal",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "button",
			  "action": {
				"type": "message",
				"label": "ดูสมาชิกกลุ่ม",
				"text": "สมาชิกกลุ่ม: %s"
			  },
			  "color": "#AAAAAA",
			  "style": "primary"
			},
			{
			  "type": "button",
			  "action": {
				"type": "message",
				"label": "เข้าร่วมกลุ่ม",
				"text": "เข้ากลุ่ม: %s"
			  },
			  "color": "#6BB583FF",
			  "style": "primary"
			}
		  ]
		}
	  }`
	jsonTmp := ""
	for i, group := range groups {
		if i == 0 {
			jsonTmp += fmt.Sprintf(jsonGroups[1:], group.Name, group.QueryLineusers().CountX(ctx), group.Name, group.Name)
		} else {
			jsonTmp += fmt.Sprintf(jsonGroups, group.Name, group.QueryLineusers().CountX(ctx), group.Name, group.Name)
		}
		
	}

	jsonStr = fmt.Sprintf(jsonStr, jsonTmp)
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	flexMessage := linebot.NewFlexMessage("Groups", flexContainer)

	return flexMessage	
}

func GetMemberGroup(group *ent.Group, ctx context.Context) *linebot.FlexMessage {
	jsonStr := `{
		"type": "bubble",
		"direction": "ltr",
		"header": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "รายชื่อสมาชิก กลุ่ม: %s",
			  "weight": "bold",
			  "size": "lg",
			  "align": "center",
			  "contents": []
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "margin": "md",
			  "contents": [
				{
				  "type": "text",
				  "text": "จำนวนสมาชิก: ",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "%d",
				  "contents": []
				}
			  ]
			},
			{
			  "type": "separator",
			  "margin": "md",
			  "color": "#00C301FF"
			}%s
		  ]
		}
	  }`

	tmpStr := ""
	memberStr := `,
	{
	  "type": "box",
	  "layout": "baseline",
	  "margin": "sm",
	  "contents": [
		{
		  "type": "text",
		  "text": "No. %d",
		  "contents": []
		},
		{
		  "type": "text",
		  "text": "%s",
		  "wrap": true,
		  "contents": []
		}
	  ]
	}`
	users := group.QueryLineusers().AllX(ctx)
	for i, user := range users {
		log.Println(i, user)
		tmpStr += fmt.Sprintf(memberStr, i+1, user.DisplyaName)
	}
	json := fmt.Sprintf(jsonStr, group.Name ,group.QueryLineusers().CountX(ctx),tmpStr)
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(json))
	if err != nil {
		log.Fatal(err)
	}
	flexMessage := linebot.NewFlexMessage("Member", flexContainer)
	return flexMessage
}

func GetMyGroups(line_user *ent.LineUser, ctx context.Context) *linebot.FlexMessage {
	
	jsonStr := `{
		"type": "carousel",
		"contents": [%s]
	  }`
	
	groups, err := line_user.QueryGroups().
						All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	jsonGroups := `,{
		"type": "bubble",
		"header": {
		  "type": "box",
		  "layout": "vertical",
		  "flex": 0,
		  "contents": [
			{
			  "type": "text",
			  "text": "Group : %s",
			  "weight": "bold",
			  "size": "lg",
			  "align": "center",
			  "contents": []
			}
		  ]
		},
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "box",
			  "layout": "baseline",
			  "contents": [
				{
				  "type": "text",
				  "text": "จำนวนสมาชิก",
				  "color": "#AAA9A9FF",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "%d",
				  "contents": []
				}
			  ]
			}
		  ]
		},
		"footer": {
		  "type": "box",
		  "layout": "horizontal",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "button",
			  "action": {
				"type": "message",
				"label": "ดูสมาชิกกลุ่ม",
				"text": "สมาชิกกลุ่ม: %s"
			  },
			  "color": "#AAAAAA",
			  "style": "primary"
			},
			{
			  "type": "button",
			  "action": {
				"type": "message",
				"label": "ออกจากกลุ่ม",
				"text": "ออกจากกลุ่ม: %s"
			  },
			  "color": "#F77575FF",
			  "style": "primary"
			}
		  ]
		}
	  }`
	jsonTmp := ""
	for i, group := range groups {
		if i == 0 {
			jsonTmp += fmt.Sprintf(jsonGroups[1:], group.Name, group.QueryLineusers().CountX(ctx), group.Name, group.Name)
		} else {
			jsonTmp += fmt.Sprintf(jsonGroups, group.Name, group.QueryLineusers().CountX(ctx), group.Name, group.Name)
		}
		
	}

	jsonStr = fmt.Sprintf(jsonStr, jsonTmp)
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	flexMessage := linebot.NewFlexMessage("MyGroups", flexContainer)

	return flexMessage	
}

func GetGroupMenu() *linebot.FlexMessage {
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(`{
		"type": "bubble",
		"direction": "ltr",
		"header": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "Group Menu",
			  "weight": "bold",
			  "size": "lg",
			  "align": "center",
			  "contents": []
			}
		  ]
		},
		"footer": {
		  "type": "box",
		  "layout": "vertical",
		  "spacing": "md",
		  "contents": [
			{
			  "type": "button",
			  "action": {
				"type": "message",
				"label": "ดูรายชื่อกลุ่ม",
				"text": "กลุ่ม"
			  },
			  "color": "#6BB583FF",
			  "style": "primary"
			},
			{
			  "type": "button",
			  "action": {
				"type": "message",
				"label": "ดูกลุ่มของฉัน",
				"text": "กลุ่มของฉัน"
			  },
			  "color": "#6BB583FF",
			  "style": "primary"
			}
		  ]
		}
	  }`))
	
	if err != nil {
		log.Fatal(err)
	}

	flexMessage := linebot.NewFlexMessage("GroupMenu", flexContainer)
	return flexMessage
}