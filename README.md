# Basket Service in Golang

- This API server provides endpoints to create,read,update & delete. 

##API Server technology stack is

- Server code: `golang`
- REST Server: `gin`
- Database: `PostgreSQL` with `GORM` to migrate
- ORM: `gorm v2`
- Unit Testing: `go test` and `testify`

## To Start API Server
```$ git clone https://github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ranaapay.git```

```$ cd src```

```$ go run main.go```

* **Basket API -> http://localhost:8080/swagger/index.html**


## Data Model

- Base Model
``````
type Base struct {
	Id        string    `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `gorm:"index" json:"deletedAt"`
	IsDeleted bool      `json:"isDeleted"`
}
``````
- User Model
``````
type User struct {
	Base
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      Role   `json:"role"`
}
``````
- Category Model
``````
type Category struct {
	Base
	CategoryName string    `json:"categoryName"`
	Description  string    `json:"description"`
	Product      []Product `json:"products,omitempty"`
}
``````
- Product Model
``````
type Product struct {
	Base
	ProductName string          `json:"productName"`
	Price       decimal.Decimal `json:"price"`
	StockNumber int             `json:"stockNumber"`
	UnitsOnCart int             `json:"unitsOnCart"`
	CategoryId  string          `json:"categoryId"`
	Category    Category        `json:"category,omitempty"`
}
``````
- CartDetails Model
``````
type CartDetails struct {
	Base
	ProductId        string          `json:"productId"`
	ProductQuantity  int64           `json:"productQuantity"`
	DetailTotalPrice decimal.Decimal `json:"detailTotalPrice"`
	CartId           string          `json:"cartId"`
}
``````
- Cart Model
``````
type Cart struct {
	Base
	UserId         string          `json:"userId"`
	TotalCartPrice decimal.Decimal `json:"totalCartPrice"`
	IsCompleted    bool            `json:"isCompleted"`
	CartDetails    []CartDetails
}
``````
- Order Model
``````
type Order struct {
	Base
	CartId      string `json:"cartId"`
	UserId      string `json:"userId"`
	IsCancelled bool   `json:"isCancelled"`
	Cart        Cart
}
``````

## Available API Endpoints

{{baseURI}} = /api/v1/basket

|  Method | API Endpoint  | Authentication Type | User Role | Description |
|---|---|---|---|---|
|POST| {{baseURI}}/authentication/register | - | - |Return jwt token in response for successful authentication
|POST| {{baseURI}}/authentication/login | - | - | Return jwt token in response for successful authentication
|POST| {{baseURI}}/category | Bearer token | Admin | Return true in response for successful read file and create categories.
|GET| {{baseURI}}/category  | - | - | Return a list of all categories in response| 
|POST| {{baseURI}}/product | Bearer token | Admin |Add an product in the database and return the added user id in response | 
|PUT| {{baseURI}}/product/{id} | Bearer token | Admin |Update the user and return the updated user info in response | 
|DELETE| {{baseURI}}/product/{id} | Bearer token | Admin | Delete the product and return true in response for successful| 
|GET| {{baseURI}}/product/{id} | - | - | Return the product of given product id in response| 
|GET| {{baseURI}}/product| - | - | Return list of products that users search and pagination results in response| 
|POST| {{baseURI}}/cart | Bearer token | User |Add an cart detail in the database and return the added cart detail in response | 
|GET| {{baseURI}}/cart/ | Bearer token| User | Return the cart with cart details of user in response| 
|PUT| {{baseURI}}/cart/{id} | Bearer token | User |Update the cart with cart detail and return the updated cart detail in response | 
|DELETE| {{baseURI}}/cart/{id} | Bearer token | User | Delete the cart detail and return true in response for successful| 
|GET| {{baseURI}}/order  | Bearer token | User | Return a list of all orders of user in response| 
|POST| {{baseURI}}/order | Bearer token | User |Add an order in the database and return the added order id in response | 
|DELETE| {{baseURI}}/order/{id} | Bearer token | User | Cancel the order if order date has not passed 14 days and return true in response for successful|

## TODO

- [ ] add more test
- [ ] with using cobra in configuration management separating services based on domain 
- [ ] communication with client between services
- [ ] add docker
- [ ] configure docker compose
