# Spring Flux resources

## Reactive streams

The crucial concept of Reactive Streams is processing infinite streams
of data in an asynchronous and non-blocking manner, so that the
computing resources (think CPU cores or network hosts) can be used in
parallel.

There are three key factors that make a stream reactive:
- the data is processed asynchronously,
- the backpressure mechanism is non-blocking,
- the fact that the downstream can be slower than the upstream is somehow represented in the domain model.

## Resources
- Spring web reactive: https://docs.spring.io/spring/docs/current/spring-framework-reference/web-reactive.html
- Reactive Streams in Java: https://blog.softwaremill.com/how-not-to-use-reactive-streams-in-java-9-7a39ea9c2cb3
- Reactive operators: http://reactivex.io/documentation/operators.html
- Reactive streams specification https://github.com/reactive-streams/reactive-streams-jvm
- Spring erbflux is based on https://projectreactor.io/
- Reactive SQL database driver  https://r2dbc.io/
- Testing Webflux https://docs.spring.io/spring/docs/current/spring-framework-reference/testing.html#webtestclient
- Visualization of example operators https://rxmarbles.com/#defaultIfEmpty
