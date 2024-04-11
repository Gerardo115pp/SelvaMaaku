<script>
    import Carousel from "@components/carousel/carousel.svelte";
    import { layout_images } from "@stores/layout";

    
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
                src: layout_images.COCOTEROS_HOUSE.getUrl(0.3)
            },
            {
                id: 2,
                src: layout_images.COCOTEROS_HOUSE_TWO.getUrl(0.3)
            },
            {
                id: 3,
                src: layout_images.COCOTEROS_HOUSE_THREE.getUrl(0.3)
            },
            {
                id: 4,
                src: layout_images.COCOTEROS_HOUSE_FOUR.getUrl(0.3)
            }
        ]    
    
    /*=====  End of Properties  ======*/
    
    
</script>

<section id="smk-hmp-header">
    <div id="smk-hmph-content-wrapper" class="design-content-width">
        <hgroup id="smk-hmph-headlines">
            <span class="small-headline">
                MODELS
            </span>
            <h1 id="smk-hmph-headline">
                Your house in the jungle
            </h1>
            <p id="smk-hmph-subheadline">
                This unique place where the sun shines bright every day, you will be able to contemplate incredible views of the jungle. The natural light and the shine of the minerals embedded in the stone materials of the area are the inspiration for the design of the 14 houses that make up this exclusive development.
            </p>
        </hgroup>
    </div>
    <div id="smk-hmph-carousel-wrapper">
        <Carousel 
            carousel_items={carousel_slides}
            autoplay_interval_ms={5000}
            visible_slides={3}
            is_marquee
        />
    </div>
</section>

<style>
    #smk-hmp-header {
        display: flex;
        flex-direction: column;
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

</style>