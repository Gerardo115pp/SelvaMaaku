<script>
    import SayulitaColorMap from "@components/icons/sayulita_color_map.svelte";
import WixaricaIcon from "@components/icons/wixarica_icon.svelte";

</script>
<section id="smk-lp-benefits" class:adebug={false}>
    <div id="smk-lpb-content-wrapper" class="design-content-width">
        <div id="smk-lpb-cw-text-column">
            <div class="wixarica-icons">
                <WixaricaIcon with_accent />
            </div>
            <hgroup id="smk-lpb-cw-headlines">
                <h2 id="smk-lpb-cw-headline">
                    Five minutes from downtown
                </h2>
                <div id="smk-lpb-cw-information">
                    Selva Maakú is located just five minutes from downtown Sayulita. It offers a peaceful retreat near the town’s vibrant beaches, shops, and restaurants, combining tranquility with easy access to local amenities and culture.
                </div>
            </hgroup>
        </div>
        <div id="smk-lpb-cw-sayulita-map">
            <SayulitaColorMap />
        </div>
    </div>
</section>

<style>
    #smk-lp-benefits {
        display: flex;
        flex-direction: column;
        align-items: center;
        padding: var(--spacing-5) 0;
        gap: var(--spacing-5);
    }

    #smk-lpb-content-wrapper {
        display: flex;
        container-type: inline-size;
        justify-content: space-between;
    }

    #smk-lpb-content-wrapper > * {
        max-width: 50%;
    }

    
    /*=============================================
    =            Text column            =
    =============================================*/
    
        #smk-lpb-cw-text-column {
            display: flex;
            flex-direction: column;
            align-items: flex-start;
            gap: calc(var(--spacing-3) + var(--spacing-1));
        }

        #smk-lpb-cw-headlines {
            width: min(39.4737cqw, 480px);
            display: flex;
            flex-direction: column;
            gap: var(--spacing-2);
        }

        h2#smk-lpb-cw-headline {
            font-size: var(--font-size-h3);
        }

        #smk-lpb-cw-information {
            font-size: var(--font-size-2);
        }


    
    /*=====  End of Text column  ======*/
    
    

    #smk-lpb-content-wrapper > #smk-lpb-cw-sayulita-map {
        width: 50cqw;
    }
</style>