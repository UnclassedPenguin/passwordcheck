package main

import (
  "os"
  "fmt"
  "bytes"
  "strings"
  "net/http"
  "io/ioutil"
  "crypto/sha1"
  "encoding/hex"
)


func hashIt(pass string) string {
  hash := sha1.New()
  hash.Write([]byte(pass))
  return hex.EncodeToString(hash.Sum(nil))
}

func getHashes(firstFive string) []string {
  url := "https://api.pwnedpasswords.com/range/" + firstFive

  result, err := http.Get(url)
  if err != nil {
    fmt.Println("Error in http.Get: ", err)
    os.Exit(1)
  }

  defer result.Body.Close()

  body, err := ioutil.ReadAll(result.Body)
  if err != nil {
    fmt.Println("Error reading body: ", err)
    os.Exit(1)
  }

  hashes := bytes.NewBuffer(body).String()
  hashesSplit := strings.Split(hashes, "\n")

  return hashesSplit
}

func checkHashes(hashes []string, last string) string {
  last = strings.ToUpper(last)

  for _, hash := range hashes {
    split := strings.Split(hash, ":")
    if split[0] == last {
      return hash
    }
  }
  return ""
}

func main() {

  // Get password to check from user
  var password string

  fmt.Println("What Password do you want to check?")
  fmt.Print(" > ")
  fmt.Scan(&password)

  // Hash password using sha-1 then split it into the first five
  // characters, and the remaining characters.
  hash := hashIt(password)
  firstFive := hash[0:5]
  last := hash[5:]

  // Send the first five characters of the hash to the HaveIBeenPwned api
  // and save the resulting []string to hashes
  hashes := getHashes(firstFive)

  // Check if any of the returned hashes match the last characters of our
  // hashed password, if they do that means the password was found in
  // at least one of the breaches that the database is aware of. 
  // Print the hash and number of times it was found.
  found := checkHashes(hashes, last)
  if found != "" {
    split := strings.Split(found, ":")
    fmt.Println("\nPassword found!\n")
    fmt.Println("Password:", password)
    fmt.Println("Hash:", strings.ToUpper(hash))
    fmt.Println("Found:", split[1])
  } else {
    fmt.Println("\nPassword not found. Nice!\n")
    fmt.Println("Password:", password)
    fmt.Println("Hash:", strings.ToUpper(hash))
    fmt.Println("Found: 0")
  }

}
