<script>
    import { onDestroy, onMount } from "svelte";
    import { browser } from "$app/environment";

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
         * Array of objects that contain the images of the carousel
         * @type {CarouselItem[]}
         */
        export let carousel_items = [];
        
        /*----------  Behavior props  ----------*/
                

            /**
             * Index of the current slide
             * @type {number}
             */
            export let current_slide_index = 0;

            /**
             * The amount of slides that will be shown at once. If not set then the CarouselItem.width will be used
             * but if set, the CarouselItem.width will be ignored
             * @type {number}
             */
            export let visible_slides;

            /**
             * Whether or not the autoplay is enabled
             * @type {boolean}
             * @default false
             */
            export let autoplay = false;

            /**
             * The interval in milliseconds between each slide
             * @type {number}
             * @default 1000
             */
            export let autoplay_interval_ms = 1000;

            export let is_marquee = false;
            $: if (is_marquee) {
                autoplay = false;
                autoplay_interval_ms = 0;
            }

        /*----------  End of Behavior props  ----------*/

        /**
         * The interval object that will be used to autoplay the carousel. 
         * Use it to clear the interval when the component will be destroyed
         * @type {number}
         */
        let autoplay_interval;

        
        /*----------  Styling  ----------*/
        
            /**
             * The heigh of the slides
             * @type {string}
             * @default "333px"
             */
            export let slide_height = "333px";
        

    /*=====  End of Properties  ======*/ 

    onMount(() => {
        if (autoplay) {
            // autoplay_interval = window.setInterval(handleAutoplay, autoplay_interval_ms);
        }
    });

    onDestroy(() => {
        if (browser && autoplay_interval != null) {
            window.clearInterval(autoplay_interval);
        }
    });

    
    /*=============================================
    =            Methods            =
    =============================================*/

        const handleAutoplay = () => {
            // Check if the autoplay is not enabled. if it is not, clear the interval(if we are in the browser)
            if (!autoplay) {
                if (browser) {
                    window.clearInterval(autoplay_interval);
                }

                return;
            }

            // current_slide_index = (current_slide_index + 1) % carousel_items.length; // let's see if Slidy can handle looping on its own
            current_slide_index++;
        }
    
    /*=====  End of Methods  ======*/
    
    
</script>


{#if carousel_items.length > 0}
    <div class="libery-carousel-component" 
        class:lcc-marquee-carousel={is_marquee}
    >
        <div class="lcc-carousel-mask" class:dtwo={false}>
            <ol class="lcc-carousel-slide-stack">
                {#each carousel_items as slide, index}
                    {@const is_slide_active = index === current_slide_index}
                    <li class="lcc-carousel-slide"
                        style:height="{slide_height}"
                        style:width={visible_slides ? `${100 / visible_slides}cqw` : `${slide.width}px`}    
                        class:lcc-carousel-slide--active={is_slide_active}
                        class:dthree={false}
                    >
                        <img src={slide.src} alt="Slide {index + 1}" />
                    </li>
                {/each}
                {#if is_marquee}
                    {#each carousel_items as slide, index}
                        <li class="lcc-carousel-slide lcc-carousel-slide--doppelganger"
                            style:height="{slide_height}"
                            style:width={visible_slides ? `${100 / visible_slides}cqw` : `${slide.width}px`}    
                            class:dfour={false}
                            aria-hidden="true"
                        >
                            <img src={slide.src} alt="Slide {index + 1}" />
                        </li>
                    {/each}
                {/if}
            </ol>
        </div>
    </div>
{/if}

<style>
    .libery-carousel-component {
        --carousel-gap: var(--spacing-2);
        width: 100%;
        container: inline-size;
        padding: 0 var(--spacing-3);
    }

    .lcc-carousel-mask {
        overflow: hidden;
        container-type: inline-size;
        width: 100%;
        max-width: 100cqw;
        padding: var(--spacing-3) 0;
        mask: linear-gradient(90deg, #ffffff01 2%, #ffffff8c 30%, white 90%, transparent 100%);
    }

    /* use to debug the carousel */
    .lcc-carousel-mask.dtwo {
        overflow: visible;
    }

    @keyframes slide-stack {
        to {
            transform: translateX(calc(-50% - calc(var(--carousel-gap) / 2)));
        }
    }

    ol.lcc-carousel-slide-stack {
        display: flex;
        flex-direction: row;
        width: fit-content;
        flex-wrap: nowrap;
        list-style: none;
        gap: var(--carousel-gap);
        margin: 0;
        padding: 0;
    }

    .lcc-marquee-carousel ol.lcc-carousel-slide-stack {
        animation: slide-stack 60s linear infinite;
    }

    .lcc-carousel-slide {
        flex-shrink: 0;
        box-shadow: 2px 7px 9px 2px rgba(0, 0, 0, 0.15);

        & img {
            width: 100%;
            height: 100%;
            object-fit: cover;
            pointer-events: none;
        }
    }
</style>