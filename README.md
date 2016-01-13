# poc_frontend
small frontend app that calls another, only for demo

*NOTE:* this relys on https://github.com/richardbowden/poc_bk_rannumgen

#Check out

`go get github.com/richardbowden/poc_frontend`

#Build

```
cd $GOPATH/src/github.com/richardbowden/poc_frontend
go build
```

#Run
```
cd $GOPATH/src/github.com/richardbowden/poc_frontend
RND_BACKEND_URL="http://127.0.0.1:9090" ./poc_frontend 
```

RND_BACKEND_URL is an env variable that contains the url to the backend service for the random number generator

#Access
hit the frontend @ http://127.0.0.1:8080 or ip of your machine as the service binds to all ip's.

