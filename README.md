# Go Restful API ApplyToCourier

## Features
The following functions are a set for creating this web APIs:
- Routing with [Fiber](https://github.com/gofiber/fiber)
- Migrations support with [Go-migrate](https://github.com/golang-migrate/migrate)

## Start Application
  - Clone this repository
  - Create a postgres database. See table code [here](migration/000001_create_applycourier_table.up.sql).
  - Change Data in `config/local.env`
  - Run the application: `go run main.go`

## API Routes
| Path          | Method | Request                       |  Desription                                           |                                    
| ------------- | ------ | ----------------------------- | ----------------------------------------------------- |
| /create       | GET    |  { "Creator Id": 0, "Item category": "", "Item weight": "", "First Address Phone": "", "Second Address Phone": "", "First Address": { "Street": "",  "Home": 0, "Housing": 0, "Entrance": 0,"Floor": 0, "Flat": 0, "Intercom Code": ""}, "Second Address": { "Street": "",  "Home": 0, "Housing": 0,  "Entrance": 0, "Floor": 0,  "Flat": 0, "Intercom Code": ""}}                             | Create order in response you wil get ID of your order |   
| /delete       | GET    | { "Order Id": 0 }             | Delete order by ID                                    |     
| /get-creator  | GET    | { "Creator Id": 0 }           | Get order by Creator ID                               |   
| /get-order    | GET    | { "Order Id": 768933 }        | Get order by order ID                                 |




