package main

import (
	"fmt"
	"time"
)

func main() {

	//text = "12 мин. 13 сек."

	var m, s time.Duration
	fmt.Scanf("%d мин. %d сек.", &m, &s)

	t := time.Unix(1589570165, 0).UTC().Add(m * time.Minute).Add(s * time.Second)

	fmt.Println(t.Format(time.UnixDate))

}
