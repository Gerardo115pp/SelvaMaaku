<script>
    import ArrowIcon from "@components/icons/arrow.svelte";
    import { page } from "$app/stores";
    import { onMount } from "svelte";
    import ThemeButton from "@components/buttons/theme_button.svelte";
    import { layout_images, layout_properties } from "@stores/layout";
    import { writable } from "svelte/store";
    import DropdownMenu from "./sub-components/DropdownMenu.svelte";
    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        /**
         * @typedef {Object} NavbarSections
         * @property {string} name
         * @property {string} href
         * @property {NavbarSections[]} options 
        */

        /** @type {NavbarSections[]} */
        let dropdown_sections = [
            {
                name: 'Home',
                href: '/',
                options: []
            },
            {
                name: 'Development',
                href: '/development',
                options: []
            },
            {
                name: 'Location',
                href: '/location',
                options: []
            },
            {
                name: 'Models',
                href: '/house-models',
                options: []
            },
            {
                name: 'Amenities',
                href: '/amenities',
                options: []
            }
        ]

        /**
         * Whether the dropdown menu is opened or not
         * @type {import('svelte/store').Writable<boolean>}
         */
        let is_dropdown_open = writable(false);

        /**
         * the scroll amount of the document
         * @type {number}
         */
        let scroll_amount = 0;
        
    
    /*=====  End of Properties  ======*/

    onMount(() => {
        // is_dropdown_open.set(true);
    })

    
    /*=============================================
    =            Methods            =
    =============================================*/
    
        const toggleDropdown = () => {
            is_dropdown_open.set(!$is_dropdown_open);
        }
    
    /*=====  End of Methods  ======*/
    
    
    
</script>


<svelte:window bind:scrollY={scroll_amount} />
<nav id="selvamaaku-navbar" class:scrolled={scroll_amount >= 200} class:adebug={false} class:has-dropdown-menu={$is_dropdown_open}>
    <div id="smk-navbar-content">
        <div id="mobile-burger-btn">
            <button class:menu-opened={$is_dropdown_open} on:click={toggleDropdown}>
                <svg viewBox="0 0 24 25">
                    <path d="M3 6.5H21"/>
                    <path d="M3 12.5H21"/>
                    <path d="M3 18.5H21"/>
                </svg>                
            </button>
        </div>
        <div id="sn-selvamaaku-logo">
            <img src="{layout_images.SELVAMAAKU_LOGOTYPE.getUrl(0.2)}" alt="Selva maaku logotipo"/>
        </div>
        <menu id="sn-navoptions" class:hide-on-mobile={$layout_properties.IS_MOBILE}>
            {#each dropdown_sections as nav_option}
                {@const is_active_link = $page.url.pathname === nav_option.href}
                <li class="nav-option-wrapper" class:is-active-link={is_active_link}>
                    <a href="{nav_option.href}" class="smk-link">
                        {nav_option.name}
                    </a>
                </li>
            {/each}
        </menu>
        <div id="sn-cta-section">
            <select name="language-controller" id="language-controller">
                <option value="en">EN</option>
                <option value="es">ES</option>
            </select>
            <ThemeButton 
                text="Contact us"
                is_button_one
                href="/contact"
            />
        </div>
    </div>
</nav>
<DropdownMenu
    is_open={is_dropdown_open}
    dropdown_sections={dropdown_sections}
/>

<style>
    #selvamaaku-navbar {
        --navbar-background: var(--theme-hero-background);

        width: 100dvw;
        height: var(--navbar-height);
        display: grid;
        position: fixed;
        place-items: center;
        container-type: inline-size;
        padding: 0 5.5%;
        z-index: var(--z-index-t-5);
        transition: background-color 0.3s ease;
    }

    #selvamaaku-navbar.scrolled {
        background-color: var(--navbar-background);
    }

    #selvamaaku-navbar.scrolled:hover {
        background-color: var(--navbar-background);
    }

    @supports (color: rgb( from white r g b)) {
        #selvamaaku-navbar {
            --base-bg-color: var(--theme-hero-background);
            --navbar-background: hsl(from var(--base-bg-color) h s calc(l * 0.5) / 0.3);
        }

        #selvamaaku-navbar.has-dropdown-menu {
            background-color: var(--base-bg-color) !important;
        }

        
        @media(pointer: fine) {            
            #selvamaaku-navbar.scrolled:not(.has-dropdown-menu):hover {
                background-color: hsl(from var(--base-bg-color) h s l / 0.8);   
            }
        }
    }

    #smk-navbar-content {
        width: 100%;
        display: flex;
        align-items: center;
        gap: var(--spacing-3);
    }

    
    /*=============================================
    =            burger menu            =
    =============================================*/
    
        #mobile-burger-btn {
            display: none;
        }

        @container (width <= 1080px) {
            #mobile-burger-btn {
                display: block;
            }

            #mobile-burger-btn svg {
                width: 31px;
            }

            #mobile-burger-btn svg path {
                stroke: var(--accent-orange-light);
                stroke-width: 2px;
                stroke-linecap: round;
                transform-box: fill-box;
                transition: rotate 0.4s linear, opacity 0.2s ease-in, translate 0.2s ease;
            }

            #mobile-burger-btn button {
                outline: none;
            }

            #mobile-burger-btn button.menu-opened {
                --burger-line-rotation: 41deg;
            }

            #mobile-burger-btn button.menu-opened path:first-child {
                transform-origin: left center;
                rotate: var(--burger-line-rotation);
            }

            #mobile-burger-btn button.menu-opened path:nth-child(2) {
                translate: 0 5px 0;
                opacity: 0;
            }

            #mobile-burger-btn button.menu-opened path:last-child {
                transform-origin: left center;
                rotate: calc(-1 * var(--burger-line-rotation));
            }
        }
    
    /*=====  End of burger menu  ======*/
    
    

    
    /*=============================================
    =            Logo            =
    =============================================*/
    
        #sn-selvamaaku-logo {
            width: 132px;

            & img {
                width: 100%;
            }
        }

    
    /*=====  End of Logo  ======*/
    
    

    
    /*=============================================
    =            Navoptions            =
    =============================================*/
    
        menu#sn-navoptions {
            display: flex;
            align-items: center;
            gap: calc(var(--spacing-2));
        }

        .nav-option-wrapper a {
            color: var(--light-orange-light);
            font-family: var(--font-titles);
            font-size: var(--font-size-1);
            white-space: nowrap;
        }

        .is-active-link a {
            color: var(--light-orange);
            pointer-events: none;
        }
    
    /*=====  End of Navoptions  ======*/
    
    
    
    /*=============================================
    =            cta section            =
    =============================================*/
    
        #sn-cta-section {
            display: flex;
            align-items: center;
            gap: var(--spacing-2);
            margin-left: auto;
        }
    
    /*=====  End of cta section  ======*/
    
    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @container (width <= 768px) {
            #sn-selvamaaku-logo {
                display: none;
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    
    

</style>