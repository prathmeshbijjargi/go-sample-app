package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	razorpay "github.com/razorpay/razorpay-go"
)

type PageVariables struct {
	Id string
}

func main() {
	http.HandleFunc("/", Form)
	log.Fatal(http.ListenAndServe(":8089", nil))
}

func Form(w http.ResponseWriter, r *http.Request) {

	client := razorpay.NewClient("YOUR_KEY_ID", "YOUR_KEY_SECRET")

	data := map[string]interface{}{
		"amount":          1234,
		"currency":        "INR",
		"receipt":         "reciept_123",
		"payment_capture": 1,
	}
	body, err := client.Order.Create(data, nil)

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}

	value := body["id"]

	str := value.(string)

	HomePageVars := PageVariables{ //store the date and time in a struct
		Id: str,
	}

	t, err := template.ParseFiles("form.html") //parse the html file homepage.html
	if err != nil {                            // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                  // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
