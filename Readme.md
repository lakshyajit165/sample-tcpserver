## How servers handle multiple connections?

-   TCP is the most reliable way for 2 machines to talk to each other over the network.
-   TCP Server: It's a simple process that runs on a machine and listens to a 'port'.
-   Any machine which wants to talk to our server, has to connect over the port and establish the connection.

### Step1: Start listening on the port.

-   When your process starts, pick a port and start listening to it.

### Step2: Wait for a client to connect.

-   Invoke the 'accept' system call and wait for a client to connect.
-   This is a blocking system call, and server would not proceed until some client connects.

### Step3: Read the request and send the response.

-   Once the connection is established,
-   Invoke the 'read' system call to read the request(blocking).
-   Invoke the 'write' system call to send the response(blocking).
-   Close the connection.

### Step4: Do this over and over again.

-   Put this thing inside an infinite for loop...
-   Continuously waiting for client to connect.
-   reading the request.
-   writing the response.
-   closing the connection.

### Step5: Parallel processing.

-   Once client connects, fork a thread to process and respond.
-   let 'main' thread come back to accept as quickly as possible.

### Scope of improvements.

> Spawing a new thread for every other client request will soon lead to too many threads getting created which is inefficient. Hence we can do the following optimizations.

-   Limiting the number of threads.
-   Add thread pool to save on thread creation time.
-   Connection timeout(say if a client is connected, but it doesn't send a request for a long time, we can't infinitely wait for it to send a request, because that takes up a tcp connection)
-   TCP backlog queue configuration(this is an OS level setting which determines, how many connections we would like to keep in our TCP backlog)

Reference: https://www.youtube.com/watch?v=f9gUFy-9uCM
