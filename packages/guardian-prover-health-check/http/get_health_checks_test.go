package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cyberhorsey/webutils/testutils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	guardianproverhealthcheck "github.com/taikoxyz/taiko-mono/packages/guardian-prover-health-check"
)

func Test_GetHealthChecks(t *testing.T) {
	srv := newTestServer("")

	err := srv.healthCheckRepo.Save(guardianproverhealthcheck.SaveHealthCheckOpts{
		GuardianProverID: 1,
		Alive:            true,
		ExpectedAddress:  "0x123",
		RecoveredAddress: "0x123",
		SignedResponse:   "0x123",
	})

	assert.Nil(t, err)

	assert.Equal(t, nil, err)

	tests := []struct {
		name                  string
		wantStatus            int
		wantBodyRegexpMatches []string
	}{
		{
			"success",
			http.StatusOK,
			// nolint: lll
			[]string{`[{"id":0,"guardianProverId":1,"alive":true,"expectedAddress":"0x123","recoveredAddress":"0x123","signedResponse":"0x123"}]`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := testutils.NewUnauthenticatedRequest(
				echo.GET,
				"/healthchecks",
				nil,
			)

			rec := httptest.NewRecorder()

			srv.ServeHTTP(rec, req)

			testutils.AssertStatusAndBody(t, rec, tt.wantStatus, tt.wantBodyRegexpMatches)
		})
	}
}
