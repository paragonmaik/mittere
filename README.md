<div align="center">
  <pre>
███╗   ███╗██╗████████╗████████╗███████╗██████╗ ███████╗
████╗ ████║██║╚══██╔══╝╚══██╔══╝██╔════╝██╔══██╗██╔════╝
██╔████╔██║██║   ██║      ██║   █████╗  ██████╔╝█████╗  
██║╚██╔╝██║██║   ██║      ██║   ██╔══╝  ██╔══██╗██╔══╝  
██║ ╚═╝ ██║██║   ██║      ██║   ███████╗██║  ██║███████╗
╚═╝     ╚═╝╚═╝   ╚═╝      ╚═╝   ╚══════╝╚═╝  ╚═╝╚══════╝
                                                        
  </pre>
</div>

A simple cUrl like tool built in go with the Cobra library. Mittere reads yaml and json files to make htpp requests.

## Installation

```sh
brew tap paragonmaik/clitools
brew install mittere
```

## Help

```sh
mittere -h
```

## Usage

```sh
mittere -f "path/to/file.json"
```

### Colorizing output
Default color is red.

```sh
mittere -f "path/to/file.json" -c
```
or
```sh
mittere -f "path/to/file.json" -c -C blue
```

### Options
Method and Url options take precedence over values written to the file.

```sh
mittere -u "url" -m "method" -f "path/to/file.json"
```

## File options

### (test.json)
```json
{
  "url": "",
  "method": "",
  "data": {},
  "headers": {}
}
```

### (test.yaml)
```yaml
url:
method:
data:
headers:
```

## Sample
### GET
```sh
{
  "url": "https://jsonplaceholder.typicode.com/posts",
  "method": "post",
  "data": {
    "title": "foo",
    "body": "bar",
    "userId": 1
  },
  "headers": {
    "Content-Type": "application/json",
    "Authorization": "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImFscGhhIjpmYWxzUX0.eyJpc3MiOiJEaW5vQ2hpZXNhLmdpdGh1Yi5pbyIsInN1YiI6InRhbWFyYSIsImF1ZCI6ImF1ZHJleSIsImlhdCI6MTcwMTAwNjUzOSwiZXhwIjoxNzAxMDA3MTM5fQ.exFVyedz_UQHXbCG8OJq2Qbaeg36HE7uBt1dExOcL6UDA90Rb6w4G9IAjPXLgkLFpu_918zkiprbMSYqb8lOTS2LQ5oJV-6u4rOM-HmLjbsuL0VH_Y25XZsq9RtR0iJ7Ooz2m4H6QOTUqMo9mZ9lwRmj0UIbd3skRyEUwpiCYYmh--H8e-d2HUQUj2TVua5OqkUHsCPg83U2xnTaA1-7N_pEuII32wlWrRwrtpppd0j4gxgCSFxsuETMdv0POshGZgdRsDsiYIGJL2rCjsRPxTA6fkIcQ0K3WEdQg5BpVBFmCF6utAKnZUlxWmT4vCFCwoINTDyWaCjqWLsTj2uJ5g"
  }
}

```
### POST

### PUT

### DELETE



