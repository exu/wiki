# RSocket

## One way that RSocket differs from HTTP is that it defines four interaction models:

- Fire-and-Forget: an optimization of request/response that is useful when a response is not needed, such as for non-critical event logging.
- Request/Response: when you send one request and receive one response, exactly like HTTP. Even here though, the protocol has advantages over HTTP in that it is asynchronous and multiplexed.
- Request/Stream: analogous to Request/Response returning a collection, the collection is streamed back instead of querying until complete, so for example send a bank account number, respond with a real-time stream of account transactions.
- Channel: a bi-directional stream of messages allowing for arbitrary interaction models.


## Resources

- [Spring Boot 2.2.0 M6](https://docs.spring.io/spring-boot/docs/2.2.0.M6/reference/html/spring-boot-features.html#boot-features-rsocket-requester)
- [RSocket vs gRPC](https://medium.com/netifi/differences-between-grpc-and-rsocket-e736c954e60)
