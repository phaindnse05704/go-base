package helpers

import (
	"fmt"
	"gem-exp/app/models"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandomString random a string with length n
func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// GetFileContentType get the content type of file
func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

// PrintBeautyRequest ..
func PrintBeautyRequest(r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
}

// PrintBeautyResponse ..
func PrintBeautyResponse(r *http.Response) {
	requestDump, err := httputil.DumpResponse(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
}

// Contains ..
func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

//TokenGenerator generate random token
func TokenGenerator() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func less(a, b models.Transaction) bool {
	return (a.Date < b.Date) ||
		(a.Date == b.Date && (a.ShopifyDomain < b.ShopifyDomain ||
			(a.ShopifyDomain == b.ShopifyDomain && ((a.Event < b.Event) ||
				((a.Event == b.Event) && (a.Details < b.Details))))))
}

//QuickSortTransaction quicksort
func QuickSortTransaction(arr []models.Transaction) []models.Transaction {
	if len(arr) <= 1 {
		return arr
	}
	m := int(len(arr) / 2)
	//fmt.Print(m, len(arr), " ")
	left, right := 0, len(arr)-1
	arr[m], arr[right] = arr[right], arr[m]
	for index := range arr {
		if less(arr[index], arr[right]) {
			arr[index], arr[left] = arr[left], arr[index]
			left++
		}
	}
	//fmt.Println(left)
	arr[left], arr[right] = arr[right], arr[left]
	QuickSortTransaction(arr[:left])
	QuickSortTransaction(arr[left+1:])
	return arr
}
