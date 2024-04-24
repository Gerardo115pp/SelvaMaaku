<script>
    import ThemeButton from "@components/buttons/theme_button.svelte";
    import viewport from "@components/viewport_actions/useViewportActions";
    import { layout_images, layout_properties } from "@stores/layout";

    
    /*=============================================
    =            Properties            =
    =============================================*/

        /**
         * Whether the component is visible or not
         * @type {boolean}
         */
        let component_visible = false;
    
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
    
    

</script>

<section id="smk-global-footer-section"
    on:viewportEnter={() => component_visible = true}
    on:viewportLeave={() => component_visible = false}
    class:is-visible={component_visible}
    use:viewport={{height_offset: 0.8}}
>
    <div id="smk-gfs-content-wrapper" class="design-content-width">
        <div id="smk-gfs-combination-mark-wrapper">
            <img src="{layout_images.COMBINATION_MARK.getUrl(0.13)}" alt="selvamaaku combination mark">
        </div>
        <hgroup id="smk-gfs-information">
            <h2 id="smk-gfs-headline">
                We are ready to assist you
            </h2>
            <p id="smk-gfs-subheadline">
                Let’s start talking about living in the paradise you will soon call home
            </p>
        </hgroup>
        <div id="smk-gfs-contact-cta">
            <ThemeButton text="Contact Us" href="/contact" is_button_one/>
        </div>
    </div>
    <footer id="smk-global-footer">
        <menu id="footer-navoptions" class:hide-on-mobile={$layout_properties.IS_MOBILE}>
            {#each dropdown_sections as nav_option}
                <li class="smk-footer-nav-option-wrapper">
                    <a href="{nav_option.href}" class="smk-link">
                        {nav_option.name}
                    </a>
                </li>
            {/each}
        </menu>
        <div id="footer-policies-wrapper">
            <ul id="smk-policies-list">
                <li class="smk-pl-policy-item">
                    <a href="/#" class="smk-link">
                        © Selva Máaku
                    </a>
                </li>
                <li class="smk-pl-policy-item">
                    <a href="/#" class="smk-link">
                        All rights reserved
                    </a>
                </li>
                <li class="smk-pl-policy-item">
                    <a href="/#" class="smk-link">
                        Terms and Conditions
                    </a>
                </li>
            </ul>
        </div>
    </footer>
</section>

<style>

    /*=============================================
    =            Contact footer            =
    =============================================*/

        #smk-global-footer-section {
            display: flex;
            flex-direction: column;
            container-type: inline-size;
            align-items: center;
            padding: var(--spacing-5) 0 0 0;
        }

        #smk-gfs-content-wrapper {
            display: flex;
            flex-direction: column;
            align-items: center;
            gap: calc(var(--spacing-3) + var(--spacing-1)); 
            padding: 0 0 var(--spacing-5) 0;
            opacity: 0;
        }
        
        #smk-global-footer-section.is-visible #smk-gfs-content-wrapper {
            transition: opacity 1.5s ease-in-out;
            opacity: 1;
        }

        #smk-gfs-information {
            display: flex;
            flex-direction: column;
            gap: calc(var(--spacing-2) + var(--spacing-1));
            align-items: center;
            text-align: center;
        }  

    /*=====  End of Contact footer  ======*/

    
    /*=============================================
    =            Footer            =
    =============================================*/
    
        footer#smk-global-footer {
            width: 100%;
            display: flex;
            padding: calc(var(--spacing-4) - var(--spacing-2)) var(--spacing-5);
            justify-content: space-between;
        }

        #footer-navoptions {
            display: flex;
            gap: var(--spacing-2);
            padding: 0;

            & li a {
                font-family: var(--font-titles);
                color: var(--light-orange-light);
                font-weight: 500;
            }
        }

        #footer-policies-wrapper ul {
            display: flex;
            gap: calc(1 * var(--spacing-2));
            
            & li a {
                font-family: var(--font-titles);
                color: var(--light-orange-light);
                font-weight: 500;
            }
        }

        #footer-policies-wrapper ul li:not(:last-child)::after {
            content: "•";
            margin-left: var(--spacing-2);
        }
    
    /*=====  End of Footer  ======*/
    
    
    /*=============================================
    =            Mobile            =
    =============================================*/

        @media only screen and (max-width: 880px) {
            section#smk-global-footer-section {
                row-gap: var(--spacing-4);
                padding: var(--spacing-5) 0 var(--spacing-4) 0;
            }
        }
    
        @container (width <= 768px) {
            #smk-global-footer-section #smk-gfs-content-wrapper {
                margin: 0;
                max-width: none;
                width: 100%;
                padding: 0;
            }

            #smk-gfs-content-wrapper #smk-gfs-combination-mark-wrapper {
                width: 51.0204cqw;
            }

            #smk-gfs-information h2 {
                font-size: calc(var(--font-size-h1) * 0.80);
            }

            footer#smk-global-footer {
                padding: 0;
            }

            #footer-policies-wrapper ul {
                gap: 0;
                flex-wrap: wrap;
                justify-content: center;
            }

            #footer-policies-wrapper ul li a {
                white-space: nowrap;
            }

            #footer-policies-wrapper ul li:not(:last-child)::after {
                content: '.';
                margin-left: 0;
                margin-right: 1ch;
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    

</style>