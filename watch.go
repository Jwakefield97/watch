package main 
import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"os/exec"
	"bytes"
	"log"
	"regexp"
)

func WatchFiles(dir string, regexStr string, script string) {
	files := make(map[string]time.Time)
	isFirst := true
	regex, err := regexp.Compile(regexStr)
	if err != nil {
		log.Fatal(err)  
	}
	for {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			matched := regex.MatchString(path)
			
			if matched && !info.IsDir() && info.ModTime() != files[path] {

				files[path] = info.ModTime()
				if(!isFirst) {

					//fmt.Println(path)
					cmd := exec.Command("bash", script, path)                                                                                                                                                          
					var out bytes.Buffer                                                                                                                                                                                   
					cmd.Stdout = &out                                                                                                                                                                                    
					err := cmd.Run()                                                                                                                                                                                     
					if err != nil  {                                                                                                                                                                                            
							log.Fatal(err)                                                                                                                                                                               
					}                                                                                                                                                                                                          
					fmt.Printf(out.String())

				}
			}
			return nil
		})
		if err != nil {
			fmt.Println("error occured while walking directory")
		}
		isFirst = false
		time.Sleep(1 * time.Second)
	}
}

func main() {
	fmt.Println("Usage: ./watch <directory> <regex> <script>")
	fmt.Println("Note: The shell script will be provided the file modified as an argument when the script is run.")
	dir := os.Args[1]
	regex := os.Args[2]
	script := os.Args[3]
	WatchFiles(dir,regex,script)
}