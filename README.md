# Word of Wisdom TCP Server with Proof of Work Protection

This project is a Word of Wisdom TCP server written in Go, designed to deliver wisdom quotes to clients. The server is protected against potential DDoS attacks using a Proof of Work (PoW) mechanism with a challenge-response protocol.

# Project Overview

The server listens for incoming TCP connections, where each client must solve a PoW challenge before receiving a response. Once the PoW is successfully validated, the server responds with a quote from a collection of wisdom sayings. This setup helps mitigate DDoS attacks by ensuring clients invest computational effort before each interaction.
For now it supports only one challenge type: hash inversion.

### How It Works

	1. Challenge Generation: Upon connection, the server generates two random strings: 'start' and 'prefix'. 'prefix' is a hexadecimal string that serves as a PoW requirement.
	2. Sending the Challenge: The server sends the challenge to the client.
    3. Solving the PoW: The client receives the challenge and must find the solution string that, when concatenated with the 'start', produces a SHA-256 hash that in hexadecimal representation starts with a given 'prefix'.
	4. Sending the Solution: The client sends the solution back to the server.
    5. Validating the Solution: The server validates the solution by checking if the hexadecimal representation of SHA-256 hash of 'start' + 'solution'. And comparing it with the 'prefix'.
    6. Responding with Wisdom: If the solution is correct, the server sends a random wisdom quote to the client, otherwise, it returns the error.

# Setup and Running the Project

### Prerequisites

	1. Docker: Ensure Docker is installed and running.
	2. Go: Required for local testing and development.

### Running with Docker

1. **Create Docker Network**

   ```bash
   make start-network
   ```
   This command creates a Docker network called `wisdom-net`, allowing the server and client containers to communicate with each other.

2. **Build Docker Image for the Server**

   ```bash
   make build-server-docker-image
   ```
   This command builds the Docker image for the server based on `Dockerfile-server`. The image will be named `wisdom-server-image`.

3. **Build Docker Image for the Client**

   ```bash
   make build-client-docker-image
   ```
   This command builds the Docker image for the client based on `Dockerfile-client`. The image will be named `wisdom-client-image`.

4. **Start the Server**

   ```bash
   make start-server
   ```
   This command runs the server container named `wisdom-server` on port 8080 and connects it to the `wisdom-net` network.

5. **Start the Client**

   ```bash
   make start-client
   ```
   This command runs the client container named `wisdom-client` and connects it to the `wisdom-net` network. The client will attempt to connect to the server running on port 8080.

6. **Stop the Server**

   ```bash
   make stop-server
   ```
   This command stops the running server container, `wisdom-server`.

7. **Stop the Client**

   ```bash
   make stop-client
   ```
   This command stops the running client container, `wisdom-client`.

## Example Usage

1. Create the network:

   ```bash
   make start-network
   ```

2. Build the server and client images:

   ```bash
   make build-server-docker-image
   make build-client-docker-image
   ```

3. Start the server and client in different terminals:

   ```bash
   make start-server
   make start-client
   ```

4. To stop the server and client:

   ```bash
   make stop-server
   make stop-client
   ```

