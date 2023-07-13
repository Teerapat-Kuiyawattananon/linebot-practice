## Linebot
This is a Linebot for practice that learning how to use Line Messaging API with ent golang database


## How to run
To run this Linebot project:
1. Clone the repository:
```
git clone https://github.com/Teerapat-Kuiyawattananon/linebot-practice.git
```
2. add linebot ID: `@551kqxbr`

3. copy .env.example file to .env file
```
cp .env.example .env
```
## Configuration [optional]
you can change this to your line bot config in `.env` file
```
LINE_CHANNEL_SECRET=your_line_channel_secret
LINE_CHANNEL_TOKEN=your_line_channel_token
```

4. Run Docker-compose:
```
docker-compose up -d
```

5. Run App:
```
go run start.go 
```
or
```
air 
```


6. Access to PostgreSQL
```
psql -h localhost -p 6789 -U teerapat
password admin1234
```

7. Public URL for webhook
```
ngrok http --region jp 7777
```

