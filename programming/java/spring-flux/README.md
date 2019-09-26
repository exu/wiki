# Reactive programming and Spring Flux resources [back to index](/)

## Reactive streams

The crucial concept of Reactive Streams is processing infinite streams
of data in an asynchronous and non-blocking manner, so that the
computing resources (think CPU cores or network hosts) can be used in
parallel.

There are three key factors that make a stream reactive:
- the data is processed asynchronously,
- the backpressure mechanism is non-blocking,
- the fact that the downstream can be slower than the upstream is somehow represented in the domain model.


Reactive programming is an asynchronous programming paradigm concerned with data streams and the propagation of change. This means that it becomes possible to express static (e.g. arrays) or dynamic (e.g. event emitters) data streams with ease via the employed programming language(s).
— https://en.wikipedia.org/wiki/Reactive_programming
An API for asynchronous programming with observable streams
— http://reactivex.io/

Well I am pretty sure that is not the quick intro you where looking for. I find a simpler, better definition the following :
In reactive programming, everything can be a stream of data.
— https://gist.github.com/staltz/868e7e9bc2a7b8c1f754
The device location, the system time, a list of entries from the DB/API, user clicks, everything can be a stream of data. Data from different streams can easily be combined and transformed, and in the end processed/observed by the subscribers.


## Project Reactor

### Flux

A Flux<T> is a standard Publisher<T> representing an asynchronous sequence of 0 to N emitted items, optionally terminated by either a completion signal or an error. As in the Reactive Streams spec, these 3 types of signal translate to calls to a downstream Subscriber’s onNext, onComplete or onError methods.


### Mono

A Mono<T> is a specialized Publisher<T> that emits at most one item and then optionally terminates with an onComplete signal or an onError signal.

It offers only a subset of the operators that are available for a Flux, and some operators (notably those that combine the Mono with another Publisher) switch to a Flux.
For example, Mono#concatWith(Publisher) returns a Flux while Mono#then(Mono) returns another Mono.




## Resources

### Articles

- [Spring web reactive](https://docs.spring.io/spring/docs/current/spring-framework-reference/web-reactive.html)
- [Reactive Streams in Java](https://blog.softwaremill.com/how-not-to-use-reactive-streams-in-java-9-7a39ea9c2cb3)
- [Reactive operators](http://reactivex.io/documentation/operators.html)
- [Reactive streams specification](https://github.com/reactive-streams/reactive-streams-jvm)
- [Spring webflux is based on](https://projectreactor.io/)
- [Reactive SQL database driver ](https://r2dbc.io/)
- [Testing Webflux](https://docs.spring.io/spring/docs/current/spring-framework-reference/testing.html#webtestclient)
- [Visualization of example operators](https://rxmarbles.com/#defaultIfEmpty)
- [vertx vs webflux](https://blog.rcode3.com/blog/vertx-vs-webflux/)


- [Why you should learn Reactive Programing](https://medium.com/corebuild-software/why-you-should-learn-reactive-programming-51b6ffc31425)
- [The introduction to Reactive Programming you've been missing](https://gist.github.com/staltz/868e7e9bc2a7b8c1f754)
- [Java Reactive Programming](https://www.scnsoft.com/blog/java-reactive-programming)
- [Performance](https://medium.com/@filia.aleks/microservice-performance-battle-spring-mvc-vs-webflux-80d39fd81bf0)
- [Backpressure](https://www.e4developer.com/2018/04/28/springs-webflux-reactor-parallelism-and-backpressure/)
- [Flux::create vs ::generate](https://stackoverflow.com/questions/49951060/difference-between-flux-create-and-flux-generate)

### Workshops / Coding excerctises

- [Head First Reactive With Spring And Reactor/](https://reactor.github.io/head-first-reactive-with-spring-and-reactor/)


### Tools

- [BlockHound - Blocking calls detector](https://github.com/reactor/BlockHound)
