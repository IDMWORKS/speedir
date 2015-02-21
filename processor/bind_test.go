package processor

import (
	"testing"

	"github.com/idmworks/speedir/datacontext"
	"github.com/mmitton/asn1-ber"
	"github.com/mmitton/ldap"
)

const (
	dbname = "speedir_test"
	dbuser = "speedir_test"
)

type credentials struct {
	username string
	password string
	result   int
}

var testCredentials = []credentials{
	{"admin", "admin", ldap.LDAPResultSuccess},
	{"admin", "admin2", ldap.LDAPResultInvalidCredentials},
	{"admin2", "admin", ldap.LDAPResultInvalidCredentials},
}

func TestMain(t *testing.T) {
	DC = &datacontext.DataContext{DBName: dbname, DBUser: dbuser}
	DC.InitDb()
	defer DC.CloseDb()
	DC.SeedDb()
}

func TestBuildBindResponse(t *testing.T) {
	messageID := 1
	expected := ldap.LDAPResultProtocolError
	packet := buildBindResponse(uint64(messageID), expected)
	actual, found := parseLDAPResult(packet)
	if !found {
		t.Error("BindResponse malformed")
	}
	if actual != expected {
		t.Error("BindResponse result mismatch")
	}
}

func TestGetBindResponse(t *testing.T) {
	DC = &datacontext.DataContext{DBName: dbname, DBUser: dbuser}
	DC.OpenDb()
	defer DC.CloseDb()

	for i, creds := range testCredentials {
		testGetBindResponse(t, uint64(i), creds)
	}
}

func BenchmarkGetBindResponse(b *testing.B) {
	DC = &datacontext.DataContext{DBName: dbname, DBUser: dbuser}
	DC.OpenDb()
	defer DC.CloseDb()

	creds := testCredentials[0]
	for i := 0; i < b.N; i++ {
		testGetBindResponse(b, uint64(i), creds)
	}
}

func testGetBindResponse(tb testing.TB, messageID uint64, creds credentials) {
	request := buildBindRequest(creds.username, creds.password)
	response, _ := getBindResponse(messageID, request)
	actual, found := parseLDAPResult(response)
	if !found {
		tb.Error("BindResponse malformed for", creds)
	}
	if actual != creds.result {
		tb.Error("BindResponse result mismatch for", creds)
	}
}

func buildBindRequest(username string, password string) *ber.Packet {
	bindRequest := ber.Encode(ber.ClassApplication, ber.TypeConstructed, ldap.ApplicationBindRequest, nil, "Bind Request")
	bindRequest.AppendChild(ber.NewInteger(ber.ClassUniversal, ber.TypePrimative, ber.TagInteger, 3, "Version"))
	bindRequest.AppendChild(ber.NewString(ber.ClassUniversal, ber.TypePrimative, ber.TagOctetString, username, "User Name"))
	bindRequest.AppendChild(ber.NewString(ber.ClassContext, ber.TypePrimative, 0, password, "Password"))
	return bindRequest
}

func parseLDAPResult(packet *ber.Packet) (result int, found bool) {
	if len(packet.Children) >= 2 {
		response := packet.Children[1]
		if response.ClassType == ber.ClassApplication && response.TagType == ber.TypeConstructed && len(response.Children) == 3 {
			found = true
			result = int(response.Children[0].Value.(uint64))
		}
	}
	return
}
