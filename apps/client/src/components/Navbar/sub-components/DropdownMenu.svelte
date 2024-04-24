<script>
    import { backOut, elasticIn, elasticInOut, elasticOut } from 'svelte/easing';
    import { slide } from 'svelte/transition';
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import { page } from "$app/stores";

    
    /*=============================================
    =            Properties            =
    =============================================*/


        /**
         * @typedef {Object} NavbarSections
         * @property {string} name
         * @property {string} href
         * @property {NavbarSections[]} options 
        */
    
        /**
         * The state of the dropdown menu(aka open or closed)
         * @type {import('svelte/store').Writable<boolean>}
         */ 
        export let is_open;

        /** @type {NavbarSections[]} */
        export let dropdown_sections;

        
        /*----------  Animation  ----------*/
        
            /**
             * The total animation duration in milliseconds
             * @type {number}
             */
            const animation_duration = 2000;

            /**
             * The point in the animation where the dropdown sections will start to appear
            */
            const dropdown_appear_point = animation_duration * 0.2;
    
    /*=====  End of Properties  ======*/
    
</script>

{#if $is_open}
    <menu id="smk-mobile-dropmenu-wrapper"
        style:animation-duration="{dropdown_appear_point}ms"
    >
        {#key $is_open}
            {#each dropdown_sections as ds, h}
                {@const is_active_link = $page.url.pathname === ds.href}
                <li class="smk-mdw-link" in:slide|global={{axis: 'y', duration: animation_duration * 0.33, delay: dropdown_appear_point + ((animation_duration * 0.05) * h), easing: backOut}}>
                    <div class="wixarica-icons">
                        <WixaricaIcon 
                            with_accent={is_active_link}
                        />
                    </div>
                    <a href="{ds.href}" on:click={() => is_open.set(false)}>
                        {ds.name}
                    </a>
                </li>
            {/each}
        {/key}
    </menu>
{/if}

<style>
    menu#smk-mobile-dropmenu-wrapper {
        display: flex;
        flex-direction: column;
        position: fixed;
        width: 100dvw;
        height: calc(100dvh -  var(--navbar-height));
        background-color: var(--theme-hero-background);
        top: var(--navbar-height);
        left: 0;
        align-items: flex-start;
        gap: var(--spacing-3);
        z-index: var(--z-index-t-5);
        animation-name: fade-in;
        animation-fill-mode: forwards;
        padding: var(--spacing-5) var(--spacing-2);
    }

    @keyframes fade-in {
        from {
            opacity: 0;
        }
        to {
            opacity: 0.95;
        }
    }

    menu#smk-mobile-dropmenu-wrapper li.smk-mdw-link {
        display: flex;
        gap: var(--spacing-1);
        align-items: center;
    }

    menu#smk-mobile-dropmenu-wrapper li a {
        font-family: var(--font-titles);
        font-size: var(--font-size-h1);
        color: var(--light-orange-light);
        text-decoration: none;
        line-height: 1;
    }
</style>