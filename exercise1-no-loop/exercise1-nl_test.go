package main

import "testing"

func TestUtopianTree2(t *testing.T) { 

	// tests even cycles
    cycles := utopianTree2(4)     
    if cycles != 7 {                   
        t.Fatal("4 -> 7")        
    }
    // tests odd cycles
    cycles = utopianTree2(1)     
    if cycles != 2 {                   
        t.Fatal("1 -> 2")        
    }                                
}    