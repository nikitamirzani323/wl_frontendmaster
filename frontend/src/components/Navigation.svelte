<script>
    import { link } from "svelte-spa-router";
    async function handleLogout() {
        localStorage.clear();
        window.location.href = "/";
    }
  </script>
<div class="sticky top-0 z-30 flex h-16 w-full justify-center bg-opacity-90 backdrop-blur transition-all duration-100 
    bg-[#1a73e8] text-base-content shadow-sm">
    <div class="navbar bg-base-100">
        <div class="navbar-start">
            <div class="dropdown">
              <label tabindex="0" class="btn btn-ghost lg:hidden">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h8m-8 6h16" /></svg>
              </label>
              <ul tabindex="0" class="menu menu-compact dropdown-content mt-3 p-2 shadow bg-base-100 rounded-box w-52">
                <li><a href="/">DASHBOARD</a></li>
                <li tabindex="0">
                  <a class="justify-between">
                    MASTER
                    <svg class="fill-current" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z"/></svg>
                  </a>
                  <ul class="p-2 bg-white">
                    <li><a use:link href="/admin">ADMIN MANAGEMENT</a></li>
                    <li><a use:link href="/adminrule">ADMIN RULE</a></li>
                    <li><a use:link href="/curr">CURRENCY</a></li>
                    <li><a use:link href="/log">LOG MANAGEMENT</a></li>
                  </ul>
                </li>
              </ul>
            </div>
            <a class="btn btn-ghost normal-case text-xl" href="/">WHITELABEL</a>
        </div>
        <div class="navbar-center hidden lg:flex">
          <ul class="menu menu-horizontal p-0">
            <li><a class="text-sm" href="/">DASBOARD</a></li>
            <li tabindex="0">
              <a class="text-sm">
                AGEN
                <svg class="fill-current" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24"><path d="M7.41,8.58L12,13.17L16.59,8.58L18,10L12,16L6,10L7.41,8.58Z"/></svg>
              </a>
              <ul class="p-2 bg-white">
                <li><a use:link href="/agen" class="text-sm">AGEN</a></li>
              </ul>
            </li>
            <li tabindex="0">
                <a class="text-sm">
                  MASTER
                  <svg class="fill-current" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24"><path d="M7.41,8.58L12,13.17L16.59,8.58L18,10L12,16L6,10L7.41,8.58Z"/></svg>
                </a>
                <ul class="p-2 bg-white">
                  <li><a use:link href="/log"  class="text-sm">LOG MANAGEMENT</a></li>
                  <li><a use:link href="/admin"  class="text-sm">ADMIN MANAGEMENT</a></li>
                  <li><a use:link href="/adminrule"  class="text-sm">ADMIN RULE</a></li>
                </ul>
            </li>
            <li><a class="text-sm" use:link href="/setting">SETTING</a></li>
          </ul>
        </div>
        <div class="navbar-end">
            <button
                on:click={() => {
                    handleLogout();
                }}
                class="btn btn-ghost btn-circle">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-black fill-current" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                </svg>
            </button>
        </div>
      </div>
    
</div>