import { AxiosProgressEvent } from 'axios'
import { instance } from '@/api'

export const uploadFile = (
	params: FormData,
	onUploadProgress?: (e: AxiosProgressEvent) => void
) => {
	return instance.post('/file/upload', params, {
		headers: {
			'Content-Type': 'multipart/form-data',
		},
		onUploadProgress: e => {
			onUploadProgress?.(e)
		},
	})
}

export const downloadFile = (params: { file_path: string }) => {
	return instance.get('/file/download', {
		params,
		responseType: 'blob',
	})
}
