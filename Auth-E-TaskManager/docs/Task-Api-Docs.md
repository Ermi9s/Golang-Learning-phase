# Project: Task_manager-api
### Task Manager API

Task manager api allows you to:

- Create tasks
    
- Update tasks
    
- Get tasks
    
- Delete tasks

## End-point: http://127.0.0.1:8080/task/
### Create New Task

This endpoint allows the client to create a new task.

#### Request Body

- `title` (string, required): The title of the task.
    
- `description` (string, required): The description of the task.
    

#### Response

The response will be in JSON format with the following schema:

``` json
{
    "data": {
        "id": "9",
        "title": "api documentaion",
        "description": "documenting the taskmanager api",
        "date": "2024/08/01"
    }
}

 ```
### Method: POST
>```
>http://127.0.0.1:8080/task/
>```
### Body (**raw**)

```json
{
    "title" : "api documentaion",
    "description" : "documenting the taskmanager api"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: http://127.0.0.1:8080/task/{{id}}
This endpoint allows the user to update a specific task identified by the ID in the URL.

### Request Body

- The request should include a JSON payload with the following parameters:
    
    - `title` (string): The updated title of the task.
        
    - `description` (string): The updated description of the task.
        

### Response

The response will be in JSON format with the following schema:

``` json
{
    "data": {
        "id": "1",
        "title": "api documentaion update",
        "description": "documenting the taskmanager api",
        "date": "2024/08/01"
    }
}
 ```
### Method: PUT
>```
>http://127.0.0.1:8080/task/1
>```
### Body (**raw**)

```json
{
    "title" : "api documentaion update",
    "description" : "documenting the taskmanager api"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: http://127.0.0.1:8080/task
# Task Retrieval

This endpoint allows you to retrieve a list of tasks.

## Request

### Query Parameters

- No query parameters required.
    

### Request Body

This request does not require a request body.

## Response

The response will be in JSON format and will have the following schema:

``` json
{
    "data": [
        {
            "id": "",
            "title": "",
            "description": "",
            "date": ""
        }
    ]
}

 ```

The `data` array contains objects representing individual tasks, with each task having an `id`, `title`, `description`, and `date` field.
### Method: GET
>```
>http://127.0.0.1:8080/task
>```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: http://127.0.0.1:8080/task/{{ID}}
The `GET` request retrieves the details of a specific task identified by the ID `1`. The response of this request is documented below as a JSON schema.

``` json
{
    "data": {
        "id": "9",
        "title": "api documentaion",
        "description": "documenting the taskmanager api",
        "date": "2024/08/01"
    }
}

 ```
### Method: GET
>```
>http://127.0.0.1:8080/task/1
>```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: http://127.0.0.1:8080/task/{{ID}}
The API endpoint sends an HTTP DELETE request to [http://127.0.0.1:8080/task/9](http://127.0.0.1:8080/task/9) to delete a specific task. Upon successful deletion, the API returns a response with status code 202 and content type as application/json. The response body contains a JSON schema with a "message" key, indicating the outcome of the deletion process.

``` json
{
    "message": "Deleted successfully"
}

 ```
### Method: DELETE
>```
>http://127.0.0.1:8080/task/1
>```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
_________________________________________________
Powered By: [postman-to-markdown](https://github.com/bautistaj/postman-to-markdown/)
