package main

import (
	"encoding/json"
	"fmt"
	"github.com/yutelin/maicoin-go"
)

func apiKeyExample() {
	fmt.Println("API key example")
	//Api key/secret authentication
	var apiKey = "YOUR_API_KEY"
	var apiSecret = "YOUR_API_SECRET"
	var accessToken = ""
	var client = maicoin.Client{ApiKey: apiKey, ApiSecret: apiSecret, AccessToken: accessToken}
	// fmt.Println(client)

	// fmt.Println("Prices")
	// var price, _ = client.Prices("TWD")
	// printResult(price)
	// price, _ = client.Prices("USD")
	// printResult(price)

	// fmt.Println("Currencies")
	// currencies, _ := client.Currencies()
	// printResult(currencies)

	fmt.Println("User")
	user, _ := client.User()
	printResult(user)

	// fmt.Println("Balance")
	// balance, _ := client.Balance()
	// printResult(balance)

	// fmt.Println("Receive Address")
	// receiveAddress, _ := client.ReceiveAddress("btc")
	// printResult(receiveAddress)

	// fmt.Println("Addresses")
	// addresses, _ := client.Addresses("btc")
	// printResult(addresses)

	// fmt.Println("Generate Receive Address")
	// new_addr, _ := client.GenerateReceiveAddress("btc")
	// printResult(new_addr)

	// fmt.Println("Create account pin")
	// accountPin, _ := client.CreateAccountPin("1234")
	// printResult(accountPin)

	// fmt.Println("Update account pin")
	// updatePin, _ := client.UpdateAccountPin("1234", "1234")
	// printResult(updatePin)

	// fmt.Println("Orders")
	// orders, _ := client.Orders(1, 3)
	// printResult(orders)

	// fmt.Println("Buy Order")
	// buyOrder, _ := client.BuyOrder(1.233, "btc")
	// printResult(buyOrder)

	// fmt.Println("Sell Order")
	// sellOrder, _ := client.SellOrder(0.12312312, "btc")
	// printResult(sellOrder)

	// fmt.Println("Order")
	// order, _ := client.Order("90cb4f6fc21f502aa980df069e9c5f4b8a22271040fd9832")
	// printResult(order)

	// fmt.Println("Transactions")
	// txns, _ := client.Transactions(1, 3)
	// printResult(txns)

	// fmt.Println("Transaction")
	// txn, _ := client.Transaction("9bba914cf4c9c943e8bc3c695fef364172a7af2d792ae5a8")
	// printResult(txn)

	// fmt.Println("Request Transaction")
	// txn, _ := client.RequestTransaction("yutelin@gmail.com", 0.1243455, "btc")
	// printResult(txn)

	// fmt.Println("Cancel Request")
	// txn, _ := client.CancelRequestTransaction("cdeaf83a1686956ac1fced8db559202668e34c117d4a4517")
	// printResult(txn)

	// fmt.Println("Approve Request")
	// txn, _ := client.ApproveRequestTransaction("88871b949c6a7613b0c0b03fafb559ff88b658d9a5f33883", "1234")
	// printResult(txn)

	// fmt.Println("Checkout")
	// cp := maicoin.CheckoutParam{}
	// cp.SetCheckoutData(150, "TWD", "http://my.com/return",
	// 	"http://my.com/cancel", "http://my.com/callback",
	// 	"mer_id_123", "pos_data123", "en")
	// cp.SetBuyerData("YL", "add1", "add2", "palo alto", "ca", "94305",
	// 	"abc@example.com", "6504232423", "US")
	// cp.AddItem("desc1", "code1", "123", "TWD", "true")
	// cp.AddItem("desc2", "code2", "456", "USD", "false")
	// param := cp.Build()
	// fmt.Println(param)
	// checkout, _ := client.CreateCheckout(param)
	// printResult(checkout)
}

func oauthExample() {
	fmt.Println("Access token example")
	//Access token authentication
	var apiKey = ""
	var apiSecret = ""
	var accessToken = "THE_ACCESS_TOKEN"
	var client = maicoin.Client{ApiKey: apiKey, ApiSecret: apiSecret, AccessToken: accessToken}
	fmt.Println("User")
	user, _ := client.User()
	printResult(user)
}

func main() {
	fmt.Println(maicoin.MAICOIN_API_ENDPOINT)
	apiKeyExample()
	oauthExample()
}

func printResult(v interface{}) {
	v2b, _ := json.Marshal(v)
	fmt.Printf("%T: %v\n", v, string(v2b))
}
