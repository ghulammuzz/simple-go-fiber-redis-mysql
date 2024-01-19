# Simple Go Fiber Redis MySQL

A simple Go script using Fiber, Redis, and MySQL for basic read and write operations.

## Overview

This script demonstrates a simple web application built with [Fiber](https://github.com/gofiber/fiber/v2), using Redis for caching and MySQL for data storage. It provides two endpoints: `/read` for reading data from Redis and `/write` for writing data to MySQL and Redis.

## Prerequisites

Before running the application, make sure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.16 or later)
- [Redis](https://redis.io/download)
- [MySQL](https://dev.mysql.com/downloads/)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/ghulammuzz/simple-go-fiber-redis-mysql.git
2. Change into the project directory:
    ```bash
    cd simple-go-fiber-redis-mysql
3. Install dependencies:
    ```bash
    go mod tidy
## Configuration
Adjust the configuration in the `main.go` file to match your Redis and MySQL settings.
```bash
const (
    redisKey = "product_key"
    mysql = "user:password@tcp(localhost:3306)/syncsb"
)
```

Execute a sql query 
```bash 
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);
```

## Usage
Run the application:
```bash
go run main.go
```
Visit the following endpoints in your browser or using a tool like curl:

- Read data from Redis: http://localhost:3000/read
- Write data to MySQL and Redis: Send a POST request to http://localhost:3000/write