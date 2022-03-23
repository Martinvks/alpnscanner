package utils

import "github.com/Martinvks/alpnscanner/arguments"

var ianaHTTP = []string{
	"http/0.9",
	"http/1.0",
	"http/1.1",
	"spdy/1",
	"spdy/2",
	"spdy/3",
	"h2",
	"h2c",
	"h3",
}

var ianaMisc = []string{
	"stun.turn",
	"stun.nat-discovery",
	"webrtc",
	"c-webrtc",
	"ftp",
	"imap",
	"pop3",
	"managesieve",
	"coap",
	"xmpp-client",
	"xmpp-server",
	"acme-tls/1",
	"mqtt",
	"dot",
	"ntske/1",
	"sunrpc",
	"smb",
	"irc",
	"nntp",
	"nnsp",
}

var draftHTTP2 = []string{
	"HTTP/2.0",
	"h2-17",
	"h2-16",
	"h2-15",
	"h2-14",
	"h2-13",
	"h2-12",
	"h2-11",
	"h2-10",
	"h2-09",
	"h2-08",
	"h2-07",
	"h2-06",
	"h2-05",
	"h2-04",
	"h2-03",
	"h2-02",
	"h2-01",
	"h2-00",
}

func GetProtocols(mode int) []string {
	switch mode {
	case arguments.ModeHttp:
		return ianaHTTP
	case arguments.ModeH2Draft:
		return draftHTTP2
	default:
		return append(ianaHTTP, ianaMisc...)
	}
}
