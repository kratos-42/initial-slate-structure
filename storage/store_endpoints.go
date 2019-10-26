package storage

import (
  "strings"
)

// Save endpoints data in `endpoint` struct.
func StoreEndpoints(data []string, endpoints *Endpoints) {
  endpoint := Endpoint{}

  for i := 0; i <= len(data)-1; i++ {
    switch i {
      case 0:
        endpoint.request = data[0]
        break
      case 1:
        // Evaluate `parameter path` existence
        if index := strings.Index(data[1], ":"); index != -1 {
          parameterPrefixed := strings.Replace(data[1], ":", "{", 1)
          var parameter string

          // Evaluate if "parameter path" is not end of `path`, e.g., GET /users/:id/jobs
          if slashIndex := strings.Index(parameterPrefixed[index:], "/"); slashIndex != -1 {
            insertionIndex := index + slashIndex
            parameter = parameterPrefixed[:insertionIndex] + "}" + parameterPrefixed[insertionIndex:]

          } else {
              parameter = parameterPrefixed[:len(parameterPrefixed)] + "}"
          }

          endpoint.path = parameter

        } else {
          endpoint.path = data[1][:len(data[1])]
        }
        break
      default:
        endpoint.middlewares = append(endpoint.middlewares, data[i])
        break
    }

  }

  *endpoints = append(*endpoints, endpoint)
}
