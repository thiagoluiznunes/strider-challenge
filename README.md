# strider-challenge
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
  - **docker-compose up --build -d --force-recreate**
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

## Planning

- What is the deadline time estimated to deploy this new feature in the production environment all tested?
- What was the opinion of the Product Design about it? Did you all have a previous talk about the impact and the requirements about this new functionality?
- To implement this new functionality both teams, backend and frontend need to meeting together to decide the conventional details about the communication between the frontend and backend applications.
- In the backend side we are going to redesign the database architecture, and maybe to change the to a Graph database, because we're using a OLTP database.
- We will need to create a new table called **reply_message** with a foreing_key to the **post** table, and a foreing_key to **user_table**, if we continue with the OLTP strategy.
- We will need to create a new route endpoint for the frontend application insert/retrieve the **reply_message** data.

## Critique

If I had more time, I'd like to improve:
1. Add more integration tests, because no all was made.
2. Create a tech docs using OpenAPI or Swagger.
3. Design an architecture.
4. Create a CI/CD workflow.
5. Design a AWS architecture to deploy the application in the Cloud. For example a simple architecture in the begining composed by: One Application Load Balancer, attached to a Auto Scaling Group to ECS Tasks Fargate, using a ECR image deployed by the CI/CD workflow.
6. Add observability and instrumentation using NewRelic or OpenTelemety.
7. Change the database strategy to a Graph database like AWS Neptune. Even because is better using a Graph database to social media application.

### Autor

* Thiago Luiz Pereira Nunes ([ThiagoLuizNunes](github.com/thiagoluiznunes)) thiagoluiz.dev@gmail.com

>Created by **[ThiagoLuizNunes](https://www.linkedin.com/in/thiago-luiz-507483112/)** 2022.

---