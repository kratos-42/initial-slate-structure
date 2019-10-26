package storage

import (
  "bufio"
  "fmt"
	"os"
  "reflect"
  utils "docParser/utils"
)

// Write stored data to file.
func WriteEndpoints(filename string, endpoints *Endpoints) {
  // `APPEND` flag used to avoid ruining previously work by mistake.
  fileToWrite, err := os.OpenFile(filename, os.O_CREATE | os.O_WRONLY, 0644)
  utils.CheckError(err)

  defer fileToWrite.Close()

  w := bufio.NewWriter(fileToWrite)
  endpoints_list := *endpoints

  _, err = fmt.Fprintf(w, skeleton_template)
  utils.CheckError(err)

  for i := 0; i < len(endpoints_list); i++ {
    /* NOTE:
      Writing order:
        - Endpoint subsection
        - Request and response examples
        - Query string options (middlewares):
          - Middlewares tables
          - Query string example and return status codes
    */

    v := reflect.ValueOf(endpoints_list[i])

    _, err := fmt.Fprintf(w, "## %s %s\n\n", v.FieldByName("request"), v.FieldByName("path"))
    utils.CheckError(err)

    _, err = fmt.Fprintf(w, example_template)
    utils.CheckError(err)

    _, err = fmt.Fprintf(w, query_string_options_template)
    utils.CheckError(err)

    middlewares := v.FieldByName("middlewares")

    for i := 0; i < middlewares.Len(); i++ {
      switch {
        case reflect.DeepEqual("include", middlewares.Index(i).String()):
            _, err = fmt.Fprintf(w, include_table_template)
            utils.CheckError(err)
        case reflect.DeepEqual("sort", middlewares.Index(i).String()):
          _, err = fmt.Fprintf(w, sort_table_template)
          utils.CheckError(err)

        case reflect.DeepEqual("filter", middlewares.Index(i).String()):
          _, err = fmt.Fprintf(w, filter_table_template)
          utils.CheckError(err)

        case reflect.DeepEqual("paginate", middlewares.Index(i).String()):
          _, err = fmt.Fprintf(w, paginate_table_template)
          utils.CheckError(err)

        case reflect.DeepEqual("authorization", middlewares.Index(i).String()):
          // Do nothing.

        default:
          _, err = fmt.Fprintf(w, "Not handled middleware: %+v\n", middlewares.Index(i).String())
          utils.CheckError(err)
      }
    }

    _, err = fmt.Fprintf(w, after_middlewares)
    utils.CheckError(err)

    w.Flush()
  }
}
