<script>
    import ThemeButton from "@components/buttons/theme_button.svelte";
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import TownGallery from "../sub-components/town_gallery.svelte";

</script>
<section id="smk-lp-town">
    <div id="smk-lpt-content-wrapper" class="design-content-width">
        <div id="smk-lpt-cw-text-column">
            <div id="smk-lpt-cw-wixarica-underlay">
                <div id="smk-lpt-cw-wixarica-underlay-content">
                    <div id="smk-lpt-gradient-underlay"></div>
                    <div id="smk-lpt-cw-wixarica-underlay-icon-wrapper">
                        <WixaricaIcon opacity={0.4} is_flat_color />
                    </div>
                </div>
            </div>
            <div class="wixarica-icons">
                <WixaricaIcon with_accent />
            </div>
            <hgroup id="smk-lpt-town-headlines">
                <h2 id="smk-lpt-town-headline">
                    A vibrant town 
                </h2>
                <p id="smk-lpt-town-information">
                    Living in Sayulita provides a relaxed beach lifestyle infused with rich cultural charm. It's a place where surfing, outdoor activities, and vibrant street life blend seamlessly, offering a perfect balance of adventure and tranquility in a beautiful coastal setting.
                </p>
            </hgroup>
            <ul id="smk-lpt-town-ctas-container">
                <li class="smk-lpt-town-cta">
                    <ThemeButton text="Life in Sayulita" href="/life-in-sayulita" is_button_three/>
                </li>
            </ul>
        </div>
        <div id="smk-lpt-cw-gallery-column">
            <TownGallery />
        </div>
    </div>
</section>

<style>
    #smk-lp-town {
        position: relative;
        display: flex;
        flex-direction: column;
        align-items: center;
        padding: var(--spacing-5) 0;
        z-index: var(--z-index-1);
    }

    #smk-lpt-content-wrapper {
        --column-gap: var(--spacing-1);
        --column-size: calc(50% - calc(var(--column-gap) / 2));
        display: grid;
        container-type: inline-size;
        grid-template-columns: var(--column-size) var(--column-size);
        gap: var(--column-gap);
    }

    
    /*=============================================
    =            Text Column            =
    =============================================*/
    
        #smk-lpt-cw-text-column {
            position: relative;
            display: flex;
            flex-direction: column;
            align-items: flex-start;
            gap: calc(var(--spacing-3) * 1.25);
        }

        
        /*=============================================
        =            Underlay            =
        =============================================*/
        
            #smk-lpt-cw-wixarica-underlay {
                --left-offset: -10%;
                position: absolute;
                left: var(--left-offset);
                width: calc(100% + calc(var(--left-offset) * -1));
                z-index: var(--z-index-b-1);
            }

            #smk-lpt-gradient-underlay {
                --transparent-color: hsl(from var(--body-bg-color) h s l / 0.7);
                position: absolute;
                top: 0;
                left: 0;
                width: 100%;
                height: 100%;
                background: linear-gradient(180deg, var(--body-bg-color) 0%, var(--body-bg-color) 2%, var(--transparent-color) 100%);
            }

            @supports not (color: rgb(from white r g b)) {
                #smk-lpt-gradient-underlay {
                    background: linear-gradient(180deg, var(--body-bg-color) 0%, var(--body-bg-color) 2%, transparent 100%);
                }
            }


            #smk-lpt-cw-wixarica-underlay-content {
                position: relative;
            }

            #smk-lpt-cw-wixarica-underlay-icon-wrapper {
                position: relative;
                z-index: var(--z-index-b-2);
            }
        
        /*=====  End of Underlay  ======*/

        hgroup#smk-lpt-town-headlines {
            width: min(39.4737cqw, 480px);
            display: flex;
            flex-direction: column;
            gap: var(--spacing-3);
        }

        h2#smk-lpt-town-headline {
            font-size: var(--font-size-h3);
        }
    
    /*=====  End of Text Column  ======*/
    
    #smk-lpt-cw-gallery-column {
        width: 100%;
    }
    
    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @media only screen and (max-width: 768px) {
            #smk-lp-town {
                padding: var(--spacing-4) var(--spacing-2);
            }

            #smk-lpt-content-wrapper {
                width: 100%;
                max-width: none;
                margin: 0;
                grid-template-columns: 1fr;
                gap: var(--spacing-4);
            }

            hgroup#smk-lpt-town-headlines {
                width: 100%;
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    

</style>