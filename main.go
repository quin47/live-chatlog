package main

import (
	"fmt"
	"github.com/quin47/live-chatlog/record"
)

func main() {

	fmt.Println("starting ...")

	roomIds := [...]int{544793, 938306,602,526,48840,12771281} // 我是怪异君,龙傲娇, 性巴莎拉,青衣才不是御姐呢,凉风OvQ,蒲苇冥冥

	for _,v := range roomIds {
		go record.Record(v)
	}

	select {

	}
}
