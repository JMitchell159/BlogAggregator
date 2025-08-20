# BlogAggregator
Build a Blog Aggregator in Go BootDev Project

## Requirements:
- Postgres
- Go (for installation)

## Installation
In a Linux CLI, you just run the command `go install github.com/JMitchell159/gator`

## Additional Setup
### Installing Postgres
#### MacOS with brew
```
brew install postgresql@15
```
#### Linux/WSL (Debian)
```
sudo apt update
sudo apt install postgresql postgresql-contrib
```
#### Verification
`psql --version` should return a non-error

Additionally, on Linux/WSL (Debian), run this to update the password for postgres: `sudo passwd postgres`, afterwards, there will be a prompt to enter a password, enter anything that you will be able to remember.

### Creating the Database
Enter the `psql` shell:
- Mac - `psql postgres`
- Linux/WSL (Debian) - `sudo -u postgres psql` (you will be prompted to enter the root password)
You will be taken to a prompt that looks like this:
```
postgres=#
```
in this prompt, create the gator database by typing in the following: `CREATE DATABASE gator;`

then connect to the database with the following: `\c gator`

The prompt should change to this:
```
gator=#
```
On Linux/WSL (Debian), you will have to additionally enter this to set the **database user's** password:
```
ALTER USER postgres PASSWORD '<password>';
```
Where `<password>` is any password you want to enter.

### Creating the Tables
While still in the gator shell, enter the following commands, pressing enter after each one:
```
CREATE TABLE users(id UUID PRIMARY KEY, created_at TIMESTAMP NOT NULL, updated_at TIMESTAMP NOT NULL, name TEXT UNIQUE NOT NULL);

CREATE TABLE feeds(id UUID PRIMARY KEY, created_at TIMESTAMP NOT NULL, updated_at TIMESTAMP NOT NULL, name TEXT NOT NULL, url TEXT UNIQUE NOT NULL, user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE, last_fetched_at TIMESTAMP);

CREATE TABLE feed_follows(id UUID PRIMARY KEY, created_at TIMESTAMP NOT NULL, updated_at TIMESTAMP NOT NULL, user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE, feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE, UNIQUE(user_id, feed_id));

CREATE TABLE posts(id UUID PRIMARY KEY, created_at TIMESTAMP NOT NULL, updated_at TIMESTAMP NOT NULL, title TEXT UNIQUE, url TEXT UNIQUE NOT NULL, description TEXT, published_at TIMESTAMP NOT NULL, feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE);
```

### Config
Afterwards, in your home directory, make a file called ".gatorconfig.json" and enter the following into it
```(json)
{
  "db_url":"postgres://postgres:<database_user_password>@localhost:5432/gator?sslmode=disable"
}
```
Here is also a single command to do the same thing as above: `touch ~/.gatorconfig.json && echo "{\"db_url\":\"postgres://postgres:<database_user_password>@localhost:5432/gator?sslmode=disable\"}" > ~/.gatorconfig.json`
For macOS, replace the second postgres with your username and remove the third postgres, the url should look like this: `"postgres://<user_name>:@localhost:5432/gator"`

### Installing Go
Instructions for installing Go can be found [here](https://go.dev/doc/install).

## Uses
- gator register <user_name>: registers a new user with the username (throws an error if given a duplicate name) and sets it to the current user
- gator login <user_name>: sets the current user to the one with the user_name, throws an error if user_name does not exist
- gator users: lists all of the users and their data
- gator addfeed <feed_name> <feed_url>: adds a feed to the feeds table associated with the current user and automatically makes the user follow the feed
- gator follow <feed_name>: adds a record to the feed_follows table relating the feed with the current user
- gator unfollow <feed_name>: deletes a record from the feed_follows table that relates the feed with the current user
- gator following: lists all of the feeds and their data that the current user is following
- gator feeds: lists all of the feeds in the database, along with the user that added them
- gator agg <time_interval>: inserts all items as posts from feeds that the current user is following into the posts table. Feeds are processed in a loop at the start of each interval, always prioritizing either the one that has not been checked in a while or the one with the earliest publish date that has not been checked
- gator browse [limit]: lists (limit) of the most recent posts by time published, limit defaults to 2
- gator reset: Resets the entire database (USE WITH EXTREME CAUTION)
