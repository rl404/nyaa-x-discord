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
6. Have fun.

**note: you and your bot must have at least 1 common server so you and the bot can DM each other.*

## License

MIT License

Copyright (c) 2020 Axel
