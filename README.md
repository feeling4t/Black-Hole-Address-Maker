# Black-Hole-Address-Maker

Make a legal bitcoin-like BlackHole address

#### WARN: ANY COIN SEND TO A BLACKHOLE ADDRESS WILL NEVER GET BACK 

Build
-----
    git clone https://github.com/feeling4t/Black-Hole-Address-Maker.git
    cd Black-Hole-Address-Maker
    go build makeblackhole.go

Usage:
-----
    makeblackhole prefix
The prefix string must NOT longer than 28

Example:
--------
    makeblackhole 1bitcoinDestroy
    The address is: 1bitcoinDestroy11111111111111qJkq5
    makeblackhole LiteCoinDestroy
    The address is: LiteCoinDestroy11111111111115LGn1H
    makeblackhole BLackCoinDestroy
    The address is: BLackCoinDestroy1111111111113WGuzJ
    makeblackhole 1bitcoinDestroyXXXXXXXXXXXXX
    The address is: 1bitcoinDestroyXXXXXXXXXXXXWvxdT7n
