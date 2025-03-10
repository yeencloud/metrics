package metrics

import (
	"github.com/yeencloud/lib-shared"
	"github.com/yeencloud/lib-shared/log"
)

type MetricPoint struct {
	Name string
	Tags map[string]string
}

type MetricValues map[string]any

type MetricsInterface interface {
	LogPoint(point MetricPoint, value MetricValues)

	Connect() error
}

func MetricsFromContext(ctx *shared.Context) (pointTags map[string]string, points map[string]MetricValues) {
	pointTags = map[string]string{}
	points = map[string]MetricValues{}

	ctx.Range(func(key, value interface{}) bool {
		entryKey, ok := key.(log.Path)

		if !ok {
			return true
		}

		root := entryKey.Root().String()
		metricKey := entryKey.MetricKey()

		if entryKey.IsMetricTag {
			pointTags[metricKey], ok = value.(string)
			return ok
		}

		if _, ok := points[entryKey.Root().String()]; !ok {
			points[root] = map[string]any{}
		}

		points[root][metricKey] = value

		return true
	})

	return
}

func LogsFromContext(ctx *shared.Context) (pointTags map[string]string, points []map[string]interface{}) {
	pointTags, _ = MetricsFromContext(ctx)
	points = ctx.Logs()

	return
}
