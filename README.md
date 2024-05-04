# Zius

Zius is a Golang struct validator library.

## Installation

To install Zius, you can use `go get`:

```sh
go get github.com/vinibgoulart/zius
```

## Usage

Here's a simple example of how to use Zius:

```golang
// Import the package
import "github.com/vinibgoulart/zius"

// Define your struct
type User struct {
    Name  string `zius:"required"`
    Email string `zius:"required,email"`
}

// Validate your struct
func ValidateUser(user User) error {
    return zius.Validate(user)
}
```

## Validators

Zius comes with a number of built-in validators:

- `required`: Ensures the field is not empty
- `email`: Validates that the field is a valid email address
- `number`: Validates that the field is a valid number
- `int`: Validates that the field is a valid integer
- `phone`: Validates that the field is a valid phone number
- `regex`: Validates the field against a custom regular expression
- `equals`: Validates the field against a custom value

### Custom Messages

You can define a custom message for each validator like that:

```golang
type User struct {
    Name  string `zius:"required:Name is required"`
    Email string `zius:"required:Email is required,email:Email must be a valid email"`
}
```

### Regex Validator

You can define a custom regular expression for the `regex` validator sending the regex as a parameter:

```golang
type User struct {
    Name  string `zius:"regex={^([a-zA-Z]+)$}"`
}
```

In the example above, the `Name` field will only be valid if it contains only letters.

### Equals Validator

You can define a custom value for the `equals` validator sending the value as a parameter:

```golang
type User struct {
    Name  string `zius:"equals={John, Vinicius}"`
}
```

In the example above, the `Name` field will only be valid if it is equal to `John` or `Vinicius`.

## Examples

You can find more examples of how to use Zius in the `__examples__` directory.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
