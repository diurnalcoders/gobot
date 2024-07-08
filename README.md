# Telegram Bot Template

This project implements a Telegram bot using Go, following clean architecture principles. The bot is designed to be modular, maintainable, and easily extensible.

## Project Structure

```bash
telegram-bot/
├── cmd/
│   └── bot/
│       └── main.go
├── internal/
│   ├── domain/
│   │   └── message.go
│   ├── usecase/
│   │   └── message_usecase.go
│   ├── repository/
│   │   └── message_repository.go
│   └── delivery/
│       └── telegram/
│           └── handler.go
├── pkg/
│   └── telegram/
│       └── client.go
├── config/
│   └── config.go
└── go.mod
```

- `cmd/bot/main.go`: Entry point of the application.
- `internal/`: Contains the core application code.
  - `domain/`: Defines core business logic and entities.
  - `usecase/`: Implements application-specific business rules.
  - `repository/`: Handles data persistence and retrieval.
  - `delivery/`: Manages interaction with external systems (Telegram).
- `pkg/`: Contains reusable packages.
- `config/`: Handles application configuration.

## Prerequisites

- Go 1.16 or higher
- A Telegram bot token (obtainable from BotFather)

## Setup

1. Clone the repository:

```bash
git clone https://github.com/yourusername/telegram-bot.git
cd telegram-bot
```

2. Initialize the Go module:

```bash
go mod init telegram-bot
```

3. Install dependencies:

```bash
go get github.com/go-telegram-bot-api/telegram-bot-api/v5
```

4. Set up your Telegram bot token as an environment variable:

```bash
export TELEGRAM_BOT_TOKEN=your_bot_token_here
```

## Configuration

The application uses environment variables for configuration. Make sure to set the following:

- `TELEGRAM_BOT_TOKEN`: Your Telegram bot token

## Building the Project

To build the project, run:

```bash
go build ./cmd/bot
```

This will create an executable named `bot` in your current directory.

## Running the Bot

To start the bot, simply run the built executable:

```bash
./bot
```

The bot will now be active and responding to messages on Telegram.

## Customizing Bot Behavior

To customize how your bot responds to messages, modify the `ProcessMessage` function in `internal/usecase/message_usecase.go`. This is where you can implement your bot's logic.

## Adding Features

To add new features to your bot:

1. Create new use cases in the `usecase` package.
2. Implement new repository methods if you need to store or retrieve data.
3. Update the `Handler` in `delivery/telegram/handler.go` to use new use cases.
4. Modify `main.go` to wire up new components.
