package int2vietnamese

import (
	"fmt"
	"strings"
)

var thousandLevelVal = []string{"", "nghìn", "triệu", "tỷ", "nghìn tỷ", "tỷ tỷ"}
var digit = []string{"không", "một", "hai", "ba", "bốn", "năm", "sáu", "bảy", "tám", "chín"}

func read3DigitWithBigestLevel(in int) string {
	if in < 10 {
		return digit[in]
	}

	if in < 100 {
		return read2Digit(in)
	}

	if in%100 == 0 {
		return fmt.Sprintf("%s trăm", digit[in/100])
	}

	return fmt.Sprintf("%s trăm %s", digit[in/100], read2Digit(in%100))
}

func read2Digit(in int) string {
	if in < 10 {
		return fmt.Sprintf("lẻ %s", digit[in])
	}

	if in == 10 {
		return "mười"
	}

	if in < 20 {
		switch in % 10 {
		case 5:
			return "mười lăm"
		default:
			return fmt.Sprintf("mười %s", digit[in%10])
		}
	}

	if in%10 == 0 {
		return fmt.Sprintf("%s mươi", digit[in/10])
	}

	switch in % 10 {
	case 1:
		return fmt.Sprintf("%s mươi mốt", digit[in/10])
	case 4:
		return fmt.Sprintf("%s mươi tư", digit[in/10])
	case 5:
		return fmt.Sprintf("%s mươi lăm", digit[in/10])
	default:
		return fmt.Sprintf("%s mươi %s", digit[in/10], digit[in%10])
	}

}

func read3Digit(isBigestLevel bool, in int) string {
	if isBigestLevel {
		return read3DigitWithBigestLevel(in)
	}

	if in == 0 {
		return ""
	}

	if in%100 == 0 {
		return fmt.Sprintf("%s trăm", digit[in/100])
	}

	return fmt.Sprintf("%s trăm %s", digit[in/100], read2Digit(in%100))
}

func Int2VietNamese(in int) string {
	ans := read3Digit(in/1000 == 0, in%1000)
	thousandLevel := 1
	in /= 1000
	for {
		if in == 0 {
			break
		}

		tmp := in % 1000
		resultCurrent3digit := read3Digit(in/1000 == 0, tmp)
		if len(resultCurrent3digit) > 0 {
			ans = fmt.Sprintf("%s %s %s", resultCurrent3digit, thousandLevelVal[thousandLevel], ans)
		}

		in /= 1000
		thousandLevel += 1
	}

	return strings.TrimSpace(ans)
}
