# Gator

This is a cli blog aggregator written in Go (second project in go)

It supports multiple users, each with their own feeds

In this project, I learned:

-   Working postgreSQL in Go
-   Using [sqlc](https://sqlc.dev/) to generate type-safe Go code from SQL queries
-   Managing database migrations with [goose](https://github.com/pressly/goose)

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

### 3. Build the Project (optional)

```bash
go build
```

or

```bash
go install
```

This compile the code into a binary file

### 4. Set up the config file and database connection string

create a config file in your home directory called "~/.gatorconfig.json" with the following content

```JSON
{
  "db_url": "postgres://example", // database connection string
  "current_user_name": null
}
```

this will be use to keep track of the current user, since this project can have multiple users

### 5. Run the project

if you compile the project you can just simply run the binary file

```bash
./gator <command> <arg>
```

Or, run it directly without building:

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
