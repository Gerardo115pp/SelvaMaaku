<script>
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import { layout_images } from "@stores/layout";

</script>
<div id="smk-models-gallery">
    <div id="smk-mg-wixarica-top" class="wixarica-gallery-icon">
        <WixaricaIcon opacity={0.5}  is_flat_color />
    </div>
    <div id="smk-mg-wixarica-bottom" class="wixarica-gallery-icon">
        <WixaricaIcon opacity={0.5}  is_flat_color />
    </div>
    <!-- TODO: Add an alt text to the images -->
    <img src="{layout_images.KEITT_HOUSE.getUrl(0.4)}" alt="" id="smk-mg-keitt-house" class="smk-mg-gallery-image">
    <img src="{layout_images.KEITT_HOUSE_TWO.getUrl(0.3)}" alt="" id="smk-mg-keitt-house-two" class="smk-mg-gallery-image">
    <img src="{layout_images.TOMMY_HOUSE_THREE.getUrl(0.3)}" alt="" id="smk-mg-tommy-house-three" class="smk-mg-gallery-image">
    <img src="{layout_images.TOMMY_HOUSE_FOUR.getUrl(0.4)}" alt="" id="smk-mg-tommy-house-four" class="smk-mg-gallery-image">
    <span class="gallery-label" id="smk-mg-keitt-label">
        Keitt
    </span>
    <span class="gallery-label" id="smk-mg-tommy-label">
        Tommy
    </span>
</div>

<style>
    #smk-models-gallery {
        position: relative;
        width: min(624px, 51.31578cqw);
        height: min(765px, 62.91118cqw);
        container-type: inline-size;
    }

    .smk-mg-gallery-image {
        object-fit: cover;
        object-position: center;
    }

    .wixarica-gallery-icon {
        width: 12.820%;
    }

    .gallery-label {
        font-family: var(--font-titles);
        color: var(--light-green);
        font-size: 15.3846cqw;
        font-weight: 300;
        line-height: 1;
    }

    
    /*=============================================
    =            Items positions            =
    =============================================*/
    
        #smk-mg-wixarica-top {
            position: absolute;
            left: 74.35%;
            top: 0;
        } 

        #smk-mg-wixarica-bottom {
            position: absolute;
            left: 21.47%;
            bottom: 0;
        }

        #smk-mg-keitt-house {
            position: absolute;
            left: 14.42%;
            top: 0;
            width: 52.24358%;
            height: 65.35947%;
        }

        #smk-mg-keitt-house-two {
            position: absolute;
            left: 0;
            top: 11.37%;
            width: 33.17307%;
            height: 18.69281%;
        }

        #smk-mg-tommy-house-three {
            position: absolute;
            left: 0;
            top: 54.37%;
            width: 33.17307%;
            height: 21.56862%;
        }

        #smk-mg-tommy-house-four {
            position: absolute;
            right: 0;
            bottom: 0;
            width: 44.87179%;
            height: 49.6732%;
        }

        #smk-mg-keitt-label {
            position: absolute;
            left: 63.14%;
            top: 20.65%;
        }

        #smk-mg-tommy-label {
            position: absolute;
            left: 0.8%;
            top: 68.62%;
        }


    
    /*=====  End of Items positions  ======*/
    
    
</style>