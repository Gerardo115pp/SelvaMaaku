<script>
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
import HmCharacteristicsGallery from "../sub-components/HMCharacteristicsGallery.svelte";
</script>

<section id="smk-hmp-characteristics-section">
    <div id="smk-hmpc-content-wrapper" class="design-content-width">
        <div id="smk-hmpc-cw-gallery-column">
            <HmCharacteristicsGallery />
        </div>
        <div id="smk-hmpc-cw-text-column">
            <div class="wixarica-icons">
                <WixaricaIcon 
                    with_accent
                />
            </div>
            <hgroup id="smk-hmpc-headlines">
                <h2 id="smk-hmpc-headline">
                    House characteristics
                </h2>
                <p id="smk-hmpc-subheadline">
                    All of our models includes:
                </p>
            </hgroup>
            <ul id="smk-hmpc-characteristics-container">
                <li class="smk-hmpc-characteristic">
                    <div class="wixarica-icons">
                        <WixaricaIcon is_flat_color/>
                    </div>
                    <span>
                        Living room, kitchen, dining room
                    </span>
                </li>
                <li class="smk-hmpc-characteristic">
                    <div class="wixarica-icons">
                        <WixaricaIcon is_flat_color/>
                    </div>
                    <span>
                        2 common bathrooms
                    </span>
                </li>
                <li class="smk-hmpc-characteristic">
                    <div class="wixarica-icons">
                        <WixaricaIcon is_flat_color/>
                    </div>
                    <span>
                        1 Flex (bedroom or study) with terrace
                    </span>
                </li>
                <li class="smk-hmpc-characteristic">
                    <div class="wixarica-icons">
                        <WixaricaIcon is_flat_color/>
                    </div>
                    <span>
                        1 master bedroom with terrace, bathroom and dressing room
                    </span>
                </li>
                <li class="smk-hmpc-characteristic">
                    <div class="wixarica-icons">
                        <WixaricaIcon is_flat_color/>
                    </div>
                    <span>
                        1 bedroom with bathroom and terrace
                    </span>
                </li>
                <li class="smk-hmpc-characteristic">
                    <div class="wixarica-icons">
                        <WixaricaIcon is_flat_color/>
                    </div>
                    <span>
                        1 bedroom 
                    </span>
                </li>
                <li class="smk-hmpc-characteristic">
                    <div class="wixarica-icons">
                        <WixaricaIcon is_flat_color/>
                    </div>
                    <span>
                        Terrace with pool
                    </span>
                </li>
                <li class="smk-hmpc-characteristic">
                    <div class="wixarica-icons">
                        <WixaricaIcon is_flat_color/>
                    </div>
                    <span>
                        Two car garage
                    </span>
                </li>
                <li class="smk-hmpc-characteristic">
                    <div class="wixarica-icons">
                        <WixaricaIcon is_flat_color/>
                    </div>
                    <span>
                        Natural cooling and ventilation systems
                    </span>
                </li>
                <li class="smk-hmpc-characteristic">
                    <div class="wixarica-icons">
                        <WixaricaIcon is_flat_color/>
                    </div>
                    <span>
                        Reduction of the carbon footprint by using local materials and labor
                    </span>
                </li>
                <li class="smk-hmpc-characteristic">
                    <div class="wixarica-icons">
                        <WixaricaIcon is_flat_color/>
                    </div>
                    <span>
                        Waste recycling program, compost for organic waste and separation for inorganic waste.
                    </span>
                </li>
            </ul>
        </div>
    </div>
</section>

<style>
    #smk-hmp-characteristics-section {
        display: flex;
        flex-direction: column;
        align-items: center;
        padding: var(--spacing-5) 0;
    }

    #smk-hmpc-content-wrapper {
        display: flex;
        justify-content: space-between;
        container-type: inline-size;
    }

    
    /*=============================================
    =            Text column            =
    =============================================*/
    
        #smk-hmpc-cw-text-column {
            width: min(42.1052%, 512px);
            display: flex;
            flex-direction: column;
            gap: var(--spacing-3);
        }

        hgroup#smk-hmpc-headlines {
            display: flex;
            flex-direction: column;
            gap: var(--spacing-2);
        }

        h2#smk-hmpc-headline {
            font-size: var(--font-size-h4);
        }

        ul#smk-hmpc-characteristics-container {
            display: flex;
            flex-direction: column;
            gap: var(--spacing-2);
        }

        li.smk-hmpc-characteristic {
            display: flex;
            gap: var(--spacing-2);
            align-items: center;

            & .wixarica-icons {
                flex-shrink: 0;
                width: 24px;
            }
        }
    
    /*=====  End of Text  ======*/
    
    
</style>