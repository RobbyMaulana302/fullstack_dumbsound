package handlers

import (
	resultdto "dumbsound/dto/result"
	transactiondto "dumbsound/dto/transaction"
	"dumbsound/models"
	"dumbsound/repositories"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gopkg.in/gomail.v2"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepositiry
}

type JsonTransaction struct {
	DataTransaction interface{} `json:"transaction"`
}

func HandlerTransaction(TransactionRepositiry repositories.TransactionRepositiry) *handlerTransaction {
	return &handlerTransaction{TransactionRepositiry}
}

func (h *handlerTransaction) CreateTransaction(c echo.Context) error {
	request := new(transactiondto.TransactionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	user_id := userLogin.(jwt.MapClaims)["id"].(float64)

	var transactionIsMatch = false
	var transactionID int
	for !transactionIsMatch {
		transactionID = int(time.Now().Unix())
		transactionData, _ := h.TransactionRepository.GetTransaction(transactionID)
		if transactionData.ID == 0 {
			transactionIsMatch = true
		}
	}

	start_date := time.Now()
	start_date_request := start_date.String()
	parseStartDate, _ := time.Parse("2006-01-02", start_date_request)
	toStringParseStartDate := parseStartDate.String()

	day := start_date.Day()
	month := start_date.Month() + 1
	year := start_date.Year()
	fmt.Println(day)
	fmt.Println(int(month))
	fmt.Println(year)

	due_date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	parse_due_date, _ := time.Parse("2006-01-02", due_date.String())
	due_date_request := parse_due_date.String()
	fmt.Println(due_date)
	fmt.Println(parse_due_date)
	fmt.Println(due_date_request)

	transaction := models.Transaction{
		ID:        transactionID,
		StartDate: toStringParseStartDate,
		DueDate:   due_date.String(),
		Status:    request.Status,
		Price:     59999,
		UserID:    int(user_id),
	}

	createTransaction, err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Status:  "Errro",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Status: "Success",
		Data: JsonTransaction{
			DataTransaction: createTransaction,
		},
	})
}

func (h *handlerTransaction) FindTransaction(c echo.Context) error {
	transaction, err := h.TransactionRepository.FindTransaction()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Status: "Error",
		Data: JsonTransaction{
			DataTransaction: transaction,
		},
	})
}

func (h *handlerTransaction) GetTransaction(c echo.Context) error {
	transaction_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	fmt.Println(transaction_id)

	transaction, err := h.TransactionRepository.GetTransaction(transaction_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Status: "Success",
		Data: JsonTransaction{
			DataTransaction: transaction,
		},
	})
}

func (h *handlerTransaction) GetUserTransaction(c echo.Context) error {
	userLogin := c.Get("userLogin")
	user_id := userLogin.(jwt.MapClaims)["id"].(float64)

	fmt.Println(user_id)
	userTransaction, err := h.TransactionRepository.GetUserTransaction(int(user_id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Status: "Success",
		Data: JsonTransaction{
			DataTransaction: userTransaction,
		},
	})
}

func (h *handlerTransaction) DeleteTransaction(c echo.Context) error {
	transaction_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	transaction, err := h.TransactionRepository.GetTransaction(transaction_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	deleteTransaction, err := h.TransactionRepository.DeleteTransaction(transaction, transaction_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Status: "Success",
		Data: JsonTransaction{
			DataTransaction: deleteTransaction,
		},
	})
}

func (h *handlerTransaction) GetPayment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: "Error", Message: err.Error()})
	}

	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Price),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: transaction.User.FullName,
			Email: transaction.User.Email,
		},
	}

	snapResp, _ := s.CreateTransaction(req)

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Status: "Success", Data: snapResp})

}

func (h *handlerTransaction) Notification(c echo.Context) error {
	var notificationPayload map[string]interface{}

	if err := c.Bind(&notificationPayload); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: "Error", Message: err.Error()})
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)
	order_Id, _ := strconv.Atoi(orderId)

	transaction, _ := h.TransactionRepository.GetTransaction(order_Id)

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			h.TransactionRepository.UpdateTransaction("pending", order_Id)
		} else if fraudStatus == "accept" {

			SendMail("success", transaction)
			h.TransactionRepository.UpdateTransaction("success", order_Id)
		}
	} else if transactionStatus == "settlement" {

		SendMail("success", transaction)
		h.TransactionRepository.UpdateTransaction("success", order_Id)
	} else if transactionStatus == "deny" {
		h.TransactionRepository.UpdateTransaction("failed", order_Id)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		h.TransactionRepository.UpdateTransaction("failed", order_Id)
	} else if transactionStatus == "pending" {
		h.TransactionRepository.UpdateTransaction("pending", order_Id)
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Status: "Success", Data: notificationPayload})

}

func SendMail(status string, transaction models.Transaction) {
	if status != transaction.Status && (status == "success") {

		var CONFIG_SMTP_HOST = "smtp.gmail.com"
		var CONFIG_SMTP_PORT = 587
		var CONFIG_SENDER_NAME = "DumbSound <robbym6999@gmail.com>"
		var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL_SYSTEM")
		var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD_SYSTEM")

		var ticket = strconv.Itoa(transaction.ID)
		var price = strconv.Itoa(transaction.Price)

		mailer := gomail.NewMessage()
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetHeader("To", transaction.User.Email)
		mailer.SetHeader("Subject", "Transaction Status")
		mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
    <html lang="en">
      <head>
      <meta charset="UTF-8" />
      <meta http-equiv="X-UA-Compatible" content="IE=edge" />
      <meta name="viewport" content="width=device-width, initial-scale=1.0" />
      <title>Document</title>
      <style>
        h1 {
        color: brown;
        }
      </style>
      </head>
      <body>
      <h2>Product payment :</h2>
      <ul style="list-style-type:none;">
        <li>Name : %s</li>
        <li>Total payment: Rp.%s</li>
        <li>Status : <b>%s</b></li>
      </ul>
      </body>
    </html>`, ticket, price, status))

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Mail sent! to " + transaction.User.Email)
	}

}
