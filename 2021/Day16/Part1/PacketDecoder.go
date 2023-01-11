package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main(){
	line := readInput("../input.txt")[0]
	scanner := newScanner(line)

	sum := 0
	if !scanner.IsDone(){
		sum += ProcessPacket(scanner)
	}

	fmt.Println(sum)
}

type Scanner struct {
	hex string
	bin string

	tailBin string
}

func newScanner(hex string) *Scanner{
	return &Scanner{hex, "", ""}
}

func (s *Scanner) Take(n int) string{
	var out bytes.Buffer

	for n > 0 {
		if len(s.bin) == 0{
			if len(s.hex) > 0{
				s.bin = hexToBinary(s.hex[0])
				s.hex = s.hex[1:]
			} else{
			s.bin = s.tailBin
			s.tailBin = ""
			}
		}
		//TODO: What is this?
		if n > len(s.bin) {
			out.WriteString(s.bin)
			n -= len(s.bin)
			s.bin = ""
		}else{
			out.WriteString(s.bin[:n])
			s.bin = s.bin[n:]
			n = 0
		}
	}

	return out.String()
}

func (s *Scanner) SubScanner(length int) *Scanner{
	sub := &Scanner{"", "", ""}

	if len(s.bin) >= length {
		sub.bin = s.bin[:length]
		s.bin = s.bin[length:]
		return sub
	}

	sub.bin = s.bin
	length -= len(s.bin)
	s.bin = ""
	hexLength := length / 4
	sub.hex = s.hex[:hexLength]
	length -= hexLength * 4
	s.hex = s.hex[hexLength:]

	sub.tailBin = s.Take(length)

	return sub
}

func (s Scanner) IsDone() bool{
	if s.Len() == 0{
		return true
	}

	for _,b := range s.bin{
		if b != '0' {
			return false
		}
	}

	for _, h := range s.hex {
		if h != '0' {
			return false
		}
	}

	return true
}

func (s Scanner) Len() int{
	return len(s.hex)*4 + len(s.bin) + len(s.tailBin)
}


//If Type ID == 4 then parse literal value
//	This means that we will parse VVVTTT000AAAAABBBBBCCCCC0000 until first char of packet is 0 then that is last packet

//If Type ID != 4 then parse operator
// This means that we will parse VVVTTTILLLLLLLLLLL
// If I is 0 then the next 15 bits will tell us how long the next subpackets will be
// 001 110 0 000000000011011 - [110 100 01010] [010 100 10001 00100] 0000000
//If I is 1 then it will tell how many packets will be coming
// 111 011 1 00000000011 [010 100 00001] [100 100 00010] [001 100 00011] 00000

func ProcessPacket(s * Scanner) int {
	versionBits := s.Take(3)
	version := parseBitString(versionBits)
	
	typeBits := s.Take(3)
	typeId := parseBitString(typeBits)

	if typeId == 4 {
		//TODO:
		valueBits := ""
		for bits := s.Take(5);bits[0] != '0'; bits = s.Take(5){
			valueBits += bits[1:]
		}
		// lastBits := s.Take(5)
		// valueBits += lastBits[1:]
		return version
	}

	opType := s.Take(1)[0]
	subVersionSum := 0
	if opType == '0'{
		lengthBits := s.Take(15)
		length := parseBitString(lengthBits)
		subScanner := s.SubScanner(length)

		for !subScanner.IsDone(){
			subVersionSum += ProcessPacket(subScanner)
		}
	}else{
		packetCountBits := s.Take(11)
		packetCount := parseBitString(packetCountBits)

		for i := 0; i < packetCount; i++{
			subVersionSum += ProcessPacket(s)
		}
	}

	return version + subVersionSum
}

func parseBitString(str string) int{
	val := 0 
	for _, c := range str {
		val <<= 1
		if c == '1'{
			val++
		}
	}

	return val
}

//TODO:
func hexToBinary(char byte) string {
	switch char {
	case '0':
		return "0000"
	case '1':
		return "0001"
	case '2':
		return "0010"
	case '3':
		return "0011"
	case '4':
		return "0100"
	case '5':
		return "0101"
	case '6':
		return "0110"
	case '7':
		return "0111"
	case '8':
		return "1000"
	case '9':
		return "1001"
	case 'A':
		return "1010"
	case 'B':
		return "1011"
	case 'C':
		return "1100"
	case 'D':
		return "1101"
	case 'E':
		return "1110"
	case 'F':
		return "1111"
	}
	return ""
}

func readInput(fileName string) []string {
	content, fileReadErr := os.ReadFile(fileName)
	check(fileReadErr)
	inputLines := strings.Split(string(content), "\n")
	return inputLines
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}