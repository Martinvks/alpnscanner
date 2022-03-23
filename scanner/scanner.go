package scanner

import "crypto/tls"

func Scan(addr string, protocols []string) []string {
	conf := &tls.Config{}

	var result []string
	for _, protocol := range protocols {
		if isSupported(conf, protocol, addr) {
			result = append(result, protocol)
		}
	}
	return result
}

func isSupported(conf *tls.Config, protocol string, addr string) bool {
	conf.NextProtos = []string{protocol}

	conn, err := tls.Dial("tcp", addr, conf)

	if err != nil {
		return false
	}

	defer func(conn *tls.Conn) {
		_ = conn.Close()
	}(conn)

	return protocol == conn.ConnectionState().NegotiatedProtocol
}
