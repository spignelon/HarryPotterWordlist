package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Get input file name from command line argument
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s <input_file>\n", os.Args[0])
        os.Exit(1)
    }
    inputFileName := os.Args[1]

    // Open input file
    fp_in, err := os.Open(inputFileName)
    if err != nil {
        panic(err)
    }
    defer fp_in.Close()

    // Read words from input file
    var words []string
    scanner := bufio.NewScanner(fp_in)
    for scanner.Scan() {
        words = append(words, scanner.Text())
    }

    // Open output file
    fp_out, err := os.Create("output.txt")
    if err != nil {
        panic(err)
    }
    defer fp_out.Close()

    // Generate permutations and write to output file
    for i := 0; i < len(words); i++ {
        for j := 0; j < len(words); j++ {
            for k := 0; k < len(words); k++ {
                if i != j && j != k && k != i {
                    fmt.Fprintf(fp_out, "%s %s %s\n", words[i], words[j], words[k])
                }
            }
        }
    }
}
