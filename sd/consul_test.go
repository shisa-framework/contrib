package sd

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
	"testing"

	"github.com/ansel1/merry"
	consul "github.com/hashicorp/consul/api"
	"github.com/stretchr/testify/assert"
)

const (
	testServiceName = "service"
	testAddr        = "127.0.0.1:9000"
)

func TestNewConsul(t *testing.T) {
	ac := &consul.Client{}

	c := NewConsul(ac)

	assert.NotNil(t, c)
	assert.NotNil(t, c.agent)
	assert.NotNil(t, c.health)
}

func TestConsulSDRegisterBadAddr(t *testing.T) {
	freg := &FakeconsulRegistry{}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	e := c.Register(testServiceName, &url.URL{Host: "bad addr"})

	assert.Error(t, e)
	freg.AssertServiceRegisterNotCalled(t)
}

func TestConsulSDRegisterBadPort(t *testing.T) {
	freg := &FakeconsulRegistry{}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	e := c.Register(testServiceName, &url.URL{Host: "127.0.0.1:badport"})

	assert.Error(t, e)
	freg.AssertServiceRegisterNotCalled(t)
}

func TestConsulSDRegisterServiceRegisterError(t *testing.T) {
	terr := merry.New("service registry error")
	freg := &FakeconsulRegistry{
		ServiceRegisterHook: func(*consul.AgentServiceRegistration) error {
			return terr
		},
	}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	e := c.Register(testServiceName, &url.URL{Host: testAddr})

	assert.Error(t, e)
	freg.AssertServiceRegisterCalledOnce(t)
}

func TestConsulSDRegisterSuccess(t *testing.T) {
	freg := &FakeconsulRegistry{
		ServiceRegisterHook: func(*consul.AgentServiceRegistration) error {
			return nil
		},
	}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	e := c.Register(testServiceName, &url.URL{Host: testAddr})

	assert.NoError(t, e)
	freg.AssertServiceRegisterCalledOnce(t)
}

func TestConsulSDDeregisterError(t *testing.T) {
	terr := merry.New("service deregister error")
	freg := &FakeconsulRegistry{
		ServiceDeregisterHook: func(string) error {
			return terr
		},
	}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	e := c.Deregister(testServiceName)

	assert.Error(t, e)
	freg.AssertServiceDeregisterCalledOnceWith(t, testServiceName)
}

func TestConsulSDDeregisterSuccess(t *testing.T) {
	freg := &FakeconsulRegistry{
		ServiceDeregisterHook: func(string) error {
			return nil
		},
	}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	e := c.Deregister(testServiceName)

	assert.NoError(t, e)
	freg.AssertServiceDeregisterCalledOnceWith(t, testServiceName)
}

func TestConsulSDResolveError(t *testing.T) {
	terr := merry.New("service resolve error")
	freg := &FakeconsulRegistry{}
	fres := &FakeconsulResolver{
		ServiceHook: func(string, string, bool, *consul.QueryOptions) ([]*consul.ServiceEntry, *consul.QueryMeta, error) {
			return []*consul.ServiceEntry{}, nil, terr
		},
	}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	n, e := c.Resolve(testServiceName)

	assert.Empty(t, n)
	assert.Error(t, e)
	fres.AssertServiceCalledOnceWith(t, testServiceName, "", true, nil)
}

func TestConsulSDResolveNodeAddress(t *testing.T) {
	host, sport, e := net.SplitHostPort(testAddr)
	if e != nil {
		t.Error(e)
	}
	port, e := strconv.Atoi(sport)
	if e != nil {
		t.Error(e)
	}

	n := &consul.Node{Address: host}
	s := &consul.AgentService{Port: port}
	freg := &FakeconsulRegistry{}
	fres := &FakeconsulResolver{
		ServiceHook: func(string, string, bool, *consul.QueryOptions) ([]*consul.ServiceEntry, *consul.QueryMeta, error) {
			return []*consul.ServiceEntry{{Node: n, Service: s}}, nil, nil
		},
	}
	exp := []string{testAddr}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	res, e := c.Resolve(testServiceName)

	assert.Equal(t, exp, res)
	assert.NoError(t, e)
	fres.AssertServiceCalledOnceWith(t, testServiceName, "", true, nil)
}

func TestConsulSDResolveServiceAddress(t *testing.T) {
	host, sport, e := net.SplitHostPort(testAddr)
	if e != nil {
		t.Error(e)
	}
	port, e := strconv.Atoi(sport)
	if e != nil {
		t.Error(e)
	}

	n := &consul.Node{}
	s := &consul.AgentService{Address: host, Port: port}
	freg := &FakeconsulRegistry{}
	fres := &FakeconsulResolver{
		ServiceHook: func(string, string, bool, *consul.QueryOptions) ([]*consul.ServiceEntry, *consul.QueryMeta, error) {
			return []*consul.ServiceEntry{{Node: n, Service: s}}, nil, nil
		},
	}
	exp := []string{testAddr}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	res, e := c.Resolve(testServiceName)

	assert.Equal(t, exp, res)
	assert.NoError(t, e)
	fres.AssertServiceCalledOnceWith(t, testServiceName, "", true, nil)
}

func TestConsulSDAddCheckGRPC(t *testing.T) {
	u := &url.URL{
		Scheme:   "grpc",
		Host:     "127.0.0.1:9000",
		RawQuery: "interval=5s",
	}
	freg := &FakeconsulRegistry{
		CheckRegisterHook: func(*consul.AgentCheckRegistration) error {
			return nil
		},
	}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	err := c.AddCheck(testServiceName, u)

	assert.NoError(t, err)
	freg.AssertCheckRegisterCalledOnce(t)
}

func TestConsulSDAddCheckGRPCMissingHost(t *testing.T) {
	u := &url.URL{
		Scheme: "grpc",
	}
	freg := &FakeconsulRegistry{}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	err := c.AddCheck(testServiceName, u)

	assert.Error(t, err)
	freg.AssertCheckDeregisterNotCalled(t)
}

func TestConsulSDAddCheckGRPCMissingInterval(t *testing.T) {
	u := &url.URL{
		Scheme: "grpc",
		Host:   "localhost:9000",
	}
	freg := &FakeconsulRegistry{}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	err := c.AddCheck(testServiceName, u)

	assert.Error(t, err)
	freg.AssertCheckDeregisterNotCalled(t)
}

func TestConsulSDAddCheckDockerMissingKey(t *testing.T) {
	u := &url.URL{
		Scheme: "docker",
	}
	freg := &FakeconsulRegistry{}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	err := c.AddCheck(testServiceName, u)

	assert.Error(t, err)
	freg.AssertCheckDeregisterNotCalled(t)
}

func TestConsulSDAddCheckDocker(t *testing.T) {
	u := &url.URL{
		Scheme:   "docker",
		Host:     "localhost:9000",
		RawQuery: "dockercontainerid=d858g8ergj&args=\"run command\"&interval=5s",
	}
	freg := &FakeconsulRegistry{
		CheckRegisterHook: func(*consul.AgentCheckRegistration) error {
			return nil
		},
	}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	err := c.AddCheck(testServiceName, u)

	assert.NoError(t, err)
	freg.AssertCheckRegisterCalledOnce(t)
}

func TestConsulSDAddCheckScriptMissingKey(t *testing.T) {
	u := &url.URL{
		Scheme: "script",
	}
	freg := &FakeconsulRegistry{}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	err := c.AddCheck(testServiceName, u)

	assert.Error(t, err)
	freg.AssertCheckDeregisterNotCalled(t)
}

func TestConsulSDAddCheckScript(t *testing.T) {
	u := &url.URL{
		Scheme:   "script",
		Host:     "localhost:9000",
		RawQuery: "args=\"run command\"&interval=5s",
	}
	freg := &FakeconsulRegistry{
		CheckRegisterHook: func(*consul.AgentCheckRegistration) error {
			return nil
		},
	}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	err := c.AddCheck(testServiceName, u)

	assert.NoError(t, err)
	freg.AssertCheckRegisterCalledOnce(t)
}

func TestConsulSDAddCheckTTL(t *testing.T) {
	u := &url.URL{
		Scheme:   "ttl",
		Host:     "localhost:9000",
		RawQuery: "ttl=5s",
	}
	freg := &FakeconsulRegistry{
		CheckRegisterHook: func(c *consul.AgentCheckRegistration) error {
			assert.NotNil(t, c.TTL)
			return nil
		},
	}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	err := c.AddCheck(testServiceName, u)

	assert.NoError(t, err)
	freg.AssertCheckRegisterCalledOnce(t)
}

func TestConsulSDAddCheckTCP(t *testing.T) {
	u := &url.URL{
		Scheme:   "tcp",
		Host:     "localhost:9000",
		RawQuery: "interval=5s",
	}
	freg := &FakeconsulRegistry{
		CheckRegisterHook: func(*consul.AgentCheckRegistration) error {
			return nil
		},
	}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	err := c.AddCheck(testServiceName, u)

	assert.NoError(t, err)
	freg.AssertCheckRegisterCalledOnce(t)
}

func TestConsulSDAddCheckTCPMissingHost(t *testing.T) {
	u := &url.URL{
		Scheme: "tcp",
	}
	freg := &FakeconsulRegistry{}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	err := c.AddCheck(testServiceName, u)

	assert.Error(t, err)
	freg.AssertCheckDeregisterNotCalled(t)
}

func TestConsulSDAddCheckTCPMissingKey(t *testing.T) {
	u := &url.URL{
		Scheme: "tcp",
		Host:   "localhost:9000",
	}
	freg := &FakeconsulRegistry{}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	err := c.AddCheck(testServiceName, u)

	assert.Error(t, err)
	freg.AssertCheckDeregisterNotCalled(t)
}

func TestConsulSDAddCheckHTTP(t *testing.T) {
	u := &url.URL{
		Scheme:   "http",
		Host:     "localhost:9000",
		RawQuery: "interval=5s",
	}
	freg := &FakeconsulRegistry{
		CheckRegisterHook: func(*consul.AgentCheckRegistration) error {
			return nil
		},
	}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	err := c.AddCheck(testServiceName, u)

	assert.NoError(t, err)
	freg.AssertCheckRegisterCalledOnce(t)
}

func TestConsulSDAddCheckHTTPS(t *testing.T) {
	u := &url.URL{
		Scheme:   "https",
		Host:     "localhost:9000",
		RawQuery: "interval=5s",
	}
	freg := &FakeconsulRegistry{
		CheckRegisterHook: func(*consul.AgentCheckRegistration) error {
			return nil
		},
	}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	err := c.AddCheck(testServiceName, u)

	assert.NoError(t, err)
	freg.AssertCheckRegisterCalledOnce(t)
}

func TestConsulSDAddCheckHTTPSMissingKey(t *testing.T) {
	u := &url.URL{
		Scheme: "https",
		Host:   "localhost:9000",
	}
	freg := &FakeconsulRegistry{}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	err := c.AddCheck(testServiceName, u)

	assert.Error(t, err)
	freg.AssertCheckDeregisterNotCalled(t)
}

func TestConsulSDAddCheckHTTPSMissingURL(t *testing.T) {
	u := &url.URL{
		Scheme: "https",
	}
	freg := &FakeconsulRegistry{}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	err := c.AddCheck(testServiceName, u)

	assert.Error(t, err)
	freg.AssertCheckDeregisterNotCalled(t)
}

func TestConsulSDAddCheckCheckRegisterError(t *testing.T) {
	terr := merry.New("error in consul registration")
	u := &url.URL{
		Scheme:   "https",
		Host:     "localhost:9000",
		RawQuery: "interval=5s",
	}
	freg := &FakeconsulRegistry{
		CheckRegisterHook: func(*consul.AgentCheckRegistration) error {
			return terr
		},
	}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	err := c.AddCheck(testServiceName, u)

	assert.Error(t, err)
	freg.AssertCheckRegisterCalledOnce(t)

}

func TestConsulSDRemoveChecksError(t *testing.T) {
	hcname := fmt.Sprintf("%s-healthcheck", testServiceName)
	terr := merry.New("service check deregister error")
	freg := &FakeconsulRegistry{
		CheckDeregisterHook: func(string) error {
			return terr
		},
	}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	e := c.RemoveChecks(testServiceName)

	assert.Error(t, e)
	freg.AssertCheckDeregisterCalledOnceWith(t, hcname)
}

func TestConsulSDRemoveChecksSuccess(t *testing.T) {
	hcname := fmt.Sprintf("%s-healthcheck", testServiceName)
	freg := &FakeconsulRegistry{
		CheckDeregisterHook: func(string) error {
			return nil
		},
	}
	fres := &FakeconsulResolver{}
	c := consulSD{
		agent:  freg,
		health: fres,
	}

	e := c.RemoveChecks(testServiceName)

	assert.NoError(t, e)
	freg.AssertCheckDeregisterCalledOnceWith(t, hcname)
}

func TestPopQueryString(t *testing.T) {
	k := "k"
	sv := "v"
	v := make(url.Values, 1)
	v[k] = []string{sv}

	r := popQueryString(v, k)

	assert.Equal(t, sv, r)
}

func TestPopQueryStringNotFound(t *testing.T) {
	k := "k"
	v := make(url.Values, 1)

	r := popQueryString(v, k)

	assert.Equal(t, "", r)
}

func TestPopQueryBool(t *testing.T) {
	k := "k"
	b := true
	bv := []string{fmt.Sprint(b)}
	v := make(url.Values, 1)
	v[k] = bv

	r := popQueryBool(v, k)

	assert.Equal(t, b, r)
}

func TestPopQueryBoolNotFound(t *testing.T) {
	k := "k"
	v := make(url.Values, 1)

	r := popQueryBool(v, k)

	assert.Equal(t, false, r)
}

func TestPopQuerySlice(t *testing.T) {
	k := "k"
	sv := []string{"v1", "v2"}
	v := make(url.Values, 1)
	v[k] = sv

	r := popQuerySlice(v, k)

	assert.Equal(t, sv, r)
}

func TestPopQuerySliceNotFound(t *testing.T) {
	k := "k"
	v := make(url.Values, 1)

	r := popQuerySlice(v, k)

	assert.Equal(t, []string(nil), r)
}

func TestValidateKeyExists(t *testing.T) {
	k := "k"
	sv := []string{"v1", "v2"}
	v := make(url.Values, 1)
	v[k] = sv

	r := validateKeysExist(v, k)

	assert.Nil(t, r)
}

func TestValidateKeyExistsNotFound(t *testing.T) {
	k := "k"
	v := make(url.Values, 1)

	r := validateKeysExist(v, k)

	assert.Error(t, r)
}
