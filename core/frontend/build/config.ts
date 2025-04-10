interface ServerConfig {
	https: boolean
	address: string
	host: string
}

const serverConfig: Record<string, ServerConfig> = {
	'192.168.66.66': {
		https: false,
		address: 'http://192.168.66.66:64325',
		host: '192.168.66.66',
	},
}

export default serverConfig
