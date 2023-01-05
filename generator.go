package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)
//Stolen from StackOverflow: https://stackoverflow.com/questions/12311033/extracting-substrings-in-go/56129336#56129336
func substr(input string, start int, length int) string {
    asRunes := []rune(input)
    if start >= len(asRunes) {
        return ""
    }
    if start+length > len(asRunes) {
        length = len(asRunes) - start
    }
    return string(asRunes[start : start+length])
}
//Checks email array if email already exists, returns how many times
func emailCheck(email string, list []string) int {
	count := -1
	for _, e := range list {
		if e == email {
			count++
		}
	}
	return count
}
func emailMod(email string, num string) (newemail string) {
	parts := strings.Split(email, "@")
	local := parts[0]
	domain := parts[1]
	newemail = local+num+"@"+domain
	return
}
//Creates an email with given name and format
func emailSub(name string, format string) (email string) {
	/*format list:
	(f.) - first initial lowercase
	(F.) - first initial uppercase
	(f) - first name lowercase
	(f~) - first name title case
	(F) - first name capital
	(l.) - last initial lowercase
	(L.) - last initial uppercase
	(l) - last name lowercase
	(l~) - last name title case
	(L) - last name capital
	*/
	parts := strings.Split(name, ",")
	lastName, firstName := string(parts[0]), string(parts[1])
	email = format
	email = strings.ReplaceAll(email, "(f.)", strings.ToLower(substr(firstName, 0, 1)))
	email = strings.ReplaceAll(email, "(F.)", strings.ToUpper(substr(firstName, 0, 1)))	
	email = strings.ReplaceAll(email, "(f)", strings.ToLower(firstName))
	email = strings.ReplaceAll(email, "(f~)", strings.Title(firstName))
	email = strings.ReplaceAll(email, "(F)", strings.ToUpper(firstName))
	email = strings.ReplaceAll(email, "(l.)", strings.ToLower(substr(lastName, 0, 1)))
	email = strings.ReplaceAll(email, "(L.)", strings.ToUpper(substr(lastName, 0, 1)))	
	email = strings.ReplaceAll(email, "(l)", strings.ToLower(lastName))
	email = strings.ReplaceAll(email, "(l~)", strings.Title(lastName))
	email = strings.ReplaceAll(email, "(L)", strings.ToUpper(lastName))
	return
}

func main() {
	//Get output and format file names from flags
	namesFilename := flag.String("n", "names.csv", "File that stores ")//Names File must be csv; last name, first name
	isRaw := flag.Bool("raw", false, "Writes only raw file to output, no format or --- lines")
	outputFilename := flag.String("o", "emails.txt", "File that output is written to")
	formatListFilename := flag.String("f", "format.txt", "File that program uses formats for. By default, 10 formats, but can be as many or as few as you'd like")//See above comment for format guide
	var duplicates bool
	flag.BoolVar(&duplicates, "duplicates", false, "Name file contains duplicates. If true, will add #s at end")
	flag.BoolVar(&duplicates, "d", false, "duplicates (short hand)")
	flag.Parse()

	//Gets format list
	formatData, fErr := os.Open(*formatListFilename)
	if fErr != nil{
		fmt.Println("Error opening format file")
		return
	}
	fileScanner := bufio.NewScanner(formatData)
	fileScanner.Split(bufio.ScanLines)
	var formats []string
	for fileScanner.Scan() {
		formats = append(formats, fileScanner.Text())
	}
	formatData.Close()

	//Gets name list
	nameData, nErr := os.Open(*namesFilename)
	if nErr != nil {
		fmt.Println("Error opening name file")
		return
	}
	nameScanner := bufio.NewScanner(nameData)
	nameScanner.Split(bufio.ScanLines)
	var names []string
	for nameScanner.Scan() {
		names = append(names, nameScanner.Text())
	}
	nameData.Close()

	//Opens output file
	outFile, oErr := os.Create(*outputFilename)
	if oErr != nil {
		fmt.Println("Error creating ouput file")
		return
	}
	defer outFile.Close()

	//do the funky format stuff
	_, writeErr := outFile.WriteString("\n")//Defines writeErr, could probably do this better but oh well
	emailList := make([]string, 0)
	for _, format := range formats{
		if !*isRaw{
			_, writeErr := outFile.WriteString(format + "\n")
				if writeErr != nil {
					fmt.Println("Error writing to output file")
				}
			}
		for _, name := range names{
			name = strings.ReplaceAll(name, " ", "")
			email := email_sub(name, format)
			emailList = append(emailList, email)
			if email_check(email, emailList) != 0 && duplicates {
				email = email_mod(email, fmt.Sprint(email_check(email, emailList)))
			}
			_, writeErr = outFile.WriteString(email + "\n")
			if writeErr != nil {
				fmt.Println("Error writing to output file")
			}
		}
		if !*isRaw{
			_, writeErr = outFile.WriteString("----------------------------------" + "\n")
				if writeErr != nil {
					fmt.Println("Error writing to output file")
				}
			}
	}
}
