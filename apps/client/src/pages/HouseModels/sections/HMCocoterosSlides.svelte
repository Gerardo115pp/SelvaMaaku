<script>
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import { layout_images, layout_properties } from "@stores/layout";
    import Carousel from "@components/carousel/carousel.svelte";
    import { CarouselItem } from "@components/carousel/types";
    
    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        /**
         * @type {CarouselItem[]}
         */
        const carousel_slides = [
            new CarouselItem({id: 1, src: layout_images.COCOTEROS_HOUSE.getUrl($layout_properties.IS_MOBILE ? 1.5 : 0.7)}),
            new CarouselItem({id: 2, src: layout_images.COCOTEROS_HOUSE_TWO.getUrl($layout_properties.IS_MOBILE ? 1.5 : 0.7)}),
            new CarouselItem({id: 3, src: layout_images.COCOTEROS_HOUSE_THREE.getUrl($layout_properties.IS_MOBILE ? 1.5 : 0.7)}),
        ]
    
    /*=====  End of Properties  ======*/
    
    
</script>

<section id="smk-hmp-cocoteros-slides">
    <div id="smk-hmpcs-content-wrapper" class="design-content-width">
        <div class="wixarica-icons">
            <WixaricaIcon with_accent />
        </div>
        <div id="smk-hmpcs-slides-wrapper">
            <h2 id="smk-hmpcs-carousel-title">
                Cocoteros House Model
            </h2>
            {#if !$layout_properties.IS_MOBILE}
                <Carousel 
                    carousel_items={carousel_slides} 
                    visible_slides={1}
                    slide_height="600px"
                    carousel_gap="0px"
                    autoplay_interval_ms={4000}
                    loop_slides
                    autoplay
                >
                    <svg class='smk-hmpcs-carousel-arrow' id="smk-hmpcs-carousel-arrow-left" slot="arrow_left" viewBox="0 0 48 48">
                        <rect width="48" height="48" rx="24" />
                        <path d="M31 24H17M17 24L24 31M17 24L24 17"/>
                    </svg>
                    <svg class='smk-hmpcs-carousel-arrow' id="smk-hmpcs-carousel-arrow-right" slot="arrow_right" viewBox="0 0 48 48">
                        <rect width="48" height="48" rx="24"/>
                        <path d="M31 24H17M17 24L24 31M17 24L24 17"/>
                    </svg>    
                </Carousel>
            {:else} 
                <Carousel 
                    carousel_items={carousel_slides} 
                    visible_slides={1}
                    slide_height="600px"
                    carousel_gap="0px"
                    autoplay_interval_ms={4000}
                    carousel_padding="0"
                    loop_slides
                    autoplay
                />
            {/if}
        </div>
    </div>
</section>

<style>
    #smk-hmp-cocoteros-slides {
        padding: var(--spacing-5) 0;
        container-type: inline-size;
    }

    #smk-hmpcs-content-wrapper {
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: var(--spacing-4);
    }
    
    #smk-hmpcs-slides-wrapper {
        display: flex;
        width: 100%;
        flex-direction: column;
        align-items: center;
        gap: var(--spacing-3);
    }

    .smk-hmpcs-carousel-arrow {
        width: 48px;
        height: 48px;

        & > rect {
            fill: var(--theme-color-light);
            opacity: 0.5;
            transition: opacity 0.2s ease-in;
        }

        &:hover > rect {
            opacity: 0.9;
        }

        & > path {
            fill: none;
            stroke: var(--theme-color-dark);
            stroke-width: 2;
            stroke-linecap: round;
        }
    }

    #smk-hmpcs-carousel-arrow-right {
        transform: rotate(180deg);
    }
    
    
    /*=============================================
    =            Mobile            =
    =============================================*/

        /* @media only screen and (max-width: 768px) {
            #smk-hmp-cocoteros-slides {
                padding: 0;
            }
        } */
               
    
        @container (width <= 800px) {
            #smk-hmpcs-content-wrapper {
                width: 100%;
                max-width: none;
                margin: 0;
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    
    
</style>