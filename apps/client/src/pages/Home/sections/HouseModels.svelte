<script>
    import Arrow from "@components/icons/arrow.svelte";
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import ModelsGallery from "../sub-components/models_gallery.svelte";
    import ThemeButton from "@components/buttons/theme_button.svelte";

</script>

<section id="smk-house-models-section" class:adebug={false}>
    <div id="smk-hms-content-wrapper" class="design-content-width">
        <div id="smk-hms-cw-text-column">
            <div class="accent-wixarica-icon-wrapper wixarica-icons">
                <WixaricaIcon with_accent />
            </div>
            <hgroup id="smk-hms-cw-tc-models-information">
                <h2 id="smk-hms-cw-tc-models-headline">
                    MODELS
                </h2>
                <h3 id="smk-hms-cw-tc-models-subheadline">
                    Your house<br/>in the jungle
                </h3>
                <p id="smk-hms-cw-tc-models-text">
                    This unique place where the sun shines bright every day, you will be able to contemplate incredible views of the jungle. The natural light and the shine of the minerals embedded in the stone materials of the area are the inspiration for the design of the 14 houses that make up this exclusive development.
                </p>
            </hgroup>
            <ul class="smk-hms-cw-ctas-wrapper">
                <ThemeButton
                    href="/house-models"
                    text="Models"
                />
            </ul>
        </div>
        <div id="smk-hms-cw-gallery-column">
            <ModelsGallery />
        </div>
    </div>
</section>

<style>
    #smk-house-models-section {
        display: flex;
        justify-content: center;
        padding: var(--spacing-5) 0;
    }

    #smk-hms-content-wrapper {
        display: flex;
        column-gap: var(--spacing-5);
        container-type: inline-size;
    }

    
    /*=============================================
    =            Text column            =
    =============================================*/
    
        #smk-hms-cw-text-column {
            display: flex;
            flex-direction: column;
            gap: var(--spacing-4);
        }

        hgroup#smk-hms-cw-tc-models-information {
            display: flex;
            max-width: 480px;
            flex-direction: column;
            gap: var(--spacing-3);
        }

        hgroup#smk-hms-cw-tc-models-information h2 {
            font-family: var(--font-read);
            font-size: var(--font-size-2);
            font-weight: 600;
        }

        hgroup#smk-hms-cw-tc-models-information h3 {
            font-size: var(--font-size-h3);
            font-weight: 400;
        }
    
    /*=====  End of Text column  ======*/
    
    
</style>