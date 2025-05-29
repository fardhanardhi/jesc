# jesc

**jesc** is a lightweight CLI tool that recursively unescapes deeply nested and escaped JSON strings within a JSON object. It's especially useful when you're dealing with messy JSON blobs where some fields are double, triple, or even more levels of string-encoded JSON.

## ✨ Features

- 🔄 Recursively unescapes deeply nested escaped JSON strings
- ✅ Supports unescaping both objects (`{}`) and arrays (`[]`)
- 🧪 Optional pretty-formatting for better readability
- 🛠️ Simple CLI interface
- 🐛 Great for debugging or cleaning up bad API responses

## 📦 Installation

### MacOS & Linux

```bash
brew tap fardhanardhi/jesc
brew install jesc
```

### Windows

See [releases page](https://github.com/fardhanardhi/jesc/releases).


## 🚀 Usage

#### With direct json string

```bash
jesc "<your-json-string>"
```

#### With input json file path

```bash
jesc --file "<your-json-filepath>"
```

### Optional flags

- `-f`, `--format`: Format output JSON with indentation (pretty print)
- `--output`: Output filepath

### Example

#### Input (escaped JSON inside JSON string)

```json
{
  "name": "John Doe",
  "hobbies": "[{\"name\":\"climbing\",\"isFavorite\":true}]",
  "esc": "{\"a\":3,\"b\":\"haha\",\"c\":\"{\\\"deep\\\":true}\"}"
}
```

#### Command

##### Direct json string with print output

```bash
jesc '{"name":"John Doe","hobbies":"[{\"name\":\"climbing\",\"isFavorite\":true}]","esc":"{\"a\":3,\"b\":\"haha\",\"c\":\"{\\\"deep\\\":true}\"}"}' -f
```

##### Json file with print output

```bash
jesc --input input.json -f
```

##### Json file with file output

```bash
jesc --input input.json --output output.json -f
```

#### Output

```json
{
  "name": "John Doe",
  "hobbies": [
    {
      "name": "climbing",
      "isFavorite": true
    }
  ],
  "esc": {
    "a": 3,
    "b": "haha",
    "c": {
      "deep": true
    }
  }
}
```

## 🧠 How It Works

Internally, jesc:
1. Parses the top-level JSON string.
2. Iterates each field:
    - If the value is a string that looks like JSON, it attempts to unmarshal it.
    - Repeats recursively for any successfully unmarshaled map.
3. Outputs the final cleaned and structured JSON.

## 🔨 Building

Clone the repository and build it using Go:

```bash
git clone https://github.com/fardhanardhi/jesc.git
cd jesc
go build
```

## 📄 License

The project is licensed under the [MIT License](https://github.com/fardhanardhi/jesc/blob/main/LICENSE).