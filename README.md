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
- gator agg <interval>: inserts all items as posts from feeds that the current user is following into the posts table, feeds are processed in a loop at the start of each interval, always prioritizing either the one that has not been checked in a while or the one with the earliest publish date that has not been checked
- gator browse [limit]: lists (limit) of the most recent posts by time published, limit defaults to 2
- gator reset: Resets the entire database (USE WITH EXTREME CAUTION)
