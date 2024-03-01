package mail

import (
	"crypto/rand"
	"fmt"
	"net/smtp"
	"strings"
	"time"

	"github.com/jordan-wright/email"
	"github.com/k0yote/backend-wallet/config"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type Passcode struct {
	Code   string
	Secret string
}

func IssuePassCode(config config.Config, email string) (Passcode, error) {

	var result Passcode

	period := config.PassCodeExpirePeriod
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      config.OTPIssuer,
		AccountName: email,
		Period:      uint(period),
		Digits:      otp.DigitsSix,
		SecretSize:  20,
		Secret:      []byte{},
		Algorithm:   otp.AlgorithmSHA512,
		Rand:        rand.Reader,
	})
	if err != nil {
		return result, err
	}

	passcode, err := generatePassCode(period, key.Secret())
	if err != nil {
		return result, err
	}

	valid, err := validatePassCode(period, key.Secret(), passcode)
	if err != nil {
		return result, err
	}

	if !valid {
		return result, fmt.Errorf("generated passcode is not valid")
	}

	return Passcode{
		Code:   passcode,
		Secret: key.Secret(),
	}, nil
}

func generatePassCode(period int, secret string) (string, error) {
	return totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    uint(period),
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA512,
	})
}

func validatePassCode(period int, secret string, passcode string) (bool, error) {
	return totp.ValidateCustom(passcode, secret, time.Now().UTC(), totp.ValidateOpts{
		Period:    uint(period),
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA512,
	})
}

func SendEmail(config config.Config, to string, passcode string) error {
	e := email.NewEmail()
	e.From = config.EmailSenderAddress
	e.To = []string{to}
	e.Subject = fmt.Sprintf("%v is your login code for Dummy Auth Demo", passcode)
	e.HTML = []byte(strings.Replace(mailTemplate, "%v", passcode, 1))

	smtpAuth := smtp.PlainAuth("", config.EmailSenderAddress, config.EmailSenderPassword, config.SmtpAuthAddress)
	return e.Send(config.SmtpServerAddress, smtpAuth)
}
