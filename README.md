# cc-checker-go

Simple credit card validation [Luhn algo](http://en.wikipedia.org/wiki/Luhn_algorithm)

## Installation
```bash
git clone git@github.com:johnull/cc-checker-go.git
```
## Usage
```bash
cd cc-checker-go/
```
```bash
go run main.go
```
**Input**
```bash
4111111111111111 4012888888881881 378282246310005 371449635398431
```
**Output**
```bash
Visa: 4111111111111111 [LIVE]
Visa: 4012888888881881 [LIVE]
Amex: 378282246310005 [LIVE]
Amex: 371449635398431 [LIVE]
```


