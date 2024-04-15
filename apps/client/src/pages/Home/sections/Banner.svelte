<script>
    import ImageMultiStage from "@components/Images/ImageMultiStage.svelte";
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import { layout_images } from "@stores/layout";
    import viewport from "@components/viewport_actions/useViewportActions";
    import { fade } from "svelte/transition";

    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        export let component_visible = false;
    
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

{#key component_visible}
    <section id="smk-home-banner-section" class:adebug={false} on:viewportEnter={() => setComponentVisibility(true)} use:viewport={{height_offset: 0.2}} in:fade={{duration: 1200}}>
        {#if component_visible}
            <ImageMultiStage 
                image_resource={layout_images.TOMMY_HOUSE_TWO}
                blur_unloaded
            />
            <div id="smk-hbs-information-overlay">
                <div id="smk-hbs-io-content-wrapper">
                    <div class="wixarica-icon-wrapper wixarica-icons-large">
                        <WixaricaIcon />
                    </div>
                    <h2 id="smk-hbs-io-headline">
                        A low density Urban Residential Environment
                    </h2>
                </div>
            </div>
        {/if}
    </section>
{/key}

<style>
    #smk-home-banner-section {
        position: relative;
        height: 90dvh;
    }

    #smk-hbs-information-overlay {
        position: absolute;
        display: flex;
        width: 100%;
        height: 100%;
        justify-content: center;
        top: 0;
        left: 0;
        padding: var(--spacing-5) 0;
        background-color: hsl(from var(--theme-color-dark) h s l / 0.5);
    }

    #smk-hbs-io-content-wrapper {
        display: flex;
        max-width: 797px;
        flex-direction: column;
        align-items: center;
        gap: calc(var(--spacing-3) + var(--spacing-1));
    }

    #smk-hbs-io-headline {
        text-align: center;
        line-height: 72px;
        letter-spacing: -2.2%;
    }
</style>