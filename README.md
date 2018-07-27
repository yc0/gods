# gods
Mimic java data-structure for golang


# Vector
ArrayList may incur contention due to thread unsafe
Vector data structure can handle synchronization.

Here comes the example of saling the tick
### ArrayList
```
    arraylist_test.go:116: Sale Ticket: 1
    arraylist_test.go:116: Sale Ticket: 2
    arraylist_test.go:116: Sale Ticket: <nil>
    arraylist_test.go:116: Sale Ticket: 4
    arraylist_test.go:116: Sale Ticket: 3
    arraylist_test.go:116: Sale Ticket: 6
    arraylist_test.go:116: Sale Ticket: <nil>
    arraylist_test.go:116: Sale Ticket: 7
    arraylist_test.go:116: Sale Ticket: 8
    arraylist_test.go:116: Sale Ticket: 8
```
As you can see, ticket 8 is sold twice.

### Vector
```
vector_test.go:118: Sale Ticket: 2
vector_test.go:118: Sale Ticket: 1
vector_test.go:118: Sale Ticket: 3
vector_test.go:118: Sale Ticket: 4
vector_test.go:118: Sale Ticket: 5
vector_test.go:118: Sale Ticket: 6
vector_test.go:118: Sale Ticket: 7
vector_test.go:118: Sale Ticket: 8
vector_test.go:118: Sale Ticket: 9
vector_test.go:118: Sale Ticket: 10
```
No matter how many times you run, you can get no replicated ticket sold.
