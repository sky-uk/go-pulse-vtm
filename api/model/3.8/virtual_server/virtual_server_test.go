package virtualserver

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestVirtualServer(t *testing.T) {

	name := api.SetTestResourceName("virtual_servers")
	setVirtualServer(name, t)
	getVirtualServer(name, t)
	deleteVirtualServer(name, t)
}

func setVirtualServer(name string, t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}

	resource := VirtualServer{}
	resource.Properties.Basic = Basic{
		AddClusterIP:         true,
		AddXForwarded:        true,
		AddXForwardedProto:   true,
		CloseWithRst:         true,
		ConnectTimeout:       30,
		Enabled:              true,
		ListenOnAny:          true,
		MSS:                  uint(0),
		Note:                 "Created by go-brocade-vtm test",
		Pool:                 "test-pool",
		Port:                 80,
		Protocol:             "http",
		SoNagle:              true,
		SslClientCertHeaders: "none",
		SslDecrypt:           false,
		SslHonorFallbackScsv: "use_default",
		Transparent:          false,
	}

	setTrue := true
	keepAliveTimeOut := uint(100)
	maxClientBuffer := uint(32768)
	maxServerBuffer := uint(32768)
	maxTransactionDuration := uint(12)
	timeOut := uint(100)

	resource.Properties.Connection = Connection{
		Keepalive:              &setTrue,
		KeepaliveTimeout:       &keepAliveTimeOut,
		MaxClientBuffer:        &maxClientBuffer,
		MaxServerBuffer:        &maxServerBuffer,
		MaxTransactionDuration: &maxTransactionDuration,
		Timeout:                &timeOut,
	}

	connectTimeout := uint(0)
	dataFrameSize := uint(2048)
	headerTableSize := uint(8192)
	idleTimeoutNoStreams := uint(240)
	idleTimeoutOpenStreams := uint(300)
	maxConncurrentStreams := uint(800)
	maxFrameSize := uint(16384)
	maxHeaderPadding := uint(128)
	streamWindowSize := uint(32768)

	resource.Properties.HTTP2 = HTTP2{
		ConnectTimeout:         &connectTimeout,
		DataFrameSize:          &dataFrameSize,
		Enabled:                &setTrue,
		HeaderTableSize:        &headerTableSize,
		HeadersIndexDefault:    &setTrue,
		IdleTimeoutNoStreams:   &idleTimeoutNoStreams,
		IdleTimeoutOpenStreams: &idleTimeoutOpenStreams,
		MaxConcurrentStreams:   &maxConncurrentStreams,
		MaxFrameSize:           &maxFrameSize,
		MaxHeaderPadding:       &maxHeaderPadding,
		MergeCookieHeaders:     &setTrue,
		StreamWindowSize:       &streamWindowSize,
	}

	newVirtualServer := VirtualServer{}
	err = client.Set("virtual_servers", name, resource, &newVirtualServer)
	if err != nil {
		t.Fatal("Error creating a resource: ", err)
	}
	log.Println("Created Virtual Server ", name)

	assert.Equal(t, true, newVirtualServer.Properties.Basic.AddClusterIP)
	assert.Equal(t, true, newVirtualServer.Properties.Basic.AddXForwarded)
	assert.Equal(t, true, newVirtualServer.Properties.Basic.AddXForwardedProto)
	assert.Equal(t, true, newVirtualServer.Properties.Basic.CloseWithRst)
	assert.Equal(t, uint(30), newVirtualServer.Properties.Basic.ConnectTimeout)
	assert.Equal(t, true, newVirtualServer.Properties.Basic.Enabled)
	assert.Equal(t, true, newVirtualServer.Properties.Basic.ListenOnAny)
	assert.Equal(t, uint(0), newVirtualServer.Properties.Basic.MSS)
	assert.Equal(t, "Created by go-brocade-vtm test", newVirtualServer.Properties.Basic.Note)
	assert.Equal(t, "test-pool", newVirtualServer.Properties.Basic.Pool)
	assert.Equal(t, uint(80), newVirtualServer.Properties.Basic.Port)
	assert.Equal(t, "http", newVirtualServer.Properties.Basic.Protocol)
	assert.Equal(t, true, newVirtualServer.Properties.Basic.SoNagle)
	assert.Equal(t, "none", newVirtualServer.Properties.Basic.SslClientCertHeaders)
	assert.Equal(t, false, newVirtualServer.Properties.Basic.SslDecrypt)
	assert.Equal(t, "use_default", newVirtualServer.Properties.Basic.SslHonorFallbackScsv)
	assert.Equal(t, false, newVirtualServer.Properties.Basic.Transparent)

	assert.Equal(t, true, *newVirtualServer.Properties.Connection.Keepalive)
	assert.Equal(t, uint(100), *newVirtualServer.Properties.Connection.KeepaliveTimeout)
	assert.Equal(t, uint(32768), *newVirtualServer.Properties.Connection.MaxClientBuffer)
	assert.Equal(t, uint(32768), *newVirtualServer.Properties.Connection.MaxServerBuffer)
	assert.Equal(t, uint(12), *newVirtualServer.Properties.Connection.MaxTransactionDuration)
	assert.Equal(t, uint(100), *newVirtualServer.Properties.Connection.Timeout)

	assert.Equal(t, uint(0), *newVirtualServer.Properties.HTTP2.ConnectTimeout)
	assert.Equal(t, uint(2048), *newVirtualServer.Properties.HTTP2.DataFrameSize)
	assert.Equal(t, true, *newVirtualServer.Properties.HTTP2.Enabled)
	assert.Equal(t, uint(8192), *newVirtualServer.Properties.HTTP2.HeaderTableSize)
	assert.Equal(t, true, *newVirtualServer.Properties.HTTP2.HeadersIndexDefault)
	assert.Equal(t, uint(240), *newVirtualServer.Properties.HTTP2.IdleTimeoutNoStreams)
	assert.Equal(t, uint(300), *newVirtualServer.Properties.HTTP2.IdleTimeoutOpenStreams)
	assert.Equal(t, uint(800), *newVirtualServer.Properties.HTTP2.MaxConcurrentStreams)
	assert.Equal(t, uint(16384), *newVirtualServer.Properties.HTTP2.MaxFrameSize)
	assert.Equal(t, uint(128), *newVirtualServer.Properties.HTTP2.MaxHeaderPadding)
	assert.Equal(t, true, *newVirtualServer.Properties.HTTP2.MergeCookieHeaders)
	assert.Equal(t, uint(32768), *newVirtualServer.Properties.HTTP2.StreamWindowSize)
}

func getVirtualServer(name string, t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	client.WorkWithConfigurationResources()

	virtualServer := VirtualServer{}
	err = client.GetByName("virtual_servers", name, &virtualServer)
	if err != nil {
		t.Fatal("Error getting a resource: ", err)
	}
	log.Println("Found Virtual Server: ", virtualServer)

	assert.Equal(t, true, virtualServer.Properties.Basic.AddClusterIP)
	assert.Equal(t, true, virtualServer.Properties.Basic.AddXForwarded)
	assert.Equal(t, true, virtualServer.Properties.Basic.AddXForwardedProto)
	assert.Equal(t, true, virtualServer.Properties.Basic.CloseWithRst)
	assert.Equal(t, uint(30), virtualServer.Properties.Basic.ConnectTimeout)
	assert.Equal(t, true, virtualServer.Properties.Basic.Enabled)
	assert.Equal(t, true, virtualServer.Properties.Basic.ListenOnAny)
	assert.Equal(t, uint(0), virtualServer.Properties.Basic.MSS)
	assert.Equal(t, "Created by go-brocade-vtm test", virtualServer.Properties.Basic.Note)
	assert.Equal(t, "test-pool", virtualServer.Properties.Basic.Pool)
	assert.Equal(t, uint(80), virtualServer.Properties.Basic.Port)
	assert.Equal(t, "http", virtualServer.Properties.Basic.Protocol)
	assert.Equal(t, true, virtualServer.Properties.Basic.SoNagle)
	assert.Equal(t, "none", virtualServer.Properties.Basic.SslClientCertHeaders)
	assert.Equal(t, false, virtualServer.Properties.Basic.SslDecrypt)
	assert.Equal(t, "use_default", virtualServer.Properties.Basic.SslHonorFallbackScsv)
	assert.Equal(t, false, virtualServer.Properties.Basic.Transparent)

	assert.Equal(t, true, *virtualServer.Properties.Connection.Keepalive)
	assert.Equal(t, uint(100), *virtualServer.Properties.Connection.KeepaliveTimeout)
	assert.Equal(t, uint(32768), *virtualServer.Properties.Connection.MaxClientBuffer)
	assert.Equal(t, uint(32768), *virtualServer.Properties.Connection.MaxServerBuffer)
	assert.Equal(t, uint(12), *virtualServer.Properties.Connection.MaxTransactionDuration)
	assert.Equal(t, uint(100), *virtualServer.Properties.Connection.Timeout)

	assert.Equal(t, uint(0), *virtualServer.Properties.HTTP2.ConnectTimeout)
	assert.Equal(t, uint(2048), *virtualServer.Properties.HTTP2.DataFrameSize)
	assert.Equal(t, true, *virtualServer.Properties.HTTP2.Enabled)
	assert.Equal(t, uint(8192), *virtualServer.Properties.HTTP2.HeaderTableSize)
	assert.Equal(t, true, *virtualServer.Properties.HTTP2.HeadersIndexDefault)
	assert.Equal(t, uint(240), *virtualServer.Properties.HTTP2.IdleTimeoutNoStreams)
	assert.Equal(t, uint(300), *virtualServer.Properties.HTTP2.IdleTimeoutOpenStreams)
	assert.Equal(t, uint(800), *virtualServer.Properties.HTTP2.MaxConcurrentStreams)
	assert.Equal(t, uint(16384), *virtualServer.Properties.HTTP2.MaxFrameSize)
	assert.Equal(t, uint(128), *virtualServer.Properties.HTTP2.MaxHeaderPadding)
	assert.Equal(t, true, *virtualServer.Properties.HTTP2.MergeCookieHeaders)
	assert.Equal(t, uint(32768), *virtualServer.Properties.HTTP2.StreamWindowSize)
}

func deleteVirtualServer(name string, t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	err = client.Delete("virtual_servers", name)
	if err != nil {
		t.Fatal("Error deleting a resource: ", err)
	} else {
		log.Printf("Resource %s deleted", name)
	}
}
