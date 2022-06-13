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
A dockerfile is provided to assist building and running this bot without needing to download and installing go.

```bash
docker build -t bittersweet-bot .
docker run -e BOT_TOKEN="your-bot-token" --network="host" bittersweet-bot
```