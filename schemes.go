package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

type Schemes interface {
	MakeRequest()
}

type DefaultScheme struct {
	method  string
	reqData *[]byte
	addr    string
}

func (s DefaultScheme) MakeRequest() {
	var (
		resp *http.Response
		err  error
	)

	if *data {
		resp, err = http.Post(s.addr, "application/json", bytes.NewBuffer(*s.reqData))
	} else {
		resp, err = http.Get(s.addr)
	}

	if err != nil {
		fail(err, "Error while getting a response.")
	}
	defer resp.Body.Close()

	if *headerOnly {
		for k, v := range resp.Header {
			fmt.Fprintln(os.Stdout, k, v)
		}
		os.Exit(0)
	} else {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fail(err, "Error while reading response body.")
		}
		fmt.Fprintln(os.Stdout, string(body))
		os.Exit(0)
	}
}

type WebsocketScheme struct {
	url      string
	protocol string
	origin   string
	data     *[]byte
}

func (s WebsocketScheme) MakeRequest() {
	conn, err := websocket.Dial(s.url, s.protocol, s.origin)
	if err != nil {
		fail(err, "Error while dialing websocket connection")
	}
	defer conn.Close()
	
	if *data {
		if _, err := conn.Write(*s.data); err != nil {
			fail(err, "Error while writing data to connection")
		}

		msg := make([]byte, 512)
		if _, err = conn.Read(msg); err != nil {
			fail(err, "Error while reading message from websocket connection")
		}

		fmt.Fprintln(os.Stdout, string(msg))
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		// FIXME(May cause error if the line is too long.)
		if *s.data, _, err = reader.ReadLine(); err != nil {
			fail(err, "Error while reading from console")
		}

		if _, err := conn.Write(*s.data); err != nil {
			fail(err, "Error while writing data to connection")
		}

		msg := make([]byte, 512)
		if _, err = conn.Read(msg); err != nil {
			fail(err, "Error while reading message from websocket connection")
		}

		fmt.Fprintln(os.Stdout, string(msg))
	}
}

type ImapScheme struct{}

func (s ImapScheme) MakeRequest() {

}

type FileScheme struct {
	path string
}

func (s FileScheme) MakeRequest() {
	body, err := os.ReadFile(s.path)
	if err != nil {
		fail(err, "Error while fileUriRequest")
	}
	fmt.Fprintln(os.Stdout, string(body))
}

type TelnetScheme struct{}

func (s TelnetScheme) MakeRequest() {

}
