package logic

import (
	"context"
	"testing"

	server "github.com/moznobkin/sonar-coverage/generated/sonar-coverage-server/model"
	"github.com/stretchr/testify/require"
)

func Test_MonthlyTimesheet_real(t *testing.T) {

	json.RegisterExtension(&timeDecodeExtension{})
	products, err := ReadStruct[[]server.Product]("../data/json/products.json")
	require.NoError(t, err)

	cov, err := AggregateMetrics(context.TODO(), *products)

	require.EqualValues(t, 2, len(cov.Products))

}
