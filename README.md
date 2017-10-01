# krakend-http-auth

HTTP Basic authentication midleware for the [KrakenD framework](https://github.com/devopsfaith/krakend).

**This is not intended to be use in production! It is just a initial version of a possible KrakenD component**

## Example

Check the dummy implementation in the `example` dir. It contains a simple config file and a KrakenD api-gateway with the auth `HandlerFactory` wrapping the default endpoint factory.

Build it

	$ make all

And run it

	$ ./auth -l DEBUG -c krakend.json

From a new terminal, try to acces the private endpoint with and without the auth header

	$ curl -i http://127.0.0.1:8080/private/kpacha
	HTTP/1.1 403 Forbidden
	Content-Type: text/plain; charset=utf-8
	Date: Sun, 01 Oct 2017 17:47:18 GMT
	Content-Length: 17

	wrong auth header

	$ curl -i -u foo:bar http://127.0.0.1:8080/private/kpacha
	HTTP/1.1 200 OK
	Cache-Control: public, max-age=0
	Content-Type: application/json; charset=utf-8
	X-Krakend: Version undefined
	Date: Sun, 01 Oct 2017 17:48:09 GMT
	Content-Length: 159

	{"authorizations_url":"https://api.github.com/authorizations","code_search_url":"https://api.github.com/search/code?q={query}{\u0026page,per_page,sort,order}"}
