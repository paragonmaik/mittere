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

Accessing help

```sh
mittere -h
```

## Usage

```sh
mittere -h
```

## File options

### (test.json)
```sh
{
  "url": "",
  "method": "",
  "data": {},
  "headers": {}
}
```

### (test.yaml)
```sh
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
    "Authorization": "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImFscGhhIjpmYWxzZX0.eyJpc3MiOiJEaW5vQ2hpZXNhLmdpdGh1Yi5pbyIsInN1YiI6InRhbWFyYSIsImF1ZCI6ImF1ZHJleSIsImlhdCI6MTcwMTAwNjUzOSwiZXhwIjoxNzAxMDA3MTM5fQ.exFVyedz_UQHXbCG8OJq2Qbaeg36HE7uBt1dExOcL6UDA90Rb6w4G9IAjPXLgkLFpu_918zkiprbMSYqb8lOTS2LQ5oJV-6u4rOM-HmLjbsuL0VH_Y25XZsq9RtR0iJ7Ooz2m4H6QOTUqMo9mZ9lwRmj0UIbd3skRyEUwpiCYYmh--H8e-d2HUQUj2TVua5OqkUHsCPg83U2xnTaA1-7N_pEuII32wlWrRwrtpppd0j4gxgCSFxsuETMdv0POshGZgdRsDsiYIGJL2rCjsRPxTA6fkIcQ0K3WEdQg5BpVBFmCF6utAKnZUlxWmT4vCFCwoINTDyWaCjqWLsTj2uJ5g"
  }
}

```
### POST

### PUT

### DELETE

### Options

