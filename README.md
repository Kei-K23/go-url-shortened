# Go URL Shortener

This is a simple URL shortener API written in Go using http standard library and GORM for database interaction. It allows users to shorten long URLs into more manageable and shareable links.

## Features

- Shorten long URLs into easy-to-share links.
- Retrieve original URLs from shortened links.
- Redirect users from shortened links to original URLs.
- SQLite database for storing URL mappings.

## Prerequisites

Before running the API, make sure you have the following installed:

- Go (version latest version)
- SQLite (for database storage)

## Installation

1. Clone the repository to your local machine:

```bash
git clone https://github.com/Kei-K23/go-url-shortened.git
```

2. Navigate to the project directory:

```bash
cd go-url-shortened
```

3. Install dependencies:

```bash
go mod tidy
```

## Usage

1. Start the application:

```bash
go run main.go
```
