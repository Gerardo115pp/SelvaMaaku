<script>
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import { layout_images, layout_properties } from "@stores/layout";
    import LifeGallery from "../sub-components/LifeGallery.svelte";
    import { site_language, supported_languages } from "@stores/site_content";
    import SateliteVideo from "../sub-components/satelite_video.svelte";
</script>

<section id="smk-development-page-life-sayulita-section">
    <div id="smk-dplss-content-wrapper" 
        class:design-content-width={!$layout_properties.IS_MOBILE}
    >
        <div id="smk-dplss-cw-text-column">
            <div class="wixarica-icons">
                <WixaricaIcon with_accent/>
            </div>
            <div id="smk-dplss-cw-text-wrapper">
                <h2 id="smk-dplss-cw-headline">
                    {#if $site_language !== supported_languages.SPANISH}
                        Life in Sayulita, Rewild Concept
                    {:else}
                        Vida en Sayulita, Concepto Rewild
                    {/if}
                </h2>
                <div id="smk-dplss-cw-text">
                    <p>
                        {#if $site_language !== supported_languages.SPANISH}
                            Mankind in its passage through Nature has left its footprint throughout centuries. Fortunately, the resistance of nature is amazing. We understand the REWILD concept as the opportunity to help our natural for the environment to maintain itself firm in its steps to recuperate vegetation and fauna spaces for the local wildlife.
                        {:else}
                            La humanidad en su paso por la Naturaleza ha dejado su huella a lo largo de los siglos. Afortunadamente, la resistencia de la naturaleza es asombrosa. Entendemos el concepto REWILD como la oportunidad de ayudar a nuestra naturaleza para que se mantenga firme en sus pasos para recuperar espacios de vegetación y fauna para la vida silvestre local.
                        {/if}
                    </p>
                    <p>
                        {#if $site_language !== supported_languages.SPANISH}
                            We have our own 2.2 hectare nursery (a surface equal to the surface of the development) to cultivate species of native trees, generate a sanctuary for fowl, replanting our development and others that may come in the near future.
                        {:else}
                            Contamos con nuestro propio vivero de 2.2 hectáreas (una superficie igual a la superficie del desarrollo) para cultivar especies de árboles nativos, generar un santuario para aves, replantar nuestro desarrollo y otros que puedan venir en un futuro cercano.
                        {/if}
                    </p>
                </div>
            </div>
        </div>
        <div id="smk-dplss-cw-gallery-column">
            <LifeGallery />
        </div>
    </div>
    <SateliteVideo />
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
    
    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @media only screen and (max-width: 768px) {
            #smk-development-page-life-sayulita-section {
                display: flex;
                flex-direction: column;
                align-items: center;
                padding: var(--spacing-5) 0 0 0;
            }

            #smk-dplss-content-wrapper {
                width: 100%;
                grid-template-columns: 1fr;
                padding: 0 var(--spacing-2);
                row-gap: var(--spacing-4);
            }

            #smk-dplss-cw-text-wrapper {
                width: 100%;
                display: flex;
                flex-direction: column;
                gap: var(--spacing-2);
            }

            h2#smk-dplss-cw-headline {
                font-size: var(--font-size-h1);
            }

        }
    
    /*=====  End of Mobile  ======*/
    
    

</style>