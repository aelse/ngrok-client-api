package ngrok

type Tunnel struct {
	Name      string `json:"name"`
	URI       string `json:"uri"`
	PublicURL string `json:"public_url"`
	Proto     string `json:"proto"`
	Config    struct {
		Addr    string `json:"addr"`
		Inspect bool   `json:"inspect"`
	} `json:"config"`
	Metrics struct {
		Conns struct {
			Count  int     `json:"count"`
			Gauge  int     `json:"gauge"`
			Rate1  float64 `json:"rate1"`
			Rate5  float64 `json:"rate5"`
			Rate15 float64 `json:"rate15"`
			P50    float64 `json:"p50"`
			P90    float64 `json:"p90"`
			P95    float64 `json:"p95"`
			P99    float64 `json:"p99"`
		} `json:"conns"`
		HTTP struct {
			Count  int     `json:"count"`
			Rate1  float64 `json:"rate1"`
			Rate5  float64 `json:"rate5"`
			Rate15 float64 `json:"rate15"`
			P50    float64 `json:"p50"`
			P90    float64 `json:"p90"`
			P95    float64 `json:"p95"`
			P99    float64 `json:"p99"`
		} `json:"http"`
	} `json:"metrics"`
}

type TunnelsResponse struct {
	Tunnels []Tunnel `json:"tunnels"`
	URI     string   `json:"uri"`
}
