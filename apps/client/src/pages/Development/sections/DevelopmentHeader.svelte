<script>
    import { layout_images } from "@stores/layout";
</script>

<section id="smk-development-page-header">
    <header id="smk-dp-header" class="design-content-width" role="banner">
        <hgroup id="smk-dp-header-headliners">
            <span id="smk-dp-hh-page-name" class="small-headline">
                DEVELOPMENT
            </span>
            <h1 id="smk-dp-hh-page-headline">
                Comfort and sustainability in harmony
            </h1>
            <p id="smk-dp-hh-page-subheadline">
                Selva Máaku is a 5.4 acres development that is located half a mile from the outskirts of Sayulita on the Nayarit Riviera in Mexico. It is one mile away from the Pacific Ocean and in the midst of a tropical jungle.
            </p>
        </hgroup>
    </header>
    <div id="smk-dp-image-wapper">
        <img src="{layout_images.HERO_VIDEO_COVER.getUrl(1)}" alt="">
    </div>
</section>

<style>
    #smk-development-page-header {
        display: flex;
        padding: var(--spacing-5) 0;
        flex-direction: column;
        gap: var(--spacing-5);
    }

    
    /*=============================================
    =            headlines            =
    =============================================*/
    
        hgroup#smk-dp-header-headliners {
            width: 100%;
            display: flex;
            flex-direction: column;
            align-items: center;
            gap: calc(var(--spacing-2) + var(--spacing-1));
            text-align: center;
        }

        #smk-dp-hh-page-name {
            color: var(--color-primary);
            font-weight: 400;
        }

        h1#smk-dp-hh-page-headline {
            width: min(873px, 66.666%);
            font-size: var(--font-size-h2);
            line-height: var(--line-height-10);
        }

        p#smk-dp-hh-page-subheadline {
            width: min(590px, 48.51976%);
            font-size: var(--font-size-p);
            line-height: var(--line-height-6);
            font-size: 20px;
        }
    
    /*=====  End of headlines  ======*/
    
    
</style>