# URL Shortener

- The purpose of this exercise is to write a http.Handler that will examine the path of any incoming web request and decide whether or not to redirect the visitor to a new page, similar to how a URL shortener might work.

## Supported Flags

- `yaml` - A YAML file with path:url mapping
- `json` - A JSON file with path:url mapping

### Example

```bash
➜ go build .

➜ ./url-shortener --json=./samples/urls.json

Starting server at port: :7070
```

If no file is provided, it fallsback to a hardcoded map url-path pattern matching within the code.
