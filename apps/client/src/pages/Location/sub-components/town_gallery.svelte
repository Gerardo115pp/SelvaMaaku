<script>
    import { layout_images } from "@stores/layout";
</script>

<div id="smk-lpt-gallery-wrapper">
    <img src="{layout_images.SAYULITA_TOWN_ONE.getUrl(0.3)}" alt="sayulita town" class="smk-lpt-gallery-image" id="smk-lpt-gallery-image-one">
    <img src="{layout_images.SAYULITA_TOWN_TWO.getUrl(0.3)}" alt="sayulita town" class="smk-lpt-gallery-image" id="smk-lpt-gallery-image-two">
</div>

<style>
    #smk-lpt-gallery-wrapper {
        position: relative;
        width: 100%;
        height: 647px;
        container-type: size;
    }

    .smk-lpt-gallery-image {
        position: absolute;
        width: 100%;
        height: 100%;
        object-fit: cover;
        object-position: center;
    }

    #smk-lpt-gallery-image-one {
        height: 100cqh;
        width: 52.8363cqw;
    }

    #smk-lpt-gallery-image-two {
        width: 67.74716cqw;
        height: 44.97681cqh;
        right: 0;
        top: 24.37248cqh;
    }

</style>