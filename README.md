# jwt.go

## JWT Printer written in Golang

### Usage

```bash
~ $ jwt --token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
```

Or redirect shell input

```bash
~ $ jwt < ~/my_token.jwt
```

#### Arguments

```plain
  --no-color, -c # Print without colour
  --header, -h # Print just the token header
  --payload, -p # Print just the token payload

```

### Features

#### Done

- Parsing a JWT
- Printing unstructured JSON

#### Todo

- CLI arguments (see above)
- Take a token as a parameter and print
- Ingest and print token from input redirection
- Testing

