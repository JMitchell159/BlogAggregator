# BlogAggregator
Build a Blog Aggregator in Go BootDev Project

## Requirements:
- Postgres
- Go (for installation)

## Installation
In a Linux CLI, you just run the command `go install github.com/JMitchell159/gator`

## Additional Setup
In your home directory, make a file called ".gatorconfig.json" and enter the following into it
```(json)
{
  "db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
}
```

## Uses
- gator register <user_name>: registers a new user with the username (throws an error if given a duplicate name) and sets it to the current user
- gator login <user_name>: sets the current user to the one with the user_name, throws an error if user_name does not exist
- gator users: lists all of the users and their data
- gator addfeed <feed_name> <feed_url>: adds a feed to the feeds table associated with the current user and automatically makes the user follow the feed
- gator follow <feed_name>: adds a record to the feed_follows table relating the feed with the current user
- gator unfollow <feed_name>: deletes a record from the feed_follows table that relates the feed with the current user
- gator following: lists all of the feeds and their data that the current user is following
- gator agg <interval>: aggregates posts from the current users followed feeds and inserts them into the posts table, one feed at the start of every interval
- gator browse [limit]: lists (limit) of the most recent posts by time published, limit defaults to 2
