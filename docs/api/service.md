---
title: Simple API v1.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
highlight_theme: darkula
headingLevel: 2

---

<!-- Generator: Widdershins v4.0.1 -->

<h1 id="simple-api">Simple API v1.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

<h1 id="simple-api-default">Default</h1>

## GetIsaFundList

<a id="opIdGetIsaFundList"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /isa-funds

```

```http
GET /isa-funds HTTP/1.1

```

```javascript

fetch('/isa-funds',
{
  method: 'GET'

})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

result = RestClient.get '/isa-funds',
  params: {
  }

p JSON.parse(result)

```

```python
import requests

r = requests.get('/isa-funds')

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/isa-funds', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/isa-funds");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/isa-funds", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /isa-funds`

*List of ISA funds*

A list of Funds compatible with ISA products

<h3 id="getisafundlist-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|None|

<aside class="success">
This operation does not require authentication
</aside>

## get__user_product_available

> Code samples

```shell
# You can also use wget
curl -X GET /user/product/available \
  -H 'Accept: application/json'

```

```http
GET /user/product/available HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json'
};

fetch('/user/product/available',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json'
}

result = RestClient.get '/user/product/available',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/user/product/available', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/user/product/available', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/user/product/available");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/user/product/available", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /user/product/available`

*Products available for the user*

A list of products that the user can open accounts for

> Example responses

<h3 id="get__user_product_available-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|None|

<h3 id="get__user_product_available-responseschema">Response Schema</h3>

<aside class="success">
This operation does not require authentication
</aside>

## GetUserAccountBalances

<a id="opIdGetUserAccountBalances"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /user/account/balances \
  -H 'Accept: application/json'

```

```http
GET /user/account/balances HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json'
};

fetch('/user/account/balances',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json'
}

result = RestClient.get '/user/account/balances',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/user/account/balances', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/user/account/balances', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/user/account/balances");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/user/account/balances", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /user/account/balances`

*List user's account balances*

A list of the user's account balances across all products

> Example responses

<h3 id="getuseraccountbalances-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|None|

<h3 id="getuseraccountbalances-responseschema">Response Schema</h3>

<aside class="success">
This operation does not require authentication
</aside>

## UserCheckDepositAmount

<a id="opIdUserCheckDepositAmount"></a>

> Code samples

```shell
# You can also use wget
curl -X POST /user/account/check-deposit \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```http
POST /user/account/check-deposit HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "account_id": 0,
  "amount": 0
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('/user/account/check-deposit',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json'
}

result = RestClient.post '/user/account/check-deposit',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json'
}

r = requests.post('/user/account/check-deposit', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','/user/account/check-deposit', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/user/account/check-deposit");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/user/account/check-deposit", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /user/account/check-deposit`

*Check if user can invest the desired amount*

Checks if the user's desired deposit amount is within the allowed limits

> Body parameter

```json
{
  "account_id": 0,
  "amount": 0
}
```

<h3 id="usercheckdepositamount-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|true|none|
|» account_id|body|integer|true|none|
|» amount|body|number|true|none|

> Example responses

<h3 id="usercheckdepositamount-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|None|

<h3 id="usercheckdepositamount-responseschema">Response Schema</h3>

<aside class="success">
This operation does not require authentication
</aside>

## OpenRetailAccount

<a id="opIdOpenRetailAccount"></a>

> Code samples

```shell
# You can also use wget
curl -X POST /user/account/open-retail \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```http
POST /user/account/open-retail HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "product_uuid": "string",
  "fund_uuids": [
    {
      "fund_uuid": "string",
      "weight": 0
    }
  ]
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('/user/account/open-retail',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json'
}

result = RestClient.post '/user/account/open-retail',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json'
}

r = requests.post('/user/account/open-retail', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','/user/account/open-retail', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/user/account/open-retail");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/user/account/open-retail", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /user/account/open-retail`

*Opens a new account*

Opens a new account for the specified product and user

> Body parameter

```json
{
  "product_uuid": "string",
  "fund_uuids": [
    {
      "fund_uuid": "string",
      "weight": 0
    }
  ]
}
```

<h3 id="openretailaccount-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|true|none|
|» product_uuid|body|string|true|none|
|» fund_uuids|body|[object]|true|none|
|»» fund_uuid|body|string|false|none|
|»» weight|body|number|false|none|

> Example responses

<h3 id="openretailaccount-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|None|

<h3 id="openretailaccount-responseschema">Response Schema</h3>

<aside class="success">
This operation does not require authentication
</aside>

