package raindrops

import "strconv"

func Convert(number int) string {

    finalresponse := ""
    
    if number % 3 == 0 {
        finalresponse+= "Pling"
    } 
    if number % 5 == 0 {
        finalresponse+= "Plang"
    } 
    if number % 7 == 0 {
        finalresponse+= "Plong"
    } 
    if finalresponse == "" {
        finalresponse+=strconv.Itoa(number)
    }

    return finalresponse
}
