<script>
    import Arrow from "@components/icons/arrow.svelte";
    import Mexico from "@components/icons/mexico.svelte";
    import SayulitaMap from "@components/icons/sayulita_map.svelte";
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
</script>

<section id="smk-home-location-section" class:adebug={false}>
    <div id="smk-hls-content-wrapper" class="design-content-width">
        <div id="smk-hls-cw-mexico-map-wrapper" class="smk-hls-bg-element">
            <Mexico
                mexico_opacity={0.4}
                map_pin_opacity={0.3}
            />
        </div>
        <div id="smk-hls-cw-sayulita-map-wrapper" class="smk-hls-bg-element">
            <SayulitaMap />
        </div>
        <div id="smk-hls-cw-content">
            <div class="smk-hls-cw-top-icon wixarica-icons">
                <WixaricaIcon with_accent />
            </div>
            <hgroup id="smk-hls-cw-sayulita-information">
                <h2 id="smk-hls-cw-sayulita-headline">
                    SAYULITA
                </h2>
                <h3 id="smk-hls-cw-sayulita-subheadline">
                    The spectacular Hidden Gem of the Riviera Nayarit
                </h3>
                <p id="smk-hls-cw-sayulita-text">
                    Sayulita, Mexico, is a vibrant and picturesque small town located on the Pacific coast in the state of Nayarit. Known for its stunning beaches, relaxed atmosphere, and rich cultural heritage.
                </p>
            </hgroup>
            <ul class="smk-hls-cw-ctas-wrapper">
                <button class="button-2">
                    <span>Location</span>
                    <Arrow />
                </button>
                <button class="button-2">
                    <span>Life in Sayulita</span>
                    <Arrow />
                </button>
            </ul>
        </div>
    </div>
</section>

<style>
    #smk-home-location-section {
        width: 100%;
        display: flex;
        justify-content: center;
        padding: var(--spacing-5) 0;
    }

    #smk-hls-content-wrapper {
        position: relative;
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        min-height: 567px;
        gap: var(--spacing-6);
    }

    
    /*=============================================
    =            Background elements            =
    =============================================*/
    
        .smk-hls-bg-element {
            visibility: visible;
        }
    
        #smk-hls-cw-mexico-map-wrapper {
            grid-column: 1 / span 1;
            display: flex;
            align-items: center;
        }
    
        #smk-hls-cw-sayulita-map-wrapper {
            grid-column: 2 / span 1;
            display: flex;
            align-items: center;
        }

    /*=====  End of Background elements  ======*/
    
    #smk-hls-cw-content {
        position: absolute;
        display: flex;
        width: 100%;
        inset: 0 0;
        flex-direction: column;
        align-items: center;
        gap: var(--spacing-4);
    }
    
    /*=============================================
    =            Information             =
    =============================================*/

        hgroup#smk-hls-cw-sayulita-information {
            display: flex;
            max-width: 640px;
            flex-direction: column;
            align-items: center;
            gap: var(--spacing-3);
            text-align: center;
        }

        h2#smk-hls-cw-sayulita-headline {
            font-family: var(--font-read);
            font-size: var(--font-size-2);
            font-weight: 600;
        }
        
        h3#smk-hls-cw-sayulita-subheadline {
            font-size: var(--font-size-h3);
        }
    
    /*=====  End of Information  ======*/

    
    /*=============================================
    =            CTA controls            =
    =============================================*/
    
        ul.smk-hls-cw-ctas-wrapper {
            display: flex;
            gap: var(--spacing-2);
        }
    
    /*=====  End of CTA controls  ======*/
    
    
    
    

</style>