# Learn Go With Tests - BDD and Acceptance tests

Many developers get caught up in the tooling around BDD, in particular DSLs to make tests "read like English". There's nothing inherently wrong with this, if you have non-technical stakeholders actually reading these tests then it can be helpful. However, I've seen many teams sink a lot of time into these tests but are missing the main point behind BDD, and it becomes a very wasteful overhead.

BDD is a much broader topic than things like Cucumber. Dave Farley discusses how it related to domain-driven-design (another misunderstood topic), in particular "ubiquitous language". Farley discusses how software development is fundamentally about modelling the real world and processes. 

If a team behind a system (technical, and non-technical people) can agree upon a ubiquitous language it improves communication and allows us to work with less ambiguity. When the development team adopt this language in their code, the tests in the system can describe what reality should be.

Taking this idea further, your acceptance tests should be decoupled from incidental technical details. This way your acceptance tests will always be the "truth", unless the business changes its mind about the model your system is about; in which case changing the tests is reasonable. 

BDD, as the name suggests should be *behaviour driven*, but what does that mean in practice?

I try to set up my teams, so they have close collaboration with domain experts and stakeholders. Along with the development team they will try to understand what the next most valuable _behaviour_ we need to deliver to our customers. This is known as writing "user stories". It's important at this point not to get bogged down in **implementation details**. Think about what the user actually wants:

- Users don't want to log in. They want to purchase a book and have it delivered.
- Users don't want to have an API endpoint so that our system can book a room when our React frontend does a POST to our backend for frontend. They want to book a hotel room.

Even if you have the discipline to remain truly user focused, when you pick up that story to do, how do you start? 

In this post I will be referring to snippets from [my reference repo](https://github.com/quii/go-reference-repo) which is an example starting point of a project which pulls together a lot of the principles I use for writing web services in Go. It's just my opinion, and it won't be for everyone and won't cater for every context and need. This is not a framework, it's just ideas. Here are the kind of things it optimises for:

- Able to run acceptance, unit and integration tests locally and without fuss. Should work out of the box. `./build.sh && git push` should be safe to go to prod. Emphasis on short cycle times, small-batch work for fast feedback-loops and reduce risk. 
- Package up in to a small Docker image for deployment.
- As testable as reasonably possible. High confidence the system works without manual checks.
- Obvious where to make a start.
- Not opinionated of how things work within `internal`. If you want to do a layered architecture, hexagonal, ports & adapters, it doesn't matter too much. 
- To keep things testable, the code inside `cmd` should be very minimal.

## How to start developing a user story

Developers usually make life hard for themselves by not working methodically. They'll start worrying about technical details, writing lots of code but not validating it actually meets the user need, or the **behaviour** that you're supposed to be delivering. 

Growing Object-Oriented Software guided by tests prescribes starting with some kind of black-box test as your starting point. By stating how you want your system to behave from the start, and aggressively only writing the code you need to make it pass; you'll have a north star to govern your efforts. 

When a pair of developers pick up a story in this repo, if it's for a distinctly new functionality, they start with an **acceptance test**. If it's a further iteration on existing functionality, they should be able to add/change the unit tests within that area to drive out the change.

Acceptance tests live in their own package in the root.

```go
type GreetingSystemAdapter interface {
	Greet(name string) (greeting string, err error)
}

func GreetingAcceptanceTest(t *testing.T, system GreetingSystemAdapter) {
	t.Run("greets people in a friendly manner", func(t *testing.T) {
		is := is.New(t)

		greeting, err := system.Greet("Pepper")
		is.NoErr(err)
		is.Equal(greeting, "Hello, Pepper!")
	})
}
```

Interfaces allow us to decouple our test from implementation detail. All it describes is **the behaviour we want from our system**. Interfaces allow us to describe the _what_, not the _how_, which means they are a great tool here. The behaviour could be realised by a HTTP API, a website, a command line tool, or a Slackbot, anything that we can invoke and observe programmatically. 

We're not using some kind of "BDD framework" like Cucumber or Ginko, we don't need to. The test clearly states what behaviour we need, in plain Go code.

To use this test, you create a `GreetingSystemAdapter` to interface with the "system under test" (SUT). This will depend on the technical nature of the SUT. For a "real" website, this would involve making a `GreetingSystemAdapter` that could use Selinium under-the-hood, to drive a web-browser against our site. For this example though, we'll just hit the web pages and scrape the responses.

## Our adapter

Our adapter is our way of interfacing with our system. In our case it will make HTTP calls to a web server to try and get the `Greet` behaviour we want. 

```go
type APIClient struct {
	baseURL string
	httpClient *http.Client
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: 5 * time.Second},
	}
}

func (a *APIClient) Greet(name string) (string, error) {
	url := a.baseURL + "/greet/"+name

	res, err := a.httpClient.Get(url)
	if err != nil {
		return "", fmt.Errorf("problem reaching %s, %w", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status %d from %q", res.StatusCode, url)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// snipped - methods for checking the system is up and running
```

### "Test in prod"

A useful property of our adapter is that it can be pointed at a different `baseURL` according to our needs. It is extremely desirable to be able to run your acceptance tests not only locally, but against other environments like a QA environment and live. 

To have this property may require extra engineering effort. For instance if you're working on a retail system and you wish to check that customers can order books, you'll need the notion of a "test" mode for a customer that wont actually take a payment and deliver the book. 

This may seem wasteful, but it is 100x cheaper and less stressful than the alternative, which is being afraid to ship your code and relying on manual testing in live. 

### Putting it together

In `/black-box-tests` we use our acceptance test and adapter to run our test against our system.

```go
func TestGreetingApplication(t *testing.T) {
	client := acceptance_tests.NewAPIClient(getBaseURL(t))

	if err := client.WaitForAPIToBeHealthy(five_retries); err != nil {
		t.Fatal(err)
	}

	acceptance_tests.GreetingAcceptanceTest(t, client)
}
```

Sometimes the timing of bringing up the application can make tests flaky so the adapter gives us a way of polling the system to check it is up (`WaitForAPIToBeHealthy`) before trying to run a test. 

From here, we snap our acceptance test together with our adapter to exercise the behaviour of the SUT.

### `getBaseURL(t)`

An important principle I have for my tests is that they should be self-contained, and set themselves up. I do not wish to manually spin up databases, applications, e.t.c. just to run a test. 

`getBaseURL` checks the environment for a `BASE_URL` value. If it finds one, then we're running it against a deployed environment like production, so from there it can work as usual. 

If there isn't one, it's assumed it's running locally. It uses a combination of `docker-compose` and [test-containers](https://www.testcontainers.org) to build the docker image of our system and bring up a container of the SUT.

If the system needs to rely on other systems like databases, we can also model them in docker-compose, so they'll also be spun up for the test. In addition, we can spin these systems up for integration tests using the same mechanism. 

Not only does it make running and writing the test simple & convienient, but it also gives us a huge degree of confidence the system works. My teams ship Docker images to be deployed, so if we can see the Docker image builds and exhibits the behaviours we defined, then we know we can ship. 

## Making it pass

Making the acceptance test pass for a new behaviour can often be quite involved, even just for the happy path:

- New routing
- New handlers
- New service-type layer to deal with the business logic
- Working with 3rd party systems like databases
- Tests for the above

People can get in to trouble with acceptance tests where they are red for a long time. For this reason, try to find ways to "cheat" and get it passing as quickly as possible. Once you have the test passing you've proved a lot of things

- Your Dockerfile is setup correctly
- Your routing is correct
- Dependencies are wired up
- Your adapter is correct

There are a few techniques to move quickly and get the test passing:

- Be ruthless with scope. Stick to the happy path and don't get distracted by edge-cases.
- If you're relying on 3rd party systems, consider using a fake version at first. For instance instead of a Postgres implementation of storage, knock together an in-memory version.

Once you have your test passing you now have the license to iterate further and with confidence. At this point, you can more or less forget about the acceptance tests, and you can probably drive out the additional behaviour you need via unit tests, which are simpler and faster to work with. 

## Further iterations

The top-down, behaviour-focused approach means you get the "happy path" working, and you can validate that all the "plumbing" code of creating new routers and services is done and is serving a real need. Further iterations can typically be done via adding more unit tests and working in the domain code. 