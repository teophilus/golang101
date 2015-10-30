package main

import "testing"

func TestUtopianTree(t *testing.T) { 

	// tests even cycles
    cycles := utopianTree(4)     
    if cycles != 7 {                   
        t.Fatal("4 -> 7")        
    }
    // tests odd cycles
    cycles = utopianTree(1)     
    if cycles != 2 {                   
        t.Fatal("1 -> 2")        
    }                                
}

func BenchmarkTree(b *testing.B) {
    for i := 0; i < b.N; i++ {
        utopianTree(10)
    }
}    