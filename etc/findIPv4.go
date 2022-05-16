package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

func findIP(input string) string {
	// 1~255까지 나올수있는 경우의 수를 변수에 저장 ( 8비트의 최고의 숫자는 255이기때문 )
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	// 1.1.1.1 ~ 255.255.255.255 즉 IPv4주소를 표현하는 변수
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	// 정규 표현식 regexp 사용
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Printf("usage: %s logFile \n", filepath.Base(arguments[0]))
		os.Exit(1)
	}

	for _, filename := range arguments[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %s\n", err)
			os.Exit(-1)
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			// bufio.ReadString을 이용해 한줄씩 읽음
			line, err := r.ReadString('\n')
			// EOF 즉 문서의 끝이면 Break
			if err == io.EOF {
				break
				// Error 발생시 Break
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
				break
			}
			// FindIP 함수 호출
			ip := findIP(line)
			// net.ParseIP는 IPv4의 주소가 유효한지 다시 한번 확인한다.
			trial := net.ParseIP(ip)
			if trial.To4() == nil {
				continue
			} else {
				fmt.Println(ip)
			}
		}
	}
}
