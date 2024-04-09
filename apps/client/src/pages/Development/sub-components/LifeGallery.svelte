<script>
    import { layout_images } from "@stores/layout";
</script>

<div class="smk-dp-sayulita-life-wrapper">
    <img src="{layout_images.JUNGLE_ONE.getUrl(0.22)}" alt="sayulita jungle tree" class="sayulita-jungle">
    <img src="{layout_images.MANGO_ISOLOGO.getUrl(0.3)}" aria-hidden="true" alt="selvamaaku isologo" class="mango-seal-underlay-wrapper">
</div>

<style>
    .smk-dp-sayulita-life-wrapper {
        position: relative;
        container-type: inline-size;
        width: 100%;
        height: max-content;
    }
    
    .sayulita-jungle {
        width: 52.07667cqw;
        height: 103.3546cqw;
    }

    .mango-seal-underlay-wrapper {
        position: absolute;
        width: 73.96166cqw;
        height: 70.171cqw;
        top: 50%;
        right: 0;
        transform: translateY(-50%);
        z-index: var(--z-index-b-2);
    }
</style>