# Utility commands

### Prerequistie

Master Account Key : one account which has Balance, this account will be used to send tx to other accounts, 


## Default config

### endpoint 
Default endpoint is https://api.herdius.com, to change endpoint just add -endpoint flag 

```
go run cmdtest.go transaction 2 -endpoint http://myherdiusserver.com 

```

### masteraccountkey

Do do series of transactions we need one account which has the balance, masteraccountkey holds the secret key of that accountt, to pass masteraccountkey use flag masteraccount

```
go run cmdtest.go transaction 2  

```

### Output
```
2019/10/16 12:02:37 Account response:  {"nonce":5,"address":"HDzLGL98C4vKtVWb3qzm92C2LX2V5kNhXR","balance":7500,"storageRoot":"10DCAFC3F97DF37C2E91357FFF13E4675372B949E1442D1C00CEAD6F8A9806DA","publicKey":"8YHKuCEDvZ+MEyEyQOA/4NAk6Q8SeB/vpdqv30mqODMaQQaMFaE=","erc20Address":"","FirstExternalAddress":{},"ExternalNonce":0,"LastBlockHeight":0,"EBalances":{}}

2019/10/16 12:02:37 Last Nonce 5
2019/10/16 12:02:37 Account Address : HDzLGL98C4vKtVWb3qzm92C2LX2V5kNhXR
2019/10/16 12:02:37 {"tx_id":"HTxc9bbaf7e6dd23c28bdb2a3c49a2b93f1c7e42724","pending":1,"status":"success"}

2019/10/16 12:02:37 ---------
2019/10/16 12:02:37 Tx executed {HTxc9bbaf7e6dd23c28bdb2a3c49a2b93f1c7e42724 1 0 success  {} [] 0}
2019/10/16 12:02:37 Account Address : HDzLGL98C4vKtVWb3qzm92C2LX2V5kNhXR
2019/10/16 12:02:37 {"tx_id":"HTxfadc073c0a42f1d9726b8784753f0c20469e8217","pending":1,"queued":1,"status":"success"}

2019/10/16 12:02:37 ---------
2019/10/16 12:02:37 Tx executed {HTxfadc073c0a42f1d9726b8784753f0c20469e8217 1 1 success  {} [] 0}
2019/10/16 12:02:37 Account Address : HDzLGL98C4vKtVWb3qzm92C2LX2V5kNhXR
2019/10/16 12:02:37 {"tx_id":"HTx5f08de4ee182b47b151967ec59494c0ce73549d5","pending":1,"queued":2,"status":"success"}

2019/10/16 12:02:37 ---------
2019/10/16 12:02:37 Tx executed {HTx5f08de4ee182b47b151967ec59494c0ce73549d5 1 2 success  {} [] 0}
```



## Doing Transactions

her 2 reprenst the number of transation need to be done

```
go run cmdtest.go transaction 2

```

# Account

# Get Address using Secret

```
go run cmdtest.go account address -s secretkey

Output
Address:  HSYxVdh3cW1MY56dLrQoHX2yAw7BtYgPSL

```


# Register Account

-t is type of asset i.e HER/BTC/ETH all herdius supported account
-s is secret key of the account 

```
go run cmdtest.go account -t BTC -s secretkey
```

#### Output
```
2019/10/16 11:56:58 Last Nonce 0
2019/10/16 11:56:58 Account Address : HSYxVdh3cW1MY56dLrQoHX2yAw7BtYgPSL
2019/10/16 11:56:58 {HSYxVdh3cW1MY56dLrQoHX2yAw7BtYgPSL AgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA  category:"crypto" symbol:"HER" network:"Herdius" value:15  Register account dnmwiCcRLoQOWqgji8TOXTmtMEETrut3JFjmzjQLYQ51Ims1Zih/Je3+BqmY7EOKcsGP0sSE1WcaBExHuWTJrA== update   map[] {} [] 0}
2019/10/16 11:56:58 {"tx_id":"HTxe7c1528d1d98d3ae4e678b8db0e2b5cf2a7e0986","status":"success"}

2019/10/16 11:56:58 ---------
2019/10/16 11:56:58 {HTxe7c1528d1d98d3ae4e678b8db0e2b5cf2a7e0986 0 0 success  {} [] 0} <nil>
```

# Register External Address

-t is type of asset i.e HER/BTC/ETH all herdius supported account
-s is secret key of the account 
-e is External Account Address

```
go run cmdtest.go account  external -t BTC -s secretkey -e 

```

#### Output
```
2019/10/16 11:56:58 Last Nonce 0
2019/10/16 11:56:58 Account Address : HSYxVdh3cW1MY56dLrQoHX2yAw7BtYgPSL
2019/10/16 11:56:58 {HSYxVdh3cW1MY56dLrQoHX2yAw7BtYgPSL AgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA  category:"crypto" symbol:"HER" network:"Herdius" value:15  Register account dnmwiCcRLoQOWqgji8TOXTmtMEETrut3JFjmzjQLYQ51Ims1Zih/Je3+BqmY7EOKcsGP0sSE1WcaBExHuWTJrA== update   map[] {} [] 0}
2019/10/16 11:56:58 {"tx_id":"HTxe7c1528d1d98d3ae4e678b8db0e2b5cf2a7e0986","status":"success"}

2019/10/16 11:56:58 ---------
2019/10/16 11:56:58 {HTxe7c1528d1d98d3ae4e678b8db0e2b5cf2a7e0986 0 0 success  {} [] 0} <nil>
```

