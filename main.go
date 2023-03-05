package main

import (
    "fmt"
    "io"
    "os"
    "path"
    "bufio"
    "strings"
    "net/http"
)

func main() {
    var in string
    var userFilename string
    reader := bufio.NewReader(os.Stdin)
    
    for {
        fmt.Print("Enter a URL or 'Exit' : ")
        fmt.Scanln(&in)

        if in == "Exit" || in == "exit" {
            break
        }
        
        response, err := http.Head(in)
        if err != nil {
            fmt.Println("==============================================")
            fmt.Println("Error when input a URL > ")
            fmt.Println(err)
            fmt.Println("==============================================")
            continue
        }
        defer response.Body.Close()

        if response.StatusCode != 200 {
            fmt.Println("==============================================")
            fmt.Println("Error when input a URL > ")
            fmt.Println("Status code : ", response.StatusCode)
            fmt.Println("==============================================")
            continue
        }

        cd := response.Header.Get("Content-Disposition")
        if !strings.Contains(cd, "attachment") {
            fmt.Println("==============================================")
            fmt.Println("Error when input a URL > ")
            fmt.Println("File type : ", response.Header.Get("Content-Type"))
            fmt.Println("==============================================")
            continue
        }

        filename := strings.Split(cd, "filename=")[1]
        filename = strings.Replace(filename, "\"", "", -1)
        fmt.Println("==============================================")
        fmt.Println("Filename  : ", filename)
        fmt.Println("File type : ", response.Header.Get("Content-Type"))
        fmt.Println("==============================================")

        fmt.Print("Enter a downloaded filename [OPTIONAL] : ")
        userFilename, _ = reader.ReadString('\n')
        userFilename = strings.TrimSpace(userFilename)

        if userFilename != "" {
            extension := path.Ext(filename)
            filename = userFilename + extension
        }

        fmt.Println("Downloading........")
        response, err = http.Get(in)
        if err != nil {
            fmt.Println("==============================================")
            fmt.Println("Error when download a file > ")
            fmt.Println(err)
            fmt.Println("==============================================")
            return
        }
        defer response.Body.Close()

        file, err := os.Create(filename)
        if err != nil {
            fmt.Println("==============================================")
            fmt.Println("Error when download a file > ")
            fmt.Println(err)
            fmt.Println("==============================================")
            return
        }
        defer file.Close()

        _, err = io.Copy(file, response.Body)
        if err != nil {
            fmt.Println("==============================================")
            fmt.Println("Error when download a file > ")
            fmt.Println(err)
            fmt.Println("==============================================")
            return
        }

        fmt.Println("Finished.....")
    }
}