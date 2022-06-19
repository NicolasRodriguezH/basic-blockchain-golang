package main

import (
	"fmt"
)

type Block struct {
	Index int `json:index`
	TimeStamp string `json:time_stamp`
	Proof int `json:proof`
	PreviousHash string `json:previous_hash`
}

type Blocks Block[]

func main()  {
	//
}