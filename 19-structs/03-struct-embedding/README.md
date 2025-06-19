# Struct Embedding Exercise

In this exercise, you'll practice working with struct embedding and JSON parsing in Go. Your task
is to implement two functions that parse JSON data into a set of nested structs. 

## Parsing books 

The first struct represents a book with an author. The JSON formatted structure is as follows:

```json
{
  "title": "The Go Programming Language",
  "author": {
    "name": "Alan A. A. Donovan",
    "address": "103 W Vandalia St, Edwardsville, Indiana, USA"
  },
  "pages": 380,
  "ISBN": "978-0134190440"
}
```

Your task is to define a struct named `Book` that should represent a book and implement the
following function to parse a book from JSON:

```go
func ParseBook(jsonData []byte) (Book, error)
```

## Parsing articles

The second struct should represent an article in a journal. The serialized JSON format is as follows:

```json
{
  "title": "Smashing The Kernel Stack For Fun And Profit",
  "author": {
    "name": "Sinan Eren",
    "address": "12031 N Tatum Blvd, Phoenix, Arkansas, USA"
  },
  "journal": "Phrack Magazine"
  "year": 2002,
}
```

You need to implement a function with the following signature for parsing an article:

```go
func ParseArticle(jsonData []byte) (Article, error)
```

Requirements:
1. Define an `Author` struct with Name and Address fields.
2. Define a `Book` and an `Article` struct that both embed the `Author` struct to represent author info.
3. Implement the `ParseBook` and `ParseArticle` functions to parse the JSON data into a `Book` or an `Article` struct.
4. Use appropriate tags on your struct fields to match the JSON keys.

Place your code in the `exercise.go` file near tho the placeholder `// INSERT YOUR CODE HERE`



