# gator

Boot.dev guided project for an RSS aggregator.

## Requirements

You need access to a Postgres server.
You need Go installed on your computer.

## Installation

Clone this repo and then run `go install` inside of it.

## Usage

### Configurations

To get started you need to make a config for it, add this to .gatorconfig.json in your home-directory:

    {
      "db_url": "http://<username>:<password>@<hostname>:><port>"
    }

Where username and password are the credentials to use to log into Postgres. Depending on how the account is setup password can be ommited.
Hostname is name of the computer that Postgres is running on, if it's the same one as you run Gator on then you can use "localhost".
Port is the port that Postgres listens on, default is 5432.

### Commands

| Command  | Parameters | Description                                                                                       |
|----------|------------|---------------------------------------------------------------------------------------------------|
| login    | username   | switch to another user                                                                            |
| register | username   | register/create a new user                                                                        |
| agg      | timespan   | Continously loop over and read the different feeds, pausing given timespan between each operation |
| addfeed  | name url   | Add a new feed given the name and URL                                                             |
| follow   | url        | Make the currently logged in user follow a feed by URL                                            |
| unfollow | url        | Make the currently logged in user unfollow a feed by URL                                          |
| browse   | amount     | Retrieve the amount of posts, if none given it default to 2                                       |
