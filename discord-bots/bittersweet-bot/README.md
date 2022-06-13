# Purpose

The bittersweet-bot interacts with a Discord channel. The Bot will respond to the caller with a complement or insult depending on the command interpreted.

```sh
!hello
!insult
!compliment
!surprise
!help
```

# Dockerfile
A Dockerfile is provided to assist building and running this bot without needing to download and installing go.

To build the bot yourself use the following:
```bash
docker build -t bittersweet-bot .
```

The [bittersweet-bot](https://github.com/edgehew/GoProjects/pkgs/container/goprojects%2Fbittersweet-bot) is available as a package so you do not need to build it yourself. You can just pull it down with the following if you allow unsigned repositories, otherwise you can build it yourself by cloning the repo.

```bash
docker pull ghcr.io/edgehew/goprojects/bittersweet-bot:latest
docker run -e BOT_TOKEN="your-bot-token" --network="host" bittersweet-bot
```