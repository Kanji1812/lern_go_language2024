package main
import "fmt"
type deck [] string

func main(){
	cards:= NewDack()

	fmt.Println(cards)
}

func NewDack() deck{
	cards:= deck{}
	cardSuits := []string{"Spandes","Dimonds","Hearts","one","two","mita","harsh","hita","harit","Om","poojan"}
	cardValue := []string {"one","two","three","foure","Five"}
	for _ , shout := range cardSuits{
		for _ , value := range cardValue{
			cards = append(cards,"The Shot is : "+shout+"The Value of value is : "+value)
		}
	}
	return cards
}