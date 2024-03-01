package mail

import (
	"testing"

	"github.com/k0yote/backend-wallet/config"
	"github.com/stretchr/testify/require"
)

func issuePassCode(t *testing.T, config config.Config, email string) Passcode {
	passCode, err := IssuePassCode(config, email)
	require.NoError(t, err)

	return passCode
}

func TestPasscodeValidate(t *testing.T) {
	config, err := config.LoadConfig()
	require.NoError(t, err)

	p := issuePassCode(t, config, "aaa@aaa.com")

	b, err := validatePassCode(config.PassCodeExpirePeriod, p.Secret, p.Code)
	require.NoError(t, err)

	require.True(t, b)
}

func TestSendMail(t *testing.T) {
	config, err := config.LoadConfig()
	require.NoError(t, err)

	to := "jongm.yu@gmail.com"

	p := issuePassCode(t, config, to)

	require.NoError(t, SendEmail(config, to, p.Code))
}
