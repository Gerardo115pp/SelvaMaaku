<script>
    import SayulitaColorMap from "@components/icons/sayulita_color_map.svelte";
import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import { site_language, supported_languages } from "@stores/site_content";

</script>

<section id="smk-lp-benefits" class:adebug={false}>
    <div id="smk-lpb-content-wrapper" class="design-content-width">
        <div id="smk-lpb-cw-text-column">
            <div class="wixarica-icons">
                <WixaricaIcon with_accent />
            </div>
            <hgroup id="smk-lpb-cw-headlines">
                <h2 id="smk-lpb-cw-headline">
                    {#if $site_language !== supported_languages.SPANISH}
                        Five minutes from downtown
                    {:else}
                        A cinco minutos del centro
                    {/if}
                </h2>
                <div id="smk-lpb-cw-information">
                    {#if $site_language !== supported_languages.SPANISH}
                        Selva Maakú is located just five minutes from downtown Sayulita. It offers a peaceful retreat near the town’s vibrant beaches, shops, and restaurants, combining tranquility with easy access to local amenities and culture.
                    {:else}
                        Selva Maakú se encuentra a solo cinco minutos del centro de Sayulita. Ofrece un refugio tranquilo cerca de las vibrantes playas, tiendas y restaurantes del pueblo, combinando la tranquilidad con un fácil acceso a las comodidades y la cultura locales.
                    {/if}
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

    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @media only screen and (max-width: 768px) {
            #smk-lp-benefits {
                padding: var(--spacing-4) var(--spacing-2);
            }

            #smk-lpb-content-wrapper {
                width: 100%;
                max-width: none;
                margin: 0;
                flex-direction: column;
                gap: var(--spacing-4);
            }

            #smk-lpb-cw-text-column {
                max-width: 100%;
                width: 100%;
            }

            hgroup#smk-lpb-cw-headlines {
                width: 100%;
            }

            #smk-lpb-content-wrapper #smk-lpb-cw-sayulita-map {
                width: 100cqw;
                max-width: none;
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    
</style>