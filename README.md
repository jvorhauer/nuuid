# Commandline UUID generator

Generate fresh UUIDs with versions 1, 3, 4, 5, 6 or 7 with ease

## Prerequisites

1. `Git` hsa been installed and is available on the PATH, verify with `git --version`
2. `Go` has been installed and is available on the PATH, verify with `go version`
3. The `go` binary directory has been added to the PATH, should be something like `$HOME/go/bin`

If the above list is succesfully and positively verified, then you can start with

## Installation

First clone this repo, or fork and the clone the fork locally:

```shell
git clone https://github.com/jvorhauer/nuuid.git
```

Then run

```shell
go build
```

and

```shell
go install
```

This will have installed the `uuid` executable program in the `go` bin directory! Now, go celebrate and generate some nice UUIDs!

## Usage

Name of the executable is `uuid`, the only parameter to be passed is the version of UUID you would like to generate:

```shell
uuid --version N
```

where `N` can be 1, 3, 4, 5, 6 or 7

## Feedback

Love some feedback! Leave constructive critisism, feature requests and bug reports in the
{GitHub Issues}(https://github.com/jvorhauer/nuuid/issues) part of [this project](https://github.com/jvorhauer/nuuid).

Happy generating!
