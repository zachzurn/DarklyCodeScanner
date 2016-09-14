
package main

import (
    "testing"
)

func TestScan(*testing.T){
    infectedFiles := scanFiles([]CodeScanner{ new (PhpScanner) })
    
    //check if clean files came back
}