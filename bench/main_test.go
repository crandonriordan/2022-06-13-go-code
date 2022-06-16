/* Alta3 Research | RZFeeser
   Mixing tests with our benchmark tests. */
   
package main

import (
    "fmt"
    "testing"
)

// additional test (non benchmark)
func TestCalculate(t *testing.T) {
    fmt.Println("Test Calculate")
    expected := 4
    result := Calculate(2)
    if expected != result {
        t.Error("Failed")
    }
}

// the only benchmark test
func BenchmarkCalculate(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Calculate(2)
    }
}

// additional test (non benchmark)
func TestOther(t *testing.T) {
    fmt.Println("Testing something else")
    fmt.Println("This shouldn't run with -run=calc")
}

