import { loadEnv } from '@rsbuild/core'
import serverConfig from './config'

/**
 * @description 获取环境变量
 * @param key
 * @returns
 */
export function getEnv(key: string) {
	const { parsed } = loadEnv()
	return parsed[key]
}

/**
 * @description 获取服务器配置
 * @returns
 */
export function getServer() {
	const name = getEnv('SERVER_NAME')
	return serverConfig[name]
}

/**
 * @description 获取代理地址
 * @returns
 */
export const getProxyAddress = () => {
	const server = getServer()
	return `http${server.https ? 's' : ''}://${server.ip}:${server.port}`
}

/**
 * @description 获取是否使用 HTTPS
 * @returns
 */
export function getHttps() {
	const server = getServer()
	return server.https === true
}
