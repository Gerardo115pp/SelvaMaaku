<script>
    import Carousel from "@components/carousel/carousel.svelte";
    import { layout_images, layout_properties } from "@stores/layout";
    import { site_language, supported_languages } from "@stores/site_content";

    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        /**
         * @typedef {Object} CarouselItem
         * @property {number} id
         * @property {number} width
         * @property {number} height
         * @property {string} src
         */
    
        /**
         * @type {CarouselItem[]}
         */
        const carousel_slides = [
            {
                id: 1,
                src: layout_images.COCOTEROS_HOUSE.getUrl($layout_properties.IS_MOBILE ? 1.3 : 0.3)
            },
            {
                id: 2,
                src: layout_images.COCOTEROS_HOUSE_TWO.getUrl($layout_properties.IS_MOBILE ? 1.3 : 0.3)
            },
            {
                id: 3,
                src: layout_images.COCOTEROS_HOUSE_THREE.getUrl($layout_properties.IS_MOBILE ? 1.3 : 0.3)
            },
            {
                id: 4,
                src: layout_images.COCOTEROS_HOUSE_FOUR.getUrl($layout_properties.IS_MOBILE ? 1.3 : 0.3)
            }
        ] 
    
    /*=====  End of Properties  ======*/
    
    
</script>

<section id="smk-hmp-header">
    <div id="smk-hmph-content-wrapper" class="design-content-width">
        <hgroup id="smk-hmph-headlines">
            <span class="small-headline">
                {#if $site_language !== supported_languages.SPANISH}
                    MODELS
                {:else}
                    MODELOS
                {/if}
            </span>
            <h1 id="smk-hmph-headline">
                {#if $site_language !== supported_languages.SPANISH}
                    Your house in the jungle
                {:else}
                    Tu casa en la selva
                {/if}
            </h1>
            <p id="smk-hmph-subheadline">
                {#if $site_language !== supported_languages.SPANISH}
                    This unique place where the sun shines bright every day, you will be able to contemplate incredible views of the jungle. The natural light and the shine of the minerals embedded in the stone materials of the area are the inspiration for the design of the 14 houses that make up this exclusive development.
                {:else}
                    Este lugar único donde el sol brilla con fuerza todos los días, podrás contemplar increíbles vistas de la selva. La luz natural y el brillo de los minerales incrustados en los materiales pétreos de la zona son la inspiración para el diseño de las 14 casas que conforman este exclusivo desarrollo.
                {/if}
            </p>
        </hgroup>
    </div>
    <div id="smk-hmph-carousel-wrapper">
        <Carousel 
            carousel_items={carousel_slides}
            carousel_padding={$layout_properties.IS_MOBILE ? "0" : "0 var(--spacing-3)"}
            visible_slides={$layout_properties.IS_MOBILE ? 1 : 3}
            use_mask={!$layout_properties.IS_MOBILE}
            autoplay_interval_ms={5000}
            is_marquee
        />
    </div>
</section>

<style>
    #smk-hmp-header {
        display: flex;
        flex-direction: column;
        container-type: inline-size;
        align-items: center;
        padding: var(--spacing-5) 0;
        gap: var(--spacing-5);
    }

    
    /*=============================================
    =            Headlines            =
    =============================================*/
    
        #smk-hmph-content-wrapper {
            display: flex;
            flex-direction: column;
            align-items: center;
        }

        #smk-hmph-headlines {
            width: 100%;
            display: flex;
            flex-direction: column;
            align-items: center;
            gap: var(--spacing-3);
            text-align: center;
        }

        #smk-hmph-headline {
            width: min(873px, 66.666%);
            font-size: var(--font-size-h2);
            line-height: var(--line-height-10);
        }

        #smk-hmph-subheadline {
            width: min(590px, 48.51976%);
            font-size: var(--font-size-p);
            line-height: var(--line-height-5);
        }
    
    
    /*=====  End of Headlines  ======*/
    
    #smk-hmph-carousel-wrapper {
        width: 100%;
    }

    
    /*=============================================
    =            Mobile            =
    =============================================*/

        @media only screen and (max-width: 768px) {
            #smk-hmp-header {
                gap: var(--spacing-4);
            }
        }
    
        @container (width <= 1200px) {

        }

        @container (width <= 1200px) {
            #smk-hmph-content-wrapper {
                width: 100%;
                max-width: none;
                margin: 0;
                padding: 0 var(--spacing-2);
            }

            h1#smk-hmph-headline {
                font-size: var(--font-size-h1);
                line-height: 1.2;
            }

            p#smk-hmph-subheadline {
                width: 100%;
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    

</style>