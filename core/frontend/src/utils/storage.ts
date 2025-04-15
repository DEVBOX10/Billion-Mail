import { isObject } from './is'

/**
 * Get data from localStorage
 * @param { string } key Key name
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
 * Store data in localStorage
 * @param { string } key Key name
 * @param { string } value Value
 * @return void
 */
export const setLocalStorage = (key: string, value?: string | number) => {
	localStorage.setItem(key, `${value}`)
}

/**
 * Delete one or multiple data from localStorage
 * @param { string | string[] } key Key name(s)
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
 * Set cookie
 * @param { string } key Key name
 * @param { string } value Value
 * @return void
 */
export const setCookie = (key: string, value: string) => {
	const day = 30
	const date = new Date()
	date.setTime(date.getTime() + day * 24 * 60 * 60 * 1000)
	document.cookie = key + '=' + escape(value) + ';expires=' + date.toUTCString()
}
