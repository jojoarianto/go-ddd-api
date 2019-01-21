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

#### /api/v1/news?limit={limit}&page={page}

- `GET` : Get all news with pagination limit and page

### Usage Examples

> Documentation api with insomnia https://intip.in/InsomniaApiKumparanTa

Get all news, URL GET `/api/v1/news`

```bash
curl --request GET \
  --url https://go-ddd-api.appspot.com/api/v1/news
```

Get all news filter by status['draft', 'publish', 'deleted'], URL GET `/api/v1/news?status={status}`

```bash
curl --request GET \
  --url https://go-ddd-api.appspot.com/api/v1/news?status=draft
```

Get all news filter by topic, URL GET `/api/v1/news/{topic-slug}`

```bash
curl --request GET \
  --url https://go-ddd-api.appspot.com/api/v1/news/liputan-khusus
```

Get all news with pagination, URL GET `/api/v1/news?limit={limit}&page={page}`

```bash
curl --request GET \
  --url https://go-ddd-api.appspot.com/api/v1/news?limit=2&page=1
```

Get all topics
```bash
curl --request GET \
  --url https://go-ddd-api.appspot.com/api/v1/topic
```

Create a topic, URL POST `/api/v1/topic`
```bash
curl --request POST \
  --url https://go-ddd-api.appspot.com/api/v1/topic \
  --header 'content-type: application/json' \
  --data '{
	"name":"Liputan Khusus",
	"slug":"liputan-khusus"
}'
```

Create a news, URL POST `/api/v1/news`
```bash
curl --request POST \
  --url https://go-ddd-api.appspot.com/api/v1/news \
  --header 'content-type: application/json' \
  --data '{
	"title": "Bonnie Triyana: TNI Razia Buku Upaya Desukarnoisasi",
	"slug": "memberangus-buku-memberangus-ilmu-1547439849539914993",
	"content": "30 November 1957, kunjungan Presiden Sukarno di Perguruan Cikini, Jakarta, atas undangan guru mendadak jadi tragedi. Hujan granat mendarat ketika ia berjalan keluar dari sekolah dua anaknya itu, Megawati Soekarnoputri dan Guruh Soekarnoputra. Dua pengawal, Oding Suhendar dan Sudiyo, merangkul Sukarno pergi menyelamatkan diri. Kedua anak Sukarno sudah lebih dulu diamankan.",
	"status": "publish",
	"Topic": [
		{
			"ID": 2,
			"CreatedAt": "2019-01-19T03:12:32Z",
			"UpdatedAt": "2019-01-19T03:12:32Z",
			"DeletedAt": null,
			"name": "Liputan Khusus",
			"slug": "liputan-khusus",
			"News": null
		}
	]
}'
```

Delete a news, URL DELETE `/api/v1/news/2`
```bash
curl --request DELETE \
  --url https://go-ddd-api.appspot.com/api/v1/news/2
```

Delete a topic, URL DELETE `/api/v1/topic/2`
```bash
curl --request DELETE \
  --url https://go-ddd-api.appspot.com/api/v1/topic/3
```

Update a news, URL PUT `/api/v1/topic/2`
```bash
curl --request PUT \
  --url https://go-ddd-api.appspot.com/api/v1/news/2 \
  --header 'content-type: application/json' \
  --data '{
	"title": "[draft] Memberangus Buku, Memberangus Ilmu",
	"slug": "memberangus-buku-memberangus-ilmu-1547439849539914993",
	"content": "Buku-buku yang disita itu berjudul Kronik ‘65: Catatan Hari Per Hari Peristiwa G30S Sebelum dan Sesudahnya, Jasmerah: Pidato-pidato Spektakuler Bung Karno Sepanjang Massa, dan Mengincar Bung Besar: Tujuh Upaya Pembunuhan Bung Karno. Tak ada satu pun judul buku yang memuat kata “PKI” atau “komunis” seperti yang dituduhkan",
	"status": "draft",
	"Topic": [
		{
			"ID": 2,
			"CreatedAt": "2019-01-19T03:12:32Z",
			"UpdatedAt": "2019-01-19T03:12:32Z",
			"DeletedAt": null,
			"name": "Liputan Khusus",
			"slug": "liputan-khusus",
			"News": null
		}
	]
}'
```

Update a topic, URL PUT `/api/v1/topic/2`
```bash
curl --request PUT \
  --url https://go-ddd-api.appspot.com/api/v1/topic/3 \
  --header 'content-type: application/json' \
  --data '{
	"name":"Sepak Bola Nasional",
	"slug":"sepak-bola-national"
}'
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
- GORM Pagination Extension : https://github.com/biezhi/gorm-paginator
- Toml : https://github.com/BurntSushi/toml
- Deploy GoApp on GCP GAE https://medium.com/google-cloud/deploying-your-go-app-on-google-app-engine-5f4a5c2a837
