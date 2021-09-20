# Go RESTful API

## API Standards

* [RESTful URLs](https://github.com/ujklm23/restful_api#restful-urls)
* [HTTP Verbs](https://github.com/ujklm23/restful_api#http-verbs)
* [Error Handling](https://github.com/ujklm23/restful_api#error-handling)
* [Versions](https://github.com/ujklm23/restful_api#versions)
* [Request & Response Examples](https://github.com/ujklm23/restful_api#request--response-examples)
* [Architecture](https://github.com/ujklm23/restful_api#architecture)

### RESTful URLs

___

#### Good URL Example

* List of notes:
    * Get http://localhost:8080/api/v1/notes?apikey=<<apikey>>

#### Bad URL example

* Non-plural noun:
    * http://localhost:8080/api/v1/note?apikey=<<apikey>>

### HTTP Verbs

___

HTTP verbs, or methods, should be used in compliance with their definitions under the HTTP/1.1 standard. The action
taken on the representation will be contextual to the media type being worked on and its current state. Here's an
example of how HTTP verbs map to create, read, update, delete operations in a particular context:

| HTTP METHOD | GET            | POST       | PUT         | DELETE |
| ----------- | --------------- | --------- | ----------- | ------ |
| CRUD OP     | READ          | Create      | UPDATE      | DELETE |
| /notes      | Get all notes | Create new notes | |  |
| /notes/[id] | Get note by id | | If exists, update Note; If not, error | Delete note by id |

### Error Handling

Error responses should include a common HTTP status code, message for the developer, message for the end-user (when
appropriate), internal error code (corresponding to some specific internally determined ID), links where developers can
find more info. For example:

```json
{
  "status_code": 400,
  "status": "BAD REQUEST",
  "data": "Key: 'NoteCreateRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
}
```

* 200 - OK
* 400 - Bad Request
* 401 - Unauthorized
* 500 - Internal Server Error

___

### Versions

* Never release an API without a version number
* Versions should be integers, not decimal numbers, prefixed with ‘v’. For example
    * Good: v1, v2, v3
    * Bad: v-1.1, v1.2, 1.3

___

### Request & Response Examples

```API Key = nyeri-sendi```

#### API Resources

* [GET /notes?apikey=`<<apikey>>`](https://github.com/ujklm23/restful_api#get-notes)
* [GET /notes/[id]?apikey=`<<apikey>>`](https://github.com/ujklm23/restful_api#get-notesid)
* [POST /notes?apikey=`<<apikey>>`](https://github.com/ujklm23/restful_api#post-notes)
* [PUT /notes/[id]?apikey=`<<apikey>>`](https://github.com/ujklm23/restful_api#put-notesid)
* [DELETE /notes/[id]?apikey=`<<apikey>>`](https://github.com/ujklm23/restful_api#delete-notesid)

#### GET /notes

Example http://localhost:8080/api/v1/notes

Response body:

```json

{
  "status_code": 200,
  "status": "OK",
  "data": [
    {
      "id": 4,
      "name": "Rudy's birthday",
      "content": "August 10th, I'm going to make her..."
    },
    {
      "id": 5,
      "name": "The Pirate Bad",
      "content": "Nothing is Stolen, Everything is Shared"
    },
    {
      "id": 6,
      "name": "lorem ipsum",
      "content": "dolor sit amet, consectetur adipiscing elit."
    },
    {
      "id": 7,
      "name": "SSG",
      "content": "Server Side Generator"
    }
  ]
}
```

### GET /notes/[id]

Example: http://localhost:8080/api/v1/notes/5

Response body:

```json
{
  "status_code": 200,
  "status": "OK",
  "data": {
    "id": 5,
    "name": "The Pirate Bad",
    "content": "Nothing is Stolen, Everything is Shared"
  }
}
```

### POST /notes

Example: POST - http://localhost:8080/api/v1/notes

Request body:

```json
{
  "name": "I Wan't to",
  "content": "You're on my mind Been there all night I've been missing seeing my midnight queen Come have a drink"
}
```

Response body:

```json
{
  "status_code": 200,
  "status": "OK",
  "data": {
    "id": 90,
    "name": "I Wan't to",
    "content": "You're on my mind Been there all night I've been missing seeing my midnight queen Come have a drink"
  }
}
```

### PUT /notes/[id]

Example: PUT - http://localhost:8080/api/v1/notes/7

Request Body

```json
{
  "name": "Hello World",
  "content": "print('Hello World')"
}
```

Response Body

```json
{
  "status_code": 200,
  "status": "OK",
  "data": {
    "id": 7,
    "name": "Hello World",
    "content": "print('Hello World')"
  }
}
```

### DELETE /notes/[id]

Example: DELETE - http://localhost:8080/api/v1/notes/4

Response Body

```json
{
  "status_code": 200,
  "status": "OK",
  "data": "Success Delete id: 4"
}
```

### Architecture

Controller -> Service -> Repository