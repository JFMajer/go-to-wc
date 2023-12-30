# Word Count in Go
This is a simple command-line application written in Go that counts the number of words, lines or bytes in a text input.

## Usage
To use this application, you can pass in a text file as input and specify what you want to count: words, lines or bytes.

The flags are as follows:

- `-l` : Count lines
- `-b` : Count bytes

If no flags are specified, the application will count words by default.

### Example usage:
```
go run main.go < input.txt
```
This will count the number of words in the file input.txt.

```
go run main.go -l < input.txt
```
This will count the number of lines in the file input.txt.

```
go run main.go -b < input.txt
```
This will count the number of bytes in the file input.txt.

## Testing
The application includes a test suite with three test cases: one for counting words, one for counting lines, and one for counting bytes.

To run the tests, use the go test command:

go test
## Implementation
The application uses the bufio package to scan the input, and the flag package to handle command-line flags. The counting is done with a simple loop that increments a counter for each word, line or byte scanned, depending on the flags specified.