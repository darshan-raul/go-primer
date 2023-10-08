# Go Primer

This guide can be used to get decently skilled with Golang


## TDD:

Steps:
    - Write a test
    - Make the compiler pass
    - Run the test, see that it fails and check the error message is meaningful
    - Write enough code to make the test pass
    - Refactor

By ensuring your tests are fast and setting up your tools so that running tests is simple you can get in to a state of flow when writing your code.
By not writing tests you are committing to manually checking your code by running your software which breaks your state of flow and you won't be saving yourself any time, especially in the long run

## Gotchas:

- You can launch the docs locally by running `godoc -http :8000`. If you go to `localhost:8000/pkg` you will see all the packages installed on your system.
    - If you don't have godoc command, then maybe you are using the newer version of Go (1.14 or later) which is no longer including godoc. You can manually install it with go install golang.org/x/tools/cmd/godoc@latest.
- The official site to visit any public package docs in golang is `https://pkg.go.dev`. If you are coming from the python background then this is similar to `pypi.org`
    - Here are the steps to publish: https://go.dev/doc/modules/publishing
- Testing:
    - It is important to question the value of your tests. It should not be a goal to have as many tests as possible, but rather to have as much confidence as possible in your code base. Having too many tests can turn in to a real problem and it just adds more overhead in maintenance. Every test has a cost.

## Reading list

- [ ] https://go.dev/doc/effective_go
- [ ] https://github.com/Pungyeon/clean-go-article
- [ ] https://github.com/golang-standards/project-layout
- [ ] https://github.com/dgryski/go-perfbook/tree/master
- [ ] https://thewhitetulip.gitbook.io/bo/
- [ ] https://github.com/avelino/awesome-go
- [ ] https://github.com/thangchung/go-coffeeshop 
- [ ] https://github.com/grpc-ecosystem/grpc-gateway

- [ ] https://medium.com/@ninucium/parallelism-and-concurrency-in-go-how-it-works-in-real-computing-systems-part-1-a680443ad8bd
- [ ] https://medium.com/@ninucium/parallelism-and-concurrency-in-go-how-it-works-in-real-computing-systems-part-2-b423288d047d



## References used:

- [Best free book to learn Golang from DigitalOcean](https://www.digitalocean.com/community/books/how-to-code-in-go-ebook)
- https://gobyexample.com/
- [Learn go with tests](https://quii.gitbook.io/learn-go-with-tests/)
    - to start writing robust, well-tested systems in Go
- [Learn Go with 11 Projects](https://www.youtube.com/watch?v=jFfo23yIWac) 
- [Logging guide](https://www.datadoghq.com/blog/go-logging/)
- [Bootcamp](https://playbook.one2n.in/go-bootcamp)
