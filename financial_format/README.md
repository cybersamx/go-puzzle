# Number to Text in Financial Format

Implement a function that returns the text representation of an input number in financial format. For example:

```go
input := 1234567
output := itoa(input)

// Output: "1,234,567"
fmt.Println(output)
```

There are 2 approaches to implementing this:

* Mathematically by breaking a number into a series of number (components) representing each digit (in the radix of 10). We then insert a comma as we iterate through each number.
* Use string manipulation by first converting into a string and then insert a comma as we iterate through each character.
