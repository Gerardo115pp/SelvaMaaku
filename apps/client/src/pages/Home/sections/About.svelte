<script>
    import AboutMangoGallery from "../sub-components/about_mango_gallery.svelte";
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import ArrowIcon from "@components/icons/arrow.svelte";
    import { layout_images, layout_properties } from "@stores/layout";
    import ThemeButton from "@components/buttons/theme_button.svelte";
    import viewport from "@components/viewport_actions/useViewportActions";
    import { fade, slide } from "svelte/transition";
    import { expoIn } from "svelte/easing";
    import { site_language, supported_languages } from "@stores/site_content";

    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        /** 
         * Whether the component is visible or not
         * @type {boolean}
         */
        let visible = false;


        /**
         * The text to display in the gallery
         * @type {string}
         */
        let gallery_text = $site_language !== supported_languages.SPANISH ? "The word Máaku means MANGO in the Wixárika language that is spoken by the original towns in the area." : 
            "La palabra Máaku significa MANGO en la lengua Wixárika que es hablada por los pueblos originarios de la zona."; 
    
    /*=====  End of Properties  ======*/

    
    /*=============================================
    =            Methods            =
    =============================================*/
    
        const setComponentVisibility = is_visible => {
            visible = is_visible;
        }
    
    /*=====  End of Methods  ======*/
    
    
    
    

</script>

<section id="smk-home-about-section" 
    class:adebug={false} 
    on:viewportEnter={() => setComponentVisibility(true)}
    use:viewport
>
    <div id="smk-has-content-wrapper">
        <div class="smk-has-cw-text-column" >
            {#if visible}            
                <div class="accent-wixarica-icon-wrapper" in:fade={{ delay:300, duration:1500}}>
                    <WixaricaIcon with_accent/>
                </div>
                <h2 id="smk-has-cw-tc-about-headline" in:fade={{ delay:500, duration:1500}}>
                    {#if $site_language !== supported_languages.SPANISH}
                        Comfort and sustainability in harmony.
                    {:else}
                        Comodidad y sostenibilidad en armonía.
                    {/if}
                </h2>
                <p id="smk-has-cw-tc-about-text" in:fade={{ delay:700, duration:1500}}>
                    {#if $site_language !== supported_languages.SPANISH}
                        Selva Máaku is a 5.4 acres development that is located half a mile from the outskirts of Sayulita on the Nayarit Riviera in Mexico. It is one mile away from the Pacific Ocean and in the midst of a tropical jungle.
                    {:else}
                        Selva Máaku es un desarrollo de 2.2 hectáreas que se encuentra a 800 metros de las afueras de Sayulita en la Riviera Nayarit en México. Está a 1.6 kilómetros del Océano Pacífico y en medio de una selva tropical.
                    {/if}
                </p>
                <div class="smk-has-cw-tc-about-cta-wrapper" in:fade={{ delay: 1000, duration:1500}}>
                    <ThemeButton 
                        href="/development"
                        text={$site_language !== supported_languages.SPANISH ? "Development" : "Desarrollo"}
                    />
                </div>
            {/if}
        </div>
        <div id="smk-has-cw-gallery-column">
            <div class="smk-has-cw-gc-gallery-wrapper">
                <AboutMangoGallery 
                    {gallery_text}
                />
            </div>
            <div class="smk-has-cw-gc-mango-underlay">
                <img src="{layout_images.MANGO_ISOTYPE.getUrl(!($layout_properties.IS_MOBILE) ? 0.1 : 0.6)}" alt="mango logo" aria-hidden="true">
            </div>
        </div>
    </div>
</section>


<style>
    #smk-home-about-section {
        display: flex;
        flex-direction: column;
        align-items: center;
        padding: calc(var(--spacing-5) - var(--spacing-1)) 0;
    }

    #smk-has-content-wrapper {
        --content-gap: var(--spacing-5);

        width: var(--design-content-width);
        container-type: inline-size;
        display: flex;
        column-gap: var(--content-gap);
    }

    
    /*=============================================
    =            Text column            =
    =============================================*/
    
        .smk-has-cw-text-column {
            display: flex;
            width: calc(calc(100cqw - var(--content-gap)) * 0.4);
            flex-direction: column;
            justify-content: center;
            gap: var(--spacing-3);
        }
 
        .accent-wixarica-icon-wrapper {
            width: 48px;
        }

        h2#smk-has-cw-tc-about-headline {
            font-size: var(--font-size-h3);
        }
    
    /*=====  End of Text column  ======*/

    
    /*=============================================
    =            Gallery column            =
    =============================================*/
    
        #smk-has-cw-gallery-column {
            --gallery-square-size: calc(calc(100cqw - var(--content-gap)) * 0.6);
            display: grid;
            position: relative;
            width: var(--gallery-square-size);
            height: var(--gallery-square-size);
            place-items: center;
        }

        .smk-has-cw-gc-gallery-wrapper {
            width: 100%;
            height: 100%;
            position: absolute;
            top: 0;
            left: 0;
        }

        .smk-has-cw-gc-mango-underlay {
            width: max-content;
        }

        .smk-has-cw-gc-mango-underlay img {
            width: 17.1527cqw;
            user-select: none;
            pointer-events: none;
        }
    
    /*=====  End of Gallery column  ======*/
    
    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @media only screen and (max-width: 768px) {
            #smk-has-content-wrapper {
                flex-direction: column;
                row-gap: var(--spacing-4);
            }

            .smk-has-cw-text-column{
                width: 100cqw;
            }

            h2#smk-has-cw-tc-about-headline {
                font-size: var(--font-size-h1);
            }

            #smk-has-cw-gallery-column {
                --gallery-square-size: 100cqw;
            }

            .smk-has-cw-gc-mango-underlay img {
                width: 57.1527cqw;
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    
    
    



</style>