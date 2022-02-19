# github.com/thiagoluiznunes/strider-challenge
---
**Goal**: Build out a RESTful API and corresponding backend system to handle the features detailed above. This RESTful API would communicate with single page JS app. This API you build should enable all the features on both of the pages.

Tools: Golang | Docker | Docker-compose

### Requirements ###

It is necessary to install previously Docker and Docker-compose

* **[Docker 20.10.x](https://docs.docker.com)** :white_check_mark:
* **[Docker compose 1.29.x](https://docs.docker.com/compose/)** :white_check_mark:

### Project installation ###

**Obs.: The following instructions was performed in the macOS BigSur distribution**

1 - Unzip the file Thiago-Nunes-1-0-web-back-end.zip in your choose directory:
  - user@user:~/your_directory/$ **unzip command Thiago-Nunes-1-0-web-back-end.zip**
  - make sure the tcp ports of number :8080, :3306 are available

2 - In the root project directory run the followed command:
  - **obs**: it is expected the Docker is running
  - **docker-compose up -d --force-recreate**
  - endpoint to access the API is http://localhost:8080

3 - The execution of integration tests could be performed using the Go/Golang command as followed:
  - it is needed to install previously Go/Golang to run the tests, it is suggested to install the latest version
```
go test ./...
```

### API Routes ###
|   Action                                |  Riquered  | Role  |  Method  | URL
|   --------------------------------------|------------| ----- |----------|--------------
|   RETRIEVE ALL POSTS                    |            |       | `GET`    | /homepage/?switch=all
|   RETRIEVE ALL FOLLOWING POSTS          |            |       | `GET`    | /homepage/?switch=following
|   ADD NEW POST                          |            |       | `POST`   | /homepage/post
|   RETRIEVE MOCK USER                    |            |       | `GET`    | /user/?user_id=1
|   MOCK USER IS FOLLOWING ANOTHER USER   |            |       | `POST`   | /user/follow
|   MOCK USER IS UNFOLLOWING ANOTHER USER |            |       | `POST`   | /user/unfollow
|   HELTH CHECK                           |            |       | `GET`    | /helth

#### RETRIEVE ALL POSTS ####
* REQUEST
```
GET /homepage/?switch=all
```
* RESPONSE
```json
{
    "posts": [
        {
            "uuid": "9d2b1a55-8aae-4f52-85b4-9e5b67ddf66f",
            "type": "original",
            "text": "text",
            "user_id": 1,
            "updated_at": "2022-02-17T18:23:05Z",
            "created_at": "2022-02-17T18:23:05Z"
        }
    ]
}
```

#### RETRIEVE ALL FOLLOWING POSTS ####
* REQUEST
```
GET /homepage/?switch=following
```
* RESPONSE
```json
{
    "posts": [
        {
            "uuid": "9d2b1a55-8aae-4f52-85b4-9e5b67ddf66f",
            "type": "original",
            "text": "text",
            "user_id": 1,
            "updated_at": "2022-02-17T18:23:05Z",
            "created_at": "2022-02-17T18:23:05Z"
        }
    ]
}
```

#### ADD NEW ORIGINAL POST ####
* REQUEST
```
POST /homepage/post
```
```json
{
    "user_id": 1,
    "type": "original",
    "text": "text original post"
}
```
* RESPONSE
```json
"OK"
```

#### ADD NEW REPOST POST ####
* REQUEST
```
POST /homepage/post
```
```json
{
    "user_id": 1,
    "type": "repost",
    "text": "text repost post"
}
```
* RESPONSE
```json
"OK"
```

#### ADD NEW QUOTE POST ####
* REQUEST
```
POST /homepage/post
```
```json
{
    "user_id": 1,
    "type": "quote",
    "text": "text quote post"
}
```
* RESPONSE
```json
"OK"
```

#### RETRIEVE MOCK USER ####
* REQUEST
```
GET /user/?user_id=1
```
* RESPONSE
```json
{
    "user_name": "thiagonunes",
    "number_of_followers": 0,
    "number_of_following": 0,
    "number_of_posts": 0,
    "created_at": "2022-02-18T18:21:25.206862-03:00"
}
```

#### MOCK USER IS FOLLOWING ANOTHER USER ####
* REQUEST
```
GET /user/follow
```
* REQUEST
```json
{
    "user_id": 1,
    "following_user_id": 2
}
```
* RESPONSE
```json
"OK"
```

#### MOCK USER IS UNFOLLOWING ANOTHER USER ####
* REQUEST
```
GET /user/unfollow
```
* REQUEST
```json
{
    "user_id": 1,
    "following_user_id": 2
}
```
* RESPONSE
```json
"OK"
```

#### HELTH CHECK ####
* REQUEST
```
GET /health
```
* RESPONSE
```json
"OK"
```

### Autor

* Thiago Luiz Pereira Nunes ([ThiagoLuizNunes](github.com/thiagoluiznunes)) thiagoluiz.dev@gmail.com

>Created by **[ThiagoLuizNunes](https://www.linkedin.com/in/thiago-luiz-507483112/)** 2022.

---