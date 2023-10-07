const WebSocket = require("ws");

const socket = new WebSocket("ws://localhost:8080/ws");

socket.onopen = (event) => {
  console.log("WebSocket connected");
  // Send a message to the server, if needed
  process.stdin.on("data", (data) => {
    const input = data.toString().trim();
    socket.send(input);
    if (input === "exit") {
      console.log("Exiting the WebSocket client.");
      socket.close();
      process.exit(0);
    }
  });
};

socket.onmessage = (event) => {
  const message = event.data;

  // Show message and timestamp
  console.log(`(${new Date().toLocaleTimeString()}) ${message}`);
};

socket.onclose = (event) => {
  console.log("Closing connection...");
};

// Keep the script running indefinitely (you can implement an exit condition)
process.stdin.resume();
