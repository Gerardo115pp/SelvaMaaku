<script>
    import AmenitiesGallery from "@pages/Home/sub-components/amenities_gallery.svelte";
    import { layout_images } from "@stores/layout";
    import ThemeButton from "@components/buttons/theme_button.svelte";
</script>

<section id="smk-ap-amenities-section">
    <div id="smk-ap-as-content-wrapper" class="design-content-width">
        <div class="mango-seal-underlay-wrapper">
            <img src="{layout_images.MANGO_ISOLOGO.getUrl(0.3)}" alt="" aria-hidden="true">
        </div>
        <div id="smk-ap-as-gallery-wrapper">
            <AmenitiesGallery />
        </div>
    </div>
</section>

<style>
    #smk-ap-amenities-section {
        display: flex;
        flex-direction: column;
        align-items: center;
        padding: var(--spacing-5) 0;
    }
    
    #smk-ap-as-content-wrapper {
        position: relative;
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: var(--spacing-4);
        container-type: inline-size;
    }

    /*=============================================
    =            Mango underlay            =
    =============================================*/
    
        .mango-seal-underlay-wrapper {
            position: absolute;
            width: max-content;
            right: 0.2cqw;
            top: -10.69cqw;
            z-index: var(--z-index-b-4);

            & img {
                width: 47.45065cqw;
            }
        }
    
    /*=====  End of Mango underlay  ======*/
    
</style>