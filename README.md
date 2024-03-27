```text
  --------------
  \            /
   \          /
    \        /
     \      /
      \    /
       \  /
        \/		         ~   -------   |
	/\        |	     |	  \ / 	   |   |
       /  \	  |	     |	   |	       |
      /	   \	  |	     |	   |	       |
     /	    \	  |	     |	   |	       |
    /        \	  |	     |	   |	       |
   /	      \	  |	     |     |	       |
  /  	       \  |	     |     |	       |
  --------------  ------------	   -           ----------
```

Xurl is a simple curl alternative, built using golang. The main purpose of Xurl is to fetch and send data to stdout to make it easy to chain/pipe it into other cli tools to do what you want.

# Usage - 

```bash
xurl [flags] [data] [address]
```

Simply provide the URL to fetch the data from that address.
## Schemes
Xurl plans to support a variety of Uri schemes. The currently supported schemes are-
- Http/Https - `http://google.com`
- File - `file://{path-to-file}`
- Websockets - `ws://localhost:3000/ws`

*Xurl determines the scheme by observing the prefix of the address, separated by the  `://` operator*

### Default Scheme
**Http/Https** is the default Uri scheme. You can also just use `www.`, which will xurl will replace by `http://` automatically. 

Example
```bash
xurl www.google.com
```

#### Supported Methods

***GET***
```bash
xurl http://google.com
```

***POST***
In order to send a post request, you would have to use the *-data* flag. This would be accompanied by the path of your data file, prefixed by `@`

```bash
xurl --data @path.json http://localhost:3000/post
```

### Websocket

Example 
```bash
xurl ws://localhost:3000/ws
```

This creates a websocket connection and opens up an interactive shell. You can use this shell to send further messages, each message being read as a single line input.

 > [!warning]
>  Redirection `>` currently does not work with websockets.

## HeadersOnly flag
In case you only want to fetch the response headers, use the `-headersOnly` flag
```bash
xurl --headersOnly www.google.com
```

***
> [!note]
>  This project is created solely for learning purposes. Not recommended for professional use.
