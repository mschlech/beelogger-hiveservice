# beelogger-service
beelogger-service


# start the docker container
CONTAINER_NAME = docker images
docker run -t -i -p 3000:3000 -ti --rm --init <containerName>

# create a mongodb container
docker run --name mongohive -v $(pwd):/data/db -p 27017:27017 -d mongo:latest