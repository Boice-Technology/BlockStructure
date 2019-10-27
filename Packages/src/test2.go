package main

import ("flag";
        "fmt";
        )

func main(){
	var ip = flag.Int("flagname",1,"help me !")
	flag.Parse()
    fmt.Println(*ip)
}