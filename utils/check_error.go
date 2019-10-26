package utils

// Check error function.
func CheckError(e error) {
  if e != nil {
    panic(e)
  }
}
