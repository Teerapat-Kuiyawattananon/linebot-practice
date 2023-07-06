## Linebot
This is a Linebot for practice that learning how to use Line Messaging API with ent golang database


## How to run
To run this Linebot project locally, please follow these steps:
1. Clone the repository:
```
git clone https://github.com/Teerapat-Kuiyawattananon/linebot-practice.git
```
2. add linebot ID: `@551kqxbr`

3. Run Docker-compose:
```
docker-compose up -d
```

4. Run App:
```
go run start.go 
```
or
```
air 
```


5. Access to PostgreSQL
```
psql -h localhost -p 6789 -U teerapat
password admin1234
```

6. Public URL for webhook
```
ngrok http --region jp 7777
```

