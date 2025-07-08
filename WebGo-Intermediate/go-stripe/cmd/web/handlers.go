package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ashparshp/go-stripe/internal/cards"
	"github.com/ashparshp/go-stripe/internal/models"
	"github.com/go-chi/chi/v5"
)

// Home displays the home page
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "home", &templateData{}); err != nil {
		app.errorLog.Println(err)
	}
}

// VirtualTerminal displays the virtual terminal page
func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "terminal", &templateData{}, "stripe-js"); err != nil {
		app.errorLog.Println(err)
	}
}

type TransactionalData struct {
	FirstName       string
	LastName        string
	Email           string
	PaymentIntentID string
	PaymentMethodID string
	PaymentAmount   int
	PaymentCurrency string
	LastFour        string
	ExpiryMonth     int
	ExpiryYear      int
	BankReturnCode  string
}

// GetTransactionalData retrieves the transactional data from the request
func (app *application) GetTransactionalData(r *http.Request) (TransactionalData, error) {
	var txnData TransactionalData
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Println("Error parsing form:", err)
		return txnData, err
	}

	firstName := r.Form.Get("first_name")
	lastName := r.Form.Get("last_name")
	email := r.Form.Get("email")
	paymentIntent := r.Form.Get("payment_intent")
	paymentMethod := r.Form.Get("payment_method")
	paymentAmount := r.Form.Get("payment_amount")
	paymentCurrency := r.Form.Get("payment_currency")

	amount, err := strconv.Atoi(paymentAmount)
	if err != nil {
		app.errorLog.Println("Error converting payment amount:", err)
		return txnData, err
	}

	card := cards.Card{
		Secret: app.config.stripe.secret,
		Key:    app.config.stripe.key,
	}

	pi, err := card.RetrievePaymentIntent(paymentIntent)
	if err != nil {
		app.errorLog.Println(err)
		return txnData, err
	}

	pm, err := card.GetPaymentMethod(paymentMethod)
	if err != nil {
		app.errorLog.Println(err)
		return txnData, err
	}

	lastFour := pm.Card.Last4
	expiryMonth := pm.Card.ExpMonth
	expiryYear := pm.Card.ExpYear

	txnData = TransactionalData{
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		PaymentIntentID: paymentIntent,
		PaymentMethodID: paymentMethod,
		PaymentAmount:   amount,
		PaymentCurrency: paymentCurrency,
		LastFour:        lastFour,
		ExpiryMonth:     int(expiryMonth),
		ExpiryYear:      int(expiryYear),
		BankReturnCode:  pi.Charges.Data[0].ID,
	}
	return txnData, nil
}

// PaymentSucceeded displays the receipt page
func (app *application) PaymentSucceeded(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	// read posted data
	widgetID, err := strconv.Atoi(r.Form.Get("product_id"))
	if err != nil {
		app.errorLog.Println("Error converting widget ID:", err)
		return
	}

	txnData, err := app.GetTransactionalData(r)
	if err != nil {
		app.errorLog.Println("Error getting transactional data:", err)
		return
	}

	// create a new customer
	customerID, err := app.SaveCustomer(txnData.FirstName, txnData.LastName, txnData.Email)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	// create a new transaction
	txn := models.Transaction{
		Amount:              txnData.PaymentAmount,
		Currency:            paymentCurrency,
		LastFour:            lastFour,
		ExpiryMonth:         int(expiryMonth),
		ExpiryYear:          int(expiryYear),
		BankReturnCode:      pi.Charges.Data[0].ID,
		PaymentIntent:       paymentIntent,
		PaymentMethod:       paymentMethod,
		TransactionStatusID: 2,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	transactionID, err := app.SaveTransaction(txn)
	if err != nil {
		app.errorLog.Println("Error saving transaction:", err)
		return
	}

	// create a new order
	order := models.Order{
		WidgetID:      widgetID,
		TransactionID: transactionID,
		CustomerID:    customerID,
		StatusID:      1,
		Quantity:      1,
		Amount:        amount,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	_, err = app.SaveOrder(order)
	if err != nil {
		app.errorLog.Println("Error saving order:", err)
		return
	}

	data := make(map[string]interface{})
	data["email"] = email
	data["pi"] = paymentIntent
	data["pm"] = paymentMethod
	data["pa"] = paymentAmount
	data["pc"] = paymentCurrency
	data["last_four"] = lastFour
	data["expiry_month"] = expiryMonth
	data["expiry_year"] = expiryYear
	data["bank_return_code"] = pi.Charges.Data[0].ID
	data["first_name"] = firstName
	data["last_name"] = lastName

	// should write this data to session, and then redirect user to new page?
	app.Session.Put(r.Context(), "receipt", data)
	http.Redirect(w, r, "/receipt", http.StatusSeeOther)
}

// Receipt displays the receipt page
func (app *application) Receipt(w http.ResponseWriter, r *http.Request) {
	data := app.Session.Get(r.Context(), "receipt").(map[string]interface{})
	if data == nil {
		http.Error(w, "No receipt data found", http.StatusNotFound)
		return
	}
	app.Session.Remove(r.Context(), "receipt")

	if err := app.renderTemplate(w, r, "receipt", &templateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println(err)
	}

}

// SaveCustomer saves a new customer to the database
func (app *application) SaveCustomer(firstName, lastName, email string) (int, error) {
	// Create a new customer in the database
	customer := models.Customer{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	id, err := app.DB.InsertCustomer(customer)
	if err != nil {
		app.errorLog.Println("Error inserting customer:", err)
		return 0, err
	}

	return id, nil
}

// SaveTransaction saves a new transaction to the database
func (app *application) SaveTransaction(txn models.Transaction) (int, error) {
	// Create a new transaction in the database
	id, err := app.DB.InsertTransaction(txn)
	if err != nil {
		app.errorLog.Println("Error inserting transaction:", err)
		return 0, err
	}

	return id, nil
}

// SaveOrder saves a new order to the database
func (app *application) SaveOrder(order models.Order) (int, error) {
	// Create a new order in the database
	id, err := app.DB.InsertOrder(order)
	if err != nil {
		app.errorLog.Println("Error inserting order:", err)
		return 0, err
	}
	return id, nil
}

// ChargeOnce displays the page to buy one widget
func (app *application) ChargeOnce(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	widgetID, _ := strconv.Atoi(id)

	widget, err := app.DB.GetWidget(widgetID)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	data := make(map[string]interface{})
	data["widget"] = widget

	if err := app.renderTemplate(w, r, "buy-once", &templateData{
		Data: data,
	}, "stripe-js"); err != nil {
		app.errorLog.Println(err)
	}
}
