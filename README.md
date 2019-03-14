scut
-----

A [scut](https://www.dictionary.com/browse/scut) is a Rabbit's tail. This is `tail` for RabbitMQ!

## Installing

You have a few installation options! After installing you can verify everything worked by running `scut --version`.

### Binary Installation

Download your binary of choice from the [Releases](https://github.com/kpurdon/scut/releases).

### Homebrew

Run `brew tap kpurdon/tap && brew install scut`.

## Usage

Let's say you have a topic exchange named `events` and want to see all messages published to that exchange with the binding key of `events.foo`. Run the following, new messages will be printed to stdout:

```
scut --exchange-name=events --binding-key=events.foo
{"foo":"bar"}
{"foo":"bar"}
{"foo":"bar"}
```

## Contributing

If you have a feature request open an issue to discuss it first.

1. Fork the repo.
2. Create a feature branch.
3. WORK. WORK. WORK.
4. Open a Pull Request.
5. Profit! (after some peer review of course)
