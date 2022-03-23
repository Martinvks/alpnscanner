# alpnscanner

Takes a file with hostnames and scans for supported protocols using [Application-Layer Protocol Negotiation](https://datatracker.ietf.org/doc/html/rfc7301).
By default, it scans for IANA's [list of ALPN protocol ids](https://www.iana.org/assignments/tls-extensiontype-values/alpn-protocol-ids.csv).

## Installation
```
go install github.com/Martinvks/alpnscanner@latest
```

## Usage
```
$ cat hosts.txt
example.com
example.no
example.se
$ alpnscanner -h hosts.txt
example.com:443 [http/1.1 h2]
example.no:443 []
example.se:443 [http/1.1]
```

## Required flags
```
  -h string
    	hosts file path
```

## Optional flags
```
  -c int
    	concurrency (default 10)
  -m string
    	mode specifies which protocols to scan for. Must be one of "iana", "http" or "h2draft" (default "iana")
  -p int
    	port (default 443)
```

The `iana` mode scans for the [protocols listed here](https://www.iana.org/assignments/tls-extensiontype-values/alpn-protocol-ids.csv).  
The `http` mode uses the same list, but only scans for HTTP protocols.  
The `h2draft` mode scans for draft versions of HTTP/2.