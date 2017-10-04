package sslServerKey

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"log"
	"testing"
	"github.com/stretchr/testify/assert"
)

var csr string = `
-----BEGIN CERTIFICATE REQUEST-----
MIIEvDCCAqQCAQAwGzEZMBcGA1UEAwwQdGVzdC5leGFtcGxlLmNvbTCCAiIwDQYJ
KoZIhvcNAQEBBQADggIPADCCAgoCggIBAOMc9Kz6ksGyae4/u+/ni1RuyFotUlwD
MYhKXRBo/3/IFuX781Z+3FzGqltR597JkcIxgFpIe9RGTerFy3bVBF7KHXWuVJXb
prVR2OvBlZi8BIvYTj9VWTXCLyJy2/08oxg1MFXUZDmlL1aE4lMHnhl2vlZiBRgs
V9Jbs2dl+UoKRDR6uN5ElZ2OOZJzRKtU2fObZU6KTBjQUNXkq/6BGDZirWT1xNQC
6io7wacmx1CWXg6ca5jt9sGipIA/e96SPEF2m5V+kuCy2TgBirlAMTIKJZDdTHb9
6bn1PB+GngXT+Yme/MnkR7soX1HPpfBKNY+Ruo82+8H83ttixRLJetpUmN7ddUfD
SEhR081vdcFtRNeSsbp/vzlacTWbRiCcj6HQNGab1qVDtDPFCYkVZyLFXTZ5L9kD
EEhIa/EDu8W3agX5gRFZ8q9JiBJOY5ZWNsJYScQYs9NHrJe0FmcFZ+ylJcNI4lSc
qgLNrDZ6UQFdcvjwWVh6S55+B0W4Cy259dkK5bsk3AySKeAAUTeCC+//62YN59VQ
OSiXI2JrUvjGKpbdWIsecoQQ2tzRNKjvlmFnnq1Aeh3zaGkzxWT/4WMx8icLP9RP
/ltkOOcs+gNOyUeuLqJZrrnXO/TprW9P9lchyc1H085yfrQtrMhDrYe5FyQsWvMl
jj9KhbyMJj+/AgMBAAGgXDBaBgkqhkiG9w0BCQ4xTTBLMAkGA1UdEwQCMAAwCwYD
VR0PBAQDAgXgMDEGA1UdEQQqMCiCEHRlc3QuZXhhbXBsZS5jb22CFHd3dy50ZXN0
LmV4YW1wbGUuY29tMA0GCSqGSIb3DQEBCwUAA4ICAQCVeX4doSPWT4MywBYf4JF0
d99KQ0KszdiDGU5WFlfs9zuBbKX/JW4zog7PT4Xc3tOC2zBoVlVLXqqbEKF/H3bU
Gid9izbOdQgzThpX1KNieEHh3X7807JaFg7RYPxrzOSMdvG3qNGUM4Hf5Fbv93ik
wfIn6c0hPcS+jhiQKngjSWpcu+2JHcsfxrkDD2aOnR2yotiDFviMcjeURF23Fl1m
jFduWt+OMJFDWi0RQx2v9cGCeKmDdSku5O0UmV63Djo7CsTgcSl39JT1NB0gUMxp
Ti+TT6ts1rMHeoupIbcRoaeJ7+vQ/ymj1gWii/C3S+n00N6RVVozP3BflRzrXiwa
STFCuHwcwkbW5IyU4w3yABkSUT9wnSU8CpvJYqKfHYQHIlg/H2y6JoJLvxiyaIXf
FNPo+HWFRPo/Ap5P4Nnf0ZKB96WrbwHh2dlkRgzp6bb/xCvseSurlsrm/WXcDt88
BL/IIoWZYPyx/dbAhbKdkthmU6T7ZuxgexZistjWmjP9YiJ1/A5WKo6vsmhlF8Sd
6BfGPIQ3EsB3ZLiF+v8ZPtUgR6Oa2qu2kcw8ZmwIzyt682NqjuI1U+SQz7V7dPVH
aqCKk4zu9mHJsndHlupJO99oxzhGtPAHf5lOuIF8X323WgwTzG0vzJNW/sPsw72d
OE/+FhDYt7wCqngrV4HVKQ==
-----END CERTIFICATE REQUEST-----
`

var certificate string = `
-----BEGIN CERTIFICATE-----
MIIFNTCCAx2gAwIBAgIJAOlCmaDJjTLEMA0GCSqGSIb3DQEBCwUAMBsxGTAXBgNV
BAMMEHRlc3QuZXhhbXBsZS5jb20wHhcNMTcxMDA0MTAxMjMxWhcNMTgxMDA0MTAx
MjMxWjAbMRkwFwYDVQQDDBB0ZXN0LmV4YW1wbGUuY29tMIICIjANBgkqhkiG9w0B
AQEFAAOCAg8AMIICCgKCAgEA4xz0rPqSwbJp7j+77+eLVG7IWi1SXAMxiEpdEGj/
f8gW5fvzVn7cXMaqW1Hn3smRwjGAWkh71EZN6sXLdtUEXsodda5UldumtVHY68GV
mLwEi9hOP1VZNcIvInLb/TyjGDUwVdRkOaUvVoTiUweeGXa+VmIFGCxX0luzZ2X5
SgpENHq43kSVnY45knNEq1TZ85tlTopMGNBQ1eSr/oEYNmKtZPXE1ALqKjvBpybH
UJZeDpxrmO32waKkgD973pI8QXablX6S4LLZOAGKuUAxMgolkN1Mdv3pufU8H4ae
BdP5iZ78yeRHuyhfUc+l8Eo1j5G6jzb7wfze22LFEsl62lSY3t11R8NISFHTzW91
wW1E15Kxun+/OVpxNZtGIJyPodA0ZpvWpUO0M8UJiRVnIsVdNnkv2QMQSEhr8QO7
xbdqBfmBEVnyr0mIEk5jllY2wlhJxBiz00esl7QWZwVn7KUlw0jiVJyqAs2sNnpR
AV1y+PBZWHpLnn4HRbgLLbn12QrluyTcDJIp4ABRN4IL7//rZg3n1VA5KJcjYmtS
+MYqlt1Yix5yhBDa3NE0qO+WYWeerUB6HfNoaTPFZP/hYzHyJws/1E/+W2Q45yz6
A07JR64uolmuudc79Omtb0/2VyHJzUfTznJ+tC2syEOth7kXJCxa8yWOP0qFvIwm
P78CAwEAAaN8MHowHQYDVR0OBBYEFEBh7lMVv3cslHuQmeFg3pF9aFB0MEsGA1Ud
IwREMEKAFEBh7lMVv3cslHuQmeFg3pF9aFB0oR+kHTAbMRkwFwYDVQQDDBB0ZXN0
LmV4YW1wbGUuY29tggkA6UKZoMmNMsQwDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0B
AQsFAAOCAgEA4i5U4H9gp9aYbkAi0hoE+Iuw7pMtKIiXZRqz0CzeCJIUOWS634Cn
AQmKQYVh8YCBLg117Wmjg/5MQZ8uX6OTK1OazQQUjlVf7947fvxcjAfizl+DDS+k
tHW8z7i3kuz9w+DOcXH9yB78Ln/1vu83Q1RX8u8KPZVYy2ieIJr2FhHh8yj/Opk1
sotnZ6NZZg/e/AFyg28Pa1iLvhqvUUkKhXlbOh/DoT7J7vpNoZ5U9dXNTBITHzo3
COgjyeYulTwYvz1ej2i4fBHfdskGu4hERsulknnsj5ODg1uQ6cPXX/A/kiJE28z4
joywMorubDnetjocV8UOHO0yVMiFqRcCsbZOFp60wmKlJZT9DElvVopVWYy4kkNL
VKz6H3pFtMCcN9B2V05NWfL0bA6/ex4iVE6qkLE97AOS43PmJKnN/Dh6ScfEsAtK
mD7drq9uj2UV7juDQYQdIarvySwyIXGS20G6E+KIFfsMh9axy+WCNg+F43hV6l3y
j6XHdxVkbd6RhCsuPEzvjdOXZPx2wTcR6Rwks1yJeYpJBYVZ79rtqQN22vlbVxbp
fsBpN2irCanEmYFunUwr0yfBUqksW484OEmMKCISB99PxMp9DK8fqsI/86MbMD8P
+bAfRKSrCUKeSq6zBwpidap2k7HEx5Nd9KZddeHqgfc4mqDR769CjHM=
-----END CERTIFICATE-----
`

var privateKey string = `
-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----
`

func TestSSLServerKey(t *testing.T) {

	name := api.SetTestResourceName("ssl_server_key_")
	setSSLServerKey(name, t)
	getSSLServerKey(name, t)
	deleteSSLServerKey(name, t)
}

func setSSLServerKey(name string, t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}

	resource := SSLServerKey{}
	resource.Properties.Basic.Note = "go-brocade-vtm test certificate"
	resource.Properties.Basic.Request = csr
	resource.Properties.Basic.Public = certificate
	resource.Properties.Basic.Private = privateKey

	newSSLServerKey := SSLServerKey{}
	err = client.Set("ssl/server_keys", name, resource, &newSSLServerKey)
	if err != nil {
		t.Fatal("Error creating a resource: ", err)
	}
	log.Println("Created SSL Server Key ", name)

	assert.Equal(t, "go-brocade-vtm test certificate", newSSLServerKey.Properties.Basic.Note)
	assert.Equal(t, csr, newSSLServerKey.Properties.Basic.Request)
	assert.Equal(t, certificate, newSSLServerKey.Properties.Basic.Public)
	// We don't test for the private key as the API doesn't return it
}

func getSSLServerKey(name string, t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	client.WorkWithConfigurationResources()

	sslServerKey := SSLServerKey{}
	err = client.GetByName("ssl/server_keys", name, &sslServerKey)
	if err != nil {
		t.Fatal("Error getting resource: ", err)
	}
	log.Println("Found SSL Server Key resource: ", name)

	assert.Equal(t, "go-brocade-vtm test certificate", sslServerKey.Properties.Basic.Note)
	assert.Equal(t, csr, sslServerKey.Properties.Basic.Request)
	assert.Equal(t, certificate, sslServerKey.Properties.Basic.Public)
	// We don't test for the private key as the API doesn't return it
}

func deleteSSLServerKey(name string, t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}

	err = client.Delete("ssl/server_keys", name)
	if err != nil {
		t.Fatal("Error deleting resource: ", err)
	} else {
		log.Printf("Resource %s deleted", name)
	}
}
