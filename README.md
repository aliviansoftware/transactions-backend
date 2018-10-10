# Transactions Backend (Stub)

Provides user management for the application

| Endpoint               |          Params            |        Description           |
|------------------------|:--------------------------:|:----------------------------:|
| GET /transactions      | none                       | gives a list of transactions |
| GET /transactions/{id} | the transactions id        | gives a specific transaction |

## Example

### Get a list of transaction
```javascript
GET http://127.0.0.1:1377/transactions 
```
A 200 status will be returned with a JSON object of all dummy transactions

### Get information of a specific transaction
```javascript
GET http://127.0.0.1:1377/transactions/{id}
```
A 200 status will be returned with the details of a specific transaction

A 404 will be returned if the transaction doesnt exist

### Creation of a transaction is still in development
