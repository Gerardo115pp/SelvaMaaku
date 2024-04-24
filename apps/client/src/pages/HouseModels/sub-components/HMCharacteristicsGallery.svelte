<script>
    import { layout_images } from "@stores/layout";
    import ThemeButton from "@components/buttons/theme_button.svelte";
</script>

<div class="smk-hmpc-gallery-wrapper">
    <img src="{layout_images.COCOTEROS_HOUSE_FIVE.getUrl(0.4)}" alt="cocoteros house living room" id="smk-hmpc-gallery-item-1" class="smk-hmpc-gallery-item">
    <img src="{layout_images.COCOTEROS_HOUSE_FOUR.getUrl(0.3)}" alt="cocoteros house kitchen" id="smk-hmpc-gallery-item-2" class="smk-hmpc-gallery-item">
    <ThemeButton text="Open 360° View" is_button_one href="https://armada.com.mx/sitio/clientes/selva-maaku/visor-360/" open_in_new_tab/>
</div>

<style>
    .smk-hmpc-gallery-wrapper {
        --hmpc-gallery-width: 46.46381cqw;
        position: relative;
        width: min(565px, var(--hmpc-gallery-width));
        height: min(629px, calc(var(--hmpc-gallery-width) * 1.113274));
        container: inline-size;
    }

    .smk-hmpc-gallery-wrapper > * {
        position: absolute;
        object-fit: cover;
        object-position: center;
    }

    #smk-hmpc-gallery-item-1 {
        top: 0;
        left: 0;
        width: 67.43362%;
        height: 100%;
    }

    #smk-hmpc-gallery-item-2 {
        top: 18.44197%;
        right: 0;
        width: 44.77876%;
        height: 48.01271%;
    }

    :global(.smk-hmpc-gallery-wrapper > :not(img)) {
        position: absolute;
        top: 79.17329%;
        left: 55.22123%;
        z-index: var(--z-index-t-1);
    }

    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @container (width <= 800px) {
            .smk-hmpc-gallery-wrapper {
                --hmpc-gallery-width: 75cqw;
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    


</style>