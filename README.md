# Nyaa-X-Discord

Discord bot that will notify you via DM if there is new [Nyaa](https://nyaa.si) update according to your filter/query.

This bot is created because my favorite anime fansubs group has disbanded and I want to keep up-to-date with new anime episodes in one place.

## Requirement

- [Discord bot](https://discordpy.readthedocs.io/en/latest/discord.html) and its token
- [MongoDB](https://www.mongodb.com/)
- [Go](https://golang.org/)
- [Docker](https://docker.com) + [Docker compose](https://docs.docker.com/compose/) (optional)

## Steps

1. Git clone this repo.
    ```bash
    git clone github.com/rl404/nyaa-x-discord
    ```
2. Rename `sample.env` to `.env` and modify according to your configuration.
3. Run.
    ```bash
    go run .
    # or
    docker-compose up
    ```
4. Invite the bot to your server.
5. DM the bot `!ping`.
6. You will get welcome message from the bot.
    ```
    Welcome rl404
    It looks like this is your first time using this bot.

    This bot will help you keeping track of Nyaa update according to your   query/filter.

    How to Start
    1. Set filter.
    2. Set category.
    3. Set query.
    4. Turn on subscription.
    5. Wait a bit and I will notify you if there is a new update.

    !help to see all command list.
    ```
6. Have fun.

**note: you and your bot must have at least 1 common server so you and the bot can DM each other.*

## Bot Commands

```
!help
Show all command list.

!filter
Get filter names and their id.

!filter set <filter_id>
Set filter type for your query. filter_id is from !filter.

!category
Get category names and their id.

!category set <category_id>
Set category type for your query. category_id is from !category.

!query
Get your query string list and their id.

!query add <string> [string...]
Add your query string. You can have more than 1 query. For example:

    Contain naruto
    !query add naruto

    Contain one, piece and horriblesubs
    !query add one piece horriblesubs

    Contain bleach but not 720p
    !query add bleach -720p


!query delete <query_id> [query_id...]
Delete 1 or more your query strings. query_id is from !query.

!subscribe
Get your subscription status.

!subscribe <on|off>
Turn on or off bot subscription.
```

## License

MIT License

Copyright (c) 2020 Axel
