<script>
	import Router from "svelte-spa-router";
	import { wrap } from "svelte-spa-router/wrap";
	import Navigation from '../src/components/Navigation.svelte'
	import NotFound from '../src/pages/NotFound.svelte'
	import Login from '../src/pages/Login.svelte'
	import Home from '../src/pages/Home.svelte'
	import Curr from '../src/pages/curr/Curr.svelte'
	import Admin from '../src/pages/admin/Admin.svelte'
	import Adminrule from '../src/pages/adminrule/Adminrule.svelte'
	import Log from '../src/pages/log/Log.svelte'
	import Pasaran from '../src/pages/pasaran/Pasaran.svelte'
	import Domain from '../src/pages/domain/Domain.svelte'
	import Company from '../src/pages/company/Company.svelte'
	import Report from '../src/pages/report/Report.svelte'
	import Setting from '../src/pages/setting/Setting.svelte'
	import Tailwindcss from './Tailwindcss.svelte'
	export let path_api = "";
	export let font_size = "";

  	let token = localStorage.getItem("token");
	let routes = "";
	let isNav = false;
  	if (token == "" || token == null) {
    // isNav = true;
		routes = {
			"/": wrap({
        		props: {
					path_api: path_api,
				},
				component: Login,
			}),
			"*": NotFound,
		};
	} else {
		isNav = true;
			routes = {
				"/": wrap({
					props: {
						path_api: path_api,
						font_size: font_size,
					},
					component: Home,
				}),
				"/admin": wrap({
						props: {
							path_api: path_api,
							font_size: font_size,
						},
						component: Admin,
					}),
				"/adminrule": wrap({
						props: {
							path_api: path_api,
							font_size: font_size,
						},
						component: Adminrule,
					}),
				"/curr": wrap({
						props: {
							path_api: path_api,
							font_size: font_size,
						},
						component: Curr,
					}),
				"/log": wrap({
						props: {
							path_api: path_api,
							font_size: font_size,
						},
						component: Log,
					}),
				"/pasaran": wrap({
						props: {
							path_api: path_api,
							font_size: font_size,
						},
						component: Pasaran,
					}),
				"/domain": wrap({
						props: {
							path_api: path_api,
							font_size: font_size,
						},
						component: Domain,
					}),
				"/company": wrap({
						props: {
							path_api: path_api,
							font_size: font_size,
						},
						component: Company,
					}),
				"/report": wrap({
						props: {
							path_api: path_api,
							font_size: font_size,
						},
						component: Report,
					}),
				"/setting": wrap({
						props: {
							path_api: path_api,
							font_size: font_size,
						},
						component: Setting,
					}),
				"*": NotFound,
			};
  }
</script>
{#if isNav}
<main class="flex flex-col space-x-2 h-screen w-full bg-[#f0f2f5]">
	<Navigation />
	<div class="justify-center ">
    	<div class="w-full  mt-5">
      		<Router {routes} />
    	</div>
  	</div>
</main>
{:else}
<div class="container mx-auto">
  <Router {routes} />
</div>
{/if}
<Tailwindcss />