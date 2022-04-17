import App from './App.svelte'

const app = new App({
  target: document.getElementById('app'),
  props: {
		path_api: "/",
		// path_api: "http://localhost:2022/",
		font_size: "text-[12px] lg:text-[13px]"
	}
})

export default app
