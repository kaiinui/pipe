pipe
====

A simple web app to pipe GET / POST request to fluentd.

### via POST-ing JSON

```
POST /p

{
    "some": "data"
}
```

Passes posted JSON data to fluentd.

To pass multiple events to fluentd, just do as follows.

```
POST /p

[
    {
        "some": "foo"
    },
    {
        "some": "bar"
    }
]
```

### via GET-ing `.gif`

For CORS, logging via GET-ing `img` is provided.

```html
<img src='https://pipe.example.com/p.gif?data={\"some\":\"data\"}' />
```

Results following request

`GET /p.gif?data={\"some\":\"data\"}`

which logs following JSON data.

```json
{
    "some": "data"
}
```

Specification
---

### Endpoints

#### `POST /p`

- Params: NO PARAMS REQUIRED
- Body: 

```yaml
$: Event
```

- Response:
  - 201 Created
    - Body: NO_CONTENT
  - 422 Unprocessable Entity
    - Body: NO_CONTENT

#### `GET /p.gif`

- Params:
  - `data`: URL encoded JSON that conforms to [Event](#event) format
- Response:
  - Body: 1x1 transparent GIF image

### Entity

#### Event

```
$: object|array
```

#### Example

```json
{
    "foo": "bar"
}
```

```json
[
  {
    "foo": "bar"
  },
  {
    "bar": "foo"
  }
]
```
