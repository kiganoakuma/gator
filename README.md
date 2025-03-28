# Gator - RSS Feed Aggregator CLI

Gator is a command-line tool for aggregating, managing, and browsing RSS feeds. This document provides information about available commands and their usage.

## Available Commands

### User Management

| Command | Description | Usage |
|---------|-------------|-------|
| `register` | Create a new user account | `gator register <username> <password>` |
| `login` | Login to an existing account | `gator login <username> <password>` |
| `reset` | Reset your password | `gator reset <username> <new_password>` |
| `users` | List all registered users | `gator users` |

### Feed Management

| Command | Description | Usage |
|---------|-------------|-------|
| `addfeed` | Add a new RSS feed to the system | `gator addfeed <feed_url> <feed_name>` |
| `feeds` | List all available feeds | `gator feeds` |
| `follow` | Follow a feed to receive updates | `gator follow <feed_id>` |
| `following` | List all feeds you're following | `gator following` |
| `unfollow` | Unfollow a feed | `gator unfollow <feed_id>` |

### Content Browsing

| Command | Description | Usage |
|---------|-------------|-------|
| `browse` | Browse posts from followed feeds | `gator browse [limit]` (default: 2) |

### System Commands

| Command | Description | Usage |
|---------|-------------|-------|
| `agg` | Start the feed aggregator service | `gator agg <time_between_requests>` (e.g., `5m`, `30s`) |

## Examples

### Getting Started

1. Register a new account:
   ```
   gator register johndoe password123
   ```

2. Login to your account:
   ```
   gator login johndoe password123
   ```

3. Add a feed:
   ```
   gator addfeed https://news.example.com/rss.xml "Example News"
   ```

4. Follow the feed:
   ```
   gator follow 123e4567-e89b-12d3-a456-426614174000
   ```

5. Start aggregating feeds (updates every 5 minutes):
   ```
   gator agg 5m
   ```

6. Browse posts from followed feeds:
   ```
   gator browse 5
   ```

## Notes

- You must be logged in to use commands that require authentication (`addfeed`, `follow`, `following`, `unfollow`, `browse`).
- The `agg` command runs continuously until interrupted (Ctrl+C).
- Time durations for the `agg` command can be specified in seconds (`s`), minutes (`m`), or hours (`h`).
- When browsing posts, you can specify how many posts to display at once.

## Configuration

Gator uses a configuration file to store database connection details and user session information. This is automatically managed by the application.
