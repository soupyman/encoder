package main

import (
	"crypto/md5"
	"flag" 
	"fmt"
	"os"
	"io"
	)


func main() {
	in := flag.String("i", "", "input file name")
	out := flag.String("o", "", "output file name")
	cipher := flag.String("c", "alphago", "cipher code")
	
	flag.Parse()

	md5Ctx := md5.New()
    md5Ctx.Write([]byte(*cipher))
    code := md5Ctx.Sum(nil)

	f, err := os.Stat(*in)
	if err != nil { panic(err) } 
	total := int(f.Size())

    fmt.Printf("input file info: %s %s %d \n", f.Name() , f.Mode(), f.Size())

    fin,err := os.Open(*in)
    if err != nil { panic(err) } 
    fout,err := os.Create(*out)
    if err != nil { panic(err) } 
    
    defer fin.Close()
    defer fout.Close()

    buf := make([]byte, 1024)
    idx := 0

    for {
        ni, err := fin.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if ni == 0 { break }

        for i:=0;i<len(code);i++ {
	   		buf[i*10+5] = buf[i*10+5] ^ code[i];
        }

        no, err := fout.Write(buf[:ni])
        if err != nil {
            panic(err)
        } else if no != ni {
            panic("error in writing")
        }
        idx += ni

        fmt.Printf("\r %d completed.", (idx * 100)/total)
    }

    
    fmt.Println("\n job finished.")


}

