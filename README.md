Golang Websockets
=================

This project is a code exercise following Jacob Martin's article [Practical
Golang: Using websockets][article].

[article]: https://jacobmartins.com/2016/03/07/practical-golang-using-websockets/

I like the article, because it shows a very minimal example of how to write a
Golang application that allows a client to connect using websockets.  There are
only two source files doing all the work.

Dependencies
============

The websocket library used is the [Gorilla WebSocket][lib] library.  You can
fetch this dependency by running:

[lib]:https://github.com/gorilla/websocket

```sh
go get github.com/gorilla/websocket
```

Running
=======

To run this example, you can simply run the following:

```sh
go run main.go
```

Then navigate your browser to <http://127.0.0.1:8000>.

You should see a "Subscribe" button and a text area with "MyTextArea" already in
the field.

When you press the "Subscribe" button you'll see the following eventually fill
the text area:

```
MyTextArea
{"Name":"Bill","Age":0}
{"Name":"Bill","Age":2}
{"Name":"Bill","Age":4}
{"Name":"Bill","Age":6}
{"Name":"Bill","Age":8}
{"Name":"Bill","Age":10}
{"Name":"Bill","Age":12}
{"Name":"Bill","Age":14}
{"Name":"Bill","Age":16}
{"Name":"Bill","Age":18}
{"Name":"Bill","Age":20}
{"Name":"Bill","Age":22}
{"Name":"Bill","Age":24}
{"Name":"Bill","Age":26}
{"Name":"Bill","Age":28}
{"Name":"Bill","Age":30}
{"Name":"Bill","Age":32}
{"Name":"Bill","Age":34}
{"Name":"Bill","Age":36}
{"Name":"Bill","Age":38}
Connection closed
```

You'll also see the following logged in your terminal:

```
2017/02/17 17:33:37 Client subscribed
2017/02/17 17:34:19 Client unsubscribed
```

You can stop the server by pressing `Ctrl+C` in your terminal.

For more details explaining the code, please see [Jacob Martin's
article][article].
