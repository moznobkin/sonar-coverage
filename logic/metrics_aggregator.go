package logic

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/antihax/optional"
	client "github.com/moznobkin/sonar-coverage/generated/sonar-client"
	server "github.com/moznobkin/sonar-coverage/generated/sonar-coverage-server/model"
)

func AggregateMetrics(ctx context.Context, req []server.Product) (*server.Coverage, error) {
	result := &server.Coverage{Date: time.Now()}
	for _, p := range req {
		pc := server.ProductCoverage{Name: p.Name}
		for _, r := range p.Repos {
			res, err := requestRepoMetrics(ctx, r.Id, r.Branch)
			if err == nil {
				rc := &server.RepoCoverage{Id: r.Id, Branch: r.Branch}
				for _, m := range res.Component.Measures {
					if m.Metric == "coverage" {
						i, err := strconv.ParseFloat(m.Value, 64)
						if err == nil {
							rc.Coverage = i
						}
					}
					if m.Metric == "lines_to_cover" {
						i, err := strconv.Atoi(m.Value)
						if err == nil {
							rc.LinesToCover = int32(i)
						}
					}
					if m.Metric == "tests" {
						i, err := strconv.Atoi(m.Value)
						if err == nil {
							rc.UnitTests = int32(i)
						}
					}
				}
				pc.Repos = append(pc.Repos, *rc)
			} else {
				result.Errors = append(result.Errors, r.Id+"/"+r.Branch+" error "+err.Error())
			}
		}
		result.Products = append(result.Products, pc)

	}
	return result, nil
}

func requestRepoMetrics(ctx context.Context, repoId, branch string) (*client.ProjectMetrics, error) {
	proxyUrl, err := url.Parse("http://192.168.191.10:8888")
	httpClient := &http.Client{
		Timeout:   15 * time.Second,
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
	}
	cfg := &client.Configuration{
		BasePath:   "https://sonar.vimpelcom.ru/api",
		HTTPClient: httpClient,
	}
	c := client.NewAPIClient(cfg)

	authCtx := context.WithValue(context.Background(), client.ContextBasicAuth, client.BasicAuth{
		UserName: "squ_2e4c4496649dbbd3f55d05d7d8ae9aa24c0ba994",
	})
	out, _, err := c.CoverageApi.MeasuresComponentGet(authCtx, repoId, &client.CoverageApiMeasuresComponentGetOpts{Branch: optional.NewString(branch), MetricKeys: optional.NewString("ncloc,complexity,violations, coverage, lines_to_cover, lines, tests")})

	return &out, err
}
