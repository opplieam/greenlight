# Greenlight

Learning from Let's go further


| Method | URL Pattern         | Handler             | Actions                      |
|--------|---------------------|---------------------|------------------------------|
| Get    | /v1/healthcheck     | healthcheckHandler  | Show application information |
| Post   | /v1/movies          | createMovieHandler  | Create movie                 |
| Get    | /v1/movies          | listMoviesHandler   | Show details of all movies   |
| Get    | /v1/movies/:id      | showMovieHandler    | Show detail of movie :id     |
| Put    | /v1/movies/:id      | updateMovieHandler  | Update detail of movie :id   |
| Delete | /v1/movies/:id      | deleteMovieHandler  | Delete movie :id             |
| Post   | /v1/users           | registerUserHandler | Register a new user          |
| Put    | /v1/users/activated | activateUserHandler | Activate a specific user     |
