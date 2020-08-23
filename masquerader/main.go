package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "io"
    "path/filepath"
    "os/exec"
    "time"
     "strconv"
)
func copy(src, dst string) (int64, error) {
        sourceFileStat, err := os.Stat(src)
        if err != nil {
                return 0, err
        }

        if !sourceFileStat.Mode().IsRegular() {
                return 0, fmt.Errorf("%s is not a regular file", src)
        }

        source, err := os.Open(src)
        if err != nil {
                return 0, err
        }
        defer source.Close()

        destination, err := os.Create(dst)
        if err != nil {
                return 0, err
        }
        defer destination.Close()
        nBytes, err := io.Copy(destination, source)
        return nBytes, err
}
func execute_it(cmdSource string){
   t := time.Now()
   formattedTime := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

   cmd := exec.Command(cmdSource)
   cmd.Stdout = os.Stdout
   err := cmd.Start()
   if err != nil {
      log.Fatal(err)
   }
   t = time.Now()
   formattedTime = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

   fmt.Printf(cmdSource + " execution started at %s (process_id: %d)\n\n", formattedTime,cmd.Process.Pid)
   logger.Println(cmdSource + " execution started at " + formattedTime + " (process_id: " + strconv.Itoa(cmd.Process.Pid) + ")\n\n")
   time.Sleep(4 * time.Second)
}
func listonly() {
    // generate winbins.txt using dir /b /s c:\windows\*.exe > winbins.txt
    file, err := os.Open("winbins.txt")

    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      path:= scanner.Text()
      file := filepath.Base(path)
      fmt.Println(file)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func queryandexecute() {
    /// generate winbins.txt using dir /b /s c:\windows\*.exe > winbins.txt
    file, err := os.Open("winbins.txt")
    des := "c:\\windows\\temp\\"
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      path:= scanner.Text()
      file := filepath.Base(path)
      fmt.Println(file)
      copy("source.exe",des+file)
      execute_it(des+file)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
var (
outfile, _ = os.Create("masqueradingexecution.log") // update path for your needs
logger      = log.New(outfile, "", 0)
)
func main(){
  queryandexecute()
}

