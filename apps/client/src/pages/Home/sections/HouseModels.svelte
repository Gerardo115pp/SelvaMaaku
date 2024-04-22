<script>
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import ModelsGallery from "../sub-components/models_gallery.svelte";
    import ThemeButton from "@components/buttons/theme_button.svelte";
    import viewport from "@components/viewport_actions/useViewportActions";
    import { fade, slide } from "svelte/transition";
    import { elasticOut, sineOut } from "svelte/easing";

    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        let component_visible = false;

        
        /*----------  Styling  ----------*/
        
            /**
             * The duration of the entry animation in milliseconds
             * @type {number}
             */
            const entry_animation_duration = 1500;

            /**
             * The transition easing function
             * @type {Function}
             */
            const easing_function = sineOut;
    
    /*=====  End of Properties  ======*/
    
    
    /*=============================================
    =            Methods            =
    =============================================*/
    
        /**
         * Sets the component visibility
         * @param {boolean} is_visible
        */
        const setComponentVisibility = (is_visible) => {
            component_visible = is_visible;
        }
    
    /*=====  End of Methods  ======*/
    
    
    

</script>

<section id="smk-house-models-section" 
    class:adebug={false}
    class:is-visible={component_visible}
    style:--entry-animation-duration="{entry_animation_duration}ms"
    on:viewportEnter={() => setComponentVisibility(true)}
    on:viewportLeave={() => setComponentVisibility(false)}
    use:viewport={{height_offset: 0.30}}
>
    <div id="smk-hms-content-wrapper" class="design-content-width">
        <div id="smk-hms-cw-text-column">
            {#if component_visible}
                <div class="accent-wixarica-icon-wrapper wixarica-icons" transition:slide={{axis: 'y', delay: entry_animation_duration / 5, easing: easing_function, duration: entry_animation_duration * 0.7}}>
                    <WixaricaIcon with_accent />
                </div>
                <hgroup id="smk-hms-cw-tc-models-information">
                    <h2 id="smk-hms-cw-tc-models-headline" transition:fade={{ delay: entry_animation_duration / 4, easing: easing_function, duration: entry_animation_duration * 0.7}}>
                        MODELS
                    </h2>
                    <h3 id="smk-hms-cw-tc-models-subheadline" transition:fade={{delay: entry_animation_duration / 3, easing: easing_function, duration: entry_animation_duration * 0.7}}>
                        Your house<br/>in the jungle
                    </h3>
                    <p id="smk-hms-cw-tc-models-text" transition:fade={{delay: entry_animation_duration / 2, duration: entry_animation_duration * 0.7}}>
                        This unique place where the sun shines bright every day, you will be able to contemplate incredible views of the jungle. The natural light and the shine of the minerals embedded in the stone materials of the area are the inspiration for the design of the 14 houses that make up this exclusive development.
                    </p>
                </hgroup>
                <ul class="smk-hms-cw-ctas-wrapper" transition:fade={{axis: 'x', delay: entry_animation_duration, easing: easing_function, duration: entry_animation_duration * 0.7}}>
                    <ThemeButton
                        href="/house-models"
                        text="Models"
                    />
                </ul>
            {/if}
        </div>
        <div id="smk-hms-cw-gallery-column">
            <ModelsGallery
                gallery_visible={component_visible}
                {entry_animation_duration}
            />
        </div>
    </div>
</section>

<style>
    #smk-house-models-section {
        display: flex;
        justify-content: center;
        padding: var(--spacing-5) 0;
    }

    #smk-hms-content-wrapper {
        display: flex;
        column-gap: var(--spacing-5);
        container-type: inline-size;
    }

    
    /*=============================================
    =            Text column            =
    =============================================*/
    
        #smk-hms-cw-text-column {
            display: flex;
            flex-direction: column;
            min-width: 40%;
            gap: var(--spacing-4);
            opacity: 0;
        }

        #smk-house-models-section.is-visible #smk-hms-cw-text-column {
            opacity: 1;
            transition: opacity calc(var(--entry-animation-duration) * 0.3) cubic-bezier(1, 0.71, 0.17, 1);
        }

        hgroup#smk-hms-cw-tc-models-information {
            display: flex;
            max-width: 480px;
            flex-direction: column;
            gap: var(--spacing-3);
        }

        hgroup#smk-hms-cw-tc-models-information h2 {
            font-family: var(--font-read);
            font-size: var(--font-size-2);
            font-weight: 600;
        }

        hgroup#smk-hms-cw-tc-models-information h3 {
            font-size: var(--font-size-h3);
            font-weight: 400;
        }
    
    /*=====  End of Text column  ======*/
    
    #smk-hms-cw-gallery-column {
        opacity: 0;
    }

    #smk-house-models-section.is-visible #smk-hms-cw-gallery-column {
        opacity: 1;
        transition: opacity calc(var(--entry-animation-duration) * 0.5) cubic-bezier(1, 0.71, 0.17, 1);
    }

    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @media only screen and (max-width: 768px) {
            #smk-hms-content-wrapper {
                flex-direction: column;
                row-gap: var(--spacing-4);
            }

            hgroup#smk-hms-cw-tc-models-information h3 {
                font-size: var(--font-size-h1);
            }
        } 
    
    /*=====  End of Mobile  ======*/
    
    
</style>