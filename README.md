# barkanekbot-go
BarkAnekBot is a bot for Telegram, which broadcast jokes

## Build and run
Firstly, setup environment variables:
- `BOT_TOKEN`: Telegram Bot API token

Then build project and run barkanekbot-go binary!

If you use docker - you can execute this command instead of building project manually:
```bash
$ docker build -t barkanekbot .
$ docker run -d -e BOT_TOKEN barkanekbot
```

Or using docker-compose:
```bash
$ docker-compose up -d
```
