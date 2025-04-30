import { instance } from '@/api'

export const getOverviewInfo = (params: {
	domain: string
	start_time: number
	end_time: number
}) => {
	return instance.get('/overview', { params })
}

export const getFailedList = (params: { domain: string; start_time: number; end_time: number }) => {
	return instance.get('/overview/failed', { params })
}
