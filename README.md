# GO DDD API

Kumparan Backend Technical Assessment, create REST API with domain driven approach (DDD) using Golang, GORM (Object Relational Mapping), and MySQL

> Demo get all news api https://go-ddd-api.appspot.com/api/v1/news

## Installation & Run

First, Make sure you have set up \$GOPATH.

```bash
# Download this project
go get github.com/jojoarianto/go-ddd-api

# It's take several minute to download project
```

Set project environment and run

```bash
# move to project directory
cd $GOPATH/src/github.com/jojoarianto/go-ddd-api

# copy and rename config.toml.example
cp config.toml.example config.toml

# set config.toml to your env
user="your_db_username"
password="your_db_password"
host="your_db_host"
port="your_db_port"
dbname="your_db_name"

# run golang project
go run main.go

# API Endpoint : http://localhost:8000/api/v1/
```

## Design

- Application
  - Write business logic
    - news.go (GetNews, GetAllNews, &...)
    - topic.go (GetTopic, GetAllTopic, &...)
- Domain
  - Define interface
    - repository interface for infrastructure
  - Define struct
    - Entity struct that represent mapping to data model
      - news.go
      - topic.go
- Infrastructure
  - Implements repository interface
    - news_repository.go
    - topic_repository.go
- Interfaces
  - HTTP handler

## Required Features

- `Manajement news` user can manage data news (CRUD)
- `Manajement topic` user can manage data topic (CRUD)
- `Relational model betwean news & topic` many to many (one news can contains multiple topic, one topic has multiple news)
- `filter by news status` filter news by it's status ['draft', 'deleted', 'publish']
- `filter by news topic` filter news by a topic (forinstance: politik)

## URL ENDPOINT

#### /api/v1/news

- `GET` : Get all news
- `POST` : Create a news

#### /api/v1/news/{news_id}

- `GET` : Get a news by id
- `PUT` : Update a news by id
- `DELETE` : Delete a news by id

#### /api/v1/topic

- `GET` : Get all topic
- `POST` : Create a topic

#### /api/v1/topic/{news_id}

- `GET` : Get a topic by id
- `PUT` : Update a topic by id
- `DELETE` : Delete a topic by id

#### /api/v1/news?status={status}

- `GET` : Get all news filter by news.status

#### /api/v1/news/{topic-slug}

- `GET` : Get all news filter by topic

### Usage Examples

Get all news, URL GET `/api/v1/news`

```bash
curl -i -H "Accept: application/json" https://go-ddd-api.appspot.com/api/v1/news
```

Get all news filter by status['draft', 'publish', 'deleted'], URL GET `/api/v1/news?status={status}`

```bash
curl -i -H "Accept: application/json" https://go-ddd-api.appspot.com/api/v1/news?status=draft
```

Get all news filter by topic, URL GET `/api/v1/news/{topic-slug}`

```bash
curl -i -H "Accept: application/json" https://go-ddd-api.appspot.com/api/v1/news/berita
```

## Product Items Backlog

- [x] **Mandatory:** Create REST API News & Topic CRUD
  - [x] News
    - [x] Get all
    - [x] Get by id
    - [x] Create
    - [x] Update
    - [x] Delete
  - [x] Topic
    - [x] Get all topic
    - [x] Get by id
    - [x] Create
    - [x] Update
    - [x] Delete
- [x] **Mandatory:** Create Filter
  - [x] Filter by status news
  - [x] Filter by topic
- [ ] **Mandatory:** API Functional Test
- [x] **Opsional:** Deploy to (heroku/aws/azure/digital ocean)
- [x] **Opsional:** Database setup migration schema DB

## References & Library

- DDD Skeleton : https://github.com/takashabe/go-ddd-sample
- Httprouter : https://github.com/julienschmidt/httprouter
- GORM Documentation : http://doc.gorm.io
- Toml : https://github.com/BurntSushi/toml
- Deploy GoApp on GCP GAE https://medium.com/google-cloud/deploying-your-go-app-on-google-app-engine-5f4a5c2a837
