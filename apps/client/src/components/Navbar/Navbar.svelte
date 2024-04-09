<script>
    import ArrowIcon from "@components/icons/arrow.svelte";
    
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
                href: '/ubicacion',
                options: []
            },
            {
                name: 'Models',
                href: '/modelos',
                options: []
            },
            {
                name: 'Amenities',
                href: '/amenidades',
                options: []
            },
            {
                name: 'Life in Sayulita',
                href: '/vive-en-sayulita',
                options: []
            }
        ]
    
    /*=====  End of Properties  ======*/
    
    
</script>

<nav id="selvamaaku-navbar" class:adebug={false}>
    <div id="smk-navbar-content">
        <div id="sn-selvamaaku-logo">
            <img src="/resources/images/logos/selva_maaku_logo.webp" alt="Selva maaku logotipo"/>
        </div>
        <menu id="sn-navoptions">
            {#each dropdown_sections as nav_option}
                <li class="nav-option-wrapper">
                    <a href="{nav_option.href}">
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
            <button id="contact-cta-btn" class="button-1">
                <span>Contact us</span>
                <ArrowIcon />
            </button>
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
            white-space: nowrap;
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

        #sn-cta-section button#contact-cta-btn span {
            white-space: nowrap;
        }
    
    /*=====  End of cta section  ======*/
    
    

</style>