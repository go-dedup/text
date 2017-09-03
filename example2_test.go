//package main

package text_test

import (
	"fmt"

	"github.com/go-dedup/text"
)

var Doc2words = text.GetWordsFactory(text.Decorators(
	text.SplitCamelCase,
	text.ToLower,
	text.RemovePunctuation,
	text.Compact,
))

// for standalone test, change package to `main` and the next func def to,
// func main() {
func ExampleDoc2Words() {
	for _, d := range testDoc {
		fmt.Printf("%v\n", Doc2words(string(d)))
	}

	// Output:
	// [ford f 150 lariat d o no t bu y truck has been in the shop 50 days so far it has had a vibration since day one and ford cannot get rid of it the have done everything possible to the underside of this truck and it is 11000km automatic]
	// [2016 ford mustang 2016 ford mustang white with black stripes this car is in showroom shape and it only has 14000kms this beast has never been in an accident nor does it have one scratch on the body i purchased 20 14000km automatic]
	// [2013 ford fiesta sedan 22116 kms body is in perfect condition no mechanical problems oil change and maintenance package done in march 17 registered inspection done in april 16 $10000 firm sales tax is extra call 22120km automatic]
	// [2015 ford explorer sport su v crossover this vehicle is a real beauty and a pleasure to drive it is in excellent condition and has been store inside since purchased in 2015 it has not been driven in winter other then to go for service 18600km automatic]
	// [2013 ford fiesta sedan 22116 kms body is in perfect condition no mechanical problems oil change and maintenance package done in march 17 registered inspection done in april 16 $10000 firm sales tax is extra call 22120km automatic]
	// [2015 ford explorer sport su v crossover this vehicle is a real beauty and a pleasure to drive it is in excellent condition and has been store inside since purchased in 2015 it has not been driven in winter other then to go for service 18600km automatic]
	// [ford f 150 lariat d o no t bu y truck has been in the shop 50 days so far it has had a vibration since day one and ford cannot get rid of it the have done everything possible to the underside of this truck and it is 11000km automatic]
	// [2016 ford mustang 2016 ford mustang white with black stripes this car is in showroom shape and it only has 14000kms this beast has never been in an accident nor does it have one scratch on the body i purchased 20 14000km automatic]
}

var testDoc = [][]byte{
	[]byte("Ford F-150. Lariat DO NOT BUY. Truck has been in the shop 50 days so far. It has had a vibration since day one and Ford cannot get rid of it. The have done everything possible to the underside of this truck and it is… 11,000km | Automatic"),
	[]byte("2016 Ford Mustang 2016 Ford Mustang white with black stripes, this car is in showroom shape and it only has 14,000kms. this beast has never been in an accident nor does it have one scratch on the body. i purchased 20… 14,000km | Automatic"),
	[]byte("2013 Ford Fiesta Sedan - 22,116 kms Body is in perfect condition. No mechanical problems. Oil change and maintenance package done in March/17. Registered inspection done in April/16. $10,000 firm (sales tax is extra). Call … 22,120km | Automatic"),
	[]byte("2015 Ford Explorer Sport SUV, Crossover This vehicle is a real beauty and a pleasure to drive. It is in excellent condition and has been store inside since purchased in 2015. It has not been driven in winter other then to go for service.!… 18,600km | Automatic"),
	[]byte("2013 Ford Fiesta Sedan - 22,116 kms Body is in perfect condition. No mechanical problems. Oil change and maintenance package done in March/17. Registered inspection done in April/16. $10,000 firm (sales tax is extra). Call … 22,120km | Automatic"),
	[]byte("2015 Ford Explorer Sport SUV, Crossover This vehicle is a real beauty and a pleasure to drive. It is in excellent condition and has been store inside since purchased in 2015. It has not been driven in winter other then to go for service.!… 18,600km | Automatic"),
	[]byte("Ford F-150. Lariat DO NOT BUY. Truck has been in the shop 50 days so far. It has had a vibration since day one and Ford cannot get rid of it. The have done everything possible to the underside of this truck and it is… 11,000km | Automatic"),
	[]byte("2016 Ford Mustang 2016 Ford Mustang white with black stripes, this car is in showroom shape and it only has 14,000kms. this beast has never been in an accident nor does it have one scratch on the body. i purchased 20… 14,000km | Automatic"),
}
