# post
curl -X POST localhost:3000/api/posts -H 'Content-Type: application/json' -d '{"Title":"newtitle","Body":"newbody"}'

# update
curl -X POST localhost:3000/api/posts/1 -H 'Content-Type: application/json' -d '{"Title":"updatetitle","Body":"updatebody"}'

# get all
curl localhost:3000/api/posts

# get by id
curl localhost:3000/api/posts/1

# delete
curl -X DELETE localhost:3000/api/posts/1
