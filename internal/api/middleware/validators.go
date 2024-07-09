package middleware

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"payment-payments-api/pkg/umdw"
	"regexp"
	"strconv"
	"strings"
)

var PasswordValidation = umdw.VerificationKeyFunction{
	Func:   func(val interface{}) bool { return len(val.(string)) >= 6 },
	ErrMsg: "Password invalid. It must be at least 6 characters.",
}

var EmailValidation = umdw.VerificationKeyFunction{
	Func: func(val interface{}) bool {
		e := "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
		re := regexp.MustCompile(e)
		return re.MatchString(val.(string))
	},
	ErrMsg: "Email invalid. Ref: example@gmail.com",
}

var BoolValidation = umdw.VerificationKeyFunction{
	Func:   func(val interface{}) bool { _, ok := val.(bool); return ok },
	ErrMsg: "Boolean invalid. Ref: true or false",
}

var NumberPositiveValidation = umdw.VerificationKeyFunction{
	Func:   func(val interface{}) bool { n, _ := val.(float64); return n > 0 },
	ErrMsg: "Number invalid. Ref: N > 0",
}

var ObjectIDValidation = umdw.VerificationKeyFunction{
	Func: func(val interface{}) bool {
		_, err := primitive.ObjectIDFromHex(val.(string))
		return err == nil
	},
	ErrMsg: "ObjectID invalid. Ref: 62a53f7fd733463b73fdf655",
}

var TimeValidation = umdw.VerificationKeyFunction{
	Func: func(val interface{}) bool {
		s, ok := val.(string)
		if !ok || len(s) != 5 || s[2] != ':' {
			return false
		}
		if _, err := strconv.Atoi(string([]byte{s[0], s[1]})); err != nil {
			return false
		}
		if _, err := strconv.Atoi(string([]byte{s[3], s[4]})); err != nil {
			return false
		}
		return true
	},
	ErrMsg: "Time invalid. Ref. 15:37",
}

var RutValidation = umdw.VerificationKeyFunction{
	Func: func(val interface{}) bool {
		rut, ok := val.(string)
		if !ok || len(rut) <= 3 {
			return false
		}
		alphaNumericRegex := regexp.MustCompile("[^A-Za-z0-9]+")
		rutRegex := regexp.MustCompile("^[0-9]+K?$")
		if rut[len(rut)-2] != uint8(0x2d) || rut[len(rut)-6] != uint8(0x2e) || rut[len(rut)-10] != uint8(0x2e) {
			return false
		}
		rut = alphaNumericRegex.ReplaceAllString(rut, "")
		rut = strings.ToUpper(strings.TrimSpace(rut))
		if len(rut) > 9 {
			return false
		}
		if !rutRegex.MatchString(rut) {
			return false
		}
		for len(rut) < 9 {
			rut = "0" + rut
		}
		r, _ := strconv.Atoi(rut[:len(rut)-1])
		m, s := 0, 1
		for ; r != 0; r /= 10 {
			s = (s + r%10*(9-m%6)) % 11
			m++
		}
		if s != 0 {
			s += 47
		} else {
			s = 75
		}
		dv := rut[len(rut)-1]
		return rune(dv) == rune(s)
	},
	ErrMsg: "RUT invalid. Ref. 11.111.111-1",
}

var OptionValidation = func(options []string) umdw.VerificationKeyFunction {
	return umdw.VerificationKeyFunction{
		Func: func(val interface{}) bool {
			o := val.(string)
			for _, option := range options {
				if o == option {
					return true
				}
			}
			return false
		},
		ErrMsg: "Option invalid. Ref. " + strings.Join(options, ", "),
	}
}
