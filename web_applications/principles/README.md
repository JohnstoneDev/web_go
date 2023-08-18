# WEB WORKING PRINCIPLES

## HTTP Protocol

<p>
	HTTP is the protocol that is used to facilitate communication between browser and web server. It is based on the TCP protocol and usually uses port 80 on the side of the web server. It is a protocol that utilizes the request-response model -clients send requests and servers respond. According to the HTTP protocol, clients always setup new connections and send HTTP requests to servers.
	Servers are not able to connect to clients proactively, or establish callback connections. The connection between a client and a server can be closed by either side. For example, you can cancel your download request and HTTP connection and your browser will disconnect from the server before you finish downloading.
</p>

<p>

	The HTTP protocol is stateless, which means the server has no idea about the relationship between the two connections even though they are both from same client. To solve this problem, web applications use cookies to maintain the state of connections.

	Because the HTTP protocol is based on the TCP protocol, all TCP attacks will affect HTTP communications in your server. Examples of such attacks are SYN flooding, DoS and DDoS attacks.
</p>


### HTTP request package (browser information)

Request packages all have three parts: request line, request header, and body. There is one blank line between header and body.

<code>
	GET /domains/example/ HTTP/1.1      ==> request line: request method, URL, protocol and its version
	<br>
	Host：www.iana.org             ==> domain name
	User-Agent：Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.4 (KHTML, like Gecko) Chrome/22.0.1229.94 Safari/537.4            ==> browser information
	Accept：text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8    ==> mime that clients can accept
	Accept-Encoding：gzip,deflate,sdch     ==> stream compression
	Accept-Charset：UTF-8,*;q=0.5     ==> character set in client side
	// blank line
	// body, request resource arguments (for example, arguments in POST)
</code>


### HTTP is stateless and Connection: keep-alive

<p>
	The term stateless doesn't mean that the server has no ability to keep a connection. It simply means that the server doesn't recognize any relationships between any two requests.

	Keep-alive is used by default. If clients have additional requests, they will use the same connection for them.

<p>
	Notice that Keep-alive cannot maintain one connection forever; the application running in the server determines the limit with which to keep the connection alive for, and in most cases you can configure this limit.
</p>

</p>