package maillog_stat

import (
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"sort"
	"time"
)

// Overview maillog data overview structure
type Overview struct{}

// NewOverview new overview instance
func NewOverview() *Overview {
	return &Overview{}
}

// filterAndPrepareTimeSection filter and provide time section
func (o *Overview) filterAndPrepareTimeSection(startTime, endTime int64) (int64, int64) {
	if startTime > 0 && endTime < 0 {
		endTime = time.Now().Unix()
	}

	if startTime > endTime {
		panic(public.Lang("end_time must greater than start_time"))
	}

	return startTime, endTime
}

// buildBaseQuery build basic query
func (o *Overview) buildBaseQuery(campaignID int64, domain string, startTime, endTime int64) *gdb.Model {
	query := g.DB().Model("mailstat_send_mails sm")
	query.LeftJoin("mailstat_senders s", "sm.postfix_message_id=s.postfix_message_id")
	query.Where("s.postfix_message_id is not null")

	if domain != "" {
		query.Where("s.sender LIKE ?", "%@"+domain)
	}

	if startTime > 0 {
		query.Where("sm.log_time_millis > ?", startTime*1000-1)
	}

	if endTime > 0 {
		query.Where("sm.log_time_millis < ?", endTime*1000+1)
	}

	return query
}

// Overview overview the maillog
func (o *Overview) Overview(campaignID int64, domain string, startTime, endTime int64) map[string]interface{} {
	startTime, endTime = o.filterAndPrepareTimeSection(startTime, endTime)

	return map[string]interface{}{
		"dashboard":         o.overviewDashboard(campaignID, domain, startTime, endTime),
		"mail_providers":    o.overviewProviders(campaignID, domain, startTime, endTime),
		"send_mail_chart":   o.chartSendMail(campaignID, domain, startTime, endTime),
		"bounce_rate_chart": o.chartBounceRate(campaignID, domain, startTime, endTime),
		"open_rate_chart":   o.chartOpenRate(campaignID, domain, startTime, endTime),
		"click_rate_chart":  o.chartClickRate(campaignID, domain, startTime, endTime),
	}
}

// overviewDashboard dashboard data
func (o *Overview) overviewDashboard(campaignID int64, domain string, startTime, endTime int64) map[string]interface{} {
	query := o.buildBaseQuery(campaignID, domain, startTime, endTime)

	query.LeftJoin("mailstat_opened o", "sm.postfix_message_id=o.postfix_message_id")
	query.LeftJoin("mailstat_clicked c", "sm.postfix_message_id=c.postfix_message_id")

	query.Fields("count(*) as sends")
	query.Fields("coalesce(sum(case when status='sent' and dsn like '2.%' then 1 else 0 end), 0) as delivered")
	query.Fields("count(distinct o.postfix_message_id) as opened")
	query.Fields("count(distinct c.postfix_message_id) as clicked")
	query.Fields("coalesce(sum(case when status='bounced' then 1 else 0 end), 0) as bounced")

	result, err := query.One()
	if err != nil {
		g.Log().Error(context.Background(), err)
		return nil
	}

	aggregate := map[string]interface{}{
		"sends":     0,
		"delivered": 0,
		"opened":    0,
		"clicked":   0,
		"bounced":   0,
	}

	for k, v := range result {
		if _, ok := aggregate[k]; !ok {
			continue
		}
		aggregate[k] = v.Int()
	}

	if aggregate["sends"].(int) > 0 {
		aggregate["delivery_rate"] = public.Round(float64(aggregate["delivered"].(int))/float64(aggregate["sends"].(int))*100, 2)
	} else {
		aggregate["delivery_rate"] = 0
	}

	if aggregate["sends"].(int) > 0 {
		aggregate["bounce_rate"] = public.Round(float64(aggregate["bounced"].(int))/float64(aggregate["sends"].(int))*100, 2)
		aggregate["open_rate"] = public.Round(float64(aggregate["opened"].(int))/float64(aggregate["sends"].(int))*100, 2)
		aggregate["click_rate"] = public.Round(float64(aggregate["clicked"].(int))/float64(aggregate["sends"].(int))*100, 2)
	} else {
		aggregate["bounce_rate"] = 0
		aggregate["open_rate"] = 0
		aggregate["click_rate"] = 0
	}

	return aggregate
}

// overviewProviders statistic mail provider
func (o *Overview) overviewProviders(campaignID int64, domain string, startTime, endTime int64) []map[string]interface{} {
	startTime, endTime = o.filterAndPrepareTimeSection(startTime, endTime)

	query := o.buildBaseQuery(campaignID, domain, startTime, endTime)

	query = query.LeftJoin("mailstat_opened o", "sm.postfix_message_id=o.postfix_message_id")
	query = query.LeftJoin("mailstat_clicked c", "sm.postfix_message_id=c.postfix_message_id")

	query = query.Fields("sm.mail_provider")
	query = query.Fields("count(*) as sends")
	query = query.Fields("coalesce(sum(case when status='sent' and dsn like '2.%' then 1 else 0 end), 0) as delivered")
	query = query.Fields("count(distinct o.postfix_message_id) as opened")
	query = query.Fields("count(distinct c.postfix_message_id) as clicked")
	query = query.Fields("coalesce(sum(case when status='bounced' then 1 else 0 end), 0) as bounced")

	query = query.Group("sm.mail_provider")

	results, err := query.All()
	if err != nil {
		g.Log().Error(context.Background(), err)
		return nil
	}

	aggregate := map[string]interface{}{
		"sends":     0,
		"delivered": 0,
		"opened":    0,
		"clicked":   0,
		"bounced":   0,
	}

	m := make(map[string]map[string]interface{})

	for _, result := range results {
		provider := result["mail_provider"].String()
		if _, ok := m[provider]; !ok {
			m[provider] = make(map[string]interface{})
			for k, v := range aggregate {
				m[provider][k] = v
			}
		}

		for k, v := range result {
			if _, ok := aggregate[k]; !ok {
				continue
			}
			m[provider][k] = v.Int()
		}
	}

	var lst []map[string]interface{}
	for k, item := range m {
		item["mail_provider"] = k
		if item["sends"].(int) > 0 {
			item["delivery_rate"] = public.Round(float64(item["delivered"].(int))/float64(item["sends"].(int))*100, 2)
			item["bounce_rate"] = public.Round(float64(item["bounced"].(int))/float64(item["sends"].(int))*100, 2)
			item["open_rate"] = public.Round(float64(item["opened"].(int))/float64(item["sends"].(int))*100, 2)
			item["click_rate"] = public.Round(float64(item["clicked"].(int))/float64(item["sends"].(int))*100, 2)
		} else {
			item["delivery_rate"] = 0
			item["bounce_rate"] = 0
			item["open_rate"] = 0
			item["click_rate"] = 0
		}
		lst = append(lst, item)
	}

	// descending by sends
	sort.Slice(lst, func(i, j int) bool {
		return lst[i]["sends"].(int) > lst[j]["sends"].(int)
	})

	return lst
}

func (o *Overview) fillChartData(data []map[string]interface{}, fillItem map[string]interface{}, fillType, fillKey string, startTime, endTime int64) []map[string]interface{} {
	switch fillType {
	case "daily":
		return o.fillChartDataDaily(data, fillItem, fillKey, startTime, endTime)
	case "hourly":
		return o.fillChartDataHourly(data, fillItem, fillKey)
	default:
		return data
	}
}

func (o *Overview) fillChartDataHourly(data []map[string]interface{}, fillItem map[string]interface{}, fillKey string) []map[string]interface{} {
	for _, item := range data {
		item[fillKey] = gconv.Int(item[fillKey])
	}

	if len(data) == 24 {
		return data
	}

	m := make(map[int]map[string]interface{})
	for _, item := range data {
		m[gconv.Int(item[fillKey])] = item
	}

	var lst []map[string]interface{}
	for i := 0; i < 24; i++ {
		if item, ok := m[i]; ok {
			lst = append(lst, item)
			continue
		}

		item := make(map[string]interface{}, len(fillItem))
		for k, v := range fillItem {
			item[k] = v
		}
		item[fillKey] = i
		lst = append(lst, item)
	}

	return lst
}

func (o *Overview) fillChartDataDaily(data []map[string]interface{}, fillItem map[string]interface{}, fillKey string, startTime, endTime int64) []map[string]interface{} {
	if startTime < 0 || endTime < 0 {
		return data
	}

	if startTime > endTime {
		return data
	}

	m := make(map[string]map[string]interface{})
	for _, item := range data {
		item[fillKey] = gconv.Int64(item[fillKey])
	}

	var lst []map[string]interface{}
	for i := startTime; i <= endTime; i += 86400 {
		dayDateObj := time.Unix(i, 0)
		dayDate := dayDateObj.Format("2006-01-02")
		dayTime := time.Date(dayDateObj.Year(), dayDateObj.Month(), dayDateObj.Day(), 0, 0, 0, 0, time.Local).Unix()

		if item, ok := m[dayDate]; ok {
			item[fillKey] = dayTime
			lst = append(lst, item)
			continue
		}

		item := make(map[string]interface{}, len(fillItem))
		for k, v := range fillItem {
			item[k] = v
		}
		item[fillKey] = dayTime
		lst = append(lst, item)
	}

	return lst
}

// sendMailDashboard dashboard data for send mail
func (o *Overview) sendMailDashboard(campaignID int64, domain string, startTime, endTime int64) map[string]interface{} {
	startTime, endTime = o.filterAndPrepareTimeSection(startTime, endTime)

	query := o.buildBaseQuery(campaignID, domain, startTime, endTime)

	query = query.Fields("count(*) as sends")
	query = query.Fields("coalesce(sum(case when status='sent' and dsn like '2.%' then 1 else 0 end), 0) as delivered")
	query = query.Fields("count(*) - coalesce(sum(case when status='sent' and dsn like '2.%' then 1 else 0 end), 0) as failed")

	result, err := query.One()
	if err != nil {
		g.Log().Error(context.Background(), err)
		return nil
	}

	aggregate := map[string]interface{}{
		"sends":     0,
		"delivered": 0,
		"failed":    0,
	}

	for k, v := range result {
		aggregate[k] = v.Int()
	}

	if aggregate["sends"].(int) > 0 {
		aggregate["delivery_rate"] = public.Round(float64(aggregate["delivered"].(int))/float64(aggregate["sends"].(int))*100, 2)
		aggregate["failure_rate"] = public.Round(float64(aggregate["failed"].(int))/float64(aggregate["sends"].(int))*100, 2)
	} else {
		aggregate["delivery_rate"] = 0
		aggregate["failure_rate"] = 0
	}

	return aggregate
}

func (o *Overview) prepareChartData(startTime, endTime int64) (string, string) {
	columnType := "daily"
	secs := endTime - startTime
	xAxisField := "extract(epoch from to_char(to_timestamp(sm.log_time_millis / 1000), 'YYYY-MM-DD')::timestamp)::integer as x"

	if secs < 86400 {
		columnType = "hourly"
		xAxisField = "to_char(to_timestamp(sm.log_time_millis / 1000), 'HH24')::integer as x"
	}

	return columnType, xAxisField
}

func (o *Overview) chartSendMail(campaignID int64, domain string, startTime, endTime int64) map[string]interface{} {
	startTime, endTime = o.filterAndPrepareTimeSection(startTime, endTime)

	query := o.buildBaseQuery(campaignID, domain, startTime, endTime)

	columnType, xAxisField := o.prepareChartData(startTime, endTime)

	query = query.Fields(xAxisField)
	query = query.Fields("count(*) as sends")
	query = query.Fields("coalesce(sum(case when status='sent' and dsn like '2.%' then 1 else 0 end), 0) as delivered")
	query = query.Fields("count(*) - coalesce(sum(case when status='sent' and dsn like '2.%' then 1 else 0 end), 0) as failed")

	query = query.Group("x")
	query = query.Order("x")

	rs, err := query.All()
	if err != nil {
		g.Log().Error(context.Background(), err)
		return nil
	}

	results := make([]map[string]interface{}, 0, len(rs))
	for _, item := range rs {
		results = append(results, item.Map())
	}

	fillItem := map[string]interface{}{
		"sends":     0,
		"delivered": 0,
		"failed":    0,
	}

	return map[string]interface{}{
		"column_type": columnType,
		"dashboard":   o.sendMailDashboard(campaignID, domain, startTime, endTime),
		"data":        o.fillChartData(results, fillItem, columnType, "x", startTime, endTime),
	}
}

func (o *Overview) chartBounceRate(campaignID int64, domain string, startTime, endTime int64) map[string]interface{} {
	startTime, endTime = o.filterAndPrepareTimeSection(startTime, endTime)

	query := o.buildBaseQuery(campaignID, domain, startTime, endTime)

	columnType, xAxisField := o.prepareChartData(startTime, endTime)

	query = query.Fields(xAxisField)
	query = query.Fields(gdb.Raw("case when count(*) > 0 then round(1.0 * coalesce(sum(case when status='bounced' then 1 else 0 end), 0) / count(*) * 100.0, 2) else 0.0 end as bounce_rate"))

	query = query.Group("x")
	query = query.Order("x")

	rs, err := query.All()
	if err != nil {
		g.Log().Error(context.Background(), err)
		return nil
	}

	results := make([]map[string]interface{}, 0, len(rs))
	for _, item := range rs {
		results = append(results, item.Map())
	}

	fillItem := map[string]interface{}{
		"bounce_rate": 0.0,
	}

	return map[string]interface{}{
		"column_type": columnType,
		"data":        o.fillChartData(results, fillItem, columnType, "x", startTime, endTime),
	}
}

func (o *Overview) chartOpenRate(campaignID int64, domain string, startTime, endTime int64) map[string]interface{} {
	startTime, endTime = o.filterAndPrepareTimeSection(startTime, endTime)

	query := o.buildBaseQuery(campaignID, domain, startTime, endTime)

	columnType, xAxisField := o.prepareChartData(startTime, endTime)

	query = query.LeftJoin("mailstat_opened o", "sm.postfix_message_id=o.postfix_message_id")

	query = query.Fields(xAxisField)
	query = query.Fields("case when coalesce(sum(case when status='sent' and dsn like '2.%' then 1 else 0 end), 0) > 0 then round(1.0 * count(distinct o.postfix_message_id) / coalesce(sum(case when status='sent' and dsn like '2.%' then 1 else 0 end), 0) * 100, 2) else 0.0 end as open_rate")

	query = query.Group("x")

	rs, err := query.All()
	if err != nil {
		g.Log().Error(context.Background(), err)
		return nil
	}

	results := make([]map[string]interface{}, 0, len(rs))
	for _, item := range rs {
		results = append(results, item.Map())
	}

	fillItem := map[string]interface{}{
		"open_rate": 0.0,
	}

	return map[string]interface{}{
		"column_type": columnType,
		"data":        o.fillChartData(results, fillItem, columnType, "x", startTime, endTime),
	}
}

func (o *Overview) chartClickRate(campaignID int64, domain string, startTime, endTime int64) map[string]interface{} {
	startTime, endTime = o.filterAndPrepareTimeSection(startTime, endTime)

	query := o.buildBaseQuery(campaignID, domain, startTime, endTime)

	columnType, xAxisField := o.prepareChartData(startTime, endTime)

	query = query.LeftJoin("mailstat_clicked c", "sm.postfix_message_id=c.postfix_message_id")

	query = query.Fields(xAxisField)
	query = query.Fields("case when coalesce(sum(case when status='sent' and dsn like '2.%' then 1 else 0 end), 0) > 0 then round(1.0 * count(distinct c.postfix_message_id) / coalesce(sum(case when status='sent' and dsn like '2.%' then 1 else 0 end), 0) * 100, 2) else 0.0 end as click_rate")

	query = query.Group("x")

	rs, err := query.All()
	if err != nil {
		g.Log().Error(context.Background(), err)
		return nil
	}

	results := make([]map[string]interface{}, 0, len(rs))
	for _, item := range rs {
		results = append(results, item.Map())
	}

	fillItem := map[string]interface{}{
		"click_rate": 0.0,
	}

	return map[string]interface{}{
		"column_type": columnType,
		"data":        o.fillChartData(results, fillItem, columnType, "x", startTime, endTime),
	}
}
