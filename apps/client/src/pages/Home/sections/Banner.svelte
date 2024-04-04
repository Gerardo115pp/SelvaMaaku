<script>
    import ImageMultiStage from "@components/Images/ImageMultiStage.svelte";
    import { layout_images } from "@stores/layout";
</script>

<section id="smk-home-banner-section">
    <ImageMultiStage 
        image_resource={layout_images.TOMMY_HOUSE_TWO}
        blur_unloaded
    />
</section>

<style>
    #smk-home-banner-section {
        min-height: 720px;
    }
</style>