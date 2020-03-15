# DeFiReserveMapper
Used for mapping DeFi pool tokens to their related reserves for caching in use by beta.mycrypto.com


##### Requires:
`go version go1.13.5`

### ToDo:
~- [ ] Figure out bottlenecking requests so we don't spam the node and get blocked~ _I don't think we need to do this actually_
- [x] Handle compound assets.
- [x] Handle uniswap assets.
- [ ] Figure out caching on AWS.
- [x] Figure out setting up handling of update intervals.

### To run:
`cd app && go build && ./app`

We still need to handle this info to cache it, but this will print out [here](https://github.com/MyCryptoHQ/DeFiReserveMapper/blob/master/outputFile.json) an object that looks like this:
```
{
    "4f96a9e6-bf30-54d0-90c0-3d6e7d7042f2"[UUID OF COMPOUND POOL TOKEN]: {
        "type": "compound",
        "lastUpdated": 1584237806,
        "reserveRates": [
            {
                "assetId": "356a192b-7913-504c-9457-4d18c28d46e6",
                "rate": "0.020011525881693628"
            }
        ]
    },
    "50b83702-2652-5e12-8585-cd8014641b74"[UUID OF UNISWAP POOL TOKEN]: {
        "type": "uniswap",
        "lastUpdated": 1584233100,
        "reserveRates": [
            {
                "assetId": "356a192b-7913-504c-9457-4d18c28d46e6",
                "rate": "1.1344950676786816"
            },
            {
                "assetId": "9454eb02-dda1-53d1-85f0-f6faa6c89267",
                "rate": "130.9760441345517"
            }
        ]
    }
}
```

Implements: https://docs.google.com/document/d/1fHQyxPfOpQtptS9pqSKW5B7axm10bCq8hsHF6WTL8jg
