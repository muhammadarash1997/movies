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
           ```bash
           $ use movie
           ```
        - Run this following sql DDL command:
           ```bash
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
