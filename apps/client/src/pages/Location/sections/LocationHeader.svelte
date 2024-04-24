<script>
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import Mexico from "@components/icons/mexico.svelte";
    import SayulitaMap from "@components/icons/sayulita_map.svelte";
</script>

<section id="smk-lp-header">
    <div id="smk-lph-content-wrapper" class="design-content-width">
        <div id="smk-lph-content" >
            <div class="wixarica-icons">
                <WixaricaIcon
                    with_accent
                />
            </div>
            <h1 id="smk-lph-location-name" class="small-headline">
                SAYULITA
            </h1>
            <h2 id="smk-lph-location-slogan">
                The spectacular Hidden Gem of the Riviera Nayarit
            </h2>
            <p id="smk-lph-location-description">
                Sayulita, Mexico, is a vibrant and picturesque small town located on the Pacific coast in the state of Nayarit. Known for its stunning beaches, relaxed atmosphere, and rich cultural heritage.
            </p>
            <div id="smk-lph-underlay-wrapper">
                <WixaricaIcon is_flat_color opacity={0.3}/> 
            </div>
        </div>
    </div>
    <div id="smk-lph-maps-wrapper" class:adebug={false}>
        <div id="smk-lph-maps" class="design-content-width">
            <div id="smk-lph-maps-mexico-map" class="smk-lph-maps-area">
                <Mexico 
                    mexico_opacity={0.4}
                    map_pin_opacity={0.3}
                />
            </div>
            <div id="smk-lph-maps-sayulita-map" class="smk-lph-maps-area">
                <SayulitaMap />
            </div>
        </div>
    </div>
</section>

<style>
    #smk-lp-header {
        display: flex;
        flex-direction: column;
        align-items: center;
        container-type:inline-size;
        padding: var(--spacing-5) 0;
        gap: var(--spacing-5);
    }

    
    /*=============================================
    =            headlines            =
    =============================================*/
    
        #smk-lph-content-wrapper {
            display: flex;
            flex-direction: column;
            align-items: center;
        }

        #smk-lph-content {
            position: relative;
            width: min(65.78947%, 800px);
            display: flex;
            flex-direction: column;
            align-items: center;
            gap: calc(var(--spacing-2) + var(--spacing-1));
            text-align: center;
        }

        #smk-lph-location-slogan {
            font-size: calc(var(--font-size-h2) * 0.99);
            line-height: var(--line-height-10);
            letter-spacing: -2.2%;
            font-weight: 400;
        }

        #smk-lph-underlay-wrapper {
            position: absolute;
            top: 0;
            left: 0;
            width: 100%;
            z-index: var(--z-index-b-3);
            transform-origin: center;
            translate: 0 -25%;
            padding: 6%;
        }
    
    /*=====  End of headlines  ======*/

    
    /*=============================================
    =            Maps            =
    =============================================*/
    
        #smk-lph-maps-wrapper {
            width: 100%;
            padding: var(--spacing-4) 0;
        }
        
        #smk-lph-maps {
            display: flex;
            container-type: inline-size;
            justify-content: space-between;
        }

        .smk-lph-maps-area {
            display: grid;
            width: 45cqw;
            height: 50cqw;
            place-items: center;
        }
            
    
    /*=====  End of Maps  ======*/
    
    
    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @container (width <= 768px) {
            #smk-lph-content {
                width: 100%;
                max-width: none;
                margin: 0;
            }

            #smk-lph-location-slogan {
                font-size: calc(var(--font-size-h1) * 0.99);
                line-height: 1.2;
                letter-spacing: -2.2%;
                font-weight: 400;
            }

            #smk-lph-underlay-wrapper {
                position: absolute;
                top: 50%;
                left: 50%;
                width: 127.55cqw;
                z-index: var(--z-index-b-3);
                transform-origin: center;
                opacity: 0.3;
                translate: -50% -50% 0;
            }
        
            #smk-lph-maps-wrapper {
                width: 100%;
                padding: var(--spacing-4) 0;
            }
            
            #smk-lph-maps {
                display: flex;
                width: 100%;
                max-width: none;
                margin: 0;
                flex-direction: column;
                justify-content: flex-start;
                align-items: center;
                gap: var(--spacing-3);
            }

            .smk-lph-maps-area {
                display: grid;
                width: 80cqw;
                height: auto;
                place-items: center;
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    

</style>