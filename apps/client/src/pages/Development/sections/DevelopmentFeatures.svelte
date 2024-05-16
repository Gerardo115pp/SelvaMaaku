<script>
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import { layout_images, layout_properties } from "@stores/layout";
    import { site_language, supported_languages } from "@stores/site_content";

    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        /**
         * Development properties
         * @typedef {Object} DevelopmentProperties
         * @property {string} en_name - The name of the development
         * @property {string} es_name - The name of the development in Spanish
         * @property {string} en_description - The description of the development in English
         * @property {string} es_description - The description of the development in Spanish
         */

        let development_properties = [
            {
                en_name: "Homes Available",
                es_name: "Casas Disponibles",
                en_description: "Homes available for pre-sale in the Selva Maaku community.",
                es_description: "Casas disponibles en preventa en la comunidad de Selva Maaku."
            },
            {
                en_name: "Cuztumizer",
                es_name: "Cuztumizer",
                en_description: "Customize your dream home with our Cuztumizer feature.",
                es_description: "Personaliza tu casa de ensueño con nuestra función Cuztumizer."
            },
            {
                en_name: "Location",
                es_name: "Ubicación",
                en_description: "Conveniently located 700 meters from downtown and just 1 mile from Sayulita Beach.",
                es_description: "Convenientemente ubicado a 700 metros del centro y a solo 1 milla de la playa de Sayulita."
            },
            {
                en_name: "Lifestyle",
                es_name: "Estilo de Vida",
                en_description: "Live the Selva Máaku lifestyle where every feels like a vacación.",
                es_description: "Vive el estilo de vida de Selva Máaku donde todo se siente como unas vacaciones."
            }

        ];
    
    /*=====  End of Properties  ======*/
    
    

</script>

<section id="smk-development-page-features-section">
    <header id="smk-development-page-features-header"
        class:design-content-width={!$layout_properties.IS_MOBILE}
    >
        <div class="wixarica-icons">
            <WixaricaIcon with_accent/>
        </div>
        <hgroup id="smk-dp-header-headliners">
            <h2 id="smk-development-page-features-headline">
                {#if $site_language !== supported_languages.SPANISH}
                    A low density Urban Residential Environment
                {:else}
                    Un entorno residencial urbano de baja densidad
                {/if}
            </h2>
            <p id="smk-development-page-features-subheadline">
                {#if $site_language !== supported_languages.SPANISH}
                    The architectural project contemplates only 14 houses, that along with the already existing 12 provide generous spaces for ideal gardens to cultivate green areas and reforest this space with Tommy Atkins and Keiit Mango varieties, besides Papelillo Rojo, Prickly Pear, Neem, Eardrop Tree, amongst other native species in the area.
                {:else}
                    El proyecto arquitectónico contempla únicamente 14 casas, que junto con las 12 ya existentes brindan espacios generosos para jardines ideales para cultivar áreas verdes y reforestar este espacio con variedades de Mango Tommy Atkins y Keiit, además de Papelillo Rojo, Nopal, Neem, Oreja de Burro, entre otras especies nativas de la zona.
                {/if}
            </p>
        </hgroup>
    </header>
    <div id="smk-development-features-wrapper" class="design-content-width">
        <ul id="smk-dp-features-container" class:adebug={false}>
            {#each development_properties as dp, h}
                <li class="smk-dp-feature-item" id="smk-dp-feature-item-{h+1}">
                    <div class="wixarica-icons">
                        <WixaricaIcon is_flat_color opacity={0.5}/>
                    </div>
                    <hgroup>
                        <h3 class="smk-fi-feature-name">
                            {#if $site_language !== supported_languages.SPANISH}
                                {dp.en_name}
                            {:else}
                                {dp.es_name}
                            {/if}
                        </h3>
                        <p class="smk-fi-feature-description">
                            {#if $site_language !== supported_languages.SPANISH}
                                {dp.en_description}
                            {:else}
                                {dp.es_description}
                            {/if}
                        </p>
                    </hgroup>
                </li>
            {/each}
            <li class="smk-dp-features-image-wrapper"><img src="{layout_images.HOUSE_MODEL_ONE.getUrl(0.3)}" alt="selvamaaku house model"></li>
        </ul>
    </div>
</section>

<style>
    #smk-development-page-features-section {
        display: flex;
        flex-direction: column;
        align-items: center;
        container-type: inline-size;
        padding: var(--spacing-5) 0;
        row-gap: var(--spacing-5);
    }
    
    /*=============================================
    =            headlines            =
    =============================================*/
    
        header#smk-development-page-features-header {
            width: min(812px, 66.77635%);
            display: flex;
            flex-direction: column;
            align-items: center;
            gap: var(--spacing-3);
        }

        hgroup#smk-dp-header-headliners {
            width: 100%;
            display: flex;
            flex-direction: column;
            align-items: center;
            gap: var(--spacing-2);
            text-align: center;
        }

        h2#smk-development-page-features-headline {
            font-size: var(--font-size-h3);
            line-height: var(--line-height-9);
        }
    
    /*=====  End of headlines  ======*/

    
    /*=============================================
    =            features            =
    =============================================*/
    
        ul#smk-dp-features-container {
            width: 100%;
            height: 640px;
            display: grid;
            grid-template: 50% 50% / repeat(3, 1fr);
        }

        li.smk-dp-feature-item {
            display: flex;
            flex-direction: column;
            align-items: center;
            gap: var(--spacing-2);
            padding: var(--spacing-3);

            & hgroup {
                display: flex;
                text-align: center;
                flex-direction: column;
                gap: var(--spacing-1);
            }

            & hgroup h3 {
                font-size: var(--font-size-4);
                line-height: var(--line-height-8);
            }

            & hgroup p {
                font-size: var(--font-size-1);
            }
        }

        #smk-dp-feature-item-1 {
            grid-area: 1 / 1 / span 1 / span 1;
            justify-content: flex-end;
        }

        #smk-dp-feature-item-2 {
            grid-area: 1 / 3 / span 1 / span 1;
            justify-content: flex-end;
        }

        #smk-dp-feature-item-3 {
            grid-area: 2 / 1 / span 1 / span 1;
        }

        #smk-dp-feature-item-4 {
            grid-area: 2 / 3 / span 1 / span 1;
        }

        li.smk-dp-features-image-wrapper {
            grid-area: 1 / 2 / span 2 / span 1;

            & img {
                width: 100%;
                height: 100%;
                object-fit: cover;
                object-position: center;
            }
        }
    
    /*=====  End of features  ======*/
    
    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @container (width <= 768px) {
            header#smk-development-page-features-header {
                width: 100%;
                padding: 0 var(--spacing-2);
            }

            h2#smk-development-page-features-headline {
                font-size: calc(var(--font-size-h1) * 0.82); 
                line-height: 1.2;
            }

            ul#smk-dp-features-container {
                height: auto;
                grid-template: none / 1fr;
                grid-auto-rows: 250px;
            }
    

            #smk-dp-feature-item-1 {
                grid-area: auto / auto / auto / auto;
                justify-content: flex-end;
            }

            #smk-dp-feature-item-2 {
                grid-area: auto / auto / auto / auto;
                justify-content: flex-end;
            }

            #smk-dp-feature-item-3 {
                grid-area: auto / auto / auto / auto;
            }

            #smk-dp-feature-item-4 {
                grid-area: auto / auto / auto / auto;
            }

            li.smk-dp-features-image-wrapper {
                grid-area: 3 / auto / span 2 / auto;
            }

            li.smk-dp-features-image-wrapper img {
                width: 100%;
                height: 100%;
                object-fit: cover;
                object-position: center;
            }

        }
    
    /*=====  End of Mobile  ======*/
    
    
    
    

</style>