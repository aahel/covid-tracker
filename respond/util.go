package respond

type payload struct {
	data interface{}
	meta interface{}
}

// Holds the error status of the request informations ,
// will only be there in case of error
// swagger:model
type Meta struct {
	// Error status code
	// example: 901
	Status int `json:"status_code"`
	// Error message
	// example: No mapping exists
	Message string `json:"error_message,omitempty"`
	// For future use ,not implemented ye
	Error interface{} `json:"error,omitempty"`
}

// TODO: can be improved

// WrapPayload is used to create a generic payload for the data
// and the metadata passed
func WrapPayload(p payload) map[string]interface{} {
	x := make(map[string]interface{})
	if p.data != nil {
		x["data"] = p.data
	}
	if p.meta != nil {
		x["error"] = p.meta
	}

	return x
}
