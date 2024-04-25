<script>
    import AmenitiesGallery from "../sub-components/amenities_gallery.svelte";
    import { layout_images, layout_properties } from "@stores/layout";
    import ThemeButton from "@components/buttons/theme_button.svelte";
    import viewport from "@components/viewport_actions/useViewportActions";
    import { fade } from "svelte/transition";
    import { writable } from "svelte/store";

    
    /*=============================================
    =            Properties            =
    =============================================*/ 
    
        /**
         * Whether to run animations or not
         */
        export let animate = true;

        /**
         * Whether the component is visible or not
         * @type {import("svelte/store").Writable<boolean>}
         */
        let component_visible = writable(!animate);

        
        /*----------  Entry animation  ----------*/
        
            /**
             * Duration of the entry animation in milliseconds
             * @type {number}
             */
            const entry_animation_duration = 1500;
    
    /*=====  End of Properties  ======*/

    
    /*=============================================
    =            Methods            =
    =============================================*/
    
        const setVisibility = is_visible => {
            if ($component_visible === is_visible) return;

            component_visible.set(is_visible || !animate);
        }
    
    /*=====  End of Methods  ======*/
    
    
    
    

</script>

<section id="smk-amenities-section"
    class:is-visible={$component_visible}
    class:no-animation={!animate}
    style:--entry-animation-duration="{entry_animation_duration}ms"
>
    <div id="smk-as-content-wrapper" class="design-content-width">
        <div class="mango-seal-underlay-wrapper">
            <img src="{layout_images.MANGO_ISOLOGO.getUrl(!$layout_properties.IS_MOBILE ? 0.3 : 0.8)}" alt="" aria-hidden="true">
        </div>
        <hgroup id="smk-as-headline-wrapper"
            on:viewportEnter={() => setVisibility(true)} 
            on:viewportLeave={() => setVisibility(false)}
            use:viewport={{height_offset: 0.8}}
        >
            <h2 id="smk-as-headline" class="small-headline">
                AMENITIES
            </h2>
            <h3 id="smk-as-subheadline">
                Where nature meets luxury
            </h3>
            <p id="smk-as-description">
                Our community has been meticulously designed to cater to a diverse range of lifestyles and interests, ensuring that every day is filled with opportunities for relaxation, recreation, and socialization.
            </p>
        </hgroup>           
        <div id="smk-as-gallery-wrapper">
            <AmenitiesGallery 
                {entry_animation_duration}
                parent_visible={component_visible}
            />
        </div>
        <ul id="smk-as-ctas-container">
            <ThemeButton 
                text="Explore"
                href="/amenities"
                is_button_three
            />
        </ul>
    </div>
</section>

<style>
    #smk-amenities-section {
        display: flex;
        flex-direction: column;
        align-items: center;
        padding: var(--spacing-5) 0;
        opacity: 0;
        transition: opacity calc(var(--entry-animation-duration) * 0.3) ease-out;
    }

    #smk-amenities-section.is-visible, #smk-amenities-section.no-animation {
        opacity: 1;
    }
    
    #smk-as-content-wrapper {
        position: relative;
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: var(--spacing-4);
        container-type: inline-size;
    }

    
    /*=============================================
    =            Headline group            =
    =============================================*/
    
        hgroup#smk-as-headline-wrapper {
            display: flex;
            flex-direction: column;
            width: min(44.44dvw, 640px);
            align-items: center;
            gap: var(--spacing-3);
            text-align: center;
        }
    
    /*=====  End of Headline group  ======*/

    
    /*=============================================
    =            CTA group            =
    =============================================*/
    
        ul#smk-as-ctas-container {
            justify-content: center;
        }
    
    /*=====  End of CTA group  ======*/
    
    
    /*=============================================
    =            Mango underlay            =
    =============================================*/
    
        .mango-seal-underlay-wrapper {
            position: absolute;
            width: max-content;
            right: 0.2cqw;
            top: 10.69cqw;
            z-index: var(--z-index-b-4);
            opacity: 0;

            & img {
                width: 47.45065cqw;
            }
        }

        .is-visible .mango-seal-underlay-wrapper {
            transition: all calc(var(--entry-animation-duration) * 0.5) ease-in-out;
            animation-name: rollMango;
            animation-duration: calc(var(--entry-animation-duration) * 0.9);
            animation-delay: calc(var(--entry-animation-duration) * 1);
            animation-timing-function: cubic-bezier(0.34, 1.56, 0.64, 1);
            animation-fill-mode: forwards;
        }

        .no-animation .mango-seal-underlay-wrapper {
            opacity: 1 !important;
            animation-name: none !important;
        }

        @keyframes rollMango {
            0% {
                scale: 0.5;
                translate: -15% 0 0;
                rotate: -240deg;
                opacity: 0;
            }
            33% {
                opacity: 1;
            }
            100% {
                rotate: 0deg;
                scale: 1;
                translate: 0 0 0;
                opacity: 1;
            }
        }
    
    /*=====  End of Mango underlay  ======*/
    
    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @container (width <= 768px) {
            #smk-as-subheadline {
                font-size: var(--font-size-h1);
            }

            hgroup#smk-as-headline-wrapper {
                width: 100cqw;
            }

            .mango-seal-underlay-wrapper {
                top: 30%;
                right: -7cqw;
            }

            .mango-seal-underlay-wrapper img {  
                width: 70cqw;
            }
            
        }

        @container (width <= 500px) {
            #smk-amenities-section .mango-seal-underlay-wrapper {
                top: 40%;
            }
                
            #smk-amenities-section .mango-seal-underlay-wrapper img {  
                width: 90cqw;
            }
        }

        @container (width < 400px) {
            #smk-amenities-section .mango-seal-underlay-wrapper {
                top: 54%;
            }            
        }
    
    /*=====  End of Mobile   ======*/
    
    
    
    
    
    
    
</style>