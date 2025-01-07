
// this file is generated — do not edit it


/// <reference types="@sveltejs/kit" />

/**
 * Environment variables [loaded by Vite](https://vitejs.dev/guide/env-and-mode.html#env-files) from `.env` files and `process.env`. Like [`$env/dynamic/private`](https://svelte.dev/docs/kit/$env-dynamic-private), this module cannot be imported into client-side code. This module only includes variables that _do not_ begin with [`config.kit.env.publicPrefix`](https://svelte.dev/docs/kit/configuration#env) _and do_ start with [`config.kit.env.privatePrefix`](https://svelte.dev/docs/kit/configuration#env) (if configured).
 * 
 * _Unlike_ [`$env/dynamic/private`](https://svelte.dev/docs/kit/$env-dynamic-private), the values exported from this module are statically injected into your bundle at build time, enabling optimisations like dead code elimination.
 * 
 * ```ts
 * import { API_KEY } from '$env/static/private';
 * ```
 * 
 * Note that all environment variables referenced in your code should be declared (for example in an `.env` file), even if they don't have a value until the app is deployed:
 * 
 * ```
 * MY_FEATURE_FLAG=""
 * ```
 * 
 * You can override `.env` values from the command line like so:
 * 
 * ```bash
 * MY_FEATURE_FLAG="enabled" npm run dev
 * ```
 */
declare module '$env/static/private' {
	export const ALLUSERSPROFILE: string;
	export const ANDROID_HOME: string;
	export const APPDATA: string;
	export const APPS: string;
	export const BUILD_TOOLS_BIN: string;
	export const BUILD_TOOLS_ROOT: string;
	export const COMMONPROGRAMFILES: string;
	export const COMMONPROGRAMW6432: string;
	export const COMPUTERNAME: string;
	export const COMSPEC: string;
	export const DRIVERDATA: string;
	export const GOPATH: string;
	export const HOMEDRIVE: string;
	export const HOMEPATH: string;
	export const INCLUDE: string;
	export const JAVA_HOME: string;
	export const LIB: string;
	export const LOCALAPPDATA: string;
	export const LOGONSERVER: string;
	export const MACOSX_DEPLOYMENT_TARGET: string;
	export const MASON: string;
	export const MYVIMRC: string;
	export const NODE: string;
	export const NODE_ENV: string;
	export const npm_command: string;
	export const npm_config_local_prefix: string;
	export const npm_config_user_agent: string;
	export const npm_execpath: string;
	export const npm_lifecycle_event: string;
	export const npm_lifecycle_script: string;
	export const npm_node_execpath: string;
	export const npm_package_json: string;
	export const npm_package_name: string;
	export const npm_package_version: string;
	export const NUMBER_OF_PROCESSORS: string;
	export const NVIM: string;
	export const NVIM_LOG_FILE: string;
	export const OS: string;
	export const PATH: string;
	export const PATHEXT: string;
	export const POWERSHELL_DISTRIBUTION_CHANNEL: string;
	export const POWERSHELL_TELEMETRY_OPTOUT: string;
	export const PROCESSOR_ARCHITECTURE: string;
	export const PROCESSOR_IDENTIFIER: string;
	export const PROCESSOR_LEVEL: string;
	export const PROCESSOR_REVISION: string;
	export const PROGRAMDATA: string;
	export const PROGRAMFILES: string;
	export const PROGRAMW6432: string;
	export const PROMPT: string;
	export const PSMODULEPATH: string;
	export const PUBLIC: string;
	export const SESSIONNAME: string;
	export const SYSTEMDRIVE: string;
	export const SYSTEMROOT: string;
	export const TAURI_CLI_VERBOSITY: string;
	export const TAURI_ENV_ARCH: string;
	export const TAURI_ENV_DEBUG: string;
	export const TAURI_ENV_FAMILY: string;
	export const TAURI_ENV_PLATFORM: string;
	export const TAURI_ENV_PLATFORM_VERSION: string;
	export const TAURI_ENV_TARGET_TRIPLE: string;
	export const TEMP: string;
	export const TERM: string;
	export const TMP: string;
	export const USERDOMAIN: string;
	export const USERDOMAIN_ROAMINGPROFILE: string;
	export const USERNAME: string;
	export const USERPROFILE: string;
	export const VCTOOLSINSTALLDIR: string;
	export const VSCMD_ARG_HOST_ARCH: string;
	export const VSCMD_ARG_TGT_ARCH: string;
	export const WINDIR: string;
	export const WINDOWSSDKDIR: string;
	export const WINDOWSSDKVERSION: string;
	export const WSLENV: string;
	export const WT_PROFILE_ID: string;
	export const WT_SESSION: string;
	export const ZES_ENABLE_SYSMAN: string;
}

/**
 * Similar to [`$env/static/private`](https://svelte.dev/docs/kit/$env-static-private), except that it only includes environment variables that begin with [`config.kit.env.publicPrefix`](https://svelte.dev/docs/kit/configuration#env) (which defaults to `PUBLIC_`), and can therefore safely be exposed to client-side code.
 * 
 * Values are replaced statically at build time.
 * 
 * ```ts
 * import { PUBLIC_BASE_URL } from '$env/static/public';
 * ```
 */
declare module '$env/static/public' {
	
}

/**
 * This module provides access to runtime environment variables, as defined by the platform you're running on. For example if you're using [`adapter-node`](https://github.com/sveltejs/kit/tree/main/packages/adapter-node) (or running [`vite preview`](https://svelte.dev/docs/kit/cli)), this is equivalent to `process.env`. This module only includes variables that _do not_ begin with [`config.kit.env.publicPrefix`](https://svelte.dev/docs/kit/configuration#env) _and do_ start with [`config.kit.env.privatePrefix`](https://svelte.dev/docs/kit/configuration#env) (if configured).
 * 
 * This module cannot be imported into client-side code.
 * 
 * Dynamic environment variables cannot be used during prerendering.
 * 
 * ```ts
 * import { env } from '$env/dynamic/private';
 * console.log(env.DEPLOYMENT_SPECIFIC_VARIABLE);
 * ```
 * 
 * > In `dev`, `$env/dynamic` always includes environment variables from `.env`. In `prod`, this behavior will depend on your adapter.
 */
declare module '$env/dynamic/private' {
	export const env: {
		ALLUSERSPROFILE: string;
		ANDROID_HOME: string;
		APPDATA: string;
		APPS: string;
		BUILD_TOOLS_BIN: string;
		BUILD_TOOLS_ROOT: string;
		COMMONPROGRAMFILES: string;
		COMMONPROGRAMW6432: string;
		COMPUTERNAME: string;
		COMSPEC: string;
		DRIVERDATA: string;
		GOPATH: string;
		HOMEDRIVE: string;
		HOMEPATH: string;
		INCLUDE: string;
		JAVA_HOME: string;
		LIB: string;
		LOCALAPPDATA: string;
		LOGONSERVER: string;
		MACOSX_DEPLOYMENT_TARGET: string;
		MASON: string;
		MYVIMRC: string;
		NODE: string;
		NODE_ENV: string;
		npm_command: string;
		npm_config_local_prefix: string;
		npm_config_user_agent: string;
		npm_execpath: string;
		npm_lifecycle_event: string;
		npm_lifecycle_script: string;
		npm_node_execpath: string;
		npm_package_json: string;
		npm_package_name: string;
		npm_package_version: string;
		NUMBER_OF_PROCESSORS: string;
		NVIM: string;
		NVIM_LOG_FILE: string;
		OS: string;
		PATH: string;
		PATHEXT: string;
		POWERSHELL_DISTRIBUTION_CHANNEL: string;
		POWERSHELL_TELEMETRY_OPTOUT: string;
		PROCESSOR_ARCHITECTURE: string;
		PROCESSOR_IDENTIFIER: string;
		PROCESSOR_LEVEL: string;
		PROCESSOR_REVISION: string;
		PROGRAMDATA: string;
		PROGRAMFILES: string;
		PROGRAMW6432: string;
		PROMPT: string;
		PSMODULEPATH: string;
		PUBLIC: string;
		SESSIONNAME: string;
		SYSTEMDRIVE: string;
		SYSTEMROOT: string;
		TAURI_CLI_VERBOSITY: string;
		TAURI_ENV_ARCH: string;
		TAURI_ENV_DEBUG: string;
		TAURI_ENV_FAMILY: string;
		TAURI_ENV_PLATFORM: string;
		TAURI_ENV_PLATFORM_VERSION: string;
		TAURI_ENV_TARGET_TRIPLE: string;
		TEMP: string;
		TERM: string;
		TMP: string;
		USERDOMAIN: string;
		USERDOMAIN_ROAMINGPROFILE: string;
		USERNAME: string;
		USERPROFILE: string;
		VCTOOLSINSTALLDIR: string;
		VSCMD_ARG_HOST_ARCH: string;
		VSCMD_ARG_TGT_ARCH: string;
		WINDIR: string;
		WINDOWSSDKDIR: string;
		WINDOWSSDKVERSION: string;
		WSLENV: string;
		WT_PROFILE_ID: string;
		WT_SESSION: string;
		ZES_ENABLE_SYSMAN: string;
		[key: `PUBLIC_${string}`]: undefined;
		[key: `${string}`]: string | undefined;
	}
}

/**
 * Similar to [`$env/dynamic/private`](https://svelte.dev/docs/kit/$env-dynamic-private), but only includes variables that begin with [`config.kit.env.publicPrefix`](https://svelte.dev/docs/kit/configuration#env) (which defaults to `PUBLIC_`), and can therefore safely be exposed to client-side code.
 * 
 * Note that public dynamic environment variables must all be sent from the server to the client, causing larger network requests — when possible, use `$env/static/public` instead.
 * 
 * Dynamic environment variables cannot be used during prerendering.
 * 
 * ```ts
 * import { env } from '$env/dynamic/public';
 * console.log(env.PUBLIC_DEPLOYMENT_SPECIFIC_VARIABLE);
 * ```
 */
declare module '$env/dynamic/public' {
	export const env: {
		[key: `PUBLIC_${string}`]: string | undefined;
	}
}
