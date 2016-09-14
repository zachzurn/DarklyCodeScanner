package main

import (
    "regexp"
)

type PhpScanner struct {
    evalRegex *regexp.Regexp
    asRegex *regexp.Regexp
}

func(s *PhpScanner) Setup(){

    //Finds eval base64 statements
    r, _ := regexp.Compile(`eval\((base64|eval|\$_|\$\$|\$[A-Za-z_0-9\{]*(\(|\{|\[))`)
    s.evalRegex = r
    
    //Finds 
    r2, _ := regexp.Compile(`php \$[a-zA-Z]*=’as’;`)
    s.asRegex = r2
}

//Return true if a vulnerability was found
func(s *PhpScanner) ScanCode(code []byte,extension string) (bool, string) {
    
    match := s.evalRegex.Find(code)
    if len(match) != 0 {
        return true, "Code is possibly infected with malicious eval statement."
    }
    
    match2 := s.asRegex.Find(code)
    if len(match2) != 0 {
        return true, "Code is possibly infected with malicious obscured eval statement."
    }
    
    return false, ""

}

func(s *PhpScanner) Teardown(){
    
}

func(s *PhpScanner) Identifier() string{
    return "php";
}

func(s *PhpScanner) AcceptsExtension(extension string) bool{
    
    if extension == ".php" || extension == ".phtml" || extension == ".html" || extension == ".htm" || extension == ".tpl" { 
        return true 
    }
    
    return false 
    
}