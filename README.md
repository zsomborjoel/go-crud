# go-crud
Sample crud with gin and gorm

## How to run the application

1.	On linux just go into the app directory
2.	Run like this ./go-crud
      
### Insert:
```
curl -X POST localhost:3000/posts -H 'Content-Type: application/json' -d '{"Title":"newtitle","Body":"newbody"}'
```

### Update:
```
curl -X POST localhost:3000/posts/1 -H 'Content-Type: application/json' -d '{"Title":"updatetitle","Body":"updatebody"}'
```

### Delete:
```
curl -X DELETE localhost:3000/posts/1
```

### Get One:
      
```
curl localhost:3000/posts/1
```

### Get All:
```
curl localhost:3000/posts/
```

