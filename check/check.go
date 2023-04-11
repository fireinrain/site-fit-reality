package check

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

type SupportedChecker struct{}

// IsSupportTls13
//
//	@Description: 判断是否支持tls1.3
//	@receiver receiver
//	@param siteUrl
//	@return bool
//	@return error
func (receiver *SupportedChecker) IsSupportTls13(siteUrl string) (bool, error) {
	hostname := siteUrl
	port := "443"
	url := fmt.Sprintf("https://%s:%s/", hostname, port)

	tlsConfig := &tls.Config{
		// Only enable TLSv1.3
		MinVersion: tls.VersionTLS13,
		MaxVersion: tls.VersionTLS13,
	}

	// Create an HTTP client with the custom TLS config
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// Make a request to the server
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return false, err
	}
	defer resp.Body.Close()

	// Check if the server supports H2
	//if resp.ProtoMajor == 2 {
	//	fmt.Println("Server supports HTTP/2 (H2)")
	//} else {
	//	fmt.Println("Server does not support HTTP/2 (H2)")
	//}

	// Check if the server supports TLSv1.3
	if resp.TLS.Version == tls.VersionTLS13 {
		fmt.Println("TLSv1.3 is supported")
		return true, nil

	}
	fmt.Println("TLSv1.3 is not supported")
	return false, nil
}

// IsSupportHttp2
//
//	@Description: 判断是否支持http2
//	@receiver receiver
//	@param siteUrl
//	@return bool
//	@return error
func (receiver *SupportedChecker) IsSupportHttp2(siteUrl string) (bool, error) {
	hostname := siteUrl
	port := "443"
	url := fmt.Sprintf("https://%s:%s/", hostname, port)
	client := &http.Client{}
	resp, err := client.Get(url)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return false, err
	}

	defer resp.Body.Close()

	if resp.ProtoMajor == 2 {
		fmt.Println("HTTP/2 is supported.")
		return true, nil
	}
	fmt.Println("HTTP/2 is not supported.")
	return false, nil
}

// DoAllCheck
//
//	@Description: 执行所有检查
//	@receiver receiver
//	@param siteUrl
//	@return bool
func (receiver *SupportedChecker) DoAllCheck(siteUrl string) bool {
	tls13, err := receiver.IsSupportTls13(siteUrl)
	if err != nil {
		fmt.Println("Error:", err)
	}
	http2, err := receiver.IsSupportHttp2(siteUrl)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return tls13 && http2
}
