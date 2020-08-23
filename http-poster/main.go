package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "bytes"
    "io/ioutil"
    "path/filepath"
    "net/http"
    "time"
)

func MakeRequest(url string) {

  content, err := ioutil.ReadFile(fileSensitiveData)
  	if err != nil {
  		log.Fatal(err)
  	}

  var postStr = []byte(content)

	req, err := http.NewRequest("POST", "http://" + url + "/post", bytes.NewBuffer(postStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(body))
	log.Println(string(body))

}


func execute_it(cmdSource string){
   t := time.Now()
   formattedTime := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

   MakeRequest(cmdSource)
   t = time.Now()
   formattedTime = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

   fmt.Printf(cmdSource + " posted %s \n\n", formattedTime)
   logger.Println(cmdSource + "  posted  at " + formattedTime + "\n\n")
   time.Sleep(4 * time.Second)
}

func queryandexecute() {
    file, err := os.Open(fileDomain)

    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      path:= scanner.Text()
      file := filepath.Base(path)
      fmt.Println(file)
      execute_it(file)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
var (
outfile, _ = os.Create("poster.log") // update path for your needs
fileSensitiveData = "data.txt"
fileDomain = "domains.txt"
logger      = log.New(outfile, "", 0)
)	
func main(){
  queryandexecute()
}
