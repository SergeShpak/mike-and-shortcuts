# Solution to the Mike and Shortcuts problem

To run the solution with a given input file `input.txt` first build the code:

```bash
go build -o solution .
```

and then run

```bash
cat input.txt | ./solution 
```

To run the tests execute:

```bash
go test ./... -count=1
```

You can find several test cases in [`samples`](./samples).