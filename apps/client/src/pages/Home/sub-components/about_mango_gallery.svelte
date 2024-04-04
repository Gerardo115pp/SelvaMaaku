<script>
    import { layout_images } from "@stores/layout";
</script>

<div id="smk-mango-gallery">
    <img src="{layout_images.KEITT_HOUSE.getUrl(0.42)}" alt="house keitt" class="mango-gallery-image" id="smg-keitt-house">
    <img src="{layout_images.TOMMY_HOUSE.getUrl(0.42)}" alt="house tommy" class="mango-gallery-image" id="smg-tommy-house">
    <img  src="{layout_images.DEVELOPMENT_HOUSE_ONE.getUrl(0.42)}" alt="house in development" class="mango-gallery-image" id="smg-development-house-one">
    <p id="smg-gallery-caption">
        The word Máaku means MANGO in the Wixárika language that is spoken by the original towns in the area.
    </p>
</div>

<style>
    #smk-mango-gallery {
        display: grid;
        height: 100%;
        width: 100%;
        grid-template: repeat(21, 1fr) / repeat(21, 1fr);
    }

    .mango-gallery-image {
        width: 100%;
        height: 100%;
        object-fit: cover;
        object-position: center;
    }

    #smg-keitt-house {
        grid-column: 1 / span 8;
        grid-row: 1 / span 15;
    }

    #smg-tommy-house {
        grid-column: 13 / -1;
        grid-row: 1 / span 11;
    }

    #smg-gallery-caption {
        grid-column: 1 / span 9;
        grid-row: 18 / -1;
    }

    #smg-development-house-one {
        grid-column: 13 / -1;
        grid-row: 14 / -1;
    }


</style>