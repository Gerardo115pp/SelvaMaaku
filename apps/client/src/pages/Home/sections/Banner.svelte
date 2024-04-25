<script>
    import ImageMultiStage from "@components/Images/ImageMultiStage.svelte";
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import { layout_images, layout_properties } from "@stores/layout";
    import viewport from "@components/viewport_actions/useViewportActions";
    import { fade } from "svelte/transition";

    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        export let component_visible = false;

        /**
         * Whether the component visibility is locked. used to hide the component when not completely out of the viewport
         * @type {boolean}
         * @default false
         */
        let visibility_locked = false;
    
    /*=====  End of Properties  ======*/

    
    /*=============================================
    =            Methods            =
    =============================================*/
    
        /**
         * Sets the component visibility
         * @param {boolean} is_visible
         * @param {boolean} toggle_lock_visibility 
         */
        const setComponentVisibility = (is_visible, toggle_lock_visibility=false) => {
            // in mobile, we just show the component whenever it enters the viewport and the don't hide it ever again.
            if (component_visible && $layout_properties.IS_MOBILE) return;
             
            if (!visibility_locked || (toggle_lock_visibility && !visibility_locked)) {
                component_visible = is_visible;
            }

            if (toggle_lock_visibility) {
                visibility_locked = !visibility_locked;
            }
        }
    
    /*=====  End of Methods  ======*/
    
    
    
</script>

<section id="smk-home-banner-section" 
    class:adebug={false} 
    on:viewportEnter={() => setComponentVisibility(true)}
    on:viewportLeave={() =>  setComponentVisibility(false, visibility_locked)}
    use:viewport={{height_offset: 0.19}}

>
    {#if component_visible}
        <div id="smk-hbs-background-image-wrapper" transition:fade={{duration: 1200}}>
            <ImageMultiStage 
                image_resource={layout_images.TOMMY_HOUSE_TWO}
                image_percentage={2}
                blur_unloaded
            />
        </div>
        <div id="smk-hbs-information-overlay" in:fade={{duration: 500, delay: 300}}>
            <div id="smk-hbs-io-content-wrapper">
                <div class="wixarica-icon-wrapper wixarica-icons-large">
                    <WixaricaIcon />
                </div>
                <h2 id="smk-hbs-io-headline" on:viewportLeave={() => setComponentVisibility(false, true)}  use:viewport>
                    A low density Urban Residential Environment
                </h2>
            </div>
        </div>
    {/if}
</section>    

<style>
    #smk-home-banner-section {
        position: relative;
        height: 90dvh;
        container-type: size;
        overflow: hidden;
    }

    #smk-home-banner-section #smk-hbs-background-image-wrapper {
        position: absolute;
        width: 120cqw;
        height: 120cqh;
        top: -5%;   
        z-index: var(--z-index-1);
        
        & > img {
            width: 100%;
            height: 100%;
            object-fit: cover;
        }
    }

    @keyframes parallax-timeline {
        from {
            transform: translateY(0px) scale(1);
        }
        to {
            transform: translateY(-15%) scale(1.2);
        }
    }
    
    @supports(animation-timeline: view()) {
        #smk-home-banner-section #smk-hbs-background-image-wrapper {
            /* scroll-timeline-name: --smk-home-banner-section-timeline; */

            & > img {
                animation-name: parallax-timeline;
                /* animation-timeline: view(block 30% 90%); */
                animation-timeline: scroll(root);
                animation-range: 40% 50%;
                /* animation-timeline: --smk-home-banner-section-timeline; */
                animation-duration: 1000ms;
                /* animation-iteration-count: infinite; */
                animation-fill-mode: both;
                animation-timing-function: linear;
                transform-origin: center top;
                transition: all 0.5s ease-in-out;
            }

        }
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
        z-index: var(--z-index-5);
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

    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @media only screen and (max-width: 768px) {
            #smk-hbs-io-headline {
                text-align: center;
                line-height: 1;
                letter-spacing: -2.2%;
                font-size: var(--font-size-h1);
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    
</style>