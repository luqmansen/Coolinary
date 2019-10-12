# Coolinary REST API

## Request

### User

POST ``api/user/new``

Endpoints to create new user

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

POST ``api/user/login``

Endpoints to login for user

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
POST ``api/user/order/new``

Endpoints to user for creating new order

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
***
### Seller

POST ``api/user/new``

Endpoints to create new seller

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

POST ``api/seller/login``

Endpoints to login for seller

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


POST ``api/seller/product/new``

Endpoints to create new product for seller

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

