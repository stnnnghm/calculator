package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// MathRequest is a request of math operation
type MathRequest struct {
	Op    string  `json:"op"`
	Left  float64 `json:"left"`
	Right float64 `json:"right"`
}

// MathResponse is a response to MathRequest
type MathResponse struct {
	Error  string  `json:"error"`
	Result float64 `json:"result"`
}

// mathHandler does a math operation
func mathHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Decode request
	dec := json.NewDecoder(r.Body) // Create a new decoder
	req := &MathRequest{}          // Create an empty request

	if err := dec.Decode(req); err != nil { // Decode the request
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Do work
	resp := &MathResponse{} // Create an empty response
	switch req.Op {
	case "+":
		resp.Result = req.Left + req.Right
	case "-":
		resp.Result = req.Left - req.Right
	case "*":
		resp.Result = req.Left * req.Right
	case "/":
		if req.Right == 0.0 {
			resp.Error = "division by 0"
		} else {
			resp.Result = req.Left / req.Right
		}
	default:
		resp.Error = fmt.Sprintf("unknown operation: %s", req.Op)
	}

	// Encode and return result
	w.Header().Set("Content-Type", "application/json")
	if resp.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	enc := json.NewEncoder(w)                // Create a new encoder
	if err := enc.Encode(resp); err != nil { // Encoding the response
		// Can't return error to client here
		log.Printf("cant encode %v - %s", resp, err)
	}
}

func main() {
	http.HandleFunc("/math", mathHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
