package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		return "Oh hai."
	})

	// http://developers.amiando.com/index.php/Tracking_Webhooks
	m.Post("/amiando-server-call", func(w http.ResponseWriter, r *http.Request) (int, string) {

		eventIdentifier, paymentDiscountCode, paymentValue, paymentCurrency, email0 := r.FormValue("eventIdentifier"), r.FormValue("paymentDiscountCode"), r.FormValue("paymentValue"), r.FormValue("paymentCurrency"), r.FormValue("ticketEmail(0)")

		numberOfTickets, err := strconv.Atoi(r.FormValue("numberOfTickets"))
		if err != nil {
			fmt.Println(err)
			return 500, "Invalid numberOfTickets"
		}

		payload := make(map[string]string)

		payload["text"] = fmt.Sprintf("SOLD! %vx %v %v (%v%v) to %v", numberOfTickets, eventIdentifier, paymentDiscountCode, paymentValue, paymentCurrency, email0)
		payload["channel"] = os.Getenv("SLACK_CHANNEL")
		payload["username"] = "amiando"

		payloadJson, err := json.Marshal(payload)
		if err != nil {
			fmt.Println(err)
			return 500, "Invalid payload JSON"
		}

		fmt.Println("Posting...")

		// Finally send the post to Slack
		resp, err := http.PostForm(os.Getenv("SLACK_URL"), url.Values{"payload": {string(payloadJson)}})

		fmt.Println(resp)
		if err != nil {
			fmt.Println(err)
			return 500, "Invalid post response"
		}

		return 200, "ok"

	})

	http.ListenAndServe(":"+os.Getenv("PORT"), m)
}
