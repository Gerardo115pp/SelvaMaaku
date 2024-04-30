<script>
    import ThemeButton from "@components/buttons/theme_button.svelte";
    import Mexico from "@components/icons/mexico.svelte";
    import SayulitaMap from "@components/icons/sayulita_map.svelte";
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import viewport from "@components/viewport_actions/useViewportActions";
    import { layout_properties } from "@stores/layout";
    import { site_language, supported_languages } from "@stores/site_content";
    import { blur, fly } from "svelte/transition";

    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        /**
         * whether the component is currently visible
         * @type {boolean}
         */
        let component_visible = false;

        
        /*----------  Animation properties  ----------*/

            /**
             * Animation total duration in milliseconds
             * @type {number}
             */
            const section_animation_duration = 5000;

            /**
             * Mexico map animation duration percentage
             * @type {number}
             */
            const mexico_animation_duration = (section_animation_duration * 0.4);

            /**
             * Mexico map visible
             * @type {boolean}
             */
            let mexico_map_visible = false;

            /**
             * Sayulita map animation duration percentage
             * @type {number}
             */
            const sayulita_animation_duration = section_animation_duration * 0.4;

            /**
             * Sayulita map visible
             * @type {boolean}
             */
            let sayulita_map_visible = false;

            /**
             * Is location content visible
             * @type {boolean}
             */
            let location_content_visible = false;
            
            /**
             * Location content animation duration
             * @type {number}
             */
            const location_content_animation_duration = section_animation_duration * 0.3;

            /**
             * The animation elements that need to be dimmed
             * @type {HTMLElement[]}
             */
            const dimmed_elements = [];


        
        
    
    /*=====  End of Properties  ======*/
    

    
    /*=============================================
    =            Methods            =
    =============================================*/

        /**
         * Dims the elements that need to be dimmed
         */
        const dimAnimationElements = () => {
            dimmed_elements.forEach((element) => {
                element.style.opacity = 0.3;
            });
        }

        /**
         * Manages the appearance of the animation components
        */
        const manageAnimation = () => {
            if ($layout_properties.IS_MOBILE) return;

            mexico_map_visible = true;

            if (component_visible) {
                setTimeout(() => {
                    sayulita_map_visible = true;

                    setTimeout(() => {
                        location_content_visible = true;
                        dimAnimationElements();
                    }, sayulita_animation_duration);

                }, mexico_animation_duration);
            }
        }

        /**
         * Registers an animation element to be dimmed
         * @param {HTMLElement} node
         */
        const registerDimmedElement = (node) => {
            dimmed_elements.push(node);
        }
    
        /**
         * sets the component visibility
         * @param {boolean} is_visible
         */
        const setComponentVisibility = (is_visible) => {
            component_visible = is_visible;

            if (is_visible) {
                manageAnimation();
            }
        }
    
    /*=====  End of Methods  ======*/
    
    
    
</script>

<section id="smk-home-location-section" class:adebug={false} on:viewportEnter={() => setComponentVisibility(true)} class:is-animated={$layout_properties.IS_MOBILE} use:viewport={{height_offset: 0.5}}>
    <div id="smk-hls-content-wrapper" class="design-content-width">
        {#key mexico_map_visible}        
            <div id="smk-hls-cw-mexico-map-wrapper" class="smk-hls-bg-element" style:transition="opacity {location_content_animation_duration * 0.2}ms ease" in:fly={{x: -400, duration: mexico_animation_duration}} use:registerDimmedElement>
                {#if mexico_map_visible}
                    <Mexico
                        mexico_opacity={0.4}
                        map_pin_opacity={0.3}
                        animation_duration={mexico_animation_duration * 1.2}
                        animation_delay={mexico_animation_duration * 0.3}
                        animate
                    />
                {/if}
            </div>
        {/key}
        <div id="smk-hls-cw-sayulita-map-wrapper" class="smk-hls-bg-element" style:transition="opacity {location_content_animation_duration * 0.2}ms ease" use:registerDimmedElement>
            {#if sayulita_map_visible}
                <SayulitaMap 
                    animation_duration={sayulita_animation_duration}
                    animate
                />
            {/if}
        </div>
        {#if location_content_visible || ($layout_properties.IS_MOBILE && component_visible)}
             <div id="smk-hls-cw-content" in:blur={{amount: 1.5, duration: location_content_animation_duration}}>
                 <div class="smk-hls-cw-top-icon wixarica-icons">
                     <WixaricaIcon with_accent />
                 </div>
                 <hgroup id="smk-hls-cw-sayulita-information">
                     <h2 id="smk-hls-cw-sayulita-headline">
                         SAYULITA
                     </h2>
                     <h3 id="smk-hls-cw-sayulita-subheadline">
                        {#if $site_language !== supported_languages.SPANISH}
                             The spectacular Hidden Gem of the Riviera Nayarit
                        {:else}
                            La espectacular joya escondida de la Riviera Nayarit
                        {/if}
                     </h3>
                     <p id="smk-hls-cw-sayulita-text">
                        {#if $site_language !== supported_languages.SPANISH}
                            Sayulita, Mexico, is a vibrant and picturesque small town located on the Pacific coast in the state of Nayarit. Known for its stunning beaches, relaxed atmosphere, and rich cultural heritage.
                        {:else}
                            Sayulita, México, es un vibrante y pintoresco pueblo ubicado en la costa del Pacífico en el estado de Nayarit. Conocido por sus impresionantes playas, ambiente relajado y rica herencia cultural.
                        {/if}
                     </p>
                 </hgroup>
                 <ul class="smk-hls-cw-ctas-wrapper">
                     <ThemeButton 
                         href="/location"
                         text={$site_language !== supported_languages.SPANISH ? "Location" : "Ubicación"}
                     />
                 </ul>
             </div>
        {/if}
    </div>
</section>

<style>
    #smk-home-location-section {
        width: 100%;
        display: flex;
        justify-content: center;
        padding: var(--spacing-5) 0;
    }

    #smk-hls-content-wrapper {
        position: relative;
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        min-height: 567px;
        gap: var(--spacing-6);
    }

    
    /*=============================================
    =            Background elements            =
    =============================================*/
    
        .smk-hls-bg-element {
            visibility: visible;
        }
    
        #smk-hls-cw-mexico-map-wrapper {
            grid-column: 1 / span 1;
            display: flex;
            align-items: center;
        }
    
        #smk-hls-cw-sayulita-map-wrapper {
            grid-column: 2 / span 1;
            display: flex;
            align-items: center;
        }

    /*=====  End of Background elements  ======*/
    
    #smk-hls-cw-content {
        position: absolute;
        display: flex;
        width: 100%;
        inset: 0 0;
        flex-direction: column;
        align-items: center;
        gap: var(--spacing-4);
    }
    
    /*=============================================
    =            Information             =
    =============================================*/

        hgroup#smk-hls-cw-sayulita-information {
            display: flex;
            max-width: 640px;
            flex-direction: column;
            align-items: center;
            gap: var(--spacing-3);
            text-align: center;
        }

        h2#smk-hls-cw-sayulita-headline {
            font-family: var(--font-read);
            font-size: var(--font-size-2);
            font-weight: 600;
        }
        
        h3#smk-hls-cw-sayulita-subheadline {
            font-size: var(--font-size-h3);
        }
    
    /*=====  End of Information  ======*/

    
    /*=============================================
    =            CTA controls            =
    =============================================*/
    
        ul.smk-hls-cw-ctas-wrapper {
            display: flex;
            gap: var(--spacing-2);
        }
    
    /*=====  End of CTA controls  ======*/
    
    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @media only screen and (max-width: 768px) {
            h3#smk-hls-cw-sayulita-subheadline {
                font-size: var(--font-size-h1);
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    
    
    

</style>