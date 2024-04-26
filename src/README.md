# json-parser
A simple JSON parser application to demonstrate the Lexical and Syntax analysis stages of compilers


### If you have Go installed and set up:

Run: 
```bash
go install
```

Usage:
```
json-parser "filename" [flags]
```

For a list of available flags and examples, run:
```bash
json-parser --help
```

### If you don't have Go installed:

Usage:
```
json-parser.exe "filename" [flags]
```

For a list of available options, run:
```bash
json-parser.exe help
```

NOTE: The token type consists of the token value (the actual token) and its position in the JSON string.