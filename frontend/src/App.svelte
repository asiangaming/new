<script>
  import Router from "svelte-spa-router";
	import { wrap } from "svelte-spa-router/wrap";
  import Navbar from './lib/Navbar.svelte'
  import Footer from './lib/Footer.svelte'
  import Home from "./pages/home.svelte";
  import Result from "./pages/result.svelte";
  import NotFound from "./pages/NotFound.svelte";
    import Responsibility from "./pages/responsibility.svelte";
  export let path_api = "";
  let routes = "";
  
  routes = {
			"/": wrap({
				component: Home,
        props: {
						path_menu: "home",
				},
			}),
      "/result": wrap({
        props: {
            props: {
                path_api: path_api,
                path_menu: "result",
            },
						path_api: path_api,
            path_menu: "result",
				},
				component: Result,
			}),
      "/responsible": wrap({
        props: {
						path_api: path_api,
            path_menu: "responsible",
				},
				component: Responsibility,
			}),
			"*": NotFound,
	};
</script>

<Navbar />
<main class="p-0 lg:p-3 lg:max-w-6xl lg:mx-auto  xl:rounded-sm mt-2">
  <Router {routes} />
</main>



<style global lang="postcss">
  @import url("https://fonts.googleapis.com/css2?family=Poppins:wght@300;500;700&display=swap");
  
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
</style>
