# Coolinary REST API

## Table of Content
- [Coolinary REST API](#coolinary-rest-api)
  * [Table of Content](#table-of-content)
  * [User](#user)
    + [Create New User](#create-new-user)
    + [User Login](#user-login)
    + [User Create New Order](#user-create-new-order)
    + [User Pay The Order](#user-pay-the-order)
    + [User Cancel The Order](#user-cancel-the-order)
    + [User Skip Order Delivery](#user-skip-order-delivery)
  * [Seller](#seller)
    + [Create New Seller Account](#create-new-seller-account)
    + [Seller Login](#seller-login)
    + [Seller Create New Product](#seller-create-new-product)

## User
### Create New User

**POST** ``api/user/new`` 

Post parameters

| Property | Type | Description |
| --- | --- | --- |
| name | string | Name of the user |
| email | string | User email |
| address | string | User Address |
| password | string | User password |

Success Response Example

```json
{
    "message": "Account Has Been Created",
    "status": 200,
    "user": {
        "ID": 2,
        "CreatedAt": "2019-10-11T22:32:05.320821109+07:00",
        "UpdatedAt": "2019-10-11T22:32:05.320821109+07:00",
        "DeletedAt": null,
        "name": "luqman setyo nugroho",
        "email": "luq@man.sen",
        "password": "",
        "address": "Semarang",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjJ9.oFjjSll1HqDEM61AaZ7DCz4TQP-n9UE5nIL2lGH0SHg"
    }
}
```

### User Login
**POST** ``api/user/login``

Post parameters

| Property | Type | Description |
| --- | --- | --- |
| email | string | User email |
| password | string | User password |

Success Response Example

```json
{
    "account": {
        "ID": 2,
        "CreatedAt": "2019-10-11T22:32:05.320821+07:00",
        "UpdatedAt": "2019-10-11T22:32:05.320821+07:00",
        "DeletedAt": null,
        "name": "luqman setyo nugroho",
        "email": "luq@man.sen",
        "password": "",
        "address": "Semarang",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjJ9.oFjjSll1HqDEM61AaZ7DCz4TQP-n9UE5nIL2lGH0SHg"
    },
    "message": "Logged In",
    "status": 200
}
```
### User Create New Order
**POST** ``api/user/order/new`` 

Post parameters

| Property | Type | Description |
| --- | --- | --- |
| product_id | int | ID of selected product  |
| subscription | bool | Subscription type (optional)  |


Success Response Example
```json
{
    "message": "Order Created",
    "order": {
        "ID": 9,
        "CreatedAt": "2019-10-12T08:56:41.056108096+07:00",
        "UpdatedAt": "2019-10-12T08:56:41.056108096+07:00",
        "DeletedAt": null,
        "product_id": 1,
        "SellerID": 1,
        "buyer_id": 2,
        "delivery_time": "13.00",
        "subscription": true,
        "total_price": 600000,
        "paid": false
    },
    "status": 200
}
```
### User Pay The Order
**POST** ``api/user/order/pay/{id}``

Success Response Example

```json
{
    "message": "Payment Success",
    "order": {
        "ID": 7,
        "CreatedAt": "2019-10-12T08:53:24.276806+07:00",
        "UpdatedAt": "2019-10-12T08:53:24.276806+07:00",
        "DeletedAt": null,
        "product_id": 1,
        "SellerID": 0,
        "buyer_id": 2,
        "delivery_time": "13.00",
        "subscription": false,
        "total_price": 20000,
        "paid": false
    },
    "status": 200
}
```

Response Example if Order Already Paid
```json
{
    "message": "Order Already Paid",
    "status": 200
}
```
### User Cancel The Order
**POST** ``api/user/order/cancel/{id}``

Note : Paid order can't be canceled

Success Response Example
```json
{
    "message": "Order Canceled",
    "order": {
        "ID": 11,
        "CreatedAt": "2019-10-12T09:59:40.169588+07:00",
        "UpdatedAt": "2019-10-12T09:59:40.169588+07:00",
        "DeletedAt": "2019-10-12T09:59:46.250632443+07:00",
        "product_id": 1,
        "SellerID": 1,
        "buyer_id": 2,
        "delivery_time": "13.00",
        "subscription": true,
        "total_price": 600000,
        "paid": false
    },
    "status": 200
}
```
Response Example if Order Already Paid

```json
{
    "message": "Order Already Paid, Can't Be Cancelled",
    "status": 200
}
```
### User Skip Order Delivery
POST ``api/user/order/skiptoday/{id}``

Note : Order should've been paid if user want to skip the delivery

Success Response Example
```json
{
    "message": "Order will be sent at 2019-10-13 10:21:12.858538824 +0700 WIB",
    "order": {
        "ID": 1,
        "CreatedAt": "2019-10-12T10:19:47.324702+07:00",
        "UpdatedAt": "2019-10-12T10:19:47.324702+07:00",
        "DeletedAt": null,
        "product_id": 1,
        "SellerID": 1,
        "buyer_id": 2,
        "delivery_time": "13.00",
        "deliver_today": false,
        "subscription": true,
        "total_price": 600000,
        "paid": true
    },
    "status": 200
}
```
***
## Seller
### Create New Seller Account
**POST** ``api/user/new``

Post parameters

| Property | Type | Description |
| --- | --- | --- |
| name | string | Name of the seller |
| store_name | string | Name of the store |
| email | string | seller email |
| address | string | seller Address |
| password | string | seller password |

Success Response Example
```json
{
    "message": "Account Has Been Created",
    "seller": {
        "ID": 2,
        "CreatedAt": "2019-10-11T22:58:23.47424779+07:00",
        "UpdatedAt": "2019-10-11T22:58:23.47424779+07:00",
        "DeletedAt": null,
        "name": "Luqman GSN",
        "store_name": "TokoKu Food",
        "store_address": "Semarang",
        "email": "luq@man.sen",
        "password": "",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjJ9.oFjjSll1HqDEM61AaZ7DCz4TQP-n9UE5nIL2lGH0SHg"
    },
    "status": 200
}
```
### Seller Login
POST ``api/seller/login``

Post parameters

| Property | Type | Description |
| --- | --- | --- |
| email | string | account email |
| password | string | account password |

Success Response Example
```json
{
    "account": {
        "ID": 1,
        "CreatedAt": "2019-10-11T22:12:28.432511+07:00",
        "UpdatedAt": "2019-10-11T22:12:28.432511+07:00",
        "DeletedAt": null,
        "name": "luqman setyo nugroho",
        "email": "luqman@get.rekt",
        "password": "",
        "address": "magelang",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjF9.VOEQc2pqr74vB_44g73RF5gTWQzcWwQWh9Cs4YOZbkg"
    },
    "message": "Logged In",
    "status": 200
}
```

### Seller Create New Product
POST ``api/seller/product/new``

Post parameters

| Property | Type | Description |
| --- | --- | --- |
| product_name | string | Product Name |
| price | uint32 |  Product price |
| selling_area | string |  Product selling area |

Success Response Example
```json

{
    "message": "New Product Added",
    "product": {
        "ID": 1,
        "CreatedAt": "2019-10-12T07:59:43.36179584+07:00",
        "UpdatedAt": "2019-10-12T07:59:43.36179584+07:00",
        "DeletedAt": null,
        "product_name": "Gudeg Jogja",
        "seller_id": 1,
        "price": 20000,
        "selling_area": "Semarang"
    },
    "status": 200
}
```

