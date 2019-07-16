package tdd

import (
	"regexp"
	"strconv"
	"strings"

	humanize "github.com/dustin/go-humanize"
)

// FormatRupiah to convert number 10000 to Rp 10.000
func FormatRupiah(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(amount, 0)
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)
	return "Rp " + stringValue
}

// GetAmount for decode Rp 1.000 to 1000
func GetAmount(rupiah string) int64 {
	re := regexp.MustCompile("[0-9]+")
	getNumber := re.FindAllString(rupiah, -1)
	amountString := strings.Join(getNumber, "")
	amount, _ := strconv.ParseInt(amountString, 10, 64)
	return amount
}
