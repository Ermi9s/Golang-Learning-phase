# Project: Auth-Task-Manager

## End-point: http://127.0.0.1:8080/register/
### Register User

This endpoint allows the client to register a new user.

#### Request Body

- `username` (string, required): The username of the user to be registered.
    
- `email` (string, required): The email address of the user to be registered.
    
- `password` (string, required): The password for the user account.
    

#### Response

The response is a JSON object with the following schema:

``` json
{
  "type": "object",
  "properties": {
    "data": {
      "type": "object",
      "properties": {
        "token": {"type": "string"},
        "user": {
          "type": "object",
          "properties": {
            "id": {"type": "string"},
            "username": {"type": "string"},
            "email": {"type": "string"},
            "password": {"type": "string"},
            "is_admin": {"type": "boolean"}
          }
        }
      }
    }
  }
}

 ```

#### Example Response

``` json
{
  "data": {
    "token": "********",
    "user": {
      "id": "********",
      "username": "********",
      "email": "********",
      "password": "********",
      "is_admin": true
    }
  }
}

 ```
### Method: POST
>```
>http://127.0.0.1:8080/register/
>```
### Body (**raw**)

```json
{
    "username":"Ermias",
    "email" : "emias@gmail.com",
    "password" : "1234"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: http://127.0.0.1:8080/log-in/
### Log In

This endpoint allows the user to log in by providing their username, email, and password.

#### Request Body

- `username` (string, required): The username of the user.
    
- `email` (string, required): The email address of the user.
    
- `password` (string, required): The password of the user.
    

#### Response

The response is in JSON format and includes the following fields:

``` json
{
  "data": {
    "token": "string",
    "user": {
      "id": "string",
      "username": "string",
      "email": "string",
      "is_admin": "boolean"
    }
  }
}

 ```
### Method: POST
>```
>http://127.0.0.1:8080/log-in/
>```
### Body (**raw**)

```json
{
    "username":"ermias",
    "email" : "E@gmail.com",
    "password" : "1234"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: http://127.0.0.1:8080/promote/{{id}}
### Update User Promotion

This endpoint is used to promote a user by their ID.

#### Request

- Method: PUT
    
- URL: `http://127.0.0.1:8080/promote/{{id}}`
    
- Path variable
    
    - `id` (string, required): The ID of the user to be promoted.
        

##### Request Body

The request body should be in JSON format and include the following parameters:

- `username` (string): The username of the user.
    
- `email` (string): The email address of the user.
    
- `password` (string): The password of the user.
    

#### Response

- Status: 202
    
- Content-Type: application/json
    

##### Response Body

The response will be in JSON format and includes the following schema:

``` json
{
  "data": {
      "id": "string",
      "username": "string",
      "email": "string",
      "password": "string",
      "is_admin": true
  }
}

 ```

The `data` object contains the promoted user's information, including their ID, username, email, password, and admin status. The `Message` field provides additional information about the response, such as an error message or status.
### Method: PUT
>```
>http://127.0.0.1:8080/promote/{{id}}
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|Bearer%20eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjY2YjI3OGI2ZTIxNGZjMTE5MjBiNzdjYiIsImVtYWlsIjoicm9vdEBnbWFpbC5jb20iLCJDbGFpbXMiOm51bGx9.ZkhjDLCAuAWp75_sWp4Je-Sphz9fvhc0vnRjI-6QfI0|


### Body (**raw**)

```json
{
    "username":"root",
    "email" : "root@gmail.com",
    "password" : "12345678"
}
```

### 🔑 Authentication noauth

|Param|value|Type|
|---|---|---|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: http://127.0.0.1:8080/user/{{id}}
### Method: PUT
>```
>http://127.0.0.1:8080/user/{{id}}
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2NmIyOGZiMmQxY2YyYzMzNmUwODQ1ZjAiLCJlbWFpbCI6InJvb3RAZ21haWwuY29tIn0.x1FCkXlGrS0zdAecF4mO3ilZFNE4NEMEc1NGcaBF40U|


### Body (**raw**)

```json
        {
            "id": "66b2909d5effab904550b878",
            "username": "updated ermias",
            "email": "update@gmail.com"

        }
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: http://127.0.0.1:8080/user/{{id}}
### Retrieve User Details

This endpoint retrieves the details of a specific user identified by their unique ID.

#### Request

- Method: GET
    
- URL: `http://127.0.0.1:8080/user/{{id}}`
    

#### Response

The response will be a JSON object with the following schema:

``` json
{
    "type": "object",
    "properties": {
        "data": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user": {
                    "type": "object",
                    "properties": {
                        "id": {
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
    }
}

 ```
### Method: GET
>```
>http://127.0.0.1:8080/user/{{id}}
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2NmIyOGZiMmQxY2YyYzMzNmUwODQ1ZjAiLCJlbWFpbCI6InJvb3RAZ21haWwuY29tIn0.x1FCkXlGrS0zdAecF4mO3ilZFNE4NEMEc1NGcaBF40U|


### Body (**raw**)

```json

```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: http://127.0.0.1:8080/user
### Method: GET
>```
>http://127.0.0.1:8080/user
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2NmIyOGZiMmQxY2YyYzMzNmUwODQ1ZjAiLCJlbWFpbCI6InJvb3RAZ21haWwuY29tIn0.x1FCkXlGrS0zdAecF4mO3ilZFNE4NEMEc1NGcaBF40U|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: http://127.0.0.1:8080/user/{{id}}
### Delete User

This endpoint is used to delete a specific user by their ID.

#### Request

- Method: DELETE
    
- URL: `http://127.0.0.1:8080/user/{{id}}`
    

#### Response

The response will be in JSON format with the following schema:

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

The response will have a status code of 202 (Accepted) upon successful deletion of the user.
### Method: DELETE
>```
>http://127.0.0.1:8080/user/{{id}}
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2NmIyOGZiMmQxY2YyYzMzNmUwODQ1ZjAiLCJlbWFpbCI6InJvb3RAZ21haWwuY29tIn0.x1FCkXlGrS0zdAecF4mO3ilZFNE4NEMEc1NGcaBF40U|


### Body (**raw**)

```json

```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
_________________________________________________
Powered By: [postman-to-markdown](https://github.com/bautistaj/postman-to-markdown/)
