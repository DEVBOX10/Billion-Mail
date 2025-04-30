package maillog_stat

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// GetLastAggregateTimeMillis retrieves the last mail log aggregation time in milliseconds
func GetLastAggregateTimeMillis() int64 {
	val, err := g.DB().Model("bm_options").Where("name", "last_maillog_aggregate_time_millis").Value("value")

	if err != nil {
		g.Log().Error(context.Background(), "GetLastAggregateTimeMillis error:", err.Error())
		return 0
	}

	if val == nil {
		return 0
	}

	return val.Int64()
}

// SetLastAggregateTimeMillis sets the last mail log aggregation time in milliseconds
func SetLastAggregateTimeMillis(timeMillis int64) {
	_, err := g.DB().Model("bm_options").OnConflict("name").OnDuplicate(g.Map{
		"value": gdb.Raw("excluded.value"),
	}).Save(g.Map{
		"name":  "last_maillog_aggregate_time_millis",
		"value": timeMillis,
	})

	if err != nil {
		g.Log().Error(context.Background(), "SetLastAggregateTimeMillis error:", err.Error())
		return
	}
}

// AggregateMaillogs aggregates mail logs from the last aggregation time to the current time
func AggregateMaillogs() (err error) {
	g.Log().Debug(context.Background(), "aggregating maillogs...")

	startTimeMillis := GetLastAggregateTimeMillis()
	endTimeMillis := time.Now().UnixMilli()

	// get all active domains
	var allDomains []string
	cols, err := g.DB().Model("domain").Where("active=1").Array("domain")

	if err != nil {
		g.Log().Error(context.Background(), "Failed to get all domain", err)
		return
	}

	allDomains = make([]string, len(cols))

	for i, col := range cols {
		allDomains[i] = col.String()
	}

	if len(allDomains) == 0 {
		g.Log().Debug(context.Background(), "no domain found, skip aggregate maillogs")
		return
	}

	if startTimeMillis == 0 {
		startTimeMillis = endTimeMillis - (90 * 86400 * 1000) // default last 90 days
	}

	latestTimeMillis := startTimeMillis

	query := g.DB().Model("mailstat_senders s").
		InnerJoin("mailstat_receive_mails rm", "s.postfix_message_id=rm.postfix_message_id").
		Where("s.log_time_millis > ?", startTimeMillis).
		Where("s.log_time_millis < ?", endTimeMillis).
		Fields("rm.*")

	// build where
	domainLen := len(allDomains)
	where1 := make([]string, domainLen)
	where2 := make([]string, domainLen)
	binds := make([]interface{}, domainLen*2)

	for i, domain := range allDomains {
		where1[i] = fmt.Sprintf("s.sender LIKE ?")
		binds[i] = "%@" + domain
		where2[i] = fmt.Sprintf("rm.recipient LIKE ?")
		binds[domainLen+i] = "%@" + domain
	}

	query.Where(fmt.Sprintf("(%s) AND (%s)", gstr.Join(where1, " OR "), gstr.Join(where2, " OR ")), binds...)
	ret, err := query.All()

	if err != nil {
		g.Log().Error(context.Background(), "aggregate maillog failed", err)
		return
	}

	retLen := len(ret)
	insertData := make([]map[string]interface{}, retLen)

	for i, record := range ret {
		logTimeMillis := record["log_time_millis"].Int64()
		if latestTimeMillis < logTimeMillis {
			latestTimeMillis = logTimeMillis
		}
		insertData[i] = record.Map()
		insertData[i]["mail_provider"] = "local"
		delete(insertData[i], "log_time")
	}

	if retLen > 0 {
		_, err = g.DB().Model("mailstat_send_mails").Batch(5000).InsertIgnore(insertData)

		if err != nil {
			g.Log().Error(context.Background(), "aggregate maillog failed", err)
			return
		}
	}

	SetLastAggregateTimeMillis(latestTimeMillis)

	g.Log().Debug(context.Background(), fmt.Sprintf("aggregating maillogs >>> Done -- latest_time_millis: %d", latestTimeMillis))
	return
}

// AggregateMaillogsTask periodically aggregates mail logs
func AggregateMaillogsTask(interval time.Duration) {
	for {
		time.Sleep(interval)

		err := AggregateMaillogs()
		if err != nil {
			g.Log().Debug(context.Background(), "aggregate maillogs failed: ", err)
		}
	}
}
