import { defineConfig } from '@rsbuild/core'
import { pluginBabel } from '@rsbuild/plugin-babel'
import { pluginVue } from '@rsbuild/plugin-vue'
import { pluginVueJsx } from '@rsbuild/plugin-vue-jsx'
import { pluginSass } from '@rsbuild/plugin-sass'
import { pluginEslint } from '@rsbuild/plugin-eslint'
import { UnoCSSRspackPlugin } from '@unocss/webpack/rspack'
import { NaiveUiResolver } from 'unplugin-vue-components/resolvers'
import { pluginBasicSsl } from '@rsbuild/plugin-basic-ssl'
import Components from 'unplugin-vue-components/rspack'
import AutoImport from 'unplugin-auto-import/rspack'
import { deployPlugin } from './build/plugin/deploy'
import { getHttps, getProxyAddress, getServer, getEnv } from './build/utils'

const server = getServer()

export default defineConfig({
	plugins: [
		pluginBabel({
			include: /\.(?:jsx|tsx)$/,
		}),
		pluginVue(),
		pluginVueJsx(),
		pluginSass(),
		...(getHttps() ? [pluginBasicSsl()] : []),
		pluginEslint({
			eslintPluginOptions: {
				cwd: __dirname,
				configType: 'flat',
				extensions: ['.js', '.jsx', '.ts', '.tsx', '.vue'],
			},
		}),
		deployPlugin({
			host: server.ip,
			port: server.sshPort,
			username: server.username,
			password: server.password,
			remoteDistPath: '/opt/billion-mail/core/public/dist',
		}),
	],
	tools: {
		rspack: {
			plugins: [
				UnoCSSRspackPlugin({
					configFile: './uno.config.ts',
				}),
				AutoImport({
					dts: 'types/auto-import.d.ts',
					imports: ['vue', 'vue-router'],
					eslintrc: {
						enabled: true,
						filepath: '.eslintrc-auto-import.json',
					},
				}),
				Components({
					dts: 'types/components.d.ts',
					dirs: ['src/components/**/*'],
					extensions: ['vue', 'tsx'],
					resolvers: [NaiveUiResolver()],
				}),
			],
		},
	},
	html: {
		template: './index.html',
	},
	source: {
		alias: {
			'@': './src/',
			'@images': './src/assets/images/',
		},
		define: {
			'import.meta.env': JSON.stringify({
				SERVER: server,
				API_URL_PREFIX: getEnv('API_URL_PREFIX'),
			}),
		},
	},
	server: {
		proxy: {
			'/api': {
				target: getProxyAddress(),
				secure: false,
				changeOrigin: true,
				pathRewrite: { '^/api': '' },
			},
		},
	},
})
