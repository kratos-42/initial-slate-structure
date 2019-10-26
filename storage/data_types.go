package storage

// Endpoint struct.
type Endpoint struct {
  request string
  path string
  middlewares []string
}

type Endpoints []Endpoint
