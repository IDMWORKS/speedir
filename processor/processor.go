package processor

import (
	"bufio"
	"errors"
	"log"
	"net"

	"github.com/idmworks/speedir/datacontext"
	"github.com/mavricknz/asn1-ber"
	"github.com/mavricknz/ldap"
	"io"
)

var ErrDecodingASN1 = errors.New("Error decoding asn1-ber packet: wrong port?")

type Processor struct {
	// DC provides access to the data layer
	DC *datacontext.DataContext
	// Verbose controls the verbosity of logging
	Verbose bool
	conn    net.Conn
}

type requestHandler func(proc *Processor, messageID uint64, request *ber.Packet) error

type requestProcessor struct {
	ldapCode uint8
	handler  requestHandler
}

var requestProcessors = make([]requestProcessor, 0)

// HandleRequest handles incoming LDAPv3 requests
func (proc *Processor) HandleRequest(conn net.Conn, errChan chan error) {
	proc.conn = conn
	// continuously read from the connection
	for {
		packet, err := ber.ReadPacket(bufio.NewReader(conn))

		if err == io.EOF {
			// connection closed by client
			conn.Close()
			return
		}

		if err != nil {
			errChan <- err
			continue
		}

		if proc.Verbose && (packet != nil) {
			ber.PrintPacket(packet)
		}

		// required to catch issues like TLS/TCP port mis-matches
		if len(packet.Children) == 0 {
			errChan <- ErrDecodingASN1
			continue
		}

		if err := proc.parsePacket(packet); err != nil {
			errChan <- err
		}
	}
}

func (proc *Processor) parsePacket(packet *ber.Packet) error {
	messageID := packet.Children[0].Value.(uint64)
	request := packet.Children[1]

	if request.ClassType == ber.ClassApplication &&
		request.TagType == ber.TypeConstructed {
		var handled bool
		for _, reqProc := range requestProcessors {
			if reqProc.ldapCode == request.Tag {
				if err := reqProc.handler(proc, messageID, request); err != nil {
					return err
				}
				handled = true
			}
		}
		if !handled {
			log.Println("LDAPv3 app code not implemented:", ldap.ApplicationMap[request.Tag])
		}
	}

	return nil
}

func (proc *Processor) sendLdapResponse(packet *ber.Packet) {
	buf := packet.Bytes()

	if proc.Verbose {
		ber.PrintPacket(packet)
	}

	for len(buf) > 0 {
		n, err := proc.conn.Write(buf)
		if err != nil {
			log.Printf("Error Sending Message: %s\n", err)
			return
		}
		if n == len(buf) {
			break
		}
		buf = buf[n:]
	}
}
