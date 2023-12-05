# customerimporter

This repository consists of two packages: a `customerimporter` package that implements the task,
and the `main` package, which implements the CLI interface. 

The CLI interface accepts one obligatory argument -- a path to the CSV file. 

### Running

To run the CLI, first compile everything from the perspective of the `main` module, and then run using the test data:

```bash
$ cd main
$ go build
$ ./main ../customers.csv
```
Documentation is available as normal in `go doc`.

### Design decisions

This implementation uses sorting and then counting. There is a more obvious approach, which uses maps instead. I argue that maps are using hashes internally, which although amortise to `O(1)`, usually have a big constant complexity. It would be however easier to extend, as it would be easier to implement sorting by domain count or other data retrieval. 

I decided not to introduce any external libaries, but I used the standard library freely. 

