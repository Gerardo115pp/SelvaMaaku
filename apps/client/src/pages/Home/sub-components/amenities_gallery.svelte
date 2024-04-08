<script>
    import { layout_images } from "@stores/layout";
</script>

<div id="amenities-gallery-wrapper">
    <!-- TODO: Add alt text to these images -->
    <img src="{layout_images.PARK.getUrl(0.3)}" alt="" class="amenities-gallery-image" id="ag-park">
    <img src="{layout_images.RAINY_POOL.getUrl(0.3)}" alt="" class="amenities-gallery-image" id="ag-rainy-pool">
    <img src="{layout_images.STONE_WALL.getUrl(0.3)}" alt="" class="amenities-gallery-image" id="ag-stone-wall">
</div>

<style>
    #amenities-gallery-wrapper {
        position: relative;
        width: min(811px, 66.69407cqw);
        height: min(369px, 30.34539cqw);
    }

    .amenities-gallery-image {
        position: absolute;
        object-fit: cover;
        object-position: center;
    }

    
    /*=============================================
    =            Elements positions            =
    =============================================*/
    
        #ag-park {
            left: 0;
            top: 16.26%;
            width: 24.90752%;
            height: 83.46883%;
            z-index: var(--z-index-b-1);
        }

        #ag-rainy-pool {
            top: 0;
            left: 20.22%;
            width: 59.92601%;
            height: 86.17886%;
            z-index: var(--z-index-b-2);
        }

        #ag-stone-wall {
            right: 0;
            top: 16.26%;
            width: 22.81134%;
            height: 53.38753%;
            z-index: var(--z-index-b-1);
        }
    
    /*=====  End of Elements positions  ======*/
    
    
</style>