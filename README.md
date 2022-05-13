# Greenlight

*Greenlight* is an API for retrieving and managing information about movies.

Ultimately, *Greenlight* API will support the following endpoints and actions:

| Method | URL Pattern | Action |
| ------ | ----------- | ------ |
| GET    | /v1/healthcheck | Show application health and version information |
| GET    | /v1/movies | Show the details of all movies |
| POST   | /v1/movies | Create a new movie |
| GET    | /v1/movies/:id | Show the details of a specific movie |
| PATCH  | /v1/movies/:id | Update the details of a specific movie |
| DELETE | /v1/movies/:id | Delete a specific movie |
| POST | /v1/users | Register a new user |
| PUT | /v1/users/activated | Activate a specific user |
| PUT | /v1/users/password | Update a password for a specific user |
| POST | /v1/users/authentication | Generate a new authentication token |
| POST | /v1/users/password-reset | Generate a new password-reset token |
| GET | /debug/vars | Display application metrics |

For example,
```
$ curl -H "Authorization: Bearer RIDBIAE3AMMK57T6IAEBUGA7ZQ" localhost:4000/v1/movies/1

{
  "movie": {
  "id": 1,
  "title": "Moana",
  "year": 2016,
  "runtime": "107 mins",
  "genres": [
    "animation",
    "adventure"
  ],
  "version": 1
  }
}
```
