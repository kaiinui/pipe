pipe
====

A simple web app to pipe GET / POST request to fluentd.

- [Usage](#usage)
  - [via POST-ing JSON](#via-post-ing-json)
  - [via GET-ing .gif](#via-get-ing-gif)
- [Specification](#specification)
  - [Endpoints](#endpoints)
    - [POST /p](#post-p)
    - [GET /p.gif](#get-pgif)
  - [Entity](#entity)
    - [Event](#event)
- [References](#references)

## Usage

### via POST-ing JSON

```
POST /e
Content-Type: application/json

{
    "some": "data"
}
```

Passes posted JSON data to fluentd.

To pass multiple events to fluentd, just do as follows.

```
POST /e
Content-Type: application/json

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
<img src='https://pipe.example.com/e.gif?data={\"some\":\"data\"}' />
```

Results following request

`GET /e.gif?data={\"some\":\"data\"}`

which logs following JSON data.

```json
{
    "some": "data"
}
```

Specification
---

### Endpoints

#### `POST /e`

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

#### `GET /e.gif`

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

References
---

- Docker - Build, Ship, and Run Any App, Anywhere : https://www.docker.com/
- Fluentd | Open Source Data Collector : http://www.fluentd.org/
- InfluxDB - Open Source Time Series, Metrics, and Analytics Database : http://influxdb.com/

LICENSE
---

```
The MIT License (MIT)

Copyright (c) 2014 kaiinui

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
