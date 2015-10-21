package main

import "testing"

func TestUtopianTree(t *testing.T) { 
    size, _ := utopianTree(1, 4)     
    if size != 4 {                   
        t.Fatal("1,4 != 4 7")        
    } 

    _ ,cycles := utopianTree(1, 4)     
    if cycles != 7 {                   
        t.Fatal("1,4 != 4 7")        
    }                                
}    