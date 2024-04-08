<script>
    import Arrow from "@components/icons/arrow.svelte";
    import AmenitiesGallery from "../sub-components/amenities_gallery.svelte";
    import { layout_images } from "@stores/layout";

</script>
<section id="smk-amenities-section">
    <div id="smk-as-content-wrapper" class="design-content-width">
        <div class="mango-seal-underlay-wrapper">
            <img src="{layout_images.MANGO_ISOLOGO.getUrl(0.3)}" alt="" aria-hidden="true">
        </div>
        <hgroup id="smk-as-headline-wrapper">
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
            <AmenitiesGallery />
        </div>
        <ul id="smk-as-ctas-container">
            <button class="button-2">
                <span>  
                    Explore
                </span>
                <Arrow />
            </button>
        </ul>
    </div>
</section>

<style>
    #smk-amenities-section {
        display: flex;
        flex-direction: column;
        align-items: center;
        padding: var(--spacing-5) 0;
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

            & img {
                width: 47.45065cqw;
            }
        }
    
    /*=====  End of Mango underlay  ======*/
    
    
    
    
    
    
    
</style>