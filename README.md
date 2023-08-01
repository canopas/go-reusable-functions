# go-reusables

A comprehensive collection of essential and reusable functions for Go developers.

This repository aims to provide solid reusable functions of foundational tools that enhance Go programming experiences and streamline the development workflow.

## Features

- 🚀 Streamline development with foundational utility functions.
- 🛠 Simplify common tasks with reusable code snippets.
- 🧰 Build upon strong reusables of well-tested Go functions.
- ⚡ Boost productivity and code readability in your projects.

## Requirements

- Go v1.20

## Installation

- If you want to use the email function, install the package with: 
    ```
    go get github.com/canopas/go-reusables/email
    ```
## Usage

go-reusables offer a diverse range of utility functions that cater to various aspects of software development. Here's a quick example of how to use a email function:

```go
package main

import (
	"github.com/canopas/go-reusables/email"
)

func main() {
	email.SendAWSSESEmail(sess, data)
}
```

For detailed documentation and examples of each function, please refer to the [examples](https://github.com/canopas/go-reusables/tree/main/examples).

## Contributions and Feedback

Contributions are highly appreciated! If you find a bug, have a feature request, or want to contribute new functions, feel free to open an issue or submit a pull request.

## License

go-reusables is open-source and distributed under the [MIT License](https://github.com/canopas/go-reusables/blob/main/LICENSE).
