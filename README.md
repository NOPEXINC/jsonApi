#JSON from API

A simple go program that gets json data from an external api and displays the 
data.

## Getting Started
Make sure you have go installed, then clone to your gosource directory.

You can either compile and run the executable like this 
`go build main.go` which will compile the main.go file in an executable 
so that you can just run it like `./main`

Another way do is just to run the program 
`go run main.go` 

##NOTE
The posts data is obtained from this url `https://jsonplaceholder.typicode.com/posts`
so if you have a different place where you can get your json data, you should just make
sure your post struct matches the data returned. 

##LICENSE
MIT
