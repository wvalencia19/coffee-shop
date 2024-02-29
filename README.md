# Coffee shop

Simulating a coffee shop. A coffee shop usually has several moving pieces: grinders, brewers, different types of beans, and even the number of baristas themselves (plus how fast they can brew coffee- which itself is a function of how fast the grinders + brewers work). What is being given to you is a basic representation (that doesn't work too well in its current form) of a coffee shop, and we want to fix it up so that it behaves more like a real coffee shop.


## Execution

```
make run
```

The file `data/config.json` serves as configuration file for some aspects of the system, also as an input for orders, you can modify to add/remove grinders and brewers and place orders.

## Running tests

```
make test
```


## Key Design Decisions Explained

I believe the main decision was related to finding a way to manage grinders and brewers in a concurrent system, taking into account their availability for use. To achieve this, I followed the principle of "sharing memory by communicating" using channel pooling. This approach offers the following advantages:

* Concurrency Safety: Channels in Go inherently ensure safe access to shared resources among concurrent goroutines without the explicit lock management required by mutexes.
  
* Resource Utilization: Pooling allows for efficient management of grinder and brewer resources by queuing access requests. A resource (grinder/brewer) is only allocated when available, thus optimizing their usage and preventing bottlenecks.

* Simplicity: The code is more readable and easier to reason about. 

Other important aspect of the design:

* Dependency Injection: Utilizing dependency injection significantly reduces the coupling between components, facilitating easier modifications to the code and enhancing testability. 

* Enhancing the StartProcessingOrders function with an orderDone channel:  improves system response providing immediate user feedback on order completion,this simplifies test creation by enabling specific order tracking.






