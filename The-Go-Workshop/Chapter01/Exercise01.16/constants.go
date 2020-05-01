package main

import (
    "fmt"
)
// Create a constant which will be a global limit size
const GlobalLimit = 100
const MaxCacheSize int = 10 * GlobalLimit
const (
  CacheKeyBook  = "book_"
  CacheKeyCD    = "cd_"
)
// Declare a map that has a string for a key and a string for its values as cache.
var cache map[string]string
// Create a function to get items from the cache.
func cacheGet(key string) string {
    return cache[key]
}
// Create a function that sets items in the cache.
func cacheSet(key, val string) {
    if len(cache)+1 >= MaxCacheSize {
        return
    }
    cache[key] = val
}
// Create a function to get a book from the cache.
func GetBook(isbn string) string {
    return cacheGet(CacheKeyBook + isbn)
}
// Create a function to add a book to the cache.
func SetBook(isbn string, name string) {
    cacheSet(CacheKeyBook+isbn, name)
}
// Create a function to get a CD data from the cache.
func GetCD(sku string) string {
    return cacheGet(CacheKeyCD + sku)
}
// Create a function to add CDs to the shared cache.
func SetCD(sku string, title string) {
    cacheSet(CacheKeyCD+sku, title)
}

func main() {
    cache = make(map[string]string) // Initialize the cache by creating a map.
    // Add a book to the cache.
    SetBook("1234-5678", "Get Ready To Go.")
    // Add a CD to the cache.
    SetCD("1234-5678", "Get Ready To Go Audio Book")
    // Get and print the Book from the cache.
    fmt.Println("Book   :", GetBook("1234-5678"))
    // Get and print the CD from the cache.
    fmt.Println("CD   :", GetCD("1234-5678"))
}
