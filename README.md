# plugin-engine

The microservice architecture app plugin engine having exposed rest endpoint `/fetchData` to handle the requests to
fetch data related to virtual devices.

## Communication research

### Here requirement: communication between Go microservice and Java microservice

Some common ways to achieve communication between Go and Java microservices:

1. RESTful API: One of the most popular and language-agnostic ways of communication is through RESTful APIs. Both Go and
Java can easily implement and consume REST APIs, making it a flexible and widely used approach. You can use standard
HTTP methods (GET, POST, PUT, DELETE) to interact with each microservice's API endpoints, passing data in JSON or other
commonly used formats.

2. gRPC: gRPC is a high-performance RPC (Remote Procedure Call) framework developed by Google. It supports bi-directional
streaming and can be an efficient choice for communication between Go and Java microservices. gRPC uses Protocol
Buffers (protobufs) to define service contracts, which provides a more efficient data serialization compared to JSON.

3. Message Brokers: Using a message broker like Apache Kafka or RabbitMQ can be beneficial for asynchronous communication
between microservices. Go and Java can both publish and consume messages from the broker, enabling decoupling and
scalability of services. Events or messages can be used to communicate between microservices.

4. WebSockets: If you need real-time communication and bidirectional data exchange, WebSockets can be a good choice. Both
Go and Java have WebSocket libraries available, allowing you to create persistent connections for real-time updates.

5. gRPC-Web: If you're already using gRPC in your microservices and want to interact with the browser-based front-end, you
can consider using gRPC-Web, which allows you to call gRPC services directly from web browsers.

6. Apache Thrift: Apache Thrift is another RPC framework that supports multiple programming languages. It allows you to
define services and data structures in a platform-neutral interface definition language (IDL) and generate code for both
Go and Java to enable communication.

7. GraphQL: GraphQL is a query language for APIs that allows clients to request exactly the data they need. It can be a
good choice if you want more fine-grained control over the data exchanged between microservices.

In choosing the best communication method, consider factors such as performance, complexity, maintainability, and
compatibility with existing systems. We are using RESTful APIs for the communication between both the microservice for
simplicity.

### Details of the plugin-engine

- Run main function from main.go file to start the application.
- For local running use URL:- http://localhost:8443/fetchData?virtualDevice={virtualDevice}.
- Supported virtual devices are `firewall`, `linuxServer`, `router`, `switch`, `loadBalancer`, `windowsServer` can be
  updated into the URL.
- Application constant data for first request and randomly generated data for all other subsequent requests.

### Unit tests
- handler package unit tests written by covering most of the scenarios with code coverage more than 90%.
- sample unit tests written for plugin package and firewall package and expect similar unit tests for scenarios and code
coverage.