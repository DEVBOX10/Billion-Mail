import { RsbuildPlugin } from '@rsbuild/core'
import SftpClient from 'ssh2-sftp-client'
import chalk from 'chalk'
import ora from 'ora'

interface DeployPluginOptions {
	host: string
	port?: number
	username: string
	password: string
	localDistPath?: string
	remoteDistPath: string
}

export function deployPlugin(options: DeployPluginOptions): RsbuildPlugin {
	const { host, port = 22, username, password, localDistPath = './dist', remoteDistPath } = options

	return {
		name: 'rsbuild-deploy-plugin',
		setup(build) {
			build.onAfterBuild(async () => {
				const sftp = new SftpClient()
				let spinner

				try {
					spinner = ora(chalk.cyan('正在连接到服务器...')).start()
					await sftp.connect({ host, port, username, password })
					spinner.succeed(chalk.green('已连接到服务器。'))

					spinner = ora(chalk.cyan('正在删除服务器上的旧文件...')).start()
					await sftp.rmdir(remoteDistPath, true)
					spinner.succeed(chalk.green('服务器上的旧文件已删除。'))

					spinner = ora(chalk.cyan('正在创建新的目标目录...')).start()
					await sftp.mkdir(remoteDistPath, true)
					spinner.succeed(chalk.green('新的目标目录已创建。'))

					spinner = ora(chalk.cyan('正在上传新的打包文件...')).start()
					await sftp.uploadDir(localDistPath, remoteDistPath)
					spinner.succeed(chalk.green('新的打包文件已上传完成。'))

					sftp.end()
					spinner.succeed(chalk.green('部署完成。'))
				} catch (error) {
					if (spinner) {
						spinner.fail(chalk.red('部署过程中出现错误'))
					}
					spinner.fail(chalk.red('错误详情:'), error)
					sftp.end()
				}
			})
		},
	}
}
