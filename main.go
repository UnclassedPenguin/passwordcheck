//-------------------------------------------------------------------------------
//-------------------------------------------------------------------------------
//
// Tyler(UnclassedPenguin) Password Checker - Now in Go! 2022
//
// Author: Tyler(UnclassedPenguin)
//    URL: https://unclassed.ca
// GitHub: https://github.com/UnclassedPenguin
//
//-------------------------------------------------------------------------------
//-------------------------------------------------------------------------------

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

// Takes a string and hashes it to sha-1
func hashIt(pass string) string {
  hash := sha1.New()
  hash.Write([]byte(pass))
  return hex.EncodeToString(hash.Sum(nil))
}

// Sends a request to the Have I Been Pwned api and then organizes
// the response into a slice of strings.
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

// Checks the returned list to see if there are any matching
// The previously hashed password. If a match is found,
// returns the line. If no match is found, returns an
// empty string.
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
  // *
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

/*
 *
This is the really cool thing about how this works. You hash the password
on your end, and only send the first five characters of that hash to the 
Have I Been Pwned api. He then returns back all hashes that has those
same first five characters. You then compare on your end again if the 
remaining characters of your hashed password match any of the responses.

In this way, you can check your password without actually having to send 
it anywhere. It all stays local. Really cool. 

To learn more, and get a way better explanation than what I can give, 
check out this Computerphile video on Youtube:
 
https://www.youtube.com/watch?v=hhUb5iknVJs

*/
