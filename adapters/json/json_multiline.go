package main

import "fmt"
import "time"


func parse(queue chan string, done chan bool) {
	output :=  make([]byte, 0, 1000)
	brace_depth := 0
	for elem := range queue	{
		for j := 0; j < len(elem); j++ {
			switch elem[j] {
			case '{' : {
				output = append(output, elem[j])
				brace_depth--
			}
			case '}' : {
				output = append(output, elem[j])
				brace_depth++
				if brace_depth == 0 {
					output = append(output, '\n')
					fmt.Print(string(output))
					output = make([]byte, 0, 1000)
							
				}
			}
			case '\n' :
			default   : 
				output = append(output, elem[j])
			}
		}
	}
	done <- true
}

func source(queue chan string) {
	json := []string{"{\n", "  \"bob\": \"fred\"  ", "{\n", "  \"bob2\": \"fred2\"  ", "}\n", "}\n"}
	for j := 0; j < len(json); j++ {
            queue <- json[j]
		time.Sleep(time.Second * 3)
        }
	close(queue)
}

func main() {
	queue := make(chan string)
	done  := make(chan bool, 1)
	
	go source(queue)
	go parse(queue, done)
	<-done
}
