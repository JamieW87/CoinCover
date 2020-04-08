# CoinCover Technical Task

## Description
A Golang backend that sends API calls to a MongoDB database, both contained within Docker containers. The backend makes two API calls to the database, /user creating a user from a JSON body and /users getting all users currently stored in the database.

## Installation

### MongoDB:

There are two ways to spin up the MongoDB docker container:

- Pull the MongoDB image from Docker repository:

`docker pull mongo:4.0.4`

- Spin up a docker container using the MongoDB image. 

`docker run -d -p 27017-27019:27017-27019 --name mongodb mongo:4.0.4`

- Connect to mongo interactive terminal

`docker exec -it mongodb bash`

- Launch MongoDB shell client:

`mongo`

- Create coinCover database

`use coinCover`

- Now your MongoDB container is running in the background and you have created the necessary DB to receive the API calls.

OR, you can use the `docker-compose.yml` in the root folder to start a mongoDB container:

`docker-compose up`

### API

- Build dockerfile

`docker build . -t coin-app:latest`

- Run dockerfile, open ports and place on same network as DB:

`docker run --name coin-app --rm -p 8080:8080 --network techtask_network1 coin-app:latest`



