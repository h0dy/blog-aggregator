# gator

This is a cli blog aggregator written in Go (second project in go)
it supports multiple users and each user has it's own feeds

In this project, I learned:

-   How to work with postgreSQL in Go
-   How to work with [sqlc](https://sqlc.dev/) (tool to compile SQL to type-safe code)
-   How to work with [goose](https://github.com/pressly/goose) (tool to migrate database and manage schema)
-   How to implement caching

---

## How to Run This Project Locally

Make sure you have [**Go**](https://go.dev/doc/install) and [**PostgreSQL**](https://learn.microsoft.com/en-us/windows/wsl/tutorials/wsl-database#install-postgresql) installed on and your machine .

### 1. Clone the repository

```bash
git clone https://github.com/h0dy/blog-aggregator.git
cd blog-aggregator
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Set up the config file and database connection string

create a config file in your home directory called "~/.gatorconfig.json" with the following content

```JSON
{
  "db_url": "postgres://example", // database connection string
  "current_user_name": null
}
```

this will be use to keep track of the current user, since this project can have multiple users

### 4. Run the project

```bash
go run . <command> <arg>
```

| command   | argument\s                       | description                                                                   |
| --------- | -------------------------------- | ----------------------------------------------------------------------------- |
| register  | username                         | to register to gator                                                          |
| login     | username                         | to login to gator (you have to be registered)                                 |
| reset     | **None**                         | to clear the database (good for testing)                                      |
| users     | **None**                         | list all the users                                                            |
| agg       | time between requests (e.g. 30s) | time between each request to aggregate posts and save them                    |
| addfeed   | feed-name feed-url               | to create a new feed in gator (you have to loggedIn)                          |
| feeds     | **None**                         | list all the feeds that exists                                                |
| follow    | feed-url                         | to follow a feed if it exists (you have to loggedIn)                          |
| following | **None**                         | list all the feeds that the current user follows                              |
| unfollow  | feed-url                         | to unfollow a feed (you have to loggedIn)                                     |
| browse    | posts-limit (optional)           | lists all the posts that have been aggregated (make sure aggregate the posts) |
