package ngrok

import "time"

type Request struct {
	URI        string    `json:"uri"`
	ID         string    `json:"id"`
	TunnelName string    `json:"tunnel_name"`
	RemoteAddr string    `json:"remote_addr"`
	Start      time.Time `json:"start"`
	Duration   int       `json:"duration"`
	Request    struct {
		Method  string `json:"method"`
		Proto   string `json:"proto"`
		Headers struct {
			Accept          []string `json:"Accept"`
			AcceptEncoding  []string `json:"Accept-Encoding"`
			AcceptLanguage  []string `json:"Accept-Language"`
			Host            []string `json:"Host"`
			Referer         []string `json:"Referer"`
			SecFetchDest    []string `json:"Sec-Fetch-Dest"`
			SecFetchMode    []string `json:"Sec-Fetch-Mode"`
			SecFetchSite    []string `json:"Sec-Fetch-Site"`
			Te              []string `json:"Te"`
			UserAgent       []string `json:"User-Agent"`
			XForwardedFor   []string `json:"X-Forwarded-For"`
			XForwardedProto []string `json:"X-Forwarded-Proto"`
		} `json:"headers"`
		URI string `json:"uri"`
		Raw string `json:"raw"`
	} `json:"request"`
	Response struct {
		Status     string `json:"status"`
		StatusCode int    `json:"status_code"`
		Proto      string `json:"proto"`
		Headers    struct {
			ContentLength []string `json:"Content-Length"`
			ContentType   []string `json:"Content-Type"`
		} `json:"headers"`
		Raw string `json:"raw"`
	} `json:"response"`
}

type RequestsResponse struct {
	Requests []Request `json:"requests"`
	URI      string    `json:"uri"`
}
