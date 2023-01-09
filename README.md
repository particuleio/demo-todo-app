# Demo Todo App
3 tier demo app


# How it works ?
![diagram](/src/diagram/diagram.svg)

The app is inteded to be very simple

# Build and run the app

```bash
cd ./src

# get redis image
docker pull docker.io/library/redis:7
docker tag docker.io/library/redis:7 redis
```

``` bash
# build server image and 
docker build -t server-todo ./server
# build front image
docker build -t front-todo ./front
```

``` bash
# run three instances
docker run --rm --net=host --name=redis redis
docker run --rm --net=host --name=server localhost/server-todo
docker run --rm --net=host --name=front localhost/front-todo
```



