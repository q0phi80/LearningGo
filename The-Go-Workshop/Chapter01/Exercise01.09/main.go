package main

import (
    "fmt"
)

func main(){
    // Create a variable to hold the total.
    var total float64 = 2 * 13 //Customer purchase two items at $13 each.
    fmt.Println("Sub : ", total)
    // Add drinks
    total = total + (4 * 2.25) // Customer purchases four drinks at $2.25 each so we multiple by four and then add the drink's total to the previous total.
    fmt.Println("Sub : ", total)
    // Subtract a discount for the Customer
    total = total - 5
    fmt.Println("Sub : ", total)
    // Use multiplication to calcuate a 10% tip
    tip := total * 0.1
    fmt.Println("Tip : ", tip)
    // Add tip to the total
    total = total + tip
    fmt.Println("Total : ", total)
    // The bill will be split by two people.
    split := total / 2
    fmt.Println("Split : ", split)
    // Calculate whether the customer gets a reward based on the number of visits to the restaurant.
    // Reward for every 5th visit
    visitCount := 24 // Setup visitCount initial value
    visitCount = visitCount + 1 // Add $1 to the visit.
    // Use % to get the remainder after dividing the visitCount by $5.
    remainder := visitCount % 5
    // The customer gets a reward on every 5th visit so if the remainder is 0, then this is one of those visits.
    // We check if the remainder is 0 and print a message that they get a reward.
    if remainder == 0 {
        fmt.Println("[+] With this visit, you've earned a reward.")
    }
}
