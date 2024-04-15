<script>
    import ImageMultiStage from "@components/Images/ImageMultiStage.svelte";
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import { layout_images } from "@stores/layout";
    import viewport from "@components/viewport_actions/useViewportActions";
</script>

<section id="smk-home-banner-section" class:adebug={false}>
    <ImageMultiStage 
        image_resource={layout_images.TOMMY_HOUSE_TWO}
        blur_unloaded
    />
    <div id="smk-hbs-information-overlay">
        <div id="smk-hbs-io-content-wrapper">
            <div class="wixarica-icon-wrapper wixarica-icons-large">
                <WixaricaIcon />
            </div>
            <h2 id="smk-hbs-io-headline">
                A low density Urban Residential Environment
            </h2>
        </div>
    </div>
</section>

<style>
    #smk-home-banner-section {
        position: relative;
        height: 90dvh;
    }

    #smk-hbs-information-overlay {
        position: absolute;
        display: flex;
        width: 100%;
        height: 100%;
        justify-content: center;
        top: 0;
        left: 0;
        padding: var(--spacing-5) 0;
        background-color: hsl(from var(--theme-color-dark) h s l / 0.5);
    }

    #smk-hbs-io-content-wrapper {
        display: flex;
        max-width: 797px;
        flex-direction: column;
        align-items: center;
        gap: calc(var(--spacing-3) + var(--spacing-1));
    }

    #smk-hbs-io-headline {
        text-align: center;
        line-height: 72px;
        letter-spacing: -2.2%;
    }
</style>