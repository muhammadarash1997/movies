# Movies Service
Movies Service is backend service for managing movie data

## Technologies:
* Golang 1.19
* MySQL
* gRPC
* REST

## APIs:
* Add Movie
* Get Movie List
* Get Movie Detail
* Update Movie
* Delete Movie

## Demo APIs:
You have to do all things in the installation section before you make a demo. You can use REST or gRPC:
* Using REST
   - In order to demo all apis in this movies service using rest connection, make sure you already have postman afterwards import movies.swagger.json file in proto folder into your postman
* Using gRPC
   - In order to demo all apis in this movies service using grpc connection, make sure you already have bloomrpc afterwards import movies.proto file in proto folder into your bloomrpc

## Installation:
* Step:
    * Note:
        - Make sure docker and docker compose are already installed
    * Run application:
        - Build and run movies service and mysql, it will run in local at port 8080, running the following command:
           ```bash
           $ docker-compose up
           ```
    * Create movies table:
        - Get the running mysql container id by running the following command:
           ```bash
           $ docker ps
           ```
        - Entering the mysql container by running the following command:
           ```bash
           $ docker exec -it <your_mysql_container_id> mysql -u root -p root
           ```
        - Move to use movie database by running the following command:
           ```mysql
           $ use movie
           ```
        - Run this following sql DDL command:
           ```mysql
           $ CREATE TABLE movies (
                    `id` int NOT NULL AUTO_INCREMENT,
                    `title` varchar(100) NOT NULL,
                    `description` varchar(100) NOT NULL,
                    `rating` int NOT NULL,
                    `image` varchar(100) NOT NULL,
                    `created_at` DATETIME NOT NULL,
                    `updated_at` DATETIME NOT NULL,
                    PRIMARY KEY (`id`)
                ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
            ```
