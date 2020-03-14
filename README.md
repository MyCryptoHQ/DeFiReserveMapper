# DeFiReserveMapper
Used for mapping DeFi pool tokens to their related reserves for caching in use by beta.mycrypto.com


##### Requires:
`go version go1.13.5`

### ToDo:
- [ ] Handle compound assets.
- [ ] Handle compound-ether asset (cETH => its different because its not an erc20 token).
- [ ] Figure out caching on AWS.
- [ ] Figure out setting up timers to handle time events.

### To run:
`cd app && go build && ./app`

We still need to handle this info to cache it, but this will log out [here](https://github.com/MyCryptoHQ/DeFiReserveMapper/blob/master/app/main.go#L34) an object that looks like this:
```
{
  'UUID OF POOL TOKEN': {
    type: 'uniswap' | 'compound' | 'compound-ether'
    reserveRates: [{
      assetId: 'UUID OF RESERVE ASSET #1',
      rate:	0.000001
    },{
      assetId: 'UUID OF RESERVE ASSET #2',
      rate:	0.01223
    }]		
  }	
}
```

Implements: https://docs.google.com/document/d/1fHQyxPfOpQtptS9pqSKW5B7axm10bCq8hsHF6WTL8jg
