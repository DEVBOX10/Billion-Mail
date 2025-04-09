import { ServerConfig } from './interface'

const serverConfig: Record<string, ServerConfig> = {
	'192.168.66.66': {
		https: false,
		ip: '192.168.66.66',
		port: 64325,
		sshPort: 22,
		username: 'root',
		password: 'www.bt.cn',
	},
}

export default serverConfig
