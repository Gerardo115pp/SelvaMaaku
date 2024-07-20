import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import path from 'path';
import fs from 'fs';

const fullReloadAlways = {
	handleHotUpdate({server}) {
		server.ws.send({type: 'full-reload'});
		return [];
	}
}

export default defineConfig(async  ({ command, mode, isSsrBuild, isPreview }) => {
	const is_production = command === 'build';

	let build_config = {
		"AI_CHAT_SERVICE": JSON.stringify(process.env.AI_CHAT_SERVICE),
	}

	/** @type {import('vite').UserConfig} */
	let config = {
		server: {
			open: false,
			host: "0.0.0.0",
			port: 1366
		},
		define: {
			...build_config,
		},
		resolve: {
			alias: {
				'@libs': path.resolve(__dirname, 'src/libs'),
				'@components': path.resolve(__dirname, 'src/components'),
				'@pages': path.resolve(__dirname, 'src/pages'),
				'@svg': path.resolve(__dirname, 'src/svg'),
				'@models': path.resolve(__dirname, 'src/models'),
				"@actions": path.resolve(__dirname, 'src/actions'),
				"@events": path.resolve(__dirname, 'src/events'),
				"@stores": path.resolve(__dirname, 'src/stores'),
				"@app": path.resolve(__dirname, 'src'),
			},
		},
		plugins: [
			sveltekit(),
			// fullReloadAlways
		],
		clearScreen: true
	}

	if (!is_production) {
		config.server = {
			...config.server,
			https: {
				key: fs.readFileSync(process.env.SSL_KEY_PATH),
				cert: fs.readFileSync(process.env.SSL_CERT_PATH)
			}
		}
	}

	return config;
})