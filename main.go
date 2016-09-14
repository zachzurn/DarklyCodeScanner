package main

import (
    "fmt"
    "os"
    "path/filepath"
    "io/ioutil"
)


type CodeScanner interface {
    
    Setup()
    ScanCode(content []byte,extension string) (bool, string)
    Teardown()
    
    AcceptsExtension(extension string) bool
    Identifier() string

}

func main(){

    availableScanners := make(map[string]CodeScanner);
    activeScanners := []CodeScanner{};
    
    scannerIdentifiers := os.Args[1:]
    
    availableScanners["php"] = new(PhpScanner);
    
    //Set up active scanners based on scannerIdentifiers passed in
    for _, si := range scannerIdentifiers {
    
        if scanner, ok := availableScanners[si]; ok {
            
            activeScanners = append(activeScanners,scanner);
            
        }
        
    }
    
    scanFiles( activeScanners )
}



func scanFiles( scanners []CodeScanner ) []string{
    
    searchDir := "./"
    scannedFiles := 0
    infectedFiles := 0
    infectedFilesList := [0]string{}
    
    for _, scanner := range scanners {
        scanner.Setup();
    }
    
    
    //Scan each file using the code scanners
    err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
        if f, err := os.Stat(path); err == nil && f.Mode().IsRegular() {

            scannedFiles++
            readFile := false
            ext := filepath.Ext(path)
            
            //Check if we need to read the contents of the file by seeing if any of the scanners accept a file with this extension
            for _, scanner := range scanners {
                if scanner.AcceptsExtension(ext) { readFile = true }
            }
            
            if !readFile { return nil }
            
            b, err := ioutil.ReadFile(path)
            if err != nil {
                fmt.Printf("Error reading file %v - %v\n",path,err);
            }
            
            for _, scanner := range scanners {

                malicious, reason := scanner.ScanCode(b,filepath.Ext(path))

                if malicious == true { 
                    infectedFiles++
                    infectedFilesList = append(infectedFilesList,path);
                    fmt.Printf("[%v] %v - %v\n",scanner.Identifier(),path,reason) 
                }
                
            }
     
            
        }
        
        return nil
    })
    
    if err != nil {
        fmt.Printf("%v\n%v", "There was an error scanning the files.", err)
    }
    
    for _, scanner := range scanners {
        scanner.Teardown();
    }
    
    fmt.Printf("Finished scanning %v files. %v infected files were found.\n", scannedFiles,infectedFiles)
    
    return infectedFilesList;

}

