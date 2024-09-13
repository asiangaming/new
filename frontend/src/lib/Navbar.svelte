<script>
    import { link } from "svelte-spa-router";
    let path_menu;
    let color_tab_menu = "text-[#e7ca40]"
    let content_text_1 = "ホームページ";
    let content_text_2 = "結果";
    let content_text_3 = "責任";
    
    if (path_menu == "home"){
      color_tab_menu = "text-[#e7ca40]"
    }else{
      color_tab_menu = "text-[#20a5ab]"
    }
    
</script>
<main class="w-full max-w-6xl mx-auto">
  
  <nav class="navbar  w-full  ">
    <div class="flex-1 w-full">
      <a class="normal-case text-xl text-center" use:link href="/">
        <img class="w-[150px] lg:w-[250px] " src="images/logojapan.png" alt="静岡市 Lottery">
      </a>
    </div>
    <div class="flex-none ">
      <ul class="hidden lg:menu lg:menu-horizontal p-0">
        <li><a class="text-sm  hover:bg-transparent focus:bg-transparent"use:link href="/">
          <span class="{color_tab_menu}">{content_text_1}</span>
        </a></li>
        <li><a class="text-sm hover:bg-transparent focus:bg-transparent"use:link href="/result">
          <span class="{color_tab_menu}">{content_text_2}</span>
        </a></li>
        <li><a class="text-sm hover:bg-transparent focus:bg-transparent"use:link href="/responsible">
          <span class="{color_tab_menu}">{content_text_3}</span>
        </a></li>
      </ul>
    </div>
    <div class="dropdown dropdown-end lg:hidden">
      <!-- svelte-ignore a11y-label-has-associated-control -->
      <label tabindex="0" class="btn btn-ghost btn-circle">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 bg-white" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h7" /></svg>
      </label>
      <ul tabindex="0" class="menu menu-compact dropdown-content mt-3 p-2 shadow bg-base-content rounded-box w-52">
        <li><a class="text-xs text-white font-bold hover:bg-transparent focus:bg-transparent"use:link href="/">{content_text_1}</a></li>
        <li><a class="text-xs text-white font-bold hover:bg-transparent focus:bg-transparent"use:link href="/result">{content_text_2}</a></li>
      </ul>
    </div>
  </nav>
</main>