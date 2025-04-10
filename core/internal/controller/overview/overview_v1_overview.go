package overview

import (
	"billionmail-core/internal/service/public"
	"context"
	"fmt"
	random "math/rand"
	"time"

	"billionmail-core/api/overview/v1"
)

func (c *ControllerV1) Overview(ctx context.Context, req *v1.OverviewReq) (res *v1.OverviewRes, err error) {
	res = &v1.OverviewRes{}

	res.SetSuccess("Success")

	columnType := "hourly"

	if req.EndTime == 0 {
		req.EndTime = time.Now().Unix()
	}

	if req.EndTime < req.StartTime {
		err = fmt.Errorf("start time must be greater or equal to end time")
		return
	}

	// fake chart map
	m := map[string]map[string][]map[string]interface{}{
		"daily": {
			"send_mail":   {},
			"bounce_rate": {},
			"open_rate":   {},
			"click_rate":  {},
		},
		"hourly": {
			"send_mail": {
				{"delivered": 25, "failed": 0, "sends": 25, "x": 0},
				{"delivered": 17, "failed": 5, "sends": 22, "x": 1},
				{"delivered": 20, "failed": 1, "sends": 21, "x": 2},
				{"delivered": 19, "failed": 0, "sends": 19, "x": 3},
				{"delivered": 15, "failed": 0, "sends": 15, "x": 4},
				{"delivered": 13, "failed": 1, "sends": 14, "x": 5},
				{"delivered": 16, "failed": 0, "sends": 16, "x": 6},
				{"delivered": 11, "failed": 0, "sends": 11, "x": 7},
				{"delivered": 9, "failed": 0, "sends": 9, "x": 8},
				{"delivered": 14, "failed": 0, "sends": 14, "x": 9},
				{"delivered": 20, "failed": 1, "sends": 21, "x": 10},
				{"delivered": 22, "failed": 1, "sends": 23, "x": 11},
				{"delivered": 17, "failed": 0, "sends": 17, "x": 12},
				{"delivered": 17, "failed": 0, "sends": 17, "x": 13},
				{"delivered": 25, "failed": 1, "sends": 26, "x": 14},
				{"delivered": 29, "failed": 2, "sends": 31, "x": 15},
				{"delivered": 26, "failed": 4, "sends": 30, "x": 16},
				{"delivered": 22, "failed": 33, "sends": 55, "x": 17},
				{"delivered": 32, "failed": 3, "sends": 35, "x": 18},
				{"delivered": 0, "failed": 0, "sends": 0, "x": 19},
				{"delivered": 0, "failed": 0, "sends": 0, "x": 20},
				{"delivered": 0, "failed": 0, "sends": 0, "x": 21},
				{"delivered": 0, "failed": 0, "sends": 0, "x": 22},
				{"delivered": 0, "failed": 0, "sends": 0, "x": 23},
			},
			"bounce_rate": {
				{"bounce_rate": 0.0, "x": 0},
				{"bounce_rate": 22.73, "x": 1},
				{"bounce_rate": 4.76, "x": 2},
				{"bounce_rate": 0.0, "x": 3},
				{"bounce_rate": 0.0, "x": 4},
				{"bounce_rate": 7.14, "x": 5},
				{"bounce_rate": 0.0, "x": 6},
				{"bounce_rate": 0.0, "x": 7},
				{"bounce_rate": 0.0, "x": 8},
				{"bounce_rate": 0.0, "x": 9},
				{"bounce_rate": 4.76, "x": 10},
				{"bounce_rate": 4.35, "x": 11},
				{"bounce_rate": 0.0, "x": 12},
				{"bounce_rate": 0.0, "x": 13},
				{"bounce_rate": 3.85, "x": 14},
				{"bounce_rate": 6.45, "x": 15},
				{"bounce_rate": 3.33, "x": 16},
				{"bounce_rate": 1.82, "x": 17},
				{"bounce_rate": 0.0, "x": 18},
				{"bounce_rate": 0.0, "x": 19},
				{"bounce_rate": 0.0, "x": 20},
				{"bounce_rate": 0.0, "x": 21},
				{"bounce_rate": 0.0, "x": 22},
				{"bounce_rate": 0.0, "x": 23},
			},
			"open_rate": {
				{"open_rate": 0.0, "x": 0},
				{"open_rate": 0.0, "x": 1},
				{"open_rate": 0.0, "x": 2},
				{"open_rate": 0.0, "x": 3},
				{"open_rate": 0.0, "x": 4},
				{"open_rate": 0.0, "x": 5},
				{"open_rate": 0.0, "x": 6},
				{"open_rate": 0.0, "x": 7},
				{"open_rate": 0.0, "x": 8},
				{"open_rate": 0.0, "x": 9},
				{"open_rate": 0.0, "x": 10},
				{"open_rate": 0.0, "x": 11},
				{"open_rate": 0.0, "x": 12},
				{"open_rate": 0.0, "x": 13},
				{"open_rate": 0.0, "x": 14},
				{"open_rate": 0.0, "x": 15},
				{"open_rate": 0.0, "x": 16},
				{"open_rate": 0.0, "x": 17},
				{"open_rate": 0.0, "x": 18},
				{"open_rate": 0.0, "x": 19},
				{"open_rate": 0.0, "x": 20},
				{"open_rate": 0.0, "x": 21},
				{"open_rate": 0.0, "x": 22},
				{"open_rate": 0.0, "x": 23},
			},
			"click_rate": {
				{"click_rate": 0.0, "x": 0},
				{"click_rate": 0.0, "x": 1},
				{"click_rate": 0.0, "x": 2},
				{"click_rate": 0.0, "x": 3},
				{"click_rate": 0.0, "x": 4},
				{"click_rate": 0.0, "x": 5},
				{"click_rate": 0.0, "x": 6},
				{"click_rate": 0.0, "x": 7},
				{"click_rate": 0.0, "x": 8},
				{"click_rate": 0.0, "x": 9},
				{"click_rate": 0.0, "x": 10},
				{"click_rate": 0.0, "x": 11},
				{"click_rate": 0.0, "x": 12},
				{"click_rate": 0.0, "x": 13},
				{"click_rate": 0.0, "x": 14},
				{"click_rate": 0.0, "x": 15},
				{"click_rate": 0.0, "x": 16},
				{"click_rate": 0.0, "x": 17},
				{"click_rate": 0.0, "x": 18},
				{"click_rate": 0.0, "x": 19},
				{"click_rate": 0.0, "x": 20},
				{"click_rate": 0.0, "x": 21},
				{"click_rate": 0.0, "x": 22},
				{"click_rate": 0.0, "x": 23},
			},
		},
	}

	if req.EndTime-req.StartTime > 86400 {
		columnType = "daily"
		dayUnixTimes := public.GenDayUnixTimestampByTimeSection(req.StartTime, req.EndTime)
		m["daily"]["send_mail"] = make([]map[string]interface{}, len(dayUnixTimes))
		m["daily"]["bounce_rate"] = make([]map[string]interface{}, len(dayUnixTimes))
		m["daily"]["open_rate"] = make([]map[string]interface{}, len(dayUnixTimes))
		m["daily"]["click_rate"] = make([]map[string]interface{}, len(dayUnixTimes))

		for i, dayUnixTime := range dayUnixTimes {
			// generate fake data
			delivered := 350 + random.Intn(100)
			failed := 30 + random.Intn(50)
			sends := delivered + failed
			deliveryRate := float64(delivered) / float64(sends) * 100
			failureRate := 100.0 - deliveryRate
			bounceRate := 1.5 + float64(random.Intn(400))/100.0
			openRate := float64(random.Intn(800)) / 100.0
			clickRate := float64(random.Intn(500)) / 100.0

			m["daily"]["send_mail"][i] = map[string]interface{}{
				"delivered":     delivered,
				"delivery_rate": public.Round(deliveryRate, 2),
				"failed":        failed,
				"failure_rate":  public.Round(failureRate, 2),
				"sends":         sends,
				"x":             dayUnixTime,
			}
			m["daily"]["bounce_rate"][i] = map[string]interface{}{
				"bounce_rate": public.Round(bounceRate, 2),
				"x":           dayUnixTime,
			}
			m["daily"]["open_rate"][i] = map[string]interface{}{
				"open_rate": public.Round(openRate, 2),
				"x":         dayUnixTime,
			}
			m["daily"]["click_rate"][i] = map[string]interface{}{
				"click_rate": public.Round(clickRate, 2),
				"x":          dayUnixTime,
			}
		}
	}

	res.Data.Dashboard = map[string]interface{}{
		"bounce_rate":   3.33,
		"bounced":       14,
		"click_rate":    0.0,
		"clicked":       0,
		"delivered":     369,
		"delivery_rate": 87.65,
		"open_rate":     0.0,
		"opened":        0,
		"sends":         421,
		"failure_rate":  12.35,
		"failed":        52,
	}

	res.Data.SendMailChart = map[string]interface{}{
		"column_type": columnType,
		"dashboard": map[string]interface{}{
			"delivered":     369,
			"delivery_rate": 87.65,
			"failed":        52,
			"failure_rate":  12.35,
			"sends":         421,
		},
		"data": m[columnType]["send_mail"],
	}

	res.Data.BounceRateChart = map[string]interface{}{
		"column_type": columnType,
		"data":        m[columnType]["bounce_rate"],
	}

	res.Data.OpenRateChart = map[string]interface{}{
		"column_type": columnType,
		"data":        m[columnType]["open_rate"],
	}

	res.Data.ClickRateChart = map[string]interface{}{
		"column_type": columnType,
		"data":        m[columnType]["click_rate"],
	}

	res.Data.MailProviders = []map[string]interface{}{
		{
			"bounce_rate":   2.88,
			"bounced":       7,
			"click_rate":    0.0,
			"clicked":       0,
			"delivered":     234,
			"delivery_rate": 96.3,
			"mail_provider": "google",
			"open_rate":     0.0,
			"opened":        0,
			"sends":         243,
		},
		{
			"bounce_rate":   6.93,
			"bounced":       7,
			"click_rate":    0.0,
			"clicked":       0,
			"delivered":     81,
			"delivery_rate": 80.2,
			"mail_provider": "other",
			"open_rate":     0.0,
			"opened":        0,
			"sends":         101,
		},
		{
			"bounce_rate":   0.0,
			"bounced":       0,
			"click_rate":    0.0,
			"clicked":       0,
			"delivered":     53,
			"delivery_rate": 98.15,
			"mail_provider": "local",
			"open_rate":     0.0,
			"opened":        0,
			"sends":         54,
		},
		{
			"bounce_rate":   0.0,
			"bounced":       0,
			"click_rate":    0.0,
			"clicked":       0,
			"delivered":     0,
			"delivery_rate": 0.0,
			"mail_provider": "outlook",
			"open_rate":     0.0,
			"opened":        0,
			"sends":         22,
		},
		{
			"bounce_rate":   0.0,
			"bounced":       0,
			"click_rate":    0.0,
			"clicked":       0,
			"delivered":     1,
			"delivery_rate": 100.0,
			"mail_provider": "yahoo",
			"open_rate":     0.0,
			"opened":        0,
			"sends":         1,
		},
	}

	return
}
