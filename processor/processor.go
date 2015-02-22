package processor

import (
	"bufio"
	"log"
	"net"

	"github.com/idmworks/speedir/datacontext"
	"github.com/mmitton/asn1-ber"
)

type Processor struct {
	// DC provides access to the data layer
	DC *datacontext.DataContext
	// Verbose controls the verbosity of logging
	Verbose bool
	conn    net.Conn
}

type requestHandler func(proc *Processor, messageID uint64, request *ber.Packet)

type requestProcessor struct {
	ldapCode uint8
	handler  requestHandler
}

var requestProcessors []requestProcessor

func init() {
	requestProcessors = make([]requestProcessor, 0)
}

// HandleRequest handles incoming LDAPv3 requests
func (proc *Processor) HandleRequest(conn net.Conn) {
	proc.conn = conn
	// continuously read from the connection
	for {
		packet, err := ber.ReadPacket(bufio.NewReader(conn))

		if proc.Verbose {
			ber.PrintPacket(packet)
		}

		if err != nil {
			defer conn.Close()
			log.Println("Error reading:", err.Error())
			return
		}

		// required to catch issues like TLS/TCP port mis-matches
		if len(packet.Children) == 0 {
			defer conn.Close()
			log.Println("Error decoding asn1-ber packet: wrong port?")
			return
		}

		proc.parsePacket(packet)
	}
}

func (proc *Processor) parsePacket(packet *ber.Packet) {
	messageID := packet.Children[0].Value.(uint64)
	request := packet.Children[1]

	if request.ClassType == ber.ClassApplication &&
		request.TagType == ber.TypeConstructed {
		var handled bool
		for _, reqProc := range requestProcessors {
			if reqProc.ldapCode == request.Tag {
				reqProc.handler(proc, messageID, request)
				handled = true
			}
		}
		if !handled {
			log.Println("LDAPv3 app code not implemented:", request.Tag)
		}
	}
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

func (proc *Processor) buildLdapResponse(messageID uint64, ldapResult int) *ber.Packet {
	packet := ber.Encode(ber.ClassUniversal, ber.TypeConstructed, ber.TagSequence, nil, "LDAP Response")
	packet.AppendChild(ber.NewInteger(ber.ClassUniversal, ber.TypePrimative, ber.TagInteger, messageID, "MessageID"))
	return packet
}
