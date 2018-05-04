package vkapi

import (
  "log"
  "os"
	"strconv"
	"strings"
)

func ArrayToStr(a []int) string {
	s := []string{}
	for _, num := range a {
		s = append(s, strconv.Itoa(num))
	}
	return strings.Join(s, ",")
}

func Print(text string, debug bool) {
  flag, _ := os.LookupEnv("DEBUG")
  if !debug || debug && flag == "1" {
    log.Println(text)
  }
}
