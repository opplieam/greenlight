# Greenlight


| Method | URL Pattern     | Handler            | Actions                      |
|--------|-----------------|--------------------|------------------------------|
| Get    | /v1/healthcheck | healthcheckHandler | Show application information |
| Post   | /v1/movies      | createMovieHandler | Create movie                 |
| Get    | /v1/movies/:id  | showMovieHandler | Show detail of movie :id     |
