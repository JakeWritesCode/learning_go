# Chapter 1 Notes

This is easy to forget with an IDE, but this is how you build a binary and run it.
```bash
go build  # in the main dir
./main    # in the main dir

# If you want a different name for the binary
go build -o mybinary
```

Go has a standard linter, there's only one way to lint code. I'm super into this because it means I don't have to waste
time arguing about code style. I can just run `gofmt` and `golint` and be done with it.

```bash
# So to format stuff, just run
go fmt ./...
# This is the standard way to tell go to recursively go through the file tree.
```

What's fun is that go actually requires a semicolon at the end of each line, but never put them in.
Go does this for you.

Additionally, go has a tool called go vet which checks for syntactically correct code that is probably wrong. 
Like this:

```go
import "fmt"

fmt.Printf("My special sring is: %s")
```

There's nothing wring with this, but why bother with an f-string if you're not going to put something in it?
`go vet` will fail this.

## Makefiles

To automate the build of a Go project we use makefiles. They allow us to automate the build process and make it repeatable
across machines. 

This is effectively kinda like the `npm run <something` commands you can specify in `package.json` of a node project.

## Compiling

Go programs compile to a native binary. This is fucking dope, because once it compiles it just compiles. No faffing
about installing dependencies from some `requirements.txt` that only works on Python 3.something and then won't 
install on your M1 mac or whatever other garbage. Once it's built, it runs fam.

Chapter 1 done!

