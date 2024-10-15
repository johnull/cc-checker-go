package luhn

import (
  "strconv"
  "os"
  "bufio"
  "strings"
  "fmt"
  "regexp"
)

func Checker() {
  cardList := []string{}
  input := bufio.NewReader(os.Stdin)

  inputLine, err := input.ReadString('\n')
  if err != nil {
    fmt.Println("Something went wrong", err)
    return
  }

  inputTrim := strings.TrimSpace(inputLine)
  inputList := strings.Split(inputTrim, ",")

  // handle input separated by spaces
  if len(inputList) < 2 {
    inputList = strings.Split(inputTrim, " ")
  }

  for _, cur := range inputList {
    card := strings.TrimSpace(cur)
    cardList = append(cardList, card)
  }

  printCards(cardList)
}

func printCards(cardList []string) {
  bins := map[string]string {
    "Amex" : "^3[47][0-9]{13}$",
    "Visa" : "^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14})$",
    "ELO": `^(?:40117[8-9]|431274|438935|451416|457393|45763[1-2]|504175|627780|636297|636368|65500[0-1]|65165[2-4]|65048[5-8]|506699|5067[0-6]\d|50677[0-8]|509\d{3})\d{10}$`,
    "Visa Mastercard" : "^(5[1-5][0-9]{14}|2(22[1-9][0-9]{12}|2[3-9][0-9]{13}|[3-6][0-9]{14}|7[0-1][0-9]{13}|720[0-9]{12}))$",
    "BCGlobal" : "^(6541|6556)[0-9]{12}$",
    "Carte Blanche" : "^389[0-9]{11}$",
    "Diners Club" : "^3(?:0[0-5]|[68][0-9])[0-9]{11}$",
    "Discover" : "^65[4-9][0-9]{13}|64[4-9][0-9]{13}|6011[0-9]{12}|(622(?:12[6-9]|1[3-9][0-9]|[2-8][0-9]{3}|9[01][0-9]|92[0-5])[0-9]{10})$",
    "Insta Payment" : "^63[7-9][0-9]{13}$",
    "Korean Local" : "^9[0-9]{15}$",
    "Laser" : "^(6304|6706|6709|6771)[0-9]{12,15}$",
    "Maestro" : "^(5018|5020|5038|6304|6759|6761|6763)[0-9]{8,15}$",
    "Solo" : "^(6334|6767)[0-9]{12}|(6334|6767)[0-9]{14}|(6334|6767)[0-9]{15}$",
    "Switch" : "^(4903|4905|4911|4936|6333|6759)[0-9]{12}|(4903|4905|4911|4936|6333|6759)[0-9]{14}|(4903|4905|4911|4936|6333|6759)[0-9]{15}|564182[0-9]{10}|564182[0-9]{12}|564182[0-9]{13}|633110[0-9]{10}|633110[0-9]{12}|633110[0-9]{13}$",
    "Union Pay" : "^(62[0-9]{14,17})$",
  }

  validCardList := map[string][]string{}

  for _, card := range cardList {
    if !isValidLuhn(card) {
      continue
    }

    for cardName, cardBin := range bins {
      if match, _ := regexp.MatchString(cardBin, card); match {
        validCardList[cardName] = append(validCardList[cardName], card)
      }
    }
  }

  if len(validCardList) == 0 {
    fmt.Println("No valid card found.")
    return
  }

  for cardName, cards := range validCardList {
    for _, card := range cards {
      fmt.Printf("%s: %s [LIVE]\n", cardName, card)
    }
  }
}

 func isValidLuhn(cardList string) bool {
   sum := 0

   for i := len(cardList) - 1; i >= 0; i-- {
     cur, _ := strconv.Atoi(string(cardList[i]))

     if (len(cardList) - i) % 2 == 0 {
       cur *= 2
       sum += cur % 10 + cur / 10
     } else {
         sum += cur
     }
   }

   return sum % 10 == 0;
 }

