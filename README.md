# Bountie Technical Test
Tecnical Test for Bountie

## Prerequisites
* Docker
* Docker Compose
* Go

## Notes
Link for Longest Palindromic String Algorithm: [Via Goplay.Space](https://goplay.space/#4CieqYKd4h9)

## Setup and Testing
Do `docker-compose up` in the backend folder.

There are two endpoints that is exposed at `http://localhost:8080/`
* `POST /decimal-binary`:
```json
// Request Body:
{
  "decimal": 3
}
```
```json
// Response:
{
  "is_success": true,
  "status": 200,
  "error": null,
  "data": "1001"
}
```
* `POST /longest-palindromic-string`:
```json
// Request Body:
{
  "query": "di rumah saya ada kasur rusak"
}
```
```json
// Response
{
  "is_success": true,
  "status": 200,
  "error": null,
  "data": "kasur rusak"
}

```
