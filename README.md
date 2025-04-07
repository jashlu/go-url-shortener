## URL shortener with Go

### Steps To Run


#### 1. Install and Start Up Redis Server
Use the following to check if you have it installed: 
```shell
which redis-server
```

If installed, run the following to start up the redis server:
```shell
redis-server
``` 


With redis started, in another terminal, start the program:
```shell
go run main.go
```


#### 2. Setting Up Request Using Postman
1. Open Postman 
2. Select "POST" from the dropdown 
3. Enter the Request URL
```
http://localhost:9808/create-short-url
```
4. Set Headers 
```
Key: Content-Type
Value: application/json
```
5. Add Request Body
Make sure the format is set to JSON
```JSON
{
    "long_url": "...",
    "user_id": "..."
}
```
6. Send Request. 
7. You should receive a success message that looks like the following
```JSON
{
    "message": "short url created successfully",
    "short_url": "http://localhost:9808/..."
}
```
Check the provided short_url, it should automatically re-direct you to the original long_url