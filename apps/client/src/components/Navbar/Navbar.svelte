<script>
    import ArrowIcon from "@components/icons/arrow.svelte";
    import { page } from "$app/stores";
    import { onMount } from "svelte";
    import ThemeButton from "@components/buttons/theme_button.svelte";
    import { layout_images } from "@stores/layout";
    
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
    
    /*=====  End of Properties  ======*/

    onMount(() => {
        console.log($page);
    })
    
    
</script>

<nav id="selvamaaku-navbar" class:adebug={false}>
    <div id="smk-navbar-content">
        <div id="sn-selvamaaku-logo">
            <img src="{layout_images.SELVAMAAKU_LOGOTYPE.getUrl(0.2)}" alt="Selva maaku logotipo"/>
        </div>
        <menu id="sn-navoptions">
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

<style>
    #selvamaaku-navbar {
        width: 100vw;
        height: var(--navbar-height);
        display: grid;
        position: sticky;
        place-items: center;
        container-type: inline-size;
        padding: 0 5.5%;
        z-index: var(--z-index-t-5);
    }

    #smk-navbar-content {
        width: 100%;
        display: flex;
        align-items: center;
        gap: var(--spacing-3);
    }

    
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
    
    

</style>