import { isObject } from './is'

/**
 * 获取数据
 * @param { string } key 键名
 * @return { string }
 */
export const getLocalStorage = <T = string>(key: string, defaultVal: T = '' as T): T => {
	const val = localStorage.getItem(key)
	if (isObject<T>(defaultVal)) {
		return val != null ? JSON.parse(val) : defaultVal
	} else {
		return val != null ? (val as T) : defaultVal
	}
}

/**
 * 存储数据
 * @param { string } key 字段名
 * @param { string } value 字段值
 * @return void
 */
export const setLocalStorage = (key: string, value?: string | number) => {
	localStorage.setItem(key, `${value}`)
}

/**
 * 删除数据或删除多个数据
 * @param { string | string[] } key 键名
 * @return void
 */
export const delLocalStorage = (key: string | string[]) => {
	if (!Array.isArray(key)) key = [key]
	const removeKeys = [...key]
	removeKeys.forEach(dataKey => {
		localStorage.removeItem(dataKey)
	})
}

/**
 * 设置cookie
 * @param { string } key 键名
 * @param { string } value 值
 * @return void
 */
export const setCookie = (key: string, value: string) => {
	const day = 30
	const date = new Date()
	date.setTime(date.getTime() + day * 24 * 60 * 60 * 1000)
	document.cookie = key + '=' + escape(value) + ';expires=' + date.toUTCString()
}
