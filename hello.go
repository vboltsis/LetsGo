package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	//fmt.Println("hello world")
	fmt.Println(getTime())
}

//fibonacci
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

//random string
func randomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

//find prime numbers
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//validate email
func isValidEmail(email string) bool {
	if len(email) > 3 && strings.Contains(email, "@") && strings.Contains(email, ".") {
		return true
	}
	return false
}

//get google location from coordinates
func getLocation(lat, lon float64) string {
	return fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?latlng=%f,%f&key=AIzaSyDfYhQNzjxSfQQnjZLz1gJ5-Xm4jnHm4y4", lat, lon)
}

//get current time
func getTime() string {
	return fmt.Sprintf("%s", time.Now())
}

//binary search
func binarySearch(arr []int, l, r, x int) int {
	if r >= l {
		mid := l + (r-l)/2
		if arr[mid] == x {
			return mid
		}
		if arr[mid] > x {
			return binarySearch(arr, l, mid-1, x)
		}
		return binarySearch(arr, mid+1, r, x)
	}
	return -1
}

//hash sha256
func hashSha256(data string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
}

//salt and hash password
func saltHashPassword(password string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password+"salt")))
}

//find the largest number in an array
func largestNumber(arr []int) int {
	var largest = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
	}
	return largest
}
