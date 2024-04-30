<script>
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import HmCharacteristicsGallery from "../sub-components/HMCharacteristicsGallery.svelte";
    import { site_language, supported_languages } from "@stores/site_content";

    
    /*=============================================
    =            Properties            =
    =============================================*/

        /**
         * @typedef {Object} HouseCharacteristic
         * @property {string} english_text
         * @property {string} spanish_text
         */
    
        /**
         * The characteristcs of selva maaku's houses
         * @type {HouseCharacteristic[]}
         */
        const house_characteristics = [
            {
                english_text: "Living room, kitchen, dining room",
                spanish_text: "Sala, cocina, comedor"
            },
            {
                english_text: "2 common bathrooms",
                spanish_text: "2 baños comunes"
            },
            {
                english_text: "1 Flex (bedroom or study) with terrace",
                spanish_text: "1 Flex (recámara o estudio) con terraza"
            },
            {
                english_text: "1 master bedroom with terrace, bathroom and dressing room",
                spanish_text: "1 recámara principal con terraza, baño y vestidor"
            },
            {
                english_text: "1 bedroom with bathroom and terrace",
                spanish_text: "1 recámara con baño y terraza"
            },
            {
                english_text: "1 bedroom",
                spanish_text: "1 recámara"
            },
            {
                english_text: "Terrace with pool",
                spanish_text: "Terraza con alberca"
            },
            {
                english_text: "Two car garage",
                spanish_text: "Cochera para dos autos"
            },
            {
                english_text: "Natural cooling and ventilation systems",
                spanish_text: "Sistemas de enfriamiento y ventilación naturales"
            },
            {
                english_text: "Reduction of the carbon footprint by using local materials and labor",
                spanish_text: "Reducción de la huella de carbono al utilizar materiales y mano de obra locales"
            },
            {
                english_text: "Waste recycling program, compost for organic waste and separation for inorganic waste.",
                spanish_text: "Programa de reciclaje de residuos, composta para residuos orgánicos y separación para residuos inorgánicos."
            }
        ];
    
    /*=====  End of Properties  ======*/
    
    
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
                    {#if $site_language !== supported_languages.SPANISH}
                        House characteristics
                    {:else}
                        Características de las casas
                    {/if}
                </h2>
                <p id="smk-hmpc-subheadline">
                    {#if $site_language !== supported_languages.SPANISH}
                        All of our models includes:
                    {:else}
                        Todos nuestros modelos incluyen:
                    {/if}
                </p>
            </hgroup>
            <ul id="smk-hmpc-characteristics-container">
                {#each house_characteristics as hc}
                     <li class="smk-hmpc-characteristic">
                         <div class="wixarica-icons">
                             <WixaricaIcon is_flat_color/>
                         </div>
                         <span>
                            {#if $site_language !== supported_languages.SPANISH}
                                {hc.english_text}
                            {:else}
                                {hc.spanish_text}
                            {/if}
                         </span>
                     </li>
                {/each}
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
    
    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @media only screen and (max-width: 880px) {
            #smk-hmpc-content-wrapper {
                width: 100%;
                max-width: none;
                margin: 0;
                flex-direction: column;
                padding: 0 var(--spacing-2);
                gap: var(--spacing-4);
            }
        }

        @container (width <= 800px) {
            #smk-hmpc-cw-text-column {
                width: 100%;
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    
    
</style>