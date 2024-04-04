<script>
    import AboutMangoGallery from "../sub-components/about_mango_gallery.svelte";
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import ArrowIcon from "@components/icons/arrow.svelte";
    import { layout_images } from "@stores/layout";

</script>

<section id="smk-home-about-section" class:adebug={false}>
    <div id="smk-has-content-wrapper">
        <div class="smk-has-cw-text-column">
            <div class="accent-wixarica-icon-wrapper">
                <WixaricaIcon with_accent/>
            </div>
            <h2 id="smk-has-cw-tc-about-headline">
                Comfort and sustainability in harmony.
            </h2>
            <p id="smk-has-cw-tc-about-text">
                Selva Máaku is a 5.4 acres development that is located half a mile from the outskirts of Sayulita on the Nayarit Riviera in Mexico. It is one mile away from the Pacific Ocean and in the midst of a tropical jungle.
            </p>
            <div class="smk-has-cw-tc-about-cta-wrapper">
                <button class="smk-has-cw-tc-about-cta button-2">
                    <span>Development</span>
                    <ArrowIcon />
                </button>
            </div>
        </div>
        <div id="smk-has-cw-gallery-column">
            <div class="smk-has-cw-gc-gallery-wrapper">
                <AboutMangoGallery />
            </div>
            <div class="smk-has-cw-gc-mango-underlay">
                <img src="{layout_images.MANGO_ISOTYPE.getUrl(0.1)}" alt="mango logo" aria-hidden="true">
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
    
    
    
    



</style>