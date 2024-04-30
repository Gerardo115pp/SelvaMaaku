<script>
    import { layout_images } from "@stores/layout";
    import { site_language, supported_languages } from "@stores/site_content";
</script>

<section id="smk-ap-header">
    <header id="smk-aph-content-wrapper" role="banner">
        <hgroup id="smk-aph-cw-headlines">
            <h1 id="smk-aph-cw-headline" class="small-headline">
                {#if $site_language !== supported_languages.SPANISH}
                    AMENITIES
                {:else}
                    AMENIDADES
                {/if}
            </h1>
            <h2 id="smk-aph-cw-subheadline">
                {#if $site_language !== supported_languages.SPANISH}
                    Where nature meets luxury
                {:else}
                    Donde la naturaleza<br/> se combina con el lujo
                {/if}
            </h2>
            <p id="smk-aph-cw-description">
                {#if $site_language !== supported_languages.SPANISH}
                    Our community has been meticulously designed to cater to a diverse range of lifestyles and interests, ensuring that every day is filled with opportunities for relaxation, recreation, and socialization.
                {:else}
                    Nuestra comunidad ha sido meticulosamente diseñada para satisfacer una amplia gama de estilos de vida e intereses, asegurando que cada día esté lleno de oportunidades para la relajación, la recreación y la socialización.
                {/if}
            </p>
        </hgroup>
    </header>
    <div id="smk-apg-teaser-screen">
        <!-- TODO: A static image is a bit of a let down even if that's what appears in the design. so let's think of something to improve this section -->
        <img src="{layout_images.RAINY_POOL_TWO.getUrl(1)}" alt="">
    </div>
</section>

<style>
    #smk-ap-header {
        display: flex;
        flex-direction: column;
        padding: var(--spacing-5) 0;        
        container-type: inline-size;
        gap: var(--spacing-5);
    }

    
    /*=============================================
    =            Header            =
    =============================================*/
    
        header#smk-aph-content-wrapper {
            display: flex;
            flex-direction: column;
            align-items: center;
        }

        #smk-aph-cw-headlines {
            width: 100%;
            display: flex;
            flex-direction: column;
            align-items: center;
            gap: var(--spacing-3);
            text-align: center;
        }

        h2#smk-aph-cw-subheadline {
            font-size: var(--font-size-h2);
        }

        p#smk-aph-cw-description {
            max-width: 590px;
        }
    
    /*=====  End of Header  ======*/

    #smk-apg-teaser-screen {
        width: 100%;
        height: 640px;
        
        & img {
            width: 100%;
            height: 100%;
            object-fit: cover;
            object-position: center;
        }
    }

    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @container (width <= 800px) {
            #smk-aph-content-wrapper {
                padding: 0 var(--spacing-2);
            }

            h2#smk-aph-cw-subheadline {
                font-size: var(--font-size-h1);
            }

            #smk-apg-teaser-screen {
                height: 320px;
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    
    
    
</style>