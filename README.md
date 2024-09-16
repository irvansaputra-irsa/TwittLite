# Final Project Golang Sanbercode - Twittlite

This project includes APIs for managing posts, comments, followers, and user accounts, using Bearer Token authentication.

## Authentication
- **Token Variable**: `{{token}}`

---

## Posts API

### 1. Create a Post
- **URL**: `http://localhost:8080/api/posts`
- **Method**: `POST`
- **Body**:
    ```json
    {
        "content": string
    }
    ```
- **Authentication**: Bearer Token required

---

### 2. Get User Posts
- **URL**: `http://localhost:8080/api/posts/user/:user-id`
- **Method**: `GET`
- **Authentication**: Bearer Token required

---

### 3. Update a Post
- **URL**: `http://localhost:8080/api/posts`
- **Method**: `PUT`
- **Body**:
    ```json
    {
        "id": number,
        "content": string
    }
    ```
- **Authentication**: Bearer Token required

---

### 4. Delete a Post
- **URL**: `http://localhost:8080/api/posts/:post-id`
- **Method**: `DELETE`
- **Authentication**: Bearer Token required

---

### 5. Get Post Details
- **URL**: `http://localhost:8080/api/posts/:post-id`
- **Method**: `GET`
- **Authentication**: Bearer Token required

---

### 6. List Timeline Posts
- **URL**: `http://localhost:8080/api/posts/timeline`
- **Method**: `GET`
- **Authentication**: Bearer Token required

---

## Followers API

### 1. Follow a User
- **URL**: `http://localhost:8080/api/follows/:user-id`
- **Method**: `POST`
- **Authentication**: Bearer Token required

---

### 2. Get Following List
- **URL**: `http://localhost:8080/api/follows/follower/:user-id`
- **Method**: `GET`
- **Authentication**: Bearer Token required

---

### 3. Get Follower List
- **URL**: `http://localhost:8080/api/follows/follower/:user-id`
- **Method**: `GET`
- **Authentication**: Bearer Token required

---

## Comments API

### 1. Create a Comment
- **URL**: `http://localhost:8080/api/comments`
- **Method**: `POST`
- **Body**:
    ```json
    {
        "content": string,
        "post_id": number
    }
    ```
- **Authentication**: Bearer Token required

---

### 2. Get Post Comments
- **URL**: `http://localhost:8080/api/comments/post/:post-id`
- **Method**: `GET`
- **Authentication**: Bearer Token required

---

### 3. Delete a Comment
- **URL**: `http://localhost:8080/api/comments/:comment-id`
- **Method**: `DELETE`
- **Authentication**: Bearer Token required

---

### 4. Update a Comment
- **URL**: `http://localhost:8080/api/comments`
- **Method**: `PUT`
- **Body**:
    ```json
    {
        "id": number,
        "content": string
    }
    ```
- **Authentication**: Bearer Token required

---

### 5. Get User Comments
- **URL**: `http://localhost:8080/api/comments/user/:user-id`
- **Method**: `GET`
- **Authentication**: Bearer Token required

---

### 6. Get Comment Details
- **URL**: `http://localhost:8080/api/comments/:comment-id`
- **Method**: `GET`
- **Authentication**: Bearer Token required

---

## User API

### 1. Register a User
- **URL**: `http://localhost:8080/api/users/register`
- **Method**: `POST`
- **Body**:
    ```json
    {
        "username": string,
        "password": string,
        "email": string
    }
    ```

---

### 2. Login a User
- **URL**: `http://localhost:8080/api/users/login`
- **Method**: `POST`
- **Body**:
    ```json
    {
        "password":string,
        "email": string
    }
    ```

---

### 3. Get User Details
- **URL**: `http://localhost:8080/api/users/:user-id`
- **Method**: `GET`
- **Authentication**: Bearer Token required

---

### 4. Edit User Profile
- **URL**: `http://localhost:8080/api/users`
- **Method**: `POST`
- **Body**:
    ```json
    {
        "bio": string,
        "location": string
    }
    ```
- **Authentication**: Bearer Token required
