<script>
    import { layout_images } from "@stores/layout";
    import { site_language, supported_languages } from "@stores/site_content";
</script>

<section id="smk-development-page-header">
    <header id="smk-dp-header" class="design-content-width" role="banner">
        <hgroup id="smk-dp-header-headliners">
            <span id="smk-dp-hh-page-name" class="small-headline">
                {#if $site_language !== supported_languages.SPANISH}
                    DEVELOPMENT
                {:else}
                    DESARROLLO
                {/if}
            </span>
            <h1 id="smk-dp-hh-page-headline">
                {#if $site_language !== supported_languages.SPANISH}
                    Comfort and sustainability in harmony
                {:else}
                    Comodidad y sostenibilidad en armonía
                {/if}
            </h1>
            <p id="smk-dp-hh-page-subheadline">
                {#if $site_language !== supported_languages.SPANISH}
                    Selva Máaku is a 5.4 acres development that is located half a mile from the outskirts of Sayulita on the Nayarit Riviera in Mexico. It is one mile away from the Pacific Ocean and in the midst of a tropical jungle.
                {:else}
                    Selva Máaku es un desarrollo de 2.2 hectáreas que se encuentra a 800 metros de las afueras de Sayulita en la Riviera Nayarit en México. Está a 1.6 kilómetros del Océano Pacífico y en medio de una selva tropical.
                {/if}
            </p>
        </hgroup>
    </header>
    <div id="smk-dp-image-wapper">
        <img src="{layout_images.DEVELOPMENT_HOUSE_ONE.getUrl(1)}" alt="">
    </div>
</section>

<style>
    #smk-development-page-header {
        display: flex;
        padding: var(--spacing-5) 0;
        flex-direction: column;
        container-type: inline-size;
        gap: var(--spacing-5);
    }

    
    /*=============================================
    =            headlines            =
    =============================================*/
    
        hgroup#smk-dp-header-headliners {
            width: 100%;
            display: flex;
            flex-direction: column;
            align-items: center;
            gap: calc(var(--spacing-2) + var(--spacing-1));
            text-align: center;
        }

        #smk-dp-hh-page-name {
            color: var(--color-primary);
            font-weight: 400;
        }

        h1#smk-dp-hh-page-headline {
            width: min(873px, 66.666%);
            font-size: var(--font-size-h2);
            line-height: var(--line-height-10);
        }

        p#smk-dp-hh-page-subheadline {
            width: min(590px, 48.51976%);
            font-size: var(--font-size-p);
            line-height: var(--line-height-6);
            font-size: 20px;
        }
    
    /*=====  End of headlines  ======*/
    
    #smk-dp-image-wapper {
        width: 100%;
        height: min(640px, 44.44445cqw);
    }

    #smk-dp-image-wapper img {
        width: 100%;
        height: 100%;
        object-fit: cover;
        object-position: center;
    }

    
    /*=============================================
    =            Mobile            =
    =============================================*/

        @media only screen and (max-width: 768px) {
            #smk-development-page-header {
                padding: var(--spacing-4) 0;
                gap: var(--spacing-4    );
            }
        }
    
        @container (width <= 768px) {
            header#smk-dp-header {
                width: 100%;
                max-width: none;
                margin: 0;
                padding: 0 var(--spacing-2);
            }

            hgroup#smk-dp-header-headliners p#smk-dp-hh-page-subheadline {
                width: 100%;
            }

            h1#smk-dp-hh-page-headline {
                font-size: var(--font-size-h1);
                line-height: 1.2;
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    
    
</style>