package main

import (
  "bufio"
  "log"
	"os"
  "regexp"
  "strings"
  utils "docParser/utils"
  "time"
  . "docParser/storage"
)

func main() {
  if len(os.Args) < 2 {
    panic("You must provide a file where the API endpoints are stored.\n" +
          "Usage: > go run kratos.go inputFilename [outputFilename]")
  }

  start := time.Now()

  // Prepare OpenApi specification file
  dir, err := os.Getwd()
  utils.CheckError(err)

  var outputFilename string

  if len(os.Args) >= 3 {
    outputFilename = os.Args[2]
  } else {
    projectName := strings.Split(dir, "/")

    // e.g.: With /users/projects/foobar, it gets `foobar`.
    outputFilename = projectName[len(projectName)-1]
  }

  mdFile := outputFilename + ".md"

  // Prepare to read file
	fileToRead, err := os.Open(os.Args[1])
	utils.CheckError(err)

	defer fileToRead.Close()

  fileInfo, err := fileToRead.Stat()
  utils.CheckError(err)

  buffer := make([]byte, fileInfo.Size())
  _, err = fileToRead.Read(buffer)

  // Read line by line.
  buf := make([]byte, 2)
  scanner := bufio.NewScanner(strings.NewReader(string(buffer)))
  scanner.Split(bufio.ScanLines)
  scanner.Buffer(buf, bufio.MaxScanTokenSize)
  token := regexp.MustCompile("[ ]+")

  // Data struct to store.
  var endpoints = Endpoints{}

  for scanner.Scan() {
    if strings.Contains(scanner.Text(), "Error") {
      continue;
    }

    splitted := token.Split(scanner.Text(), len(scanner.Text()))
    StoreEndpoints(splitted, &endpoints)
  }

  // Write to `.md` file.
  WriteEndpoints(mdFile, &endpoints)

  elapsed := time.Since(start)
  log.Printf("Execution time: %s", elapsed)
}
