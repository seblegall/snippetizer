# Snippetizer

Print code snippets extracted from readme file of a given github repository.

## Context

With the growth of platforms like Github and dependencies management tools (whatever the language) more and more people share their work and open source the peace of code they have created to solve a specific problem. That's great !

But at the end... you find out that for a given problem you want to solve, their are tones of libs that claim to do the same thing but each in a different way.

When I need to choose a third-party component, I usually start by reading code example (hopefully in the README file).

The code bellow is a peace of Go code that parse Github readme file of a given repository and look for code example to show in the terminal.

It uses the markdown parser [golang-commonmark/markdown](https://github.com/golang-commonmark/markdown). Again, it took me time to find out the right vendor to do the job. You know... the one that don't do too much magic and stay simple. But not to simple because I don't want to waste time. The one that is also well documented....

Anyway, the code bellow can be run easily in command line using the `--url` flag within you pass the Github url of the repository you want to extract code snippets. It get the content of the README file, parse the markdown and print the extracted peaces of code.

If you combine this with the Go import system, it becomes really easy to quickly know how a third party library may be use.

## Installation

```sh
go get github.com/seblegall/snippetizer
```

## Usage

```sh
go run main.go --url http://github.com/{login}/{project}
```
