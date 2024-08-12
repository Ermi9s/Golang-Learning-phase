# Project: Clean-Architecture

## End-point: Register-api
### Register User

This endpoint allows the client to register a new user.

#### Request Body

- `username` (string, required): The username of the user.
    
- `email` (string, required): The email address of the user.
    
- `password` (string, required): The password for the user account.
    

#### Response

The response will be a JSON object with the following schema:

``` json
{
  "data": {
    "token": "string",
    "user": {
      "_id": "string",
      "username": "string",
      "email": "string",
      "password": "string",
      "is_admin": "boolean"
    }
  }
}

 ```

#### Example

``` json
{
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
    "user": {
      "_id": "5f64a2d8a3b4d8f1f8b9d5a2",
      "username": "example_user",
      "email": "user@example.com",
      "password": "********",
      "is_admin": false
    }
  }
}

 ```
### Method: POST
>```
>http://127.0.0.1:8080/api/register/
>```
### Body (**raw**)

```json
{
    "username" : "Ermias",
    "email" : "Ermias.ay@gmail.com",
    "password" : "1234"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: log-in-aip
### API Request Description

This endpoint allows the user to log in by providing their username, email, and password in the request body.

#### Request Body

- `username` (string): The username of the user.
    
- `email` (string): The email of the user.
    
- `password` (string): The password of the user.
    

### Response

The response is in JSON format and includes the following schema:

``` json
{
  "data": {
    "token": "",
    "user": {
      "_id": "",
      "username": "",
      "email": "",
      "password": "",
      "is_admin": true
    }
  }
}

 ```

The response includes a `token` for authentication and user information such as `_id`, `username`, `email`, `password`, and `is_admin` status.
### Method: POST
>```
>http://127.0.0.1:8080/api/log-in/
>```
### Body (**raw**)

```json
{
    "username" : "Ermias",
    "email" : "Ermias@gmail.com",
    "password" : "1234"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Create-Task-api
This API endpoint allows you to create a new task by sending an HTTP POST request to the specified URL. The request should include a JSON payload in the raw request body with the following parameters:

- "title": (string) The title of the task.
    
- "description": (string) The description of the task.
    
- "Status": (string) The status of the task.
    

Upon successful execution, the API will respond with a status code of 200 and a JSON object containing the details of the newly created task, including the following fields:

- "_id": The unique identifier of the task.
    
- "title": The title of the task.
    
- "description": The description of the task.
    
- "status": The status of the task.
    
- "date": The date of the task.
    
- "duedate": The due date of the task.
    
- "creator_id": The unique identifier of the creator of the task.
### Method: POST
>```
>http://127.0.0.1:8080/api/task/
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2NmI1ZGE3NTQ5YWJiM2M0ZTFkNjQ4M2QiLCJlbWFpbCI6IkVybWlhc0BnbWFpbC5jb20iLCJpc19hZG1pbiI6ZmFsc2V9.inv8W5yNpni4y_XVvQBeftjHggP48yAXZxPCd5EeqzU|


### Headers

|Content-Type|Value|
|---|---|
|Content-Type|application/json|


### Body (**raw**)

```json
{
    "title" : "clean task",
    "description" : "documenting the taskmanager clen arc",
    "Status" : "Active"
}
```

### 🔑 Authentication noauth

|Param|value|Type|
|---|---|---|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get-Task-by-ID
The endpoint retrieves the details of a specific task identified by the provided ID.

The response for this request can be documented as a JSON schema:

``` json
{
  "type": "object",
  "properties": {
    "data": {
      "type": "object",
      "properties": {
        "_id": {"type": "string"},
        "title": {"type": "string"},
        "description": {"type": "string"},
        "status": {"type": "string"},
        "date": {"type": "string"},
        "duedate": {"type": "string"},
        "creator_id": {"type": "string"}
      }
    }
  }
}

 ```
### Method: GET
>```
>http://127.0.0.1:8080/s-api/task/66b5df6c0b65b192b8a1276b
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2NmI1ZGE3NTQ5YWJiM2M0ZTFkNjQ4M2QiLCJlbWFpbCI6IkVybWlhc0BnbWFpbC5jb20iLCJpc19hZG1pbiI6ZmFsc2V9.inv8W5yNpni4y_XVvQBeftjHggP48yAXZxPCd5EeqzU|


### 🔑 Authentication noauth

|Param|value|Type|
|---|---|---|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get-all-Task-id
### GET /s-api/task/

This endpoint retrieves a list of tasks.

#### Request

No request body is required for this endpoint.

#### Response

The response will be in JSON format with the following schema:

``` json
{
    "data": [
        {
            "_id": "",
            "title": "",
            "description": "",
            "status": "",
            "date": "",
            "duedate": "",
            "creator_id": ""
        }
    ]
}

 ```

The response will contain an array of tasks, where each task object includes the following properties:

- `_id`: The unique identifier of the task.
    
- `title`: The title of the task.
    
- `description`: The description of the task.
    
- `status`: The status of the task.
    
- `date`: The date of the task.
    
- `duedate`: The due date of the task.
    
- `creator_id`: The unique identifier of the creator of the task.
### Method: GET
>```
>http://127.0.0.1:8080/s-api/task/
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2NmIyOGZiMmQxY2YyYzMzNmUwODQ1ZjAiLCJlbWFpbCI6InJvb3RAZ21haWwuY29tIiwiaXNfYWRtaW4iOnRydWV9.Hd-uPoKnQ-49whLa0_nTYDfmS-Bb4hxNVipAZFT5T2s|


### 🔑 Authentication noauth

|Param|value|Type|
|---|---|---|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get all users - protected (only admin)
# Retrieve Protected User Data

This endpoint makes an HTTP GET request to retrieve protected user data.

## Request

### Request URL

- Type: HTTP
    
- URL: `http://127.0.0.1:8080/protected/user/`
    

## Response

- Status: 200
    
- Content-Type: application/json
    

### Response Body

The response contains an array of user data objects, including the user ID, username, email, password, and admin status.

Example response body:

``` json
{
    "data":[
        {
            "_id":"",
            "username":"",
            "email":"",
            "password":"",
            "is_admin":true
        }
    ]
}

 ```
### Method: GET
>```
>http://127.0.0.1:8080/protected/user/
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2NmIyOGZiMmQxY2YyYzMzNmUwODQ1ZjAiLCJlbWFpbCI6InJvb3RAZ21haWwuY29tIiwiaXNfYWRtaW4iOnRydWV9.Hd-uPoKnQ-49whLa0_nTYDfmS-Bb4hxNVipAZFT5T2s|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get-User-By-Id (only owner or admin)
The endpoint retrieves user information based on the provided user ID. The response returns a JSON object with the user data.

``` json
{
  "type": "object",
  "properties": {
    "data": {
      "type": "object",
      "properties": {
        "_id": {
          "type": "string"
        },
        "username": {
          "type": "string"
        },
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "is_admin": {
          "type": "boolean"
        }
      }
    }
  }
}

 ```
### Method: GET
>```
>http://127.0.0.1:8080/s-api/user/66b60ac21d4e14dbf27374df
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2NmI2MGFjMjFkNGUxNGRiZjI3Mzc0ZGYiLCJlbWFpbCI6IkVybWlhc0BnbWFpbC5jb20iLCJpc19hZG1pbiI6ZmFsc2V9.4PAMUu53OSKlDVpuoUau2Ap6h4pR0TeBUT5Dr4f7GFs|


### 🔑 Authentication noauth

|Param|value|Type|
|---|---|---|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Update-Task-api
This endpoint makes an HTTP PUT request to update a specific task identified by the provided ID. The request body should include the updated title, description, and status of the task.

### Request Body

- `title` (string, optional): The updated title of the task.
    
- `description` (string, optional): The updated description of the task.
    
- `Status` (string, optional): The updated status of the task.
    

### Response

Upon a successful execution (Status 200), the response will be in JSON format with the updated data of the task, including the task ID, title, description, status, date, due date, and creator ID.

Example Response Body:

``` json
{
    "data": {
        "_id": "",
        "title": "",
        "description": "",
        "status": "",
        "date": "",
        "duedate": "",
        "creator_id": ""
    }
}

 ```
### Method: PUT
>```
>http://127.0.0.1:8080/s-api/task/66b5df6c0b65b192b8a1276b
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2NmI1ZGE3NTQ5YWJiM2M0ZTFkNjQ4M2QiLCJlbWFpbCI6IkVybWlhc0BnbWFpbC5jb20iLCJpc19hZG1pbiI6ZmFsc2V9.inv8W5yNpni4y_XVvQBeftjHggP48yAXZxPCd5EeqzU|


### Body (**raw**)

```json
{
    "title" : "clean task updated",
    "description" : "update comming",
    "Status" : "Done"
}
```

### 🔑 Authentication noauth

|Param|value|Type|
|---|---|---|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Promote to admin-api(protected)
The HTTP PUT request is used to promote a user to admin status. The request is sent to the endpoint [http://127.0.0.1:8080/protected/promote/](http://127.0.0.1:8080/protected/promote/66b5da7549abb3c4e1d6483d):id

### Response

The response returned with a status code of 202 and a content type of application/json. The response body follows the JSON schema below:

``` json
{
    "type": "object",
    "properties": {
        "data": {
            "type": "object",
            "properties": {
                "_id": {"type": "string"},
                "username": {"type": "string"},
                "email": {"type": "string"},
                "password": {"type": "string"},
                "is_admin": {"type": "boolean"}
            }
        }
    }
}

 ```
### Method: PUT
>```
>http://127.0.0.1:8080/protected/promote/66b5da7549abb3c4e1d6483d
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2NmIyOGZiMmQxY2YyYzMzNmUwODQ1ZjAiLCJlbWFpbCI6InJvb3RAZ21haWwuY29tIiwiaXNfYWRtaW4iOnRydWV9.Hd-uPoKnQ-49whLa0_nTYDfmS-Bb4hxNVipAZFT5T2s|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Update user-account
The endpoint makes an HTTP PUT request to update user information at the specified URL. The request payload includes a JSON object with the "username" field.

### Response

The response returns a status code of 202 and the content type is "application/json". The response body contains a JSON object with the following schema:

``` json
{
  "type": "object",
  "properties": {
    "data": {
      "type": "object",
      "properties": {
        "_id": {
          "type": "string"
        },
        "username": {
          "type": "string"
        },
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "is_admin": {
          "type": "boolean"
        }
      }
    }
  }
}

 ```
### Method: PUT
>```
>http://127.0.0.1:8080/s-api/user/66b292404e836bf08ea018ee
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2NmIyOTI0MDRlODM2YmYwOGVhMDE4ZWUiLCJlbWFpbCI6ImVtaWFzQGdtYWlsLmNvbSIsImlzX2FkbWluIjpmYWxzZX0.QAwmgWFvS6KSdcr_nonZPncjHXvuKXxaUNuDYmXfZrY|


### Body (**raw**)

```json
        {
            "username": "my man"
        }
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Delete a task
The API endpoint sends an HTTP DELETE request to [http://127.0.0.1:8080/s-api/task/:id](http://127.0.0.1:8080/s-api/task/:id) to delete a specific task. Upon successful deletion, the response returns a status code of 200 and a JSON object with a message property.

``` json
{
  "type": "object",
  "properties": {
    "message": {
      "type": "string"
    }
  }
}

 ```
### Method: DELETE
>```
>http://127.0.0.1:8080/s-api/task/66b5df6c0b65b192b8a1276b
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2NmIyOGZiMmQxY2YyYzMzNmUwODQ1ZjAiLCJlbWFpbCI6InJvb3RAZ21haWwuY29tIiwiaXNfYWRtaW4iOnRydWV9.Hd-uPoKnQ-49whLa0_nTYDfmS-Bb4hxNVipAZFT5T2s|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Delete a user
### Delete User

This endpoint is used to delete a specific user by their ID.

#### Request Body

This request does not require a request body.

#### Response

- Status: 202
    
- Content-Type: application/json
    
- {    "message": "deleted successfully"}
### Method: DELETE
>```
>http://127.0.0.1:8080/s-api/user/66b5da7549abb3c4e1d6483d
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2NmIyOGZiMmQxY2YyYzMzNmUwODQ1ZjAiLCJlbWFpbCI6InJvb3RAZ21haWwuY29tIiwiaXNfYWRtaW4iOnRydWV9.Hd-uPoKnQ-49whLa0_nTYDfmS-Bb4hxNVipAZFT5T2s|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
_________________________________________________
Powered By: [postman-to-markdown](https://github.com/bautistaj/postman-to-markdown/)
