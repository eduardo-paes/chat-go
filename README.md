# Chat Application

A real-time chat application built with Go for the backend server and JavaScript for the client-side.

## Table of Contents

- [Chat Application](#chat-application)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Prerequisites](#prerequisites)
  - [Getting Started](#getting-started)
    - [Installation](#installation)
    - [Running the Application](#running-the-application)
  - [Usage](#usage)
  - [Architecture](#architecture)
  - [Contributing](#contributing)
  - [License](#license)

## Features

- Real-time chat functionality using WebSockets.
- User registration and login.
- Sending and receiving messages.
- WebSocket server built with Go (Golang).
- Client-side interface using JavaScript.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- [Go](https://golang.org/) installed.
- [Node.js](https://nodejs.org/) and [npm](https://www.npmjs.com/) installed.
- Knowledge of Go and JavaScript programming.

## Getting Started

### Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/eduardo-paes/chat-go.git
   cd chat-go/src
   ```

2. Install the required Go dependencies:

   ```bash
   go mod download
   ```

3. Install the required JavaScript dependencies for the client:

   ```bash
   cd web
   npm install
   ```

### Running the Application

1. Start the Go WebSocket server:

   ```bash
   go run main.go
   ```

2. Start the client (JavaScript):

   ```bash
   cd web
   npm start
   ```

The application should now be running, and you can access it in your web browser at `http://localhost:8080`.

## Usage

- Open your web browser and navigate to `http://localhost:8080`.
- Register or log in to access the chat.
- Start sending and receiving real-time messages with other users.

## Architecture

The application follows a client-server architecture:

- The backend server is built with Go and uses the "gorilla/websocket" package for WebSocket communication.
- The frontend client is implemented in JavaScript, using the WebSocket API to communicate with the server.

## Contributing

Contributions are welcome! If you'd like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your changes to your fork.
5. Submit a pull request to the main repository.

Please ensure that your code follows the project's coding standards and includes appropriate documentation.

## License

This project is licensed under the [MIT License](LICENSE).
