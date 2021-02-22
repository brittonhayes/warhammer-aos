# Warhammer - Age of Sigmar API ⚔️

A REST API for public Warhammer Age of Sigmar data. Built with Go + Lambda.

![Logo](./images/logo.png)

[![Go Reference](https://pkg.go.dev/badge/github.com/brittonhayes/warhammer-aos.svg)](https://pkg.go.dev/github.com/brittonhayes/warhammer-aos)

[![Go Report Card](https://goreportcard.com/badge/github.com/brittonhayes/warhammer-aos)](https://goreportcard.com/report/github.com/brittonhayes/warhammer-aos)

[![JSON Schema](https://github.com/brittonhayes/warhammer-aos/actions/workflows/validate.yml/badge.svg)](https://github.com/brittonhayes/warhammer-aos/actions/workflows/validate.yml)

## Usage ⚡

The API is free and public. Try it out!

### 🌐 [aos-api.com/armies](https://aos-api.com/armies)

[comment]: <> (## Docs 📝)

[comment]: <> (Looking for API documentation? We've got 'em! Our docs cover everything you can request from our API, how to do it, and)

[comment]: <> (how it works on the backend for anyone curious. Check em out!)

[comment]: <> (To get an example of a JSON response, check out the [data/json]&#40;data/json&#41; directory.)

[comment]: <> (### 🌐 [aos-api.com/docs]&#40;https://github.com/brittonhayes/warhammer-aos&#41;)

---

## Development 🔧

### Setup

> If you'd like to contribute to the request handling and inner-workings of the API, you'll need a few things set up in your development environment.

1. [Install Taskfile.dev](https://taskfile.dev/#/)
2. [Setup Go](https://golang.org/doc/install)
3. Enjoy 🎉

### Commands

```shell
# Build the lambda handlers
task build

# Run the lambda handlers locally
task invoke

# Run the tests
task test
```

## Contributing

> Wanna contribute _your_  Warhammer mini data to the API?
>
> We'd love your contributions!
> We've tried to make it as easy as possible for people to add their info without needing to understand the code.
>
> Check out our [contributing page](https://github.com/brittonhayes/warhammer-aos) to get started!

## Show your support

Give a ⭐️ if this project helped you!
