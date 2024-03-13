```text
  --------------
  \            /
   \          /
    \        /
     \      /
      \    /
       \  /				                   |
        \/		             ~   -------   |
	    /\        |	     |	  \ / 	   |   |
       /  \	      |	     |	   |	       |
      /	   \	  |	     |	   |	       |
     /	    \	  |	     |	   |	       |
    /        \	  |	     |	   |	       |
   /	      \	  |	     |     |	       |
  /  	       \  |	     |     |	       |
  --------------  --------    ----	       -----------
```

Xurl is a simple curl alternative, built using golang. The main purpose of Xurl is to fetch and send data to stdout to make it easy to chain/pipe it into other cli tools to do what you want.

> [!note] This project is created solely for learning purposes. Not recommended for professional use.


# Usage - 

```bash
xurl [flags] [address]
```

Simply provide the URL to fetch the data from that address. 

> [!warning] Xurl as of now only supports the Http protocol. Other protocols would be supported in the upcoming releases.

## HeadersOnly
In case you only want to fetch the response headers, use the `-headersOnly` flag
```bash
xurl --headersOnly www.google.com
```