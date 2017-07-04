# fizzbuzz
## An interview excercise
This REST api is an interface to the classic FizzBuzz coding example.
It returns a JSON list of numbers and fizz or buzz or fizzbuzz strings.
It's operated using a http request with a GET request. You must provide the following URL paramenters.
- string1: the string for the fizz argument 
- string2: the name for the buzz argument
- int1: the value of the fizz divisor
- int2: the name of the buzz divisor
- limit: tha amount of numbers for which the fizzbuzz sequence is generated.

To start up te project:

```
go get github.com/wouterbeets/fizzbuzz
fizzbuzz
```

this will launch the server on your localhost port 8080

open this in a webrowser with fizz as the path
localhost:8080/fizz
for the api to work you need to pass the required parameters in the url

```
localhost:8080/fizz?string1=foo&string2=bar&int1=2&int2=5&limit=100
```


This project demonstrates a way of coding that I appreciate.
The naming is clear,
The architecture is simple with little or no dependencies,
It's scalable,

The code is relativly clean
It needs more testing and documentation
But this is as far as i could go with a fulltime job and 2 kids in the short timespan given.
I've coded this in around 4 hours.

I hope it's what you are after, I can comment on every design decision. I'm happy to discuss the implementation details and scaling challenges.

Cheers,
Wouter
