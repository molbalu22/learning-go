# Filtering data

Write a function that accepts two arguments: a slice of keys and a slice of indexes. The function has to return an array of filtered keys of length 10. The filtering should fulfill the following requirements:

- If the length of key slice does not match the length of the index slice, return an array of 10 empty strings.
- Remove the i-th key if the i-th index is *`even`*, and keep it otherwise.
- If the length of resultant array would be shorter than 10, then fill up the remaining elements with empty keys.

Place your code into the file `exercise.go` near the placeholder `// INSERT YOUR CODE HERE`.
