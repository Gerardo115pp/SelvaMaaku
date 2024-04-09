<script>
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import { layout_images } from "@stores/layout";
    import LifeGallery from "../sub-components/LifeGallery.svelte";
</script>

<section id="smk-development-page-life-sayulita-section">
    <div id="smk-dplss-content-wrapper" class="design-content-width">
        <div id="smk-dplss-cw-text-column">
            <div class="wixarica-icons">
                <WixaricaIcon with_accent/>
            </div>
            <div id="smk-dplss-cw-text-wrapper">
                <h2 id="smk-dplss-cw-headline">
                    Life in Sayulita, Rewild Concept
                </h2>
                <div id="smk-dplss-cw-text">
                    <p>
                        Mankind in its passage through Nature has left its footprint throughout centuries. Fortunately, the resistance of nature is amazing. We understand the REWILD concept as the opportunity to help our natural for the environment to maintain itself firm in its steps to recuperate vegetation and fauna spaces for the local wildlife.
                    </p>
                    <p>
                        We have our own 2.2 hectare nursery (a surface equal to the surface of the development) to cultivate species of native trees, generate a sanctuary for fowl, replanting our development and others that may come in the near future.
                    </p>
                </div>
            </div>
        </div>
        <div id="smk-dplss-cw-gallery-column">
            <LifeGallery />
        </div>
    </div>
    <div id="smk-dplss-saylita-map-wrapper"><img src="{layout_images.SAYULITA_MAP.getUrl(1)}" alt="sayulita map"></div>
</section>

<style>

    #smk-development-page-life-sayulita-section {
        display: flex;
        flex-direction: column;
        align-items: center;
        padding: var(--spacing-5) 0;
        gap: var(--spacing-5);
    }

    
    /*=============================================
    =            two columns content            =
    =============================================*/
    
        #smk-dplss-content-wrapper {
            display: grid;
            grid-template-columns: 50% 50%;
            container-type: inline-size;
        }

        #smk-dplss-cw-text-column {
            display: flex;
            flex-direction: column;
            align-items: flex-start;
            gap: var(--spacing-3);
        }

        #smk-dplss-cw-text-wrapper {
            width: min(39.4737cqw, 480px);
            display: flex;
            flex-direction: column;
            gap: var(--spacing-2);
        }

        h2#smk-dplss-cw-headline {
            font-size: var(--font-size-h3);
        }

        #smk-dplss-cw-text {
            display: flex;
            flex-direction: column;
            gap: calc(var(--spacing-1) * 1.75);
        }

        #smk-dplss-cw-gallery-column {
            width: 100%;
        }
    
    /*=====  End of two columns content  ======*/

    #smk-dplss-saylita-map-wrapper {
        width: 100%;
    }

    #smk-dplss-saylita-map-wrapper img {
        width: 100%;
        /* TODO: Ask the client if he rathers want the ocean to disappear on larger screens to keep the aspect ratio of the design */
        max-height: 740px;
        object-fit: cover;
        object-position: left bottom;
        pointer-events: none;
    }
    

</style>