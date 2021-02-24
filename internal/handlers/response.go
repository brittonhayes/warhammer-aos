package handlers


// Response is the default API response
// format
type Response struct {
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
}

