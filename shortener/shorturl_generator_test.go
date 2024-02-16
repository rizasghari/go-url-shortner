package shortener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const UserId = "507f1f77bcf86cd799439011"

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://developer.redis.com/howtos/solutions/?_gl=1*1wkqqrn*_gcl_au*OTA3Mjc2Mzg0LjE3MDgwODQ4OTA.*_ga*MTg3NTUzMjAxMC4xNzA4MDg0ODkw*_ga_8BKGRQKRPV*MTcwODA5MjAxOC4yLjEuMTcwODA5MjAyMS41Ny4wLjA."
	shortLink_1 := GenerateShortLink(initialLink_1, UserId)

	initialLink_2 := "https://www.makeuseof.com/go-urls-net-package/"
	shortLink_2 := GenerateShortLink(initialLink_2, UserId)

	assert.Equal(t, shortLink_1, "PWa3BKsq")
	assert.Equal(t, shortLink_2, "gTtuPJuM")
}