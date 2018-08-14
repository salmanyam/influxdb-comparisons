package influxdb

import "time"
import bulkQuerygen "github.com/influxdata/influxdb-comparisons/bulk_query_gen"

// InfluxDevopsGroupby produces Influx-specific queries for the devops groupby case.
type InfluxDevopsGroupby struct {
	InfluxDevops
}

func NewInfluxQLDevopsGroupBy(dbConfig bulkQuerygen.DatabaseConfig, start, end time.Time) bulkQuerygen.QueryGenerator {
	underlying := newInfluxDevopsCommon(InfluxQL, dbConfig, start, end).(*InfluxDevops)
	return &InfluxDevopsGroupby{
		InfluxDevops: *underlying,
	}

}

func NewFluxDevopsGroupBy(dbConfig bulkQuerygen.DatabaseConfig, start, end time.Time) bulkQuerygen.QueryGenerator {
	underlying := newInfluxDevopsCommon(Flux, dbConfig, start, end).(*InfluxDevops)
	return &InfluxDevopsGroupby{
		InfluxDevops: *underlying,
	}

}

func (d *InfluxDevopsGroupby) Dispatch(i, scaleVar int) bulkQuerygen.Query {
	q := bulkQuerygen.NewHTTPQuery() // from pool
	d.MeanCPUUsageDayByHourAllHostsGroupbyHost(q, scaleVar)
	return q
}
