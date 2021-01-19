package httpclient

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

type dataSourceTransport struct {
	headers   map[string]string
	transport *http.Transport
}

func newTransport(ds *backend.DataSourceInstanceSettings) (*dataSourceTransport, error) {
	tlsConfig, err := tlsConfig(ds)
	if err != nil {
		return nil, err
	}

	tlsConfig.Renegotiation = tls.RenegotiateFreelyAsClient

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
		Proxy:           http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
	}

	customHeaders := customHeaders(ds)

	if ds.BasicAuthEnabled {
		user := ds.BasicAuthUser
		password := ds.DecryptedSecureJSONData["basicAuthPassword"]
		customHeaders["Authorization"] = basicAuth(user, password)
	}

	return &dataSourceTransport{
		headers:   customHeaders,
		transport: transport,
	}, nil
}

func (d *dataSourceTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, value := range d.headers {
		req.Header.Set(key, value)
	}
	return d.transport.RoundTrip(req)
}

func basicAuth(user string, password string) string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+password))
}

func customHeaders(ds *backend.DataSourceInstanceSettings) map[string]string {
	headers := make(map[string]string)

	var jsonData map[string]string
	if err := json.Unmarshal(ds.JSONData, &jsonData); err != nil {
		return headers
	}

	index := 1

	for {
		keyName := fmt.Sprintf("httpHeaderName%d", index)
		keyValue := fmt.Sprintf("httpHeaderValue%d", index)

		key, ok := jsonData[keyName]
		if !ok {
			break
		}

		if key != "" {
			if value, ok := ds.DecryptedSecureJSONData[keyValue]; ok {
				headers[key] = value
			}
		}

		index++
	}

	return headers
}

func tlsConfig(ds *backend.DataSourceInstanceSettings) (*tls.Config, error) {
	var jsonData struct {
		TLSClientAuth     bool `json:"tlsAuth"`
		TLSAuthWithCACert bool `json:"tlsAuthWithCACert"`
		TLSSkipVerify     bool `json:"tlsAuthWithCACert"`
	}
	if err := json.Unmarshal(ds.JSONData, &jsonData); err != nil {
		return nil, err
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: jsonData.TLSSkipVerify,
	}

	if jsonData.TLSClientAuth || jsonData.TLSAuthWithCACert {
		decrypted := ds.DecryptedSecureJSONData
		if jsonData.TLSAuthWithCACert && len(decrypted["tlsCACert"]) > 0 {
			caPool := x509.NewCertPool()
			ok := caPool.AppendCertsFromPEM([]byte(decrypted["tlsCACert"]))
			if !ok {
				return nil, errors.New("Failed to parse TLS CA PEM certificate")
			}
			tlsConfig.RootCAs = caPool
		}

		if jsonData.TLSClientAuth {
			cert, err := tls.X509KeyPair([]byte(decrypted["tlsClientCert"]), []byte(decrypted["tlsClientKey"]))
			if err != nil {
				return nil, err
			}
			tlsConfig.Certificates = []tls.Certificate{cert}
		}
	}

	return tlsConfig, nil
}
