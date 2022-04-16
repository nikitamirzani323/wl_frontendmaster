<script>
    import { createEventDispatcher } from "svelte";
    export let button_title = "";
    export let button_class = "";
    export let button_disable_class = "";
    export let button_disable = false;
    export let button_function = "";
    let dispatch = createEventDispatcher()
</script>
{#if !button_disable}
    <button
        on:click={() => dispatch('click',button_function)}
        class="btn btn-primary {button_class} m-0 h-1 min-h-[40px] shadow-lg">
        {button_title}
    </button>
{:else}
    <button class="btn {button_disable_class}  loading">Loading</button>
{/if}