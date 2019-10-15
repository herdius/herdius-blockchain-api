# Utility commands

### Prerequistie

Master Account Key : one account which has Balance, this account will be used to send tx to other accounts, 


## Default config

### endpoint 
Default endpoint is http://locahost, to change endpoint just add -endpoint flag 

```
go run cmdtest.go transaction 2 -endpoint http://myherdiusserver.com 

```

### masteraccountkey

Do do series of transactions we need one account which has the balance, masteraccountkey holds the secret key of that accountt, to pass masteraccountkey use flag masteraccount

```
go run cmdtest.go transaction 2  -masteraccount ecja1TZR8pWr/5OI5j7km4eg+BXh+RB2oN1d4oo6yZE=

```



## Doing Transactions

her 2 reprenst the number of transation need to be done

```
go run cmdtest.go transaction 2

```

## Register  Account

-t is type of asset i.e HER/BTC/ETH all herdius supported account
-s is secret key of the account 

```
go run cmdtest.go account -t BTC -s secretkey

```

