/* 
ОПИСАНИЕ:

IT WILL DELETE ALL TIMECODE
FROM TEXT.TXT

From:
./main text.txt 

Result: answer.txt

*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Reading file.txt
	file := os.Args[1] // Только один аргумент
	if len(file) != 0 { // Если есть аргумент file.txt
		deleteTimeCode(file) // Вызов функции удаления ТаймКодов с параметром: стринговый file
	} else { // Если НЕТ аргумента file.txt
		fmt.Println("You must type argument! Example: ./main text.txt") // Сообщение об ошибке
	}
}


func deleteTimeCode(f_to_open string) {
	// Opening file from argument
	f, err := os.Open(f_to_open)
	if err != nil {
        fmt.Println(err)
    }

	lookFor := ":"
	refString := "some text from file"
	true_or_false := false

	// Creating answer.txt
	ff, err := os.Create("answer.file")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("asnwer.txt created!")
	}


	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() { // Считываем из файла построчно

		refString = fileScanner.Text()
		fmt.Println(refString)

	    true_or_false = strings.Contains(refString, lookFor)
	    if true_or_false == true && len(refString) <= 7 {
	    	continue

	    } else {
	    	_, err = ff.WriteString(refString+"\n")
	    	if err != nil {
				fmt.Println(err)
			}

	    }
		
	}
	// Игнорируем потенциальные ошибки из input.Err()
	f.Close()
	ff.Close()
}



