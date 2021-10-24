# Go HTTP service reference implementation

This is my current state-of-the-art opinion on how I like to structure Go projects for how I, and how I want my teams to work.

## High level-requirements

I want developers to practice trunk-based development [for various reasons](https://quii.dev/Reduce_WIP_by_practicing_trunk-based_development,_rather_than_pull_requests). 

The system we work on, it's structure and its internal quality, has a huge effect on the way we work and our productivity. I am very-much subscribed to [The DevOps Three Ways](https://itrevolution.com/the-three-ways-principles-underpinning-devops/) which emphasises **flow, feedback-loops and a continuous culture of improvement and learning**. Too many repos I've worked on impede flow and have poor and slow feedback loops.

It's important that developers can **safely and confidently push small, positive, high-quality changes to the system frequently through the day**

The process for making change should _roughly_ be:

- `git pull -r`
- If it's a distinctly new feature, start with an acceptance test, otherwise, a unit test to drive a further iteration of an existing feature.
- See the test fail.
- Make it pass.
- `git commit -am "added new feature"`
- Refactor.
- `git add .`
- `git commit --amend --no-edit`
- `git pull -r`
- `./build.sh && git push`

Repeat as necessary. Always bear in mind the [test pyramid](https://martinfowler.com/bliki/TestPyramid.html).

### What does it take to work that way safely?

- Modular code. Each bit of code should have a clear purpose which is cohesive and loosely coupled. If a system has lots of inappropriate and tight-coupling then developers will tread on each other's toes. 
- Enough structure & convention to make it obvious where to start work, and where to put things. 
  - But not so opinionated about a particular "way" that if a new requirement comes along that doesn't fit that model, that it requires extensive re-work.
- Excellent observability (out of scope for this repo, this is org-specific)
- **Tests**. Manual testing is unacceptable.

#### Tests!!

- Extremely fast unit tests. Developers should be re-running them constantly. In order to make small, frequent, positive changes to the system through a day you need a tight feedback-loop.
- Integration tests, ideally running against real versions of the systems our code is working with. Use docker-compose and testcontainers to orchestrate spinning up containers for the test. No manual work required to run tests, they should work out of the box.
- Acceptance tests.
  - Behaviour & domain focused.
  - Decoupled from implementation detail.
  - Can be executed against our local version, or against other environments, including live.
  - As we ship Docker images to be deployed, for the local run we should build our image and test against a running container that we intend to ship. This gives us huge confidence the system will work in production.

The tests should all be runnable locally. Having to push code to a "CI server" to get feedback on changes is too slow.

## Implementation notes

### Prereqs

- Go
- Docker
- [golangci-lint](https://golangci-lint.run)

### Specifications & acceptance tests

Specifications should be decoupled from your implementation detail. For new features they should be seen as a starting point for work where you describe "the truth" in terms of what behaviour your system should exhibit. 

```go
type GreetingSystemDriver interface {
	Greet(name string) (greeting string, err error)
}

func Greeting(t *testing.T, greetingSystem GreetingSystemDriver) {
	t.Helper()
	t.Run("greets people in a friendly manner", func(t *testing.T) {
		is := is.New(t)

		greeting, err := greetingSystem.Greet("Pepper")
		is.NoErr(err)
		is.Equal(greeting, "Hello, Pepper!")
	})
}
```

To use this test, you create a `driver` which implements the interface you need to run the test. For the black-box acceptance tests that's a [HTTP client which calls our API](https://github.com/quii/go-http-reference-impl/blob/main/black-box-tests/acceptance/greeting-api-driver.go#L19) given a `baseURL`. This means we can run them locally but also against deployed environments like live with very little effort.

You can also re-use these specifications to test your domain code too, because the criteria and rules of the domain should hold true _within_ your system too.

```go
func HelloGreeter(name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}
```

```go
func TestHelloGreeter(t *testing.T) {
    specifications.Greeting(t, specifications.GreetingSystemFunc(HelloGreeter))
}
```

### cmd

All this should be responsible for is:

- Getting configuration
- Using configuration to create the necessary dependencies for the application (calling `NewFoo` functions)

### application

Within here should live what you could roughly call your "domain" or application code. It should contain most of your data models and business logic.

Specifics as to how to organise beyond that I don't have especially strong opinions about, but the code must be:

- Modular & testable
- Have a sensible separation of concerns
- Cohesive

You should not have code that interacts with "the outside world" in here. For that you put in

### adapters

This is where other users and systems interact with `application`. Think HTTP servers, message queues, databases e.t.c.

#### HTTP

One strong opinion I do hold is around to structure HTTP servers.

`NewWebServer(config SomeConfig, dependencyA DependencyA, dependencyB, DependencyB, etc) *http.Server`

This means in `main` I can pass in configuration and real dependencies to create my server and then start `log.Fatal(server.ListenAndServe())`. It also means we can use `httptest.NewServer(NewWebServer(...))` to test our web server at a unit level too. 

`http.Handlers` should all look roughly the same.

- Parse and validate a request.
- Call `h.service.SomeUsefulThing(parsedRequest)`.
- Send a response based on what was returned above

```go
type Greeter interface {
    Greet(name string) (greeting string, err error)
}

type GreetHandler struct {
	greeter Greeter
}

func NewGreetHandler(greeter Greeter) *GreetHandler {
	return &GreetHandler{greeter: greeter}
}

func (g *GreetHandler) Greet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	greeting, err := g.greeter.Greet(vars["name"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, greeting)
}
```

This keeps handlers, skinny, simple to test, and means we can unit test our important business logic without HTTP causing noise and complexity. [I've written more about this in Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/questions-and-answers/http-handlers-revisited).

The responsibility of handling HTTP is with "HTTP handlers", but they shouldn't do much more beyond that.

### Dockerfile and Docker-compose

The Dockerfile is a fairly standard, multi-stage build image which allows us to build our code and then ship very small containers

Docker-compose allows us to declaratively define what our app depends on, which is useful for the acceptance tests when running locally but also lets us spin up say `Redis` for our integration tests (in conjunction with [`testcontainers`](https://www.testcontainers.org)).
